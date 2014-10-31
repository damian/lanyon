package main

import (
	"fmt"
)

func CopyStaticAssets(source string, dest string) error {
	err = CopyDir(source, dest)
	if err != nil {
		fmt.Println("Error copying files: ", err)
		return err
	} else {
		fmt.Println("Static directory copied.")
	}

	return nil
}
