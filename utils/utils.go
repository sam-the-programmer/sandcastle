package utils

import (
	"os"
	"strings"
)

func RemoveFlags() {
	for i, v := range os.Args {
		if strings.HasPrefix(v, "-") {
			os.Args = append(os.Args[:i], os.Args[i+2:]...)
			break
		}
	}
}
