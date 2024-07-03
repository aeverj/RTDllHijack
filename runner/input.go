package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProcessInput(inputPath, excludePattern string, peFilePath chan string) {
	info, _ := os.Stat(inputPath)
	if info.IsDir() {
		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				_ = fmt.Errorf("unable to walk the input directory: %v", err)
				return nil
			}
			if strings.HasSuffix(info.Name(), ".exe") {
				if excludePattern != "" && strings.Contains(strings.ToLower(path), strings.ToLower(excludePattern)) {
					return nil
				}
				peFilePath <- path
			}
			return nil
		})
		if err != nil {
			_ = fmt.Errorf("unable to walk the input directory: %v", err)
		}
	} else {
		if strings.HasSuffix(info.Name(), ".exe") {
			peFilePath <- inputPath
		}
	}
	close(peFilePath)
}
