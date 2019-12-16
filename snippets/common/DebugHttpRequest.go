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
	"net/http"
	"net/http/httputil"

	"github.com/spf13/viper"
)

// DebugHttpRequest - Common Utility Function for Debugging API Requests
func DebugHttpRequest(request *http.Request, debugMsg string) {
	DebugMessage("[common] func DebugHttpRequest()")
	debugMap := viper.GetStringMapString("debug")
	debug := debugMap["trace"]
	if debug == "true" {
		LogMessage("info", fmt.Sprintf("[common] DebugHttpRequest() for --> %v\n", debugMsg))
		requestDump, err := httputil.DumpRequest(request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\n\n")
		fmt.Println(string(requestDump))
		fmt.Printf("\n\n")
	}

}
