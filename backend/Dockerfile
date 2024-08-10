# Dockerfile
# TODO: マルチステージビルドにする

FROM golang:alpine

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && apk add --no-cache git

# Goモジュールをコピーして依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションのポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
