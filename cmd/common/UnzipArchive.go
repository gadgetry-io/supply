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
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// UnzipArchive - Common Utility Function for Unzipping an Archive File (.zip)
func UnzipArchive(filename string, target string) {

	// Create a reader out of the zip archive
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		var targetDir string

		if target == "" {
			targetDir = "./"
		} else {
			targetDir = target
		}

		extractedFilePath := filepath.Join(targetDir, file.Name)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {

			// Check if Directory Path Exists
			if _, err := os.Stat(filepath.Dir(extractedFilePath)); os.IsNotExist(err) {
				// Directory Path Does Not Exist
				log.Println("Directory Path Does Not Exist:", filepath.Dir(extractedFilePath))
				log.Println("Creating Path:", filepath.Dir(extractedFilePath))
				// Create Directory Path
				os.MkdirAll(filepath.Dir(extractedFilePath), 0755)
			}

			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())

		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Check if File Path Exists
			log.Println("Checking path:", extractedFilePath)
			if _, err := os.Stat(filepath.Dir(extractedFilePath)); os.IsNotExist(err) {
				// File Directory Path Does Not Exist
				log.Println("Path Does Not Exist:", filepath.Dir(extractedFilePath))
				log.Println("Creating Path:", filepath.Dir(extractedFilePath))
				// Create Directory Path
				os.MkdirAll(filepath.Dir(extractedFilePath), 0755)
			}

			// Create an output file for writing
			log.Println("Creating Output file:", extractedFilePath)
			f, err := os.Create(extractedFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(f, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
