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

package cmd

import (
	"fmt"
	"os"

	"github.com/gadgetry-io/supply/cmd/configure"
	"github.com/gadgetry-io/supply/cmd/guide"
	"github.com/gadgetry-io/supply/cmd/project"
	"github.com/gadgetry-io/supply/cmd/version"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "supply",
	Short: "supply - Be Well-Equipped",
	Long: `
	SUPPLY Command Line Interface

	supply (verb)
	- to make available for use
	- to provide what is needed or wanted
	
	supply (noun)
	- the provisions and equipment necessary for people engaged in a particular project
	- a stock of a resource from which a person or place can be provided with the necessary amount of that resource.
	
	
	
	 
		

  Handcrafted for Technology Practitioners
	
	
USAGE:
  supply <gadget> <command> [flags]

  
Examples:

  supply --help
  supply <gadget> --help
  supply <gadget> <command> --help`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add Command Packages Here
	rootCmd.AddCommand(configure.Cmd)
	rootCmd.AddCommand(guide.Cmd)
	rootCmd.AddCommand(project.Cmd)
	rootCmd.AddCommand(version.Cmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.supply/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "$HOME/.supply/config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".supply/config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
