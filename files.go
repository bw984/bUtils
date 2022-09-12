package bUtils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CopyFile(originFilepath, destinationFilepath string) error {
	var err error
	filename := filepath.Base(originFilepath)

	// Open the local file
	original, err := os.Open(originFilepath)
	if err != nil {
		log.Printf("Unable to open originFilepath: %s | Error: %s\n", originFilepath, err)
		return err
	}

	// Create new file in serverFolder
	serverFilepath := filepath.Join(destinationFilepath, filename)
	newFile, err := os.Create(serverFilepath)
	if err != nil {
		log.Printf("Unable to create originFilepath: %s | Error: %s\n", serverFilepath, err)
		return err
	}

	// Copy data from origin to destination file
	_, err = io.Copy(newFile, original)
	if err != nil {
		log.Printf("Unable to copy originFilepath: %s to destinationFilepath: %s | Error: %s", originFilepath, destinationFilepath, err)
		return err
	} else {
		log.Printf("Copied %s to %s \n", filename, destinationFilepath)
	}

	// Close files
	err = original.Close()
	if err != nil {
		log.Printf("Unable to close %s. Error: %s", destinationFilepath, err)
		return err
	}
	err = newFile.Close()
	if err != nil {
		log.Printf("Unable to close %s. Error: %s", serverFilepath, err)
		return err
	}
	return err
}
