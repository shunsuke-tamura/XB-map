FROM golang:1.19.3-alpine

RUN apk upgrade --update && \
    apk --no-cache add git

COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum
# COPY vendor/ /go/src/app/vendor
WORKDIR /go/src/app
COPY . /go/src/app/

# airをインストール
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

# gin，gorm，driverのインストール
RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

# modファイルにあるパッケージをダウンロード
RUN go mod download

# airコマンドでgoを起動
CMD ["air", "-c", ".air.toml"]