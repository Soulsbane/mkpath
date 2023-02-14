package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alexflint/go-arg"
)

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	permission, err := strconv.ParseUint(args.Permission, 8, 16)

	if err != nil {
		fmt.Println(err)
	} else {
		makeDirectoryPath(args.Path, permission)
	}
}

func makeDirectoryPath(path string, permission uint64) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.FileMode(permission))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Directories created successfully.")
		}
	} else {
		fmt.Println("Directory already exists.")
	}

}
