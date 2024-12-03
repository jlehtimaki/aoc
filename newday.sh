#!/bin/bash

IFS='-' read -r -a array <<< "$1"
YEAR="${array[0]}"
DAY="${array[1]}"

echo "Creating new day $DAY in year $YEAR"
mkdir -p $YEAR/$DAY
cp -r template/* $YEAR/$DAY/
