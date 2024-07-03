package models

type DLLInfo struct {
	DllName    string
	Arch       string
	FunName    []string
	FunOrdinal []uint32
}
