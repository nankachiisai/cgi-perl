package main

import (
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"path/filepath"
)

func main() {
	scriptDir := "cgi-bin"

	// ディレクトリがなければ作成
	if _, err := os.Stat(scriptDir); os.IsNotExist(err) {
		_ = os.Mkdir(scriptDir, 0755)
	}

	http.HandleFunc("/cgi-bin/", func(w http.ResponseWriter, r *http.Request) {
		// URLパスからスクリプト名を取得 (test.pl)
		scriptName := r.URL.Path[len("/cgi-bin/"):]
		if scriptName == "" {
			http.Error(w, "Script name required", http.StatusBadRequest)
			return
		}

		cwd, _ := os.Getwd()
		scriptPath := filepath.Join(cwd, scriptDir, scriptName)

		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			http.Error(w, "Script not found", http.StatusNotFound)
			return
		}

		// CGIハンドラー設定
		handler := cgi.Handler{
			Path: scriptPath,
			Dir:  scriptDir,
			Env:  []string{
				// Perlに必要な環境変数があればここに追加
				// "PERL5LIB=/my/lib/path", 
			},
		}

		handler.ServeHTTP(w, r)
	})

	// 静的ファイルをホスティングするためのハンドラ
	http.Handle("/", http.FileServer(http.Dir("./contents")))

	port := ":8080"
	log.Printf("Starting CGI Server on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
