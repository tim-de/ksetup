package main

import (
	"log"
	"archive/zip"
	"ksetup/throbber"
	"os"
	"io"
)

func unzip_file(filename string) {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		throbber.Delay()
		log.Printf("Error: Failed to open zip file %s: %s\n", filename, err.Error())
		return
	}
	defer reader.Close()

	for _, subfile := range reader.File {
		filehandle, err := subfile.Open()
		if err != nil {
			log.Printf("Error: Failed to open '%s' within %s: %s\n",
				subfile.Name, filename, err.Error())
			return
		}

		if subfile.FileInfo().IsDir() {
			if err := os.Mkdir(subfile.Name, 0755); err != nil {
				log.Printf("Error: Failed to create directory %s: %s\n",
					subfile.Name, err.Error())
				return
			}
			continue
		}
		destfile, err := os.Create(subfile.Name)
		if err != nil {
			log.Printf("Error: Failed to create file %s: %s\n",
				subfile.Name, err.Error())
		}
		if _, err := io.Copy(destfile, filehandle); err != nil {
			log.Printf("Error: Failed to copy file %s: %s\n",
				subfile.Name, err.Error())
		}
	}
}
