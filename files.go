package main

import (
	"io"
	"log"
	"os"
)

func CopyFile(originFilepath, destinationFilepath string) error {
	var err error

	// Open the local file
	original, err := os.Open(originFilepath)
	if err != nil {
		log.Printf("Unable to open originFilepath: %s | Error: %s\n", originFilepath, err)
		return err
	}

	// Create new file for destination
	newFile, err := os.Create(destinationFilepath)
	if err != nil {
		log.Printf("Unable to create destinationFilepath: %s | Error: %s\n", destinationFilepath, err)
		return err
	}

	// Copy data from origin to destination file
	_, err = io.Copy(newFile, original)
	if err != nil {
		log.Printf("Unable to copy originFilepath: %s to destinationFilepath: %s | Error: %s\n", originFilepath, destinationFilepath, err)
		return err
	} else {
		log.Printf("Copied %s to %s\n", originFilepath, destinationFilepath)
	}

	// Close files
	err = original.Close()
	if err != nil {
		log.Printf("Unable to close %s | Error: %s\n", destinationFilepath, err)
		return err
	}
	err = newFile.Close()
	if err != nil {
		log.Printf("Unable to close %s | Error: %s\n", destinationFilepath, err)
		return err
	}
	return err
}
