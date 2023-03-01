package services

import (
	"time"

	"github.com/rs/zerolog/log"
	"programmigpercy.tech/go-base-template/domain/modbus"
	"programmigpercy.tech/go-base-template/domain/mysql"
)

// DataCollection is a service that fetches data and stores it in a DB
type DataCollection struct {
	modbusClient modbus.Client
	database     mysql.Client
}

func NewDataCollectionService(modbusClient modbus.Client, database mysql.Client) DataCollection {
	return DataCollection{
		modbusClient: modbusClient,
		database:     database,
	}
}

// Close will cancel all connections etc
func (dc DataCollection) Close() {
	// dc.database.Close()
	// dc.modbusClient.Close()
}

// PollForData is used to fetch data from the modbus and store it in a db
func (dc DataCollection) PollForData(done chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	// For select will run until the DONE channel recieves a value
	for {
		select {
		// incase Timmer triggers an event
		case <-ticker.C:
			data, err := dc.modbusClient.FetchData()
			if err != nil {
				log.Error().Str("function", "modbusClient.FetchData").Err(err).Msg("Failed to fetch data")
				continue
			}

			if err := dc.database.InsertData(data.BatteryLevel); err != nil {
				log.Error().Str("function", "database.InsertData").Err(err).Msg("Failed to save data")
				continue
			}

			log.Info().Int("batteryLevel", data.BatteryLevel).Msg("New Data fetched")
		case <-done:
			log.Info().Msg("Closing PollForData")

			return
		}
	}

}
