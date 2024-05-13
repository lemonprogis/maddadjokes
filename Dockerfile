FROM golang:1.21 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o maddadjokes

FROM scratch
EXPOSE 8080
COPY --from=builder /app/data /data
COPY --from=builder /app/maddadjokes /server
ENTRYPOINT ["/server"]