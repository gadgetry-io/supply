package main

import (
	"fmt"

	"github.com/gadgetry-io/supply/flog"
	"github.com/gadgetry-io/supply/log"
	"github.com/spf13/cobra"
)

// CmdExample represents the example command
var CmdExample = &cobra.Command{
	Use:   "example",
	Short: "Short Description",
	Long:  "Long Description",
	RunE:  runExample,
}

type exampleFlags struct {
	name string
}

var exampleOptions exampleFlags

func init() {
	// Cobra also supports local flags, which will only run when this action is called directly.
	CmdExample.Flags().StringVar(&exampleOptions.name, "name", "", "The Name to give this example")
}

func runExample(cmd *cobra.Command, args []string) error {

	// LOG Messages (STANDARD)
	fmt.Printf("\n\n=== LOG Messages (STANDARD) ===\n")
	log.Print("Default Message")
	log.VPrint("Verbose Message")
	log.Print("Success Message")
	log.Warn("Warn Message")
	log.Error("Error Message")
	log.Debug("Debug Message")
	log.Trace("Trace Message")

	// FLOG Messages (FORMATTED w/COLOR)
	fmt.Printf("\n\n=== FLOG Messages (FORMATTED w/COLOR) ===\n")
	flog.Print("Default Message")
	flog.VPrint("Verbose Message")
	flog.Success("Success Message")
	flog.Warn("Warning Message")
	flog.Failure("Failure Message")
	flog.Error("Error Message")
	flog.Debug("Debug Message")
	flog.Trace("Trace Message")

	// FLOG Messages (CUSTOM FORMATTED w/PREFIX + COLOR)
	fmt.Printf("\n\n=== FLOG CUSTOM Messages (FORMATTED w/PREFIX + COLOR) ===\n")
	flog.Custom("prefix", "blue", "Custom Blue Message")
	flog.Custom("prefix", "green", "Custom Green Message")
	flog.Custom("prefix", "red", "Custom Red Message")
	flog.Custom("prefix", "white", "Custom White Message")
	flog.Custom("prefix", "yellow", "Custom Yellow Message")
	flog.Custom("", "maroon", "Custom Default Message")
	return nil
}
