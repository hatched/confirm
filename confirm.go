package confirm

import (
	"fmt"
)

func Ask(q string) (bool, error) {
	fmt.Print(q + " [Y/n]? ")

	var (
		confirm string
		ok      = true
	)
	fmt.Scanln(&confirm)

	if confirm == "n" {
		ok = false
	}
	return ok, nil
}
