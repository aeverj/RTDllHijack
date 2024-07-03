# PE File Parser

PE File Parser is a tool that parses PE file import tables and generates hijackable DLL source code.

## Features

- Automatically discover hijackable DLLs in a given directory
- Generate corresponding source code for discovered DLLs
- Support for choosing compiler (MinGW or MSVC)
- Exclude specific files or directories
- Provide verbose output for debugging

## Installation

```sh
git clone https://github.com/aeverj/AGHD.git
cd AGHD
go mod tidy
go build -o AGHD.exe cmd/cmd.go
```
## Usage
```sh
.\AGHD.exe -h
NAME:
   PE File Parser - Parses PE file import tables and generates hijackable DLL source code

USAGE:
   PE File Parser [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --compiler value, -c value  Compiler to use (mingw or msvc) (default: "msvc")
   --input value, -i value     Input file or directory path
   --output value, -o value    Output directory path
   --exclude value, -e value   Exclude file or directory name pattern
   --verbose, -v               Enable verbose output (default: false)
   --help, -h                  show help

```
## Parsing all C drive files
```sh
AGHD.exe -i C:\
```
## Generate source files that support the MingW compiler
```cgo
AGHD.exe -i C:\ -c mingw
```
## Exclude specific files or directories
```cgo
AGHD.exe -i C:\ -e admin
```
## Result
```sh
├─NoSigned
│  ├─C__Program Files_Common Files_microsoft shared_ink_InputPersonalization.exe
│  │      elscore.dll.cpp
│  │      elscore.dll.def
│  │      InputPersonalization.exe
│  │      ntdll.dll.cpp
│  │      ntdll.dll.def
│  │      XmlLite.dll.cpp
│  │      XmlLite.dll.def
│  │
│  ├─C__Program Files_Common Files_microsoft shared_ink_mip.exe
│  │      COMCTL32.dll.cpp
│  │      COMCTL32.dll.def
│  │      dwmapi.dll.cpp
│  │      dwmapi.dll.def
│  │      mip.exe
│  │      MSIMG32.dll.cpp
│  │      MSIMG32.dll.def
│  │      ntdll.dll.cpp
│  │      ntdll.dll.def
│  │      OLEACC.dll.cpp
│  │      OLEACC.dll.def
│  │      UxTheme.dll.cpp
│  │      UxTheme.dll.def
│  │      VERSION.dll.cpp
│  │      VERSION.dll.def
│  │
│  ├─C__Program Files_Common Files_microsoft shared_ink_ShapeCollector.exe
│  │      COMCTL32.dll.cpp
│  │      COMCTL32.dll.def
│  │      DUI70.dll.cpp
│  │      DUI70.dll.def
│  │      ntdll.dll.cpp
│  │      ntdll.dll.def
│  │      ShapeCollector.exe
│  │
│  └─C__Program Files_Common Files_microsoft shared_MSInfo_msinfo32.exe
│          ATL.DLL.cpp
│          ATL.DLL.def
│          COMCTL32.dll.cpp
│          COMCTL32.dll.def
│          MFC42u.dll.cpp
│          MFC42u.dll.def
│          msinfo32.exe
│          ntdll.dll.cpp
│          ntdll.dll.def
│          POWRPROF.dll.cpp
│          POWRPROF.dll.def
│          SLC.dll.cpp
│          SLC.dll.def
│
└─Signed
    ├─C__Program Files_Common Files_microsoft shared_ClickToRun_appvcleaner.exe
    │      appvcleaner.exe
    │      APPVMANIFEST.dll.cpp
    │      APPVMANIFEST.dll.def
    │      APPVPOLICY.dll.cpp
    │      APPVPOLICY.dll.def
    │      msi.dll.cpp
    │      msi.dll.def
    │      ntdll.dll.cpp
    │      ntdll.dll.def
    │      USERENV.dll.cpp
    │      USERENV.dll.def
    │
    ├─C__Program Files_Common Files_microsoft shared_ClickToRun_AppVShNotify.exe
    │      AppVShNotify.exe
    │      USERENV.dll.cpp
    │      USERENV.dll.def
    │
    ├─C__Program Files_Common Files_microsoft shared_ClickToRun_InspectorOfficeGadget.exe
    │      InspectorOfficeGadget.exe
    │      mscoree.dll.cpp
    │      mscoree.dll.def
    │
    └─C__Program Files_Common Files_microsoft shared_ClickToRun_IntegratedOffice.exe
            IntegratedOffice.exe
            IPHLPAPI.DLL.cpp
            IPHLPAPI.DLL.def
```
