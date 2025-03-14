//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Command that checks each file for a license
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "no folder specified\n")
		os.Exit(1)
	}
	userExceptions := os.Args[2:]
	files := []string{}
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			// skip a few folders by default
			var excludedFolders = []string{
				".git",
				".github",
			}
			for _, excludedFolder := range excludedFolders {
				if path == excludedFolder {
					return filepath.SkipDir
				}
			}
			return nil
		}

		for _, userException := range userExceptions {
			if strings.HasSuffix(path, userException) {
				return nil
			}
		}

		base := filepath.Base(path)
		// skip readme files
		if strings.HasPrefix(strings.ToLower(base), "readme") {
			return nil
		}

		// skip a few files by default
		var excludedFiles = []string{
			// license file
			"LICENSE",
			// files to ignore in git
			".gitignore",
			// mac os cache file
			".DS_Store",
			// go module files
			"go.mod",
			"go.sum",
		}
		for _, excludeFile := range excludedFiles {
			if base == excludeFile {
				return nil
			}
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		datas := string(data)

		// skip over files that are generated code
		if strings.HasPrefix(datas, "// Code generated by") ||
			strings.HasPrefix(datas, "// generated by") {
			return nil
		}

		datas = strings.Replace(datas, "//", "", 1)
		datas = strings.Replace(datas, "#", "", 1)
		datas = strings.TrimSpace(datas)
		if !strings.HasPrefix(datas, "Copyright") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if len(files) == 0 {
		return
	}
	fmt.Fprintf(os.Stderr, "ERROR The following files still needs a LICENSE: [%s]\n", strings.Join(files, ", "))
	os.Exit(1)
}
