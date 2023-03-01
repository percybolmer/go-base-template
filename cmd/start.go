package cmd

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"programmigpercy.tech/go-base-template/domain/modbus"
	"programmigpercy.tech/go-base-template/domain/mysql"
	"programmigpercy.tech/go-base-template/services"
)

var (

	// startCmd is the Start Command
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts a long running scraper of data",
		Long:  `Will fetch ddata frmo ModBus and print it`,
		RunE:  startModBusScrapers,
	}
	// modBusIP is the IP addr to modbus, an example input Parameter to the stat command
	modBusIP string
)

func init() {
	startCmd.Flags().StringVarP(&modBusIP, "modbusIP", "m", "", "The IP address of the ModBus unit")
	// set flag as required
	startCmd.MarkFlagRequired("modbusIP")
	// Add the Command to the rootCMD to enable it
	rootCmd.AddCommand(startCmd)
}

func startModBusScrapers(cmd *cobra.Command, args []string) error {
	// Handle Innput arguments
	ipAddr := net.ParseIP(modBusIP)
	// Configure a Logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msg("Welcome to an Example Base Project")
	// Connect the needed clients
	modbusClient, err := modbus.Connnect(ipAddr)
	if err != nil {
		return err
	}

	mysqlClient, err := mysql.Connect()
	if err != nil {
		return err
	}

	// Create a Service, used to Combine Clients and Data, this is where we can Couple things
	dataService := services.NewDataCollectionService(*modbusClient, mysqlClient)

	// Done is a channel used to notify exit to background processes
	done := make(chan bool)
	// Fetch data in the background each 5th second
	go dataService.PollForData(done)
	// Run until operating system Signals a SIGTerm
	quitChannel := make(chan os.Signal, 1)

	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	// Block until exit
	<-quitChannel
	// Cleanup any GoRoutines running in the background etx etc
	done <- true
	dataService.Close()
	log.Info().Msg("Exiting program, Bye!")

	return nil
}
