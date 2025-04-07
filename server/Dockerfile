FROM golang:1.24

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
COPY . .
RUN go mod download

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 5000

CMD ["app"]