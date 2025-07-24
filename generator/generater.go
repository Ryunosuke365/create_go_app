package generator

import (
	"fmt"
	"os"
	"path/filepath"
	// 埋め込み変数を明示的にimport
)

func GenerateFiles(appName string) error {
	fmt.Println("📦 プロジェクト構成を生成中...")

	dirs := []string{
		filepath.Join(appName, "cmd", appName),
		filepath.Join(appName, "internal", "router"),
		filepath.Join(appName, "internal", "middleware"),
		filepath.Join(appName, "internal", "controller"),
		filepath.Join(appName, "internal", "model"),
		filepath.Join(appName, "internal", "repository"),
		filepath.Join(appName, "internal", "service"),
		filepath.Join(appName, "internal", "config"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("❌ ディレクトリ作成失敗: %w", err)
		}
	}

	// 埋め込んだテンプレートからファイルを生成
	files := map[string]string{
		filepath.Join(appName, "cmd", appName, "main.go"):           MainTemplate,
		filepath.Join(appName, "internal", "router", "router.go"):   RouterTemplate,
		filepath.Join(appName, "internal", "middleware", "cors.go"): CorsTemplate,
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("❌ ファイル生成失敗: %w", err)
		}
	}

	fmt.Println("✅ プロジェクト生成が完了しました！")
	return nil
}
