#!/usr/bin/env bash
TARGETS="linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64 windows/arm windows/arm64 darwin/amd64 darwin/arm64"

rm -rf ./bin
for target in $TARGETS; do
    GOOS=${target%/*}
    GOARCH=${target#*/}
    echo "Building for $GOOS/$GOARCH"
    GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/$GOOS-$GOARCH/"
    (
        cd bin/$GOOS-$GOARCH/
        zip -r "../$GOOS-$GOARCH.zip" .
    )
    rm -rf "bin/$GOOS-$GOARCH/"
done