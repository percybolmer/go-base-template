package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCMD is the root of all Commands, the starting point.
// SubCommands are added to the rootCMD
var rootCmd = &cobra.Command{
	Use:   "base",
	Short: "This is a base template, an example of how to structure a project in Go",
	Long: `This project is structured with cobra

  The project uses only the Cobra Command tooling, not the Cobra CLI to generate code.

  The structure is Easy 

  the Folder cmd contains the Commads that are executable, such as a start command, or a small script
  Think of it as a bin folder.

  The domain/internal folder contains code/libraries that are used for certain domains
  such as the example ModBus Package

  Make sure to place all code that are related to ONE domain i the same package and dont share stuff across domain to avoid
  Coupling

  `,
}

// Execute simle runs the CMD of the Root Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
