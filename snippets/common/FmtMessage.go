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
	"os"
	"strings"

	auroraPackage "github.com/logrusorgru/aurora"
	isatty "github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
)

// FmtMessage - Common Utility Function for Formatting Stdout Messages
func FmtMessage(format string, message string) {
	var aurora auroraPackage.Aurora
	aurora = auroraPackage.NewAurora(isatty.IsTerminal(os.Stdout.Fd()))
	switch format {

	case "h1-white":
		fmt.Printf("\n\n========================================\n%v\n________________________________________\n\n", strings.ToUpper(message))
	case "h1-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("\n\n========================================\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h1-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("\n\n========================================\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h1-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("\n\n========================================\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h1-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("\n\n========================================\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h2-white":
		fmt.Printf("\n%v\n________________________________________\n\n", strings.ToUpper(message))
	case "h2-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h2-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h2-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h2-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("\n%v\n________________________________________\n\n").String(), strings.ToUpper(message)))
	case "h3-white":
		fmt.Printf("%v\n\n", strings.ToUpper(message))
	case "h3-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("%v\n\n").String(), strings.ToUpper(message)))
	case "h3-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("%v\n\n").String(), strings.ToUpper(message)))
	case "h3-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("%v\n\n").String(), strings.ToUpper(message)))
	case "h3-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("%v\n\n").String(), strings.ToUpper(message)))
	case "text-white":
		fmt.Printf("%v\n", message)
	case "text-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("%v\n").String(), message))
	case "text-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("%v\n").String(), message))
	case "text-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("%v\n").String(), message))
	case "text-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("%v\n").String(), message))
	case "upper-white":
		fmt.Printf("%v\n", strings.ToUpper(message))
	case "upper-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("%v\n").String(), strings.ToUpper(message)))
	case "upper-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("%v\n").String(), strings.ToUpper(message)))
	case "upper-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("%v\n").String(), strings.ToUpper(message)))
	case "upper-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("%v\n").String(), strings.ToUpper(message)))
	case "lower-white":
		fmt.Printf("%v\n", strings.ToLower(message))
	case "lower-blue":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("%v\n").String(), strings.ToLower(message)))
	case "lower-yellow":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("%v\n").String(), strings.ToLower(message)))
	case "lower-green":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("%v\n").String(), strings.ToLower(message)))
	case "lower-red":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("%v\n").String(), strings.ToLower(message)))
	case "info":
		fmt.Printf(fmt.Sprintf(aurora.BrightCyan("INFO: %v\n").String(), message))
	case "debug":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("DEBUG: %v\n").String(), message))
	case "warn":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("WARN: %v\n").String(), message))
	case "pass":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("PASS: %v\n").String(), message))
	case "fail":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("FAIL: %v\n").String(), message))
	case "success":
		fmt.Printf(fmt.Sprintf(aurora.BrightGreen("SUCCESS: %v\n").String(), message))
	case "warning":
		fmt.Printf(fmt.Sprintf(aurora.BrightYellow("WARNING: %v\n").String(), message))
	case "failure":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("FAILURE: %v\n").String(), message))
	case "error":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("ERROR: %v\n").String(), message))
	case "fatal":
		fmt.Printf(fmt.Sprintf(aurora.BrightRed("ERROR: %v\n").String(), message))
	default:
		fmt.Printf("DEFAULT: %v\n", message)
	}
}
