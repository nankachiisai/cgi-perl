# cgi-perl: Go言語でPerl CGIを実行するシンプルなサーバー

本リポジトリは、Perlで書かれたCGIスクリプトを実行するテスト環境です。
いくつかのCGIサンプルを置いています。

## 技術スタック
- Go言語 (サーバーサイド)
- Perl (テスト用CGI、アクセスカウンターCGI)

## プロジェクト構造
- `main.go`: Go言語で書かれたHTTPサーバーのエントリポイント。`/cgi-bin/` へのリクエストをCGIとして処理します。
- `cgi-bin/`: CGIスクリプトを配置するディレクトリ。
  - `cgi-bin/index.cgi`: サンプルのPerl CGIスクリプト。
  - `cgi-bin/counter/`: アクセスカウンター関連のCGIスクリプトと画像ファイルを配置するディレクトリ。
    - `cgi-bin/counter/counter.cgi`: アクセスカウンターのPerl CGIスクリプト。
    - `cgi-bin/counter/count.txt`: カウント数を記録するテキストファイル。
  - `contents/`: 静的コンテンツを配置するディレクトリ。
    - `contents/img/number_*.png`: カウンター表示用の数字画像。

## 実行方法

### 1. Goサーバーの起動
プロジェクトのルートディレクトリで以下のコマンドを実行します。

```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

### 2. テスト用CGIへのアクセス
ブラウザで以下のURLにアクセスすると、`cgi-bin/index.cgi` が実行され、結果が表示されます。

```
http://localhost:8080/cgi-bin/index.cgi
```

### 3. アクセスカウンターCGIへのアクセス
ブラウザで以下のURLにアクセスすると、`cgi-bin/counter/counter.cgi` が実行され、アクセスカウンターが表示されます。

```
http://localhost:8080/cgi-bin/counter/counter.cgi
```
