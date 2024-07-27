package corehandler

import "Enclave/core/requesthandler"

func DeclareListByDirectory(path string) error {
	var err = requesthandler.DeclareListByDirectory(path)
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
	requesthandler.EncryptFileList(password)
}

func StartDecryption(password string) int {
	var result = requesthandler.DecryptFileList(password)
	if result == 1 {
		return 1
	}
	return 0
}
