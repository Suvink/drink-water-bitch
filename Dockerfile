FROM golang:1.23.1-alpine AS build-env

RUN mkdir /app
WORKDIR /app

COPY go.mod ./

RUN go mod download

RUN addgroup -g 10014 choreo && \
    adduser --disabled-password --no-create-home --uid 10014 --ingroup choreo choreouser

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -a -installsuffix cgo -o /go/bin/drink-water-reminder -buildvcs=false

FROM alpine

COPY --from=build-env /go/bin/drink-water-reminder /go/bin/drink-water-reminder
COPY phrases.txt /phrases.txt

USER 10014

ENTRYPOINT ["/go/bin/drink-water-reminder"]
