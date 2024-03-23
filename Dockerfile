FROM golang:1.21 as builder
WORKDIR $GOPATH/src/todo
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 50051
ENTRYPOINT ["./app"]
