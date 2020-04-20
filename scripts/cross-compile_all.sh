#!/bin/sh

go clean
rm -r build

# From `go tool dist list`.
env GOOS=aix	   GOARCH=ppc64	   go build -o build/losgoi-aix-ppc64
env GOOS=android   GOARCH=386      go build -o build/losgoi-android-386
env GOOS=android   GOARCH=amd64    go build -o build/losgoi-android-amd64
env GOOS=android   GOARCH=arm      go build -o build/losgoi-android-arm
env GOOS=android   GOARCH=arm64    go build -o build/losgoi-android-arm64
env GOOS=darwin    GOARCH=386      go build -o build/losgoi-darwin-386
env GOOS=darwin    GOARCH=amd64    go build -o build/losgoi-darwin-amd64
env GOOS=darwin    GOARCH=arm      go build -o build/losgoi-darwin-arm
env GOOS=darwin    GOARCH=arm64    go build -o build/losgoi-darwin-arm64
env GOOS=dragonfly GOARCH=amd64    go build -o build/losgoi-dragonfly-amd64
env GOOS=freebsd   GOARCH=386      go build -o build/losgoi-freebsd-386
env GOOS=freebsd   GOARCH=amd64    go build -o build/losgoi-freebsd-amd64
env GOOS=freebsd   GOARCH=arm      go build -o build/losgoi-freebsd-arm
env GOOS=illumos   GOARCH=amd64    go build -o build/losgoi-illumos-amd64
env GOOS=js        GOARCH=wasm     go build -o build/losgoi-js-wasm
env GOOS=linux     GOARCH=386      go build -o build/losgoi-linux-386
env GOOS=linux     GOARCH=amd64    go build -o build/losgoi-linux-amd64
env GOOS=linux     GOARCH=arm      go build -o build/losgoi-linux-arm
env GOOS=linux     GOARCH=arm64    go build -o build/losgoi-linux-arm64
env GOOS=linux     GOARCH=mips     go build -o build/losgoi-linux-mips
env GOOS=linux     GOARCH=mips64   go build -o build/losgoi-linux-mips64
env GOOS=linux     GOARCH=mips64le go build -o build/losgoi-linux-mips64le
env GOOS=linux     GOARCH=mipsle   go build -o build/losgoi-linux-mipsle
env GOOS=linux     GOARCH=ppc64    go build -o build/losgoi-linux-ppc64
env GOOS=linux     GOARCH=ppc64le  go build -o build/losgoi-linux-ppc64le
env GOOS=linux     GOARCH=s390x    go build -o build/losgoi-linux-s390x
env GOOS=nacl      GOARCH=386      go build -o build/losgoi-nacl-386
env GOOS=nacl      GOARCH=amd64p32 go build -o build/losgoi-nacl-amd64p32
env GOOS=nacl      GOARCH=arm      go build -o build/losgoi-nacl-arm
env GOOS=netbsd    GOARCH=386      go build -o build/losgoi-netbsd-386
env GOOS=netbsd    GOARCH=amd64    go build -o build/losgoi-netbsd-amd64
env GOOS=netbsd    GOARCH=arm      go build -o build/losgoi-netbsd-arm
env GOOS=netbsd    GOARCH=arm64    go build -o build/losgoi-netbsd-arm64
env GOOS=openbsd   GOARCH=386      go build -o build/losgoi-openbsd-386
env GOOS=openbsd   GOARCH=amd64    go build -o build/losgoi-openbsd-amd64
env GOOS=openbsd   GOARCH=arm      go build -o build/losgoi-openbsd-arm
env GOOS=openbsd   GOARCH=arm64    go build -o build/losgoi-openbsd-arm64
env GOOS=plan9     GOARCH=386      go build -o build/losgoi-plan9-386
env GOOS=plan9     GOARCH=amd64    go build -o build/losgoi-plan9-amd64
env GOOS=plan9     GOARCH=arm      go build -o build/losgoi-plan9-arm
env GOOS=solaris   GOARCH=amd64    go build -o build/losgoi-solaris-amd64
env GOOS=windows   GOARCH=386      go build -o build/losgoi-windows-386.exe
env GOOS=windows   GOARCH=amd64    go build -o build/losgoi-windows-amd64.exe
env GOOS=windows   GOARCH=arm      go build -o build/losgoi-windows-arm.exe
