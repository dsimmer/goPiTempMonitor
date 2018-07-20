# goPiTempMonitor

to build for pi:
CC=arm-linux-gnueabi-gcc CGO_ENABLED=1  env GOOS=linux GOARCH=arm GOARM=5 go build