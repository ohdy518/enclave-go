package main

import (
	"Enclave/core/micro"
	"Enclave/corehandler"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var path string = ""
var isDecrypting bool = true

var password string = ""

const pathsConfigLocation string = "./paths.enclave.config.txt"
const setupConfigLocation string = "./setup.enclave.config.txt"

func main() {
	log.Print("Running")

	if password == "" {
		setupScanner()
	}

	if path == "" {
		fileListScanner()
	} else {
		var absolutePath, err = filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
			// panic(err)
		}

		exists, err := micro.Exists(path)
		if err != nil {
			log.Fatal(err)
		}
		if !exists {
			log.Fatal("Path does not exist")
			return
		}

		err = corehandler.AppendListByDirectory(absolutePath)
		if err != nil {
			log.Fatal(err)
			// panic(err)
			return
		}
	}

	if isDecrypting {
		corehandler.StartDecryption(password)
		log.Print("Done. ")
		return
	}

	// Encrypt
	corehandler.StartEncryption(password)
	log.Print("Done. ")
}

func declare(pathsList []string) {
	for _, path := range pathsList {
		// Check if path is a directory.
		if micro.IsDirectory(path) {
			// Append as directory
			err := corehandler.AppendListByDirectory(path)
			if err != nil {
				return
			}
		} else {
			// Append as file
			err := corehandler.AppendListByFile(path)
			if err != nil {
				return
			}
		}
	}
}

func fileListScanner() {
	var pathsConfigContent, err = os.ReadFile(pathsConfigLocation)
	if err != nil {
		panic(err)
	}
	var pathsConfig = string(pathsConfigContent)
	pathsConfig = strings.ReplaceAll(pathsConfig, "\r", "")
	var pathsList = strings.Split(pathsConfig, "\n")
	var temp = make([]string, 0)
	for _, path := range pathsList {
		if path != "" {
			temp = append(temp, path)
		}
	}
	pathsList = temp
	declare(pathsList)
}

func setupScanner() {
	var setupConfigContent, err = os.ReadFile(setupConfigLocation)
	if err != nil {
		panic(err)
	}
	var setupConfig = string(setupConfigContent)
	setupConfig = strings.ReplaceAll(setupConfig, "\r", "")
	var setups = strings.Split(setupConfig, "\n")
	if setups[0] == "true" {
		isDecrypting = true
	} else if setups[0] == "false" {
		isDecrypting = false
	} else {
		panic("Unrecognized setup command: " + setups[0])
	}

	if setups[1] != "" {
		password = setups[1]
	} else {
		panic("Unrecognized setup command: " + setups[1])
	}
}
