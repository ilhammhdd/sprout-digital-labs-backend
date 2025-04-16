FROM golang:1.24

WORKDIR /sprout-digital-labs-backend
COPY ./ ./
RUN go mod tidy
RUN go mod download
RUN go build -o main ./internal/cmd

CMD [ "/bin/sh" ]