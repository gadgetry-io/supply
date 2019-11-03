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

import "strings"

// GetFileExt - Common Utility Function for Getting a File Extension
func GetFileExt(filename string) string {
	DebugMessage("[common] func GetFileExt()")
	//log.Printf("Checking Extension for Template: %v\n", filename)

	// Split Filename into a Slice
	fSlice := strings.Split(filename, ".")

	// Get the Length of the Slice
	fSliceLen := len(fSlice)

	// Get the extension of the Filename from last index position in the Slice (length -1)
	fileExtension := fSlice[fSliceLen-1]

	// Return the Extension
	return fileExtension
}
