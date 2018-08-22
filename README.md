# gojisample-go
Go言語で Goji と template を使った簡単なWebアプリのサンプルです

## ビルド手順

### 作業ディレクトリ作成
```bash
$ mkdir golang
$ cd golang
```

### GOPATH 用ディレクトリの作成
```bash
$ mkdir lib
$ export GOPATH=./lib
```

### Goji のライブラリを取得する
```bash
$ go get github.com/zenazn/goji
$ go get github.com/zenazn/goji/web
```

### このプロジェクトをクローンしてビルドする
```bash
$ git clone https://github.com/extra-wasabi-jp/gojisample-go
$ cd gojisample-go
$ go build
```

### 実行する
```bash
$ ./gojisample-go
```

### アクセスする
```bash
$ curl http://localhost:8000/employees
```
