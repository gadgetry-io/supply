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
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TarArchive - Common Utility for Creating tar/gzip Compressed Artifacts
func TarArchive(target string, source string) {

	DebugMessage("[common] func TarArchive()")
	log.Printf("Creating Tar Archive %v\n", target)

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

	err := createTarArchive(target, source)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Tar Archive %v was created successfully.\n\n", target)
	}

}

func createTarArchive(target string, source string) error {

	// CREATE OUTPUT FILE
	tarfile, err := os.Create(target)
	if err != nil {
		fmt.Printf("\n\n")
		return err
	}
	defer tarfile.Close()

	// SETUP gzipWriter
	gzipWriter := gzip.NewWriter(tarfile)
	defer gzipWriter.Close()

	// SETUP tarWriter
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	file, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if file.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(file, file.Name())
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

		if file.IsDir() {
			header.Name += "/"
		} else {
			header.Size = file.Size()
		}

		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		if file.IsDir() {
			return nil
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}

		defer sourceFile.Close()
		_, err = io.Copy(tarWriter, sourceFile)
		return err
	})

	return err

}
