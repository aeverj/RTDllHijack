# PE File Parser

PE File Parser 是一个解析 PE 文件导入表并生成可劫持 DLL 源代码的工具。

## 特性

- 自动发现给定目录中的可劫持 DLL
- 根据发现的 DLL 生成对应的源代码
- 支持选择编译器（MinGW 或 MSVC）
- 可排除特定文件或目录
- 提供详细输出以便调试

## 安装

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
## 获取所有C盘可执行文件dll劫持
```sh
AGHD.exe -i C:\
```
## 生成支持MingW编译器的源文件
```sh
AGHD.exe -i C:\ -c mingw
```
## 排除特定的文件或目录
```sh
AGHD.exe -i C:\ -e admin
```
## 结果
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
