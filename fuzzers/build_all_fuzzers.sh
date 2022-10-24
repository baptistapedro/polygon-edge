#!/bin/bash

ROOT_PATH=$(pwd)
export GOROOT=/usr/local/go

# Pkgs
$GOROOT/bin/go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
$GOROOT/bin/go get github.com/dvyukov/go-fuzz/go-fuzz-dep
$GOROOT/bin/go get github.com/0xPolygon/polygon-edge

for d in ./**/;
do
	cd $(basename $d)
	
	for f in ./*.go
	do
		filename_ext=$(basename "$f")
		ext="${filename_ext##*.}"
		filename="${filename_ext%.*}"
		echo "[*] Building $filename ..."
		
		$HOME/go/bin/go-fuzz-build -libfuzzer -o $filename.a
		clang -fsanitize=fuzzer $filename.a -o $filename
		echo "[*] $filename built successfuly!"
		mv $filename $ROOT_PATH

		cd ..
	done
done
mkdir "$ROOT_PATH/bin_fuzz/"
BIN_FUZZ_PATH=$ROOT_PATH/bin_fuzz/
mv fuzz_* $BIN_FUZZ_PATH
echo "[*] Fuzzers binaries are located in the ./bin_fuzz directory"
