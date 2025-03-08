# Dockerfile
FROM golang:1.23

WORKDIR /app

# 必要な依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# hot reload 用ツール air のインストール
RUN go install github.com/air-verse/air@latest

# ポート50051を公開
EXPOSE 50051

# air によりホットリロードを実施
CMD ["air"]
