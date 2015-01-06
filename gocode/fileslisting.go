package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type InspectFileDetails struct {
	Name         string
	Path         string
	Size         int64
	LastModified time.Time
	Mode         string
}

func getHighLevelFileDetails(cfg *config) ([]InspectFileDetails, error) {
	var ifDetails []InspectFileDetails
	for _, onefile := range cfg.Files {
		fmt.Printf("filelistings.go:getHighLevelFileDetails(): Walking directory %s for files with pattern %s and isRecursive: %b.\n", onefile.Dir, onefile.Pattern, onefile.Recursive)

		r, err := regexp.Compile(onefile.Pattern)
		if err != nil {
			fmt.Printf("filelistings.go:getHighLevelFileDetails(): Regex Error: Given pattern: %v, Error: %v\n", onefile.Pattern, err)
			return nil, err
		}

		err = filepath.Walk(onefile.Dir, func(fullpath string, info os.FileInfo, walkerr error) error {
			if info.IsDir() && fullpath != onefile.Dir && !onefile.Recursive {
				fmt.Printf("filelistings.go:getHighLevelFileDetails():walkanon(): Skipping dir: %v.\n", info.Name())
				return filepath.SkipDir
			}

			if !r.MatchString(info.Name()) || info.IsDir() {

				fmt.Printf("filelistings.go:getHighLevelFileDetails():walkanon(): Rejecting: %v.\n", info.Name())
				return nil
			}

			fmt.Printf("filelistings.go:getHighLevelFileDetails():walkanon(): Accepting file: %v.\n", info.Name())
			ifDetails = append(ifDetails, InspectFileDetails{
				info.Name(),
				fullpath,
				info.Size(),
				info.ModTime(),
				info.Mode().String(),
			})
			fmt.Printf("filelistings.go:getHighLevelFileDetails():walkanon(): Add file details: %v.\n", ifDetails[len(ifDetails)-1])

			return nil
		})

		if err != nil {
			fmt.Printf("filelistings.go:getHighLevelFileDetails(): Error: %v\n", err)
			return nil, err
		}

	}

	return ifDetails, nil
}
