package core

import (
	"AGHD/models"
	"AGHD/utils"
	"github.com/saferwall/pe"
	"strings"
)

var known_dlls_list = []string{"kernel32.dll", "ucrtbase.dll", "MSCTF.dll", "SHLWAPI.dll", "WS2_32.dll", "KERNELBASE.dll", "wow64.dll", "msvcp_win.dll", "gdiplus.dll", "user32.dll", "bcrypt.dll", "COMCTL32.dll", "cfgmgr32.dll", "combase.dll", "IMM32.dll", "rpcrt4.dll", "ntdll.dll", "bcryptPrimitives.dll", "coml2.dll", "win32u.dll", "wow64cpu.dll", "COMDLG32.dll", "gdi32full.dll", "IMAGEHLP.dll", "SHELL32.dll", "sechost.dll", "WINTRUST.dll", "NORMALIZ.dll", "difxapi.dll", "Setupapi.dll", "CRYPT32.dll", "gdi32.dll", "MSVCRT.dll", "wow64win.dll", "advapi32.dll", "PSAPI.DLL", "NSI.dll", "WLDAP32.dll", "OLEAUT32.dll", "SHCORE.dll", "ole32.dll", "clbcatq.dll"}


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
