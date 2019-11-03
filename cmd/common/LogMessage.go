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
	"os"

	auroraPackage "github.com/logrusorgru/aurora"
	colorable "github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	isatty "github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
)

// LogMessage - Common Utility Function for Formating Log Messages
func LogMessage(format string, message string) {
	var aurora auroraPackage.Aurora
	aurora = auroraPackage.NewAurora(isatty.IsTerminal(os.Stdout.Fd()))
	log.SetOutput(colorable.NewColorableStdout())

	switch format {
	case "info":
		log.Printf(fmt.Sprintf(aurora.BrightCyan("INFO: %v").String(), message))
	case "debug":
		log.Printf(fmt.Sprintf(aurora.BrightYellow("DEBUG: %v").String(), message))
	case "warn":
		log.Printf(fmt.Sprintf(aurora.BrightYellow("WARN: %v").String(), message))
	case "pass":
		log.Printf(fmt.Sprintf(aurora.BrightGreen("PASS: %v").String(), message))
	case "fail":
		log.Printf(fmt.Sprintf(aurora.BrightRed("FAIL: %v").String(), message))
	case "success":
		log.Printf(fmt.Sprintf(aurora.BrightGreen("SUCCESS: %v").String(), message))
	case "error":
		log.Printf(fmt.Sprintf(aurora.BrightRed("ERROR: %v").String(), message))
	case "fatal":
		log.Fatalf(fmt.Sprintf(aurora.BrightRed("ERROR: %v").String(), message))
	default:
		log.Printf("LOG: %v", message)
	}
}
