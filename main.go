package main

import (
	"Enclave/core/micro"
	"Enclave/corehandler"
	"log"
	"path/filepath"
)

const Path string = "D:\\Personal\\Downloads"
const IsDecrypting bool = true

const Password string = "thisisatestpassword"

func main() {
	log.Print("Running")
	var absolutePath, err = filepath.Abs(Path)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	// TODO: Check if directory exists
	exists, err := micro.Exists(Path)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Fatal("Path does not exist")
		return
	}

	err = corehandler.DeclareListByDirectory(absolutePath)
	if err != nil {
		log.Fatal(err)
		// panic(err)
		return
	}

	if IsDecrypting {
		corehandler.StartDecryption(Password)
		log.Print("Done. ")
		return
	}

	// Encrypt
	corehandler.StartEncryption(Password)
	log.Print("Done. ")
}
