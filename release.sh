#!/bin/bash

version=$1

if [ -z "$version" ];
then
	echo "$0 <version>"
fi

echo "Releasing Macarena $version"
rm -rf "./build_actual/*$version*"

dir="./build_actual/macarena-$version-linux-amd64"
rm -rf $dir*

mkdir $dir -p

printf "Copying base files..."

cp ./build/UNLICENSE         $dir
cp ./build/README.md         $dir
cp ./build/run.sh            $dir
cp ./build/example.conf.json $dir

godocdown github.com/Xe/macarena/config > $dir/config.md

echo "  done"

printf "Rebuilding macarena..."

go clean ./...
go build

echo " done"

printf "Making tarball..."

cd ./build_actual
tar cf "macarena-$version-linux-amd64.tar" "macarena-$version-linux-amd64"
xz -z "macarena-$version-linux-amd64.tar"

echo "      done"

echo "Build done at ./build_actual/macarena-$version-linux-amd64.xz"
