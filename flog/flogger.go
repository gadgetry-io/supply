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

// Package flog provides a consistent Log Formatting API for the SUPPLY CLI to customize print messages to the user
package flog

import (
	"fmt"
	"log"
	"os"
	"strings"

	auroraPackage "github.com/logrusorgru/aurora"
	colorable "github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	isatty "github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
)

// Aurora Variable
var aurora auroraPackage.Aurora

// Flag Controlling Verbose Logging
var verboseEnabled bool

// Flag Controlling Verbose Logging
var debugEnabled bool

// Flag Controlling Verbose Logging
var traceEnabled bool

// Init Function to Configure Colorable Output
func init() {
	aurora = auroraPackage.NewAurora(isatty.IsTerminal(os.Stdout.Fd()))
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFlags(0)
}

// Print logs a formatted info message
func Print(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightCyan("INFO %v").String(), message))
}

// VPrint logs a formatted info message when verboseEnabled is true
func VPrint(message string) {
	if verboseEnabled {
		log.Printf(fmt.Sprintf(aurora.BrightCyan("INFO %v").String(), message))
	}
}

// Success logs a formatted success message
func Success(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightGreen("SUCCESS %v").String(), message))
}

// Warn logs a formatted warning message
func Warn(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightYellow("WARNING %v").String(), message))
}

// Failure logs a formatted failure message
func Failure(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightRed("FAILURE %v").String(), message))
}

// Error logs a formatted error message
func Error(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightRed("ERROR %v").String(), message))
}

// Fatal logs a formatted fatal message
func Fatal(message string) {
	log.Fatalf(fmt.Sprintf(aurora.BrightRed("FATAL %v").String(), message))
}

// Panic logs a formatted fatal message
func Panic(message string) {
	log.Panicf(fmt.Sprintf(aurora.BrightRed("PANIC %v").String(), message))
}

// Debug logs a formatted debug message
func Debug(message string) {
	if debugEnabled {
		log.Printf("DEBUG %v", message)
	}
}

// Trace logs a formatted trace message
func Trace(message string) {
	if traceEnabled {
		log.Printf("TRACE %v", message)
	}
}

// Custom logs a custom formatted message specifying prefix and color
func Custom(prefix string, color string, message string) {

	if prefix == "" {
		prefix = "CUSTOM"
	}

	switch color {

	case "blue":
		log.Printf(fmt.Sprintf(aurora.BrightCyan("%v %v").String(), strings.ToUpper(prefix), message))
	case "green":
		log.Printf(fmt.Sprintf(aurora.BrightGreen("%v %v").String(), strings.ToUpper(prefix), message))
	case "red":
		log.Printf(fmt.Sprintf(aurora.BrightRed("%v %v").String(), strings.ToUpper(prefix), message))
	case "white":
		log.Printf(fmt.Sprintf("%v %v", strings.ToUpper(prefix), message))
	case "yellow":
		log.Printf(fmt.Sprintf(aurora.BrightYellow("%v %v").String(), strings.ToUpper(prefix), message))
	default:
		log.Printf(fmt.Sprintf("%v %v", strings.ToUpper(prefix), message))
	}
}

// EnableVerbose turns on verbose logging
func EnableVerbose() {
	verboseEnabled = true
}

// EnableDebug turns on debug logging
func EnableDebug() {
	debugEnabled = true
}

// EnableTrace turns on trace logging
func EnableTrace() {
	traceEnabled = true
}
