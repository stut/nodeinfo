package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	outputFolder := flag.String("output-folder", ".", "the root folder for the output files")
	flag.Parse()

	os.MkdirAll(*outputFolder, 0755)
	writeHardwareInfo(*outputFolder)
	writeSystemInfo(*outputFolder)
}

func writeObject(folder string, category string, data interface{}) {
	content, _ := json.MarshalIndent(data, "", "  ")
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.json", folder, category), content, 0644)
	if err != nil {
		log.Printf("ERROR Failed to write %s/%s.json: %s\n", folder, category, err)
	}
}
