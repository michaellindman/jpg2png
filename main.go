package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("dir", "", "directory to scan")
	flag.Parse()

	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	pngheader := []byte{
		0x89, 0x50, 0x4e, 0x47, 0xd, 0xa, 0x1a, 0xa, 0x0,
		0x0, 0x0, 0xd, 0x49, 0x48, 0x44, 0x52, 0x0, 0x0}

	var count int
	for _, f := range files {
		orig := *dir + "/" + f.Name()
		new := *dir + "/" + strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())) + ".png"
		if path.Ext(f.Name()) == ".jpg" {
			file, err := ioutil.ReadFile(orig)
			if err != nil {
				log.Fatal(err)
			}
			if bytes.Equal(file[:18], pngheader) == true {
				if err := os.Rename(orig, new); err != nil {
					log.Fatal(err)
				}
				count++
			}
		}
	}
	fmt.Printf("renamed %d jpg files png\n", count)
}
