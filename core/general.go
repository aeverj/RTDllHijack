package core

import (
	"AGHD/models"
	"fmt"
	"strings"
)

func GeneCppCodeForMsvc(dllInfo models.DLLInfo) string {
	var builder strings.Builder

	// 添加头文件和注释
	builder.WriteString(fmt.Sprintf(`#include <windows.h>

// %s implementation %s

`, dllInfo.DllName, dllInfo.Arch))

	// 添加函数导出
	for _, funcName := range dllInfo.FunName {
		builder.WriteString(fmt.Sprintf(`#pragma comment(linker, "/export:%s=run")`, funcName))
		builder.WriteString("\n")

	}
	for _, funcOrdinal := range dllInfo.FunOrdinal {
		builder.WriteString(fmt.Sprintf(`#pragma comment(linker, "/export:%d=run,@%d,NONAME")`, funcOrdinal, funcOrdinal))
		builder.WriteString("\n")

	}

	// 添加函数实现和 DllMain 函数
	builder.WriteString(`
extern "C" void run() {
    // Function implementation
}

BOOL WINAPI DllMain( HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpvReserved ) {
    switch (fdwReason) {
        case DLL_PROCESS_ATTACH:
        case DLL_THREAD_ATTACH:
        case DLL_THREAD_DETACH:
        case DLL_PROCESS_DETACH:
            break;
    }
    return TRUE; // Successful DLL_PROCESS_ATTACH.
}
`)

	return builder.String()
}

func GeneCppCodeForMinGW(dllInfo models.DLLInfo) string {
	var builder strings.Builder
	builder.WriteString(`LIBRARY def
EXPORTS
`)
	// 添加函数导出
	for _, funcName := range dllInfo.FunName {
		builder.WriteString(fmt.Sprintf(`    %s=run`, funcName))
		builder.WriteString("\n")

	}
	for _, funcOrdinal := range dllInfo.FunOrdinal {
		builder.WriteString(fmt.Sprintf(`    run%d=run @%d NONAME PRIVATE`, funcOrdinal, funcOrdinal))
		builder.WriteString("\n")

	}

	builder.WriteString(strings.Repeat("=", 50))
	builder.WriteString("\n")
	// 添加头文件和注释
	builder.WriteString(fmt.Sprintf(`#include <windows.h>

// %s implementation %s
// g++ -shared -o %s -O3 -static -s -Wl,--strip-all %s.cpp %s.def

`, dllInfo.DllName, dllInfo.Arch, dllInfo.DllName, dllInfo.DllName, dllInfo.DllName))
	builder.WriteString(`
extern "C" void run() {
    // Function implementation
}


BOOL WINAPI DllMain( HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpvReserved ) {
    switch (fdwReason) {
        case DLL_PROCESS_ATTACH:
        case DLL_THREAD_ATTACH:
        case DLL_THREAD_DETACH:
        case DLL_PROCESS_DETACH:
            break;
    }
    return TRUE; // Successful DLL_PROCESS_ATTACH.
}`)

	return builder.String()
}
