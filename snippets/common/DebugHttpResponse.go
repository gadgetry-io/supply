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
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/spf13/viper"
)

// DebugHttpResponse - Common Utility Function for Debugging API Responses
func DebugHttpResponse(response *http.Response, debugMsg string) {
	DebugMessage("[common] func DebugHttpResponse()")
	debugMap := viper.GetStringMapString("debug")
	debug := debugMap["trace"]
	if debug == "debug" {
		log.Printf("INFO: [common] DebugHttpResponse() for --> %v\n", debugMsg)
		responseDump, err := httputil.DumpResponse(response, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\n\n")
		fmt.Println(string(responseDump))
		fmt.Printf("\n\n")
	}

}
