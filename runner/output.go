package runner

import (
	"os"
	"path/filepath"
	"strings"
)

func WriteMsvcDLLSource(content, dllName, outputPath string) error {
	err := os.WriteFile(filepath.Join(outputPath, dllName+".cpp"), []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}

func WriteMinGWDLLSource(content, dllName, outputPath string) error {
	defContent := strings.Split(content, strings.Repeat("=", 50))[0]
	source := strings.Split(content, strings.Repeat("=", 50))[1]
	err := os.WriteFile(filepath.Join(outputPath, dllName+".cpp"), []byte(source), 0666)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(outputPath, dllName+".def"), []byte(defContent), 0666)
	if err != nil {
		return err
	}
	return nil
}
