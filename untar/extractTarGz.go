package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ExtractTarGz(sourceFilePath, destinationDir string) {

	r, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("error")
	}

	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	prefix := destinationDir
	if !strings.HasSuffix(prefix, string(os.PathSeparator)) {
		prefix = prefix + string(os.PathSeparator)
	}

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
			case tar.TypeDir, 'V':
				if err := os.MkdirAll(prefix + header.Name, 0755); err != nil {
					log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
				}

			case tar.TypeReg:
				outFile, err := os.Create(prefix + header.Name)
				if err != nil {
					log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
				}
				if _, err := io.Copy(outFile, tarReader); err != nil {
					log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
				}
				outFile.Close()

			case tar.TypeSymlink:
				if err := os.Symlink(header.Linkname, prefix + header.Name); err != nil {
					log.Fatalf("ExtractTarGz: Symlink() failed: %s", err.Error())
				}

			default:
				log.Fatalf(
					"ExtractTarGz: uknown type: %c in %s",
					header.Typeflag,
					header.Name)
		}

	}
}