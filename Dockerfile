FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o httpbin-clone ./cmd/

EXPOSE 8080

CMD ["./httpbin-clone"]

## Go yüklü küçük bir Linux ortamı açıyor
## Projeni içine kopyalıyor
## go build ile binary üretiyor
## ./httpbin-clone çalıştırıyor
## 8080 portunu dışarı açacağını söylüyor