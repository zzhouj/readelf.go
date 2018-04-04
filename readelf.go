package main

import (
	"os"
	"strings"
	"debug/elf"
	"log"
	"fmt"
)

var (
	elfPath       string
	isShowHeader  bool
	isShowSection bool
	isShowProgram bool
	isShowAll     bool
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
				isShowAll = true
			}
			if strings.Contains(arg, "h") {
				isShowHeader = true
			}
			if strings.Contains(arg, "S") {
				isShowSection = true
			}
			if strings.Contains(arg, "S") {
				isShowProgram = true
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
	if isShowHeader || isShowAll {
		showHeader(elfFile)
	}
	if isShowSection || isShowAll {
		showSections(elfFile.Sections)
	}
}

func showHeader(file *elf.File) {
	fmt.Printf("ELF File Header:\n")
	fmt.Printf("  Class:      %s\n", file.Class)
	fmt.Printf("  Version:    %s\n", file.Version)
	fmt.Printf("  Data:       %s\n", file.Data)
	fmt.Printf("  OSABI:      %s\n", file.OSABI)
	fmt.Printf("  ABIVersion: %d\n", file.ABIVersion)
	fmt.Printf("  ByteOrder:  %s\n", file.ByteOrder)
	fmt.Printf("  Type:       %s\n", file.Type)
	fmt.Printf("  Machine:    %s\n", file.Machine)
	fmt.Printf("  Entry:      %d\n", file.Entry)
}

func showSections(sections []*elf.Section) {
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
