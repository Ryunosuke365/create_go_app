package main

import (
	"fmt"
	"os"

	"github.com/ryunosuke365/create_go_app/generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ プロジェクト名を指定してください")
		fmt.Println("Usage: create-go-app myapp")
		return
	}

	projectName := os.Args[1]
	if err := generator.GenerateFiles(projectName); err != nil {
		fmt.Printf("❌ 生成エラー: %v\n", err)
		return
	}

	fmt.Printf("✅ プロジェクト '%s' を作成しました！\n", projectName)
}
