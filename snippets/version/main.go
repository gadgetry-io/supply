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

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gadgetry-io/supply/version"
)

// Cmd represents the version command
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version number of supply cli.",
	Long:  `Get the version number of supply cli.`,
	Run:   getVersion,
}

func init() {

	// Version Sub Commands

	// Define Flags and configuration settings.

	// Persistent Flags for this command and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Local Flags which will only run for this command
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getVersion(_ *cobra.Command, _ []string) {
	fmt.Printf("supply cli version: %s\n\n", version.Version)
}
