package stalkerlauncher

import (
	"fmt"

	"github.com/kola24511/stalker-launcher/internal/utils/hash"
)

func UpdateClient() {
	directory := "./client"
	fileHash, err := hash.HashDirectory(directory)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Printf("Хэш папки: %s\n", fileHash)
}
