FROM golang AS build-env

COPY . ./src/github.com/dsimmer/goPiTempMonitor
RUN apt-get update
RUN yes | apt-get install gcc-arm-linux-gnueabi
RUN CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 env GOOS=linux GOARCH=arm GOARM=5 go build -o ./src/github.com/dsimmer/goPiTempMonitor/goPiTempMonitor ./src/github.com/dsimmer/goPiTempMonitor/...

FROM alpine

WORKDIR /
RUN mkdir /root/goPiTempMonitor
COPY --from=build-env /go/src/github.com/dsimmer/goPiTempMonitor /root/goPiTempMonitor

ENTRYPOINT /root/goPiTempMonitor/goPiTempMonitor
