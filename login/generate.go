package main

import (
	"fmt"
	"login/infra/generator"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(len(os.Args))
		os.Exit(1)
	}

	info := generator.CodeInfo{
		PublicName:  os.Args[1],
		PrivateName: os.Args[2],
		PkIFName:    "test",
		PkCTName:    "testct",
	}

	info.CreateCode()
}
