package main

import (
	"Enclave/core/requesthandler"
	"log"
)

func main() {
	err := requesthandler.DeclareListByDirectory("./test/")
	if err != nil {
		log.Fatal(err)
	}

	interr := requesthandler.DecryptFileList("test")
	if interr != 0 {
		log.Fatal(interr)
	}

	/*
		filePath := "./test/example.txt"
		outputPath := "./test/output.enc"
		password := "your-password"

		files, dirErr := directorytools.GetFilesInDirectory("./test")
		if dirErr != nil {
			log.Fatal(dirErr)
		}
		fmt.Printf("files = %#v\n", files)

		err := filehandler.EncryptFile(filePath, outputPath, password)
		if err != nil {
			log.Fatalf("Error encrypting file: %v", err)
		}
		fmt.Println("File encrypted successfully! ")

		err = filehandler.DecryptFile(outputPath, filePath, password)
		if err != nil {
			log.Fatalf("Error decrypting file: %v", err)
		}
		fmt.Println("File decrypted successfully! ")
	*/
}
