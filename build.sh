#!/bin/bash

prepare_go() {
    rm -f go.mod
    rm -f go.sum
    go mod init github.com/amnezia-vpn/amnezia-libxray
    go mod tidy
}

prepare_gomobile() {
    go install golang.org/x/mobile/cmd/gomobile@latest
    go get golang.org/x/mobile/cmd/gobind
    export PATH=~/go/bin:$PATH
    gomobile init
    go get -d golang.org/x/mobile/cmd/gomobile
}

build_apple() {
    rm -fr *.xcframework
    prepare_gomobile
    gomobile bind -target ios,iossimulator,macos -iosversion 15.0
}

build_android() {
    rm -fr *.jar
    rm -fr *.aar
    prepare_gomobile
    rm -fr assets
    mkdir -p assets/geo
    mv dat/* assets/geo
    gomobile bind -target android -androidapi 24 -javapkg=org.amnezia.vpn.protocol.xray -o libxray.aar -ldflags="-w -s -buildid=" -trimpath
}

download_geo() {
    go run main/main.go
}

echo "will build libxray for $1"
prepare_go
download_geo
if [ "$1" != "apple" ]; then
build_android
else
build_apple
fi
echo "build libxray done"
