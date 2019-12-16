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

package project

import (
	"fmt"

	"github.com/gadgetry-io/supply/cmd/common"
	"github.com/spf13/cobra"
)

// Cmd represents the configure command
var Cmd = &cobra.Command{
	Use:   "project",
	Short: "supply project commands",
	Long: `The supply "project" commands manage supply projects

USAGE:
  supply project <command> [flags]
	
Examples:

	supply project list-packages
	supply project describe-package --name
	supply project get-package --name <package_name> --
	supply project create-package --name
	supply project tar-package --name
	supply project zip-package --name`,

	Run: func(cmd *cobra.Command, args []string) {
		common.DebugMessage("[package: project] func Run()")
		fmt.Printf("Use 'supply project --help' for more information\n\n")

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
