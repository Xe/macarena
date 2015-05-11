#!/bin/sh

cd /macarena

if ! [ -e ./config.json ]
then
	echo "Please read the readme, you need a config."
	exit 1
fi

while true
do
	./macarena -conf ./config.json
	echo "I died, waiting"
	sleep 5
done
