FROM golang:1.17-alpine3.14

# add a non-root user to run our code as
RUN adduser --disabled-password --gecos "" appuser
USER appuser

WORKDIR /app
COPY go.mod .
COPY go.sum .

#ENV GOPROXY=https://proxy.golang.org,direct

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]