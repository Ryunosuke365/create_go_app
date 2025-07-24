package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ryunosuke365/create_go_app/utils"
)

var dirs = []string{
	"cmd/{{.AppName}}",
	"internal/config",
	"internal/controller",
	"internal/router",
	"internal/middleware",
	"internal/model",
	"internal/repository",
	"internal/service",
	"internal/util",
}

var files = map[string]string{
	"templates/main.tpl":   "cmd/{{.AppName}}/main.go",
	"templates/router.tpl": "internal/router/router.go",
	"templates/cors.tpl":   "internal/middleware/cors.go",
}

type TemplateData struct {
	AppName string
}

func Generate(appName string) error {
	base := "./" + appName
	data := TemplateData{AppName: appName}

	// ディレクトリ作成
	for _, d := range dirs {
		path := utils.ReplaceVars(d, data)
		if err := os.MkdirAll(filepath.Join(base, path), os.ModePerm); err != nil {
			return fmt.Errorf("dir作成失敗: %w", err)
		}
	}

	// テンプレートファイル生成
	for tplSrc, dst := range files {
		outPath := filepath.Join(base, utils.ReplaceVars(dst, data))
		if err := utils.RenderTemplate(tplSrc, outPath, data); err != nil {
			return err
		}
	}

	// go mod init の実行
	if err := runGoModInit(base, appName); err != nil {
		return fmt.Errorf("go mod init に失敗しました: %w", err)
	}

	return nil
}

func runGoModInit(basePath string, moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = basePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
