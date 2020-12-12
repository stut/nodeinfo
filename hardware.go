package main

import (
	"os"

	"github.com/jaypipes/ghw"
)

func writeHardwareInfo(folder string) {
	os.Setenv("GHW_DISABLE_WARNINGS", "1")

	baseboard, _ := ghw.Baseboard()
	writeObject(folder, "baseboard", baseboard)

	bios, _ := ghw.BIOS()
	writeObject(folder, "bios", bios)

	disks, _ := ghw.Block()
	writeObject(folder, "disks", disks)

	chassis, _ := ghw.Chassis()
	writeObject(folder, "chassis", chassis)

	cpu, _ := ghw.CPU()
	writeObject(folder, "cpu", cpu)

	gpu, _ := ghw.GPU()
	writeObject(folder, "gpu", gpu)

	memory, _ := ghw.Memory()
	writeObject(folder, "memory", memory)

	network, _ := ghw.Network()
	writeObject(folder, "network", network)

	pci, _ := ghw.PCI()
	writeObject(folder, "pci", pci)

	product, _ := ghw.Product()
	writeObject(folder, "product", product)

	topology, _ := ghw.Topology()
	writeObject(folder, "topology", topology)
}
