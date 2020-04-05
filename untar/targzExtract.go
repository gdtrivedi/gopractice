package main

import (
	"github.com/walle/targz"
)

func targzExtract(sourceFile, destinationDir string) {

	targz.Extract(sourceFile, destinationDir)
}