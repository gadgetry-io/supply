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
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// YAMLTemplate Note:
// struct fields must be public in order for unmarshal to correctly populate the data.
type YAMLTemplate struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Links       []string `yaml:"list"`
}

// ParseYAML - Common Utility Function for Parsing YAML Templates
func ParseYAML(filename string) (yt YAMLTemplate) {
	DebugMessage("[package: product][common] func ParseYAML()")
	// Logging Output
	log.Printf("Parsing YAML Template %v\n", filename)

	yamlFile, err := os.Open(filename)
	HandleError(err)

	// Defer the closing of the yaml file so we can parse it later
	defer yamlFile.Close()

	// Read Template File and Covert to Byte Slice Value
	byteValue, err := ioutil.ReadAll(yamlFile)
	HandleError(err)

	// Declare t variable as Type YamlTemplate (see struct)
	var yamlTemplate YAMLTemplate

	// Unmarshal parses the YAML-encoded data and stores the result in the template pointed to by &template
	err = yaml.Unmarshal(byteValue, &yamlTemplate)
	HandleError(err)
	return yamlTemplate

}
