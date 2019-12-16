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
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ZipArchive - Common Utility Function for Zipping an Archive File (.zip)
func ZipArchive(target string, source string) {
	DebugMessage("[common] func ZipArchive()")
	log.Printf("Creating Zip Archive %v\n", target)

	if target == "" {
		log.Fatal("ERROR: A target is required")
	} else {
		log.Printf("The Target: %v\n", target)
	}

	if source == "" {
		log.Fatal("ERROR: A source is required")
	} else {
		log.Printf("The Source: %v\n", source)
	}

	err := createZipArchive(source, target)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Zip Archive %v was created successfully.\n\n", target)
	}
}

func createZipArchive(source string, target string) error {

	zipfile, err := os.Create(target)
	if err != nil {
		fmt.Printf("\n\n")
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {

			if baseDir == "." {
				header.Name = filepath.ToSlash(filepath.Join(baseDir, path))
			} else {
				prefix := baseDir + "/"
				header.Name = filepath.ToSlash(path)
				header.Name = filepath.ToSlash(filepath.Join(strings.TrimPrefix(header.Name, prefix)))

				// ROOT DIRECTORY CHECK
				// Check if baseDir matches header.Name
				if baseDir == header.Name {
					return nil
				}
			}
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err

}
