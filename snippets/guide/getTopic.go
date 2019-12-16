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

package guide

import (
	"github.com/gadgetry-io/supply/cmd/common"
	"github.com/gadgetry-io/supply/cmd/guide/topics"
	"github.com/spf13/cobra"
)

// getTopicCmd represents the supply guide get-topic command
var getTopicCmd = &cobra.Command{
	Use:   "get-topic",
	Short: "supply guide get-topic",
	Long: `supply guide get-topic
	
    USAGE:
	  supply guide get-topic <topic_name>

	EXAMPLE:
	  supply guide get-topic --name supply-message-formatting

	`,

	Run: func(cmd *cobra.Command, args []string) {
		common.DebugMessage("[package: guide][file: getTopic.go][Cmd Run: func()]")
		name, err := cmd.Flags().GetString("name")
		common.HandleError(err)

		switch name {
		case "supply-message-formatting":
			topics.SupplyMessageFormatting()
		default:
			topics.SupplyGuide()
		}

	},
}

func init() {

	// Define Flags and configuration settings.

	// Persistent Flags for this command and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Local Flags which will only run for this command
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getTopicCmd.Flags().String("name", "", "Supply Guide Topic Name")

}
