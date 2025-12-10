# cgi-perl: Go言語でPerl CGIを実行するシンプルなサーバー

このプロジェクトは、Go言語で実装されたシンプルなHTTPサーバーを使用して、Perlで書かれたCGIスクリプトを実行するテスト環境です。懐かしのCGIサイトをGoで動かすことを目的としています。

## 技術スタック
- Go言語 (サーバーサイド)
- Perl (CGIスクリプト)

## プロジェクト構造
- `main.go`: Go言語で書かれたHTTPサーバーのエントリポイント。`/cgi-bin/` へのリクエストをCGIとして処理します。
- `cgi-bin/`: CGIスクリプトを配置するディレクトリ。
  - `cgi-bin/index.cgi`: サンプルのPerl CGIスクリプト。

## 実行方法

### 1. Goサーバーの起動
プロジェクトのルートディレクトリで以下のコマンドを実行します。

```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

### 2. CGIスクリプトへのアクセス
ブラウザで以下のURLにアクセスすると、`cgi-bin/index.cgi` が実行され、結果が表示されます。

```
http://localhost:8080/cgi-bin/index.cgi
```
