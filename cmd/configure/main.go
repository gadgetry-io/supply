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

package configure

import (
	"fmt"

	"github.com/gadgetry-io/supply/cmd/common"
	"github.com/spf13/cobra"
)

// Cmd represents the configure command
var Cmd = &cobra.Command{
	Use:   "configure",
	Short: "supply configure commands",
	Long: `The supply "configure" commands manage the $HOME/.supply/config.yaml

USAGE:
  supply configure <command> [flags]
	
Examples:

  PROFILE CONFIGURATION
  supply configure user --get-email
  supply configure user --set-email "<email>"
  
  supply configure user --get-name
  supply configure user --set-name "<name>"`,

	Run: func(cmd *cobra.Command, args []string) {
		common.DebugMessage("[package: configure] func Run()")
		fmt.Printf("Use 'supply configure --help' for more information\n\n")

	},

	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return errors.New("The 'supply' command expects a few parameters, please review usage")
	// },
}

func init() {

	// Configure Sub Commands
	//Cmd.AddCommand(userCmd)

	// Define Flags and configuration settings.

	// Persistent Flags for this command and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Local Flags which will only run for this command
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
