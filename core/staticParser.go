package core

import (
	"AGHD/models"
	"AGHD/utils"
	"github.com/saferwall/pe"
	"strings"
)

var known_dlls_list = []string{"kernel32.dll", "wow64.dll", "shell32.dll", "imagehlp.dll", "advapi32.dll", "wow64win.dll", "psapi.dll", "msctf.dll", "imm32.dll", "gdiplus.dll", "shcore.dll", "wowarmhw.dll", "setupapi.dll", "msvcrt.dll", "gdi32.dll", "xtajit.dll", "wldap32.dll", "combase.dll", "nsi.dll", "normaliz.dll", "oleaut32.dll", "user32.dll", "clbcatq.dll", "ole32.dll", "wow64cpu.dll", "sechost.dll", "difxapi.dll", "comdlg32.dll", "coml2.dll", "shlwapi.dll", "ws2_32.dll", "rpcrt4.dll", "ntdll.dll','mscoree.dll", "msvcp_win"}

func StaticParsePE(filePath string) (dllInfoList []models.DLLInfo, SignatureValid bool) {
	file, err := pe.New(filePath, &pe.Options{})
	if err != nil {
		utils.Mlogger.Error("Error opening PE file: %v", err)
		return nil, false
	}

	err = file.Parse()
	if err != nil {
		utils.Mlogger.Error("Error parsing PE file: %v", err)
		return nil, false
	}
	SignatureValid = false
	if file.IsSigned {
		SignatureValid = file.Certificates.Certificates[0].SignatureValid
	}
	if file.HasImport {
		var currDllInfo models.DLLInfo
		currDllInfo.Arch = file.PrettyOptionalHeaderMagic()
		for _, imp := range file.Imports {
			if utils.Contains(known_dlls_list, imp.Name) || strings.HasPrefix(imp.Name, "api-ms-win") {
				continue
			}
			currDllInfo.DllName = imp.Name
			currDllInfo.FunName = nil
			currDllInfo.FunOrdinal = nil
			for _, funcName := range imp.Functions {
				if !funcName.ByOrdinal {
					currDllInfo.FunName = append(currDllInfo.FunName, funcName.Name)
				} else {
					currDllInfo.FunOrdinal = append(currDllInfo.FunOrdinal, funcName.Ordinal)
				}
			}
			if len(currDllInfo.FunOrdinal) > 0 || len(currDllInfo.FunName) > 0 {
				dllInfoList = append(dllInfoList, currDllInfo)
			}
		}
	} else {
		utils.Mlogger.Error("No imports found.")
		return nil, false
	}
	return
}
