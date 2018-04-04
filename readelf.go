package main

import (
	"os"
	"strings"
	"debug/elf"
	"log"
	"fmt"
)

var (
	elfPath     string
	showHeader  bool
	showSection bool
	showProgram bool
	showAll     bool
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
			if strings.Contains(arg, "a") {
				showAll = true
			}
			if strings.Contains(arg, "h") {
				showHeader = true
			}
			if strings.Contains(arg, "S") {
				showSection = true
			}
			if strings.Contains(arg, "S") {
				showProgram = true
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
	if showHeader || showAll {
		fmt.Printf("ELF File Header:\n")
		fmt.Printf("  Class:      %s\n", elfFile.Class)
		fmt.Printf("  Version:    %s\n", elfFile.Version)
		fmt.Printf("  Data:       %s\n", elfFile.Data)
		fmt.Printf("  OSABI:      %s\n", elfFile.OSABI)
		fmt.Printf("  ABIVersion: %d\n", elfFile.ABIVersion)
		fmt.Printf("  ByteOrder:  %s\n", elfFile.ByteOrder)
		fmt.Printf("  Type:       %s\n", elfFile.Type)
		fmt.Printf("  Machine:    %s\n", elfFile.Machine)
		fmt.Printf("  Entry:      %d\n", elfFile.Entry)
	}
	if showSection || showAll {
		sections := elfFile.Sections
		fmt.Printf("ELF Sections:\n")
		fmt.Printf("  [%2s] %-25s %-15s %-8s %-8s %-8s %2s %3s %3s %3s %3s\n", "Nr",
			"Name",
			"Type",
			"Addr",
			"Off",
			"Size",
			"ES",
			"Flg",
			"Lk",
			"Inf",
			"Al")
		for i, section := range sections {
			fmt.Printf("  [%2d] %-25s %-15s %08x %08x %08x %02x %3s %3d %3d %3d\n", i,
				section.Name,
				strings.Replace(section.Type.String(), "SHT_", "", -1),
				section.Addr,
				section.Offset,
				section.Size,
				section.Entsize,
				strings.Replace(section.Flags.String(), "SHF_", "", -1),
				section.Link,
				section.Info,
				section.Addralign)
		}
	}
}
