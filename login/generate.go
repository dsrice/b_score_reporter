package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"login/infra/generator"
	"os"
	"strings"
)

var commandList = []string{
	"controller",
	"usecase",
	"repository",
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println(len(os.Args))
		os.Exit(1)
	}

	err := checkArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	info := generator.CodeInfo{
		PublicName:  os.Args[2],
		PrivateName: os.Args[3],
	}

	info.SetPackageInfo(os.Args[1])

	info.CreateCode()
}

// checkArgs
// generateコマンドのバリデーション
func checkArgs(args []string) error {
	command := args[1]
	pbName := args[2]
	prName := args[3]

	if !slices.Contains(commandList, command) {
		return fmt.Errorf("種別異常")
	}

	err := upperCheck(pbName)
	if err != nil {
		return fmt.Errorf("公開名称が大文字から始まっていない")
	}

	err = lowerCheck(prName)
	if err != nil {
		return fmt.Errorf("プライベート名称が小文字から始まっていない")
	}

	return nil
}

// upperCheck
// 先頭文字が大文字であるかの確認
func upperCheck(target string) error {
	t := target[:1]

	if t == strings.ToUpper(t) {
		return nil
	}

	return fmt.Errorf("対象の単語が大文字から始まっていない")
}

// lowerCheck
// 先頭文字が小文字であるかの確認
func lowerCheck(target string) error {
	t := target[:1]

	if t == strings.ToLower(t) {
		return nil
	}

	return fmt.Errorf("対象の単語が小文字から始まっていない")
}
