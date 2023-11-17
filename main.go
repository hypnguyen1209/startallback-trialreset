package main

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	keyPath := `Software\Microsoft\Windows\CurrentVersion\Explorer\CLSID\`

	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.READ)
	if err != nil {
		fmt.Println("error opening registry key:", err)
		return
	}
	defer key.Close()
	subkeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		fmt.Println("error reading subkeys:", err)
		return
	}
	for _, subkey := range subkeys {
		if strings.ToUpper(subkey) != subkey {
			err = registry.DeleteKey(key, subkey)
			if err != nil {
				fmt.Println("Error deleting subkey:", err)
				return
			}
			fmt.Println("trial reseted")
		}
	}
}
