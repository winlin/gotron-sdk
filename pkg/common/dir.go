package common

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("gotron-sdk GetCurrentDir() Error: %s", err.Error())
	}
	return dir
}

func GetHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		fmt.Printf("gotron-sdk GetHomeDir() Error: %s", err.Error())
	}
	return dir
}
