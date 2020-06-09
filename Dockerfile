FROM golang:1.14-alpine

# 作業ディレクトリの作成・設定
WORKDIR /gorm-initiation

# Go Modules を有効化
ENV GO111MODULE=on

COPY ./src/go.mod .
# COPY go.sum .

# go.mod 内のパッケージをダウンロード
RUN go mod download

EXPOSE 8080
