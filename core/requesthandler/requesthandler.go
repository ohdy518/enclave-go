package requesthandler

import (
	"Enclave/core/directorytools"
	"Enclave/core/filehandler"
	"Enclave/core/micro"
	"errors"
	"os"
)

const DefaultEncryptedFileExtension string = ".enclave"

var fileList []string
var encryptedFileList []string
var decryptedFileList []string
var isDeclared = false
var isEncrypted = false
var isDecrypted = false
var completedCount int = 0

func DeclareListByDirectory(directoryPath string) error {
	isDeclared = false
	fileList = make([]string, 0)
	files, err := directorytools.GetFilesInDirectory(directoryPath)
	if err != nil {
		return err
	}

	tempList := make([]string, 0)

	for _, file := range files {
		tempList = append(tempList, file)
	}

	fileList = tempList

	isDeclared = true

	return nil
}

func EncryptFileList(password string, optionalEncryptedFileExtension ...string) error {
	encryptedFileExtension := DefaultEncryptedFileExtension
	if len(optionalEncryptedFileExtension) > 0 {
		encryptedFileExtension = optionalEncryptedFileExtension[0]
	}

	if !isDeclared {
		return errors.New("not declared")
	}

	for _, file := range fileList {
		encryptedFilePath := file + encryptedFileExtension
		err := filehandler.EncryptFile(password, file, encryptedFilePath)
		if err != nil {
			return err
		}
		encryptedFileList = append(encryptedFileList, encryptedFilePath)
	}

	isEncrypted = true
	deleteFiles()

	return nil
}

func DecryptFileList(password string, optionalEncryptedFileExtension ...string) int {
	encryptedFileExtension := DefaultEncryptedFileExtension
	if len(optionalEncryptedFileExtension) > 0 {
		encryptedFileExtension = optionalEncryptedFileExtension[0]
	}

	if !isDeclared {
		return 1
	}

	for _, file := range fileList {
		decryptedFilePath := micro.SubtractString(file, encryptedFileExtension)
		err := filehandler.DecryptFile(password, file, decryptedFilePath)
		if err != nil {
			return 1
		}
		decryptedFileList = append(decryptedFileList, decryptedFilePath)
	}
	isDecrypted = true
	isEncrypted = false

	deleteFiles()

	return 0
}

func deleteFiles(optionalDeleteEncryptedFiles ...bool) int {
	deleteEncryptedFiles := false
	if len(optionalDeleteEncryptedFiles) > 0 {
		deleteEncryptedFiles = optionalDeleteEncryptedFiles[0]
	}

	if deleteEncryptedFiles {
		for _, file := range encryptedFileList {
			err := os.Remove(file)
			if err != nil {
				return 1
			}
		}
		return 0
	}

	for _, file := range fileList {
		err := os.Remove(file)
		if err != nil {
			return 1
		}
	}
	return 0
}

func GetListLength() int {
	return len(fileList)
}

func GetEncryptionCompletedCount() int {
	completedCount = len(encryptedFileList)
	return completedCount
}

func GetDecryptionCompletedCount() int {
	completedCount = len(decryptedFileList)
	return completedCount
}
