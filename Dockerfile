FROM golang:1.23.0

WORKDIR /usr/src/app
COPY . .

# RUN go mod init github.com/minmaxmar/bankapp
# RUN go get github.com/gofiber/fiber/v2