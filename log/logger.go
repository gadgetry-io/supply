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

// Package log provides a consistent Logging API for the SUPPLY CLI to print messages to the user
// Logrus has seven logging levels: Trace, Debug, Info, Warning, Error, Fatal and Panic.
package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Flag Controlling Verbose Logging
var verboseEnabled bool

// Flag Controlling Debug Logging
var debugEnabled bool

// Flag Controlling Trace Logging
var traceEnabled bool

// Init Function to configure log format
func init() {
	log.Formatter = new(logrus.TextFormatter)
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	log.Formatter.(*logrus.TextFormatter).DisableLevelTruncation = true
	log.Level = logrus.InfoLevel
	log.Out = os.Stdout
}

// Default Logging Level - Standard Logging Output for User

// Print logs a message at level Info.
func Print(args ...interface{}) {
	log.Level = logrus.InfoLevel
	log.Info(args...)
}

// Printf logs a message at level Info using the format string
func Printf(format string, args ...interface{}) {
	log.Level = logrus.InfoLevel
	log.Infof(format, args...)
}

// Verbose Logging Level - Additional Logging Output for User

// VPrint logs a message at level Info (if verbose logging is enabled)
func VPrint(args ...interface{}) {
	if verboseEnabled {
		log.Level = logrus.InfoLevel
		log.Info(args...)
	}
}

// VPrintf logs a message at level Info using the format string (if verbose logging is enabled)
func VPrintf(format string, args ...interface{}) {
	if verboseEnabled {
		log.Level = logrus.InfoLevel
		log.Infof(format, args...)
	}
}

// Warn logs a message at level Warn
func Warn(args ...interface{}) {
	log.Level = logrus.WarnLevel
	log.Warn(args...)
}

// Warnf logs a message at level Warn using the format string
func Warnf(format string, args ...interface{}) {
	log.Level = logrus.WarnLevel
	log.Warnf(format, args...)
}

// Error logs a message at level Error
func Error(args ...interface{}) {
	log.Level = logrus.ErrorLevel
	log.Error(args...)
}

// Errorf logs a message at level Error using the format string
func Errorf(format string, args ...interface{}) {
	log.Level = logrus.ErrorLevel
	log.Errorf(format, args...)
}

// Panic logs a message at level Panic
func Panic(args ...interface{}) {
	log.Level = logrus.PanicLevel
	log.Panic(args...)
}

// Panicf logs a message at level Panic using the format string
func Panicf(format string, args ...interface{}) {
	log.Level = logrus.PanicLevel
	log.Panicf(format, args...)
}

// Fatal logs a message at level Fatal
func Fatal(args ...interface{}) {
	log.Level = logrus.FatalLevel
	log.Fatal(args...)
}

// Fatalf logs a message at level Fatal using the format string
func Fatalf(format string, args ...interface{}) {
	log.Level = logrus.FatalLevel
	log.Fatalf(format, args...)
}

// Debug logs a message at level Debug
func Debug(args ...interface{}) {
	if debugEnabled {
		log.Level = logrus.DebugLevel
		log.Debug(args...)
	}
}

// Debugf logs a message at level Debug using the format string
func Debugf(format string, args ...interface{}) {
	if debugEnabled {
		log.Level = logrus.DebugLevel
		log.Debugf(format, args...)
	}
}

// Trace logs a message at level Trace
func Trace(args ...interface{}) {
	if traceEnabled {
		log.Level = logrus.TraceLevel
		log.Trace(args...)
	}
}

// Tracef logs a message at level Trace using the format string
func Tracef(format string, args ...interface{}) {
	if traceEnabled {
		log.Level = logrus.TraceLevel
		log.Tracef(format, args...)
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
