package corehandler

import (
	"Enclave/core/requesthandler"
	"log"
)

func DeclareListByDirectory(path string) error {
	var err = requesthandler.DeclareListByDirectory(path)
	if err != nil {
		return err
	}

	return nil
}

func AppendListByDirectory(path string) error {
	var err = requesthandler.AppendListByDirectory(path)
	if err != nil {
		return err
	}
	return nil
}

func AppendListByFile(path string) error {
	var err = requesthandler.AppendListByFile(path)
	if err != nil {
		return err
	}
	return nil
}

func GetListLength() int {
	return requesthandler.GetListLength()
}

func GetEncryptionCompletedCount() int {
	return requesthandler.GetEncryptionCompletedCount()
}

func GetDecryptionCompletedCount() int {
	return requesthandler.GetDecryptionCompletedCount()
}

func StartEncryption(password string) {
	var err = requesthandler.EncryptFileList(password)
	if err != nil {
		log.Fatal(err)
	}
}

func StartDecryption(password string) int {
	var result = requesthandler.DecryptFileList(password)
	if result == 1 {
		return 1
	}
	return 0
}
