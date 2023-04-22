package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv
// envファイルの読み込み処理
func LoadEnv() {
	err := godotenv.Load(fmt.Sprintf("%s/%s.env", os.Getenv("APP_ROOT"), os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Printf("読み込みエラー: %v", err)
	}
}
