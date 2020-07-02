#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

GOARM=6 GOOS=linux GOARCH=arm go build