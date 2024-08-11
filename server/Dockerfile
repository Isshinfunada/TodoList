# ビルドステージ
FROM golang:alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && apk add --no-cache git

# Goモジュールをコピーして依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# 追加：wait-for-it.shスクリプトをコピーして実行権限を付与
# 削除：COPY wait-for-it.sh /wait-for-it.sh
# 削除：RUN chmod +x /wait-for-it.sh

# アプリケーションをビルド
RUN go build -o main .

# 実行ステージ
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドされたバイナリをコピー
COPY --from=builder /app/main .
# wait-for-it.shスクリプトをコピー
# 削除：COPY --from=builder /wait-for-it.sh .

# アプリケーションのポートを公開
EXPOSE 8080

# アプリケーションを実行
# CMD ["/wait-for-it.sh", "db:5432", "--", "./main"]
CMD ["./main"]
