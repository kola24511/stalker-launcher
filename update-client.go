package stalkerlauncher

import (
	"fmt"

	"github.com/kola24511/stalker-launcher/internal/utils/hash"
)

func UpdateClient() string {
	directory := "./client"
	fileHash, err := hash.HashDirectory(directory)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}
	return fileHash
}
