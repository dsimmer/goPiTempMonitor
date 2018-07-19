FROM golang AS build-env

COPY . ./src/github.com/dsimmer/goPiTempMonitor

RUN CGO_ENABLED=0 env GOOS=linux GOARCH=arm GOARM=5 go build -o ./src/github.com/dsimmer/goPiTempMonitor/goPiTempMonitor ./src/github.com/dsimmer/goPiTempMonitor/...

FROM alpine

WORKDIR /
RUN mkdir /root/goPiTempMonitor
COPY --from=build-env /go/src/github.com/dsimmer/goPiTempMonitor /root/goPiTempMonitor

ENTRYPOINT /root/goPiTempMonitor/goPiTempMonitor
