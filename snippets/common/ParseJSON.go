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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// JSONTemplate Note:
// struct fields must be public in order for unmarshal to correctly populate the data.
type JSONTemplate struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Links       []string `json:"list"`
}

// ParseJSON - Common Utility Function for Parsing JSON Templates
func ParseJSON(filename string) (jt JSONTemplate) {
	DebugMessage("[package: product][common] func parseJSON()")
	// Logging Output
	log.Printf("Parsing JSON Template %v\n", filename)

	// Open the json template file and handle any errors
	jsonFile, err := os.Open(filename)
	HandleError(err)

	// Defer the closing of the json file so we can parse it later
	defer jsonFile.Close()

	// Read Template File and Covert to Byte Slice
	byteValue, err := ioutil.ReadAll(jsonFile)
	HandleError(err)

	// Declare jsonTemplate variable as Type JsonTemplate (see struct)
	var jsonTemplate JSONTemplate

	// Unmarshal parses the JSON-encoded data and stores the result in the template pointed to by &jsonTemplate
	err = json.Unmarshal(byteValue, &jsonTemplate)
	HandleError(err)
	return jsonTemplate

}
