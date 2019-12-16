/*
Copyright © 2019 Gadgetry.io
Project: supply
Author: Brian Hooper (@KnownTraveler)
Email: brian.hooper@gadgetry.io
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at
    https://mozilla.org/MPL/2.0/.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/gadgetry-io/supply/flog"
	"github.com/gadgetry-io/supply/log"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "supply",
	Short: "supply - Be Well-Equipped",
	Long: `
	SUPPLY Command Line Interface (CLI)

	supply (verb)
	- to make available for use
	- to provide what is needed or wanted
	
	supply (noun)
	- the provisions and equipment necessary for people engaged in a particular project
	- a stock of a resource from which a person or place can be provided with the necessary amount of that resource.

  Handcrafted for Technology Practitioners
	
	
USAGE:
  supply <command> [flags]

  
Examples:

  supply --help
	supply init
	supply list
	supply pack
	supply ship
	supply tool`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Check Log Level Environment Variable
		rootOpts.logLevel = os.Getenv("SUPPLY_LOG_LEVEL")

		// Configure Verbose Logging
		if rootOpts.verbose {
			// Enable Verbose Logging
			log.EnableVerbose()
			flog.EnableVerbose()
		}

		// Configure Debug Logging
		if rootOpts.logLevel == "debug" {
			// Enable Verbose Logging
			log.EnableVerbose()
			flog.EnableVerbose()

			// Enable Debug Logging
			log.EnableDebug()
			flog.EnableDebug()
		}

		// Configure Trace Logging
		if rootOpts.logLevel == "trace" {
			// Enable Verbose Logging
			log.EnableVerbose()
			flog.EnableVerbose()

			// Enable Debug Logging
			log.EnableDebug()
			flog.EnableDebug()

			// Enable Trace Logging
			flog.EnableTrace()
			log.EnableTrace()
		}

		log.Print("PersistentPreRun:")
		log.Printf("verboseEnabled = %v\n", rootOpts.verbose)
		log.Printf("logLevel       = %v\n", rootOpts.logLevel)
		log.Printf("nonInteractive = %v\n", rootOpts.nonInteractive)
	},
}

type rootFlags struct {
	verbose        bool
	logLevel       string
	nonInteractive bool
}

var rootOpts rootFlags

func init() {
	// Initialize Config Here
	// cobra.OnInitialize(initConfig)

	// Add Command Packages Here
	rootCmd.AddCommand(CmdExample)
	// rootCmd.AddCommand(CmdVersion)
	// rootCmd.AddCommand(CmdInit)
	// rootCmd.AddCommand(CmdPack)
	// rootCmd.AddCommand(CmdShip)
	// rootCmd.AddCommand(CmdTool)

	// Persistent flags which are global for the application.
	rootCmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "Enable Verbose Output")
	rootCmd.PersistentFlags().BoolVar(&rootOpts.nonInteractive, "non-interactive", false, "Disable Interactive Prompts")

	// Local Flags which only run when this action is called.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name "$HOME/.supply/config" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".supply/config")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
