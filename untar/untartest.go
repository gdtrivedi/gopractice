package main

import "fmt"

func main() {
	fmt.Println("Start")

	//WORKING
	//targzExtract("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted")
	//Untar1("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted")

	//NOTWORKING
	ExtractTarGz("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted")
	//ungzip("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted")
	//archiverUnarchive("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted/")
	//untar("/tmp/tarprocess/jaguar.tar.gz", "/tmp/tarprocess/extracted")

	fmt.Println("End")
}