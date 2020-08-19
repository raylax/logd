#!/usr/bin/env bash

os_list=(darwin linux windows)
arch_list=(amd64)

for os in "${os_list[@]}"
do
  for arch in "${arch_list[@]}"
  do
    CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "-w -s" -v -o "bin/logd_${os}_${arch}"
  done
done

