package service

import (
	"crypto/sha256"
	"fmt"
)

func Hash(password string) (string, error) {
	// TODO add solt
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password))), nil
}
