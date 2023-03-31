# UPS Tools
[![Go](https://github.com/rameshvarun/ups/actions/workflows/go.yml/badge.svg)](https://github.com/rameshvarun/ups/actions/workflows/go.yml)

## Installation

Builds can be downloaded from the [releases page](https://github.com/rameshvarun/ups/releases/latest). Put the executable somewhere in your PATH. If you have the Go toolchain setup, you can also run `go get github.com/rameshvarun/ups` to install from source.

## Commands

### `ups diff -b base -m modified -o output.ups`
Create a UPS patch file.

### `ups apply -b base -p patch.ups -o output`
Apply a UPS patch to a file.

## Links
- [UPS Format Specification](http://individual.utoronto.ca/dmeunier/ups-spec.pdf)
- [mGBA UPS Reader](https://github.com/mgba-emu/mgba/blob/master/src/util/patch-ups.c)
- [UPS Patcher in .NET](http://www.romhacking.net/utilities/606/)
