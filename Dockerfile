FROM golang:1.20-alpine

WORKDIR /app

COPY ./*.go ./
COPY ./go.* ./

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/cnmotd *.go 

FROM gcr.io/distroless/base-debian10

COPY --from=0 /app/bin/cnmotd /app/cnmotd
COPY entries /app/entries

WORKDIR /app

ENTRYPOINT [ "/app/cnmotd" ]