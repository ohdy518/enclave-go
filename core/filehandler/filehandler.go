package filehandler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
)

// Key generation using SHA-256
func generateKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

// EncryptFile encrypts the content of the specified file using the provided password
func EncryptFile(password string, filePath string, outputPath string) error {
	// Read file content
	plaintext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Generate AES key from password
	key := generateKey(password)

	// Create AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Create a nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Encrypt the data
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// Write the encrypted data to a file
	encFile := outputPath
	err = os.WriteFile(encFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts the content of the specified encrypted file using the provided password
func DecryptFile(password string, filePath string, outputPath string) error {
	// Read file content
	ciphertext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Generate AES key from password
	key := generateKey(password)

	// Create AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Extract the nonce and ciphertext
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Write the decrypted data to a file
	decFile := outputPath
	err = os.WriteFile(decFile, plaintext, 0644)
	if err != nil {
		return err
	}

	return nil
}
