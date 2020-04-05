package main

import (
	"github.com/mholt/archiver/v3"
)

func archiverUnarchive(sourceFile, destinationDir string) {
	archiver.Unarchive(sourceFile, destinationDir)
}