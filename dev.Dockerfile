FROM golang:1.18-alpine
RUN mkdir /app
ADD .. /app/
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]