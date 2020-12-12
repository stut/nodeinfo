package main

import (
	"github.com/zcalusic/sysinfo"
)

func writeSystemInfo(folder string) {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	writeObject(folder, "node", si.Node)
	writeObject(folder, "os", si.OS)
	writeObject(folder, "kernel", si.Kernel)
}
