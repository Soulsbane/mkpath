package main

import "github.com/alexflint/go-arg"

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	// path := "./new_directory"
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	err = os.MkdirAll(path, 0755)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("Directory created successfully.")
	// 	}
	// } else {
	// 	fmt.Println("Directory already exists.")
	// }
}
