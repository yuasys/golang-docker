# golang-docker

golangとdockerの学習環境です  

<p align="right"><i><u>活動記録　2023-09-30 sun</u></i></p>

### [dockerを使ってgoの最低限の環境構築](https://youtu.be/km9B4qGhb-o?si=GnJEgs9uFtRBxC8Q)

#### /dockerfile

```dockerfile
FROM golang:1.16.5
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app/
```

#### /docker-compose.yml

```yml
version: '3'

services:
  app:
    build: .
    tty: true
    ports:
      - '8080:8080'
    volumes:
      - ./app:/go/src/app #ローカルファイルをdockerfileで定義したディレクトリにマウントする

```

#### /app/main.go

```go
package main

import(
  "fmt"
  "html"
  "log"
  "net/http"
)

func main(){
  http.HndleFunc("/", func(w http.ResponseWriter, *http.
  Request){
    fmt.Fprintf(w, "Hi Hello, %q", html:EscapeString(r.Path)
    )
  })

  log.Fatal(http.ListenAndServe(":8080",nil))
}
```

### ターミナルでコンテナを動かす

```bash
# イメージの作成
docker-compose build

# バックグラウンドでコンテナを起動
docker-compose up -d

# appコンテナを起動しつつ、その後のコマンドを実行
docker-compose exec app go run main.go

```

#### ブラウザで確認する

http://localhost:8080

#### コンテナの停止

```bash
docker-compose down
```
