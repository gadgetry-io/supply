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

import "errors"

// RemoveStringFromSlice - Common Utility Function for Removing A String from a Slice
func RemoveStringFromSlice(slice []string, element string) ([]string, error) {

	elementIndex := indexOfStringSlice(slice, element)

	if elementIndex >= len(slice) {
		return nil, errors.New("ERROR: Service Not Found. Index Out of Range Error Searching in Services Slice")
	}
	return append(slice[:elementIndex], slice[elementIndex+1:]...), nil
}

func indexOfStringSlice(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
