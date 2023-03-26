package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func ValidateSHA256Sum(checksum string, filename string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read file")
		fmt.Println(err)
		return false
	}
	sha256sum := sha256.Sum256(data)

	bytesha256sum := sha256sum[:]

	fileChecksum := hex.EncodeToString(bytesha256sum)

	if fileChecksum != checksum {
		return false
	}

	return true
}
