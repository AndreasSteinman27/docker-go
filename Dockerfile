FROM golang:1.9
WORKDIR /go/src/bettercloud
COPY . .
RUN go get ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=0 /go/src/bettercloud/app .
CMD ["./app"]
