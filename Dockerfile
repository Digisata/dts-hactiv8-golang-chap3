FROM golang:1.19-alpine
LABEL maintener="Hanif Naufal <hnaufal123@gmail.com>"
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN swag init
RUN go build -o ./out/dist .
CMD ["./out/dist"]