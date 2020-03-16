FROM golang:1.12.3 as builder

RUN mkdir /covid19-remote-certificate-issuer
WORKDIR /covid19-remote-certificate-issuer

ADD app/go.mod .
ADD app/go.sum .

RUN go mod download

ADD app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/covid19-remote-certificate-issuer .

FROM alpine

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/covid19-remote-certificate-issuer /app/
COPY startup.sh /app/

WORKDIR /app
CMD ["sh","startup.sh"]⏎