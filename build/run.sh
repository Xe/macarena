#!/bin/sh

if ! [ -e ./config.json ]
then
	echo "Please read the readme, you need a config."
	exit 1
fi

if [ $DOCKER = "yes" ];
then
	macarena_bin=/go/bin/macarena
else
	macarena_bin=./macarena
fi

while true
do
	$macarena_bin -conf ./config.json
	echo "I died, waiting"
	sleep 5
done
