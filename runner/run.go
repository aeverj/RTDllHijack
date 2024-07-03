package runner

import (
	"AGHD/core"
	"AGHD/models"
	"AGHD/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type RunnerOpt struct {
	Input   string
	Output  string
	Exclude string
	Arch    string
	Compile string
	Verbose bool
}

func New(p *models.Parameters) (*RunnerOpt, error) {
	opt := &RunnerOpt{}
	_, err := os.Stat(p.InputPath)
	if err != nil {
		return nil, fmt.Errorf("cannot access input path: %v", err)
	}
	opt.Input = p.InputPath
	if p.OutputPath == "" {
		currentTime := time.Now()
		formattedTime := currentTime.Format("20060102_150405")
		folderName := formattedTime + "_out"
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			utils.Mlogger.Error("Error creating directory: %s\n", err.Error())
			return nil, err
		}
		opt.Output = folderName
	} else {
		opt.Output = p.OutputPath
	}
	if p.Verbose {
		utils.Mlogger = utils.NewLogger(utils.LogLevelInfo)
	} else {
		utils.Mlogger = utils.NewLogger(utils.LogLevelSuccess)
	}
	opt.Compile = p.Compiler
	opt.Exclude = p.ExcludePattern
	opt.Verbose = true
	return opt, nil
}

func (r *RunnerOpt) Run() {
	pePathChan := make(chan string)
	go ProcessInput(r.Input, r.Exclude, pePathChan)
	for filePath := range pePathChan {
		contentList, signValid := core.StaticParsePE(filePath)
		resultPath := r.Output
		if len(contentList) > 0 {
			if signValid {
				resultPath = filepath.Join(resultPath, "Signed", strings.Replace(strings.Replace(filePath, string(os.PathSeparator), "_", -1), ":", "_", -1))
			} else {
				resultPath = filepath.Join(resultPath, "NoSigned", strings.Replace(strings.Replace(filePath, string(os.PathSeparator), "_", -1), ":", "_", -1))
			}
			_, err := utils.CopyFile(filepath.Join(resultPath, filepath.Base(filePath)), filePath)
			if err != nil {
				continue
			}
		}
		for _, content := range contentList {
			switch r.Compile {
			case "msvc":
				sourceMsvc := core.GeneCppCodeForMsvc(content)
				err := WriteMsvcDLLSource(sourceMsvc, content.DllName, resultPath)
				if err != nil {
					utils.Mlogger.Error("Write cpp file failed, %s", err.Error())
					break
				}
				utils.Mlogger.Success("Success general %s", content.DllName)
				break
			case "mingw":
				sourceMinGW := core.GeneCppCodeForMinGW(content)
				err := WriteMinGWDLLSource(sourceMinGW, content.DllName, resultPath)
				if err != nil {
					utils.Mlogger.Error("Write cpp file failed, %s", err.Error())
					break
				}
				utils.Mlogger.Success("Success general %s", content.DllName)
				break
			default:
				break
			}
		}
	}
	entries, _ := os.ReadDir(r.Output)
	if len(entries) == 0 {
		utils.Mlogger.Info("No matching files found in the %s folder.", r.Input)
		_ = os.RemoveAll(r.Output)
	} else {
		utils.Mlogger.Success("Operation completed successfully! Results saved to the %s folder.", r.Output)
	}
}
