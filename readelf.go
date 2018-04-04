package main

import (
	"os"
	"strings"
	"debug/elf"
	"log"
)

var (
	elfPath      string
	showHeader   bool
	showSections bool
	showAll      bool = true
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") { // options
			if strings.Contains(arg, "h") {
				showHeader = true
				showAll = false
			}
			if strings.Contains(arg, "S") {
				showSections = true
				showAll = false
			}
		} else { // elf path
			elfPath = arg
		}
	}
	if elfPath == "" {
		log.Fatal("Usage: readelf [-hS] elfPath\n")
	}
	elfFile, err := elf.Open(elfPath)
	if err != nil {
		log.Fatal(err)
	}
	defer elfFile.Close()
	if showHeader {
		h := elfFile.FileHeader
		log.Printf("ELF File Header:\n")
		log.Printf("  Class:      %s\n", h.Class)
		log.Printf("  Version:    %s\n", h.Version)
		log.Printf("  Data:       %s\n", h.Data)
		log.Printf("  OSABI:      %s\n", h.OSABI)
		log.Printf("  ABIVersion: %d\n", h.ABIVersion)
		log.Printf("  ByteOrder:  %s\n", h.ByteOrder)
		log.Printf("  Type:       %s\n", h.Type)
		log.Printf("  Machine:    %s\n", h.Machine)
		log.Printf("  Entry:      %d\n", h.Entry)
	}
}
