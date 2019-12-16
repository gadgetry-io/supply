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

package common

import (
	"fmt"

	"github.com/spf13/viper"
)

// SetUserConfig - Common Utility Function to Set Config Value in $HOME/.supply/config.yaml
func SetUserConfig(key string, value string) {
	DebugMessage("[package: common] func SetUserConfig()")
	err := viper.ReadInConfig()
	if err != nil {
		LogMessage("error", fmt.Sprintf("Config File Missing ($HOME/.supply/config.yaml)"))
		LogMessage("fatal", "Please create $HOME/.supply/config.yaml")
	} else {
		viper.Set(key, value)
		viper.WriteConfig()
	}
}
