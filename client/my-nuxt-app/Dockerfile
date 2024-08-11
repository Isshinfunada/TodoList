FROM node:alpine

# yarnをインストールします
RUN apk add --no-cache yarn

WORKDIR /app

# プロジェクトの依存関係をインストール
COPY package.json yarn.lock ./
RUN yarn install

# プロジェクトのソースコードをコピー
COPY . .

# アプリケーションのポートを公開
EXPOSE 3000

# 開発サーバーを起動
CMD ["yarn", "dev"]