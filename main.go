package main

import (
	"Enclave/corehandler"
	"path/filepath"
)

const Path string = "./test/"
const IsDecrypting bool = true

const Password string = "thisisatestpassword"

func main() {
	var absolutePath, err = filepath.Abs(Path)
	if err != nil {
		panic(err)
	}

	// TODO: Check if directory exists

	err = corehandler.DeclareListByDirectory(absolutePath)
	if err != nil {
		return
	}

	if IsDecrypting {
		corehandler.StartDecryption(Password)
		return
	}

	// Encrypt
	corehandler.StartEncryption(Password)
}
