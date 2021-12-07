#!/bin/bash

IFS='-' read -r -a array <<< "$1"
YEAR="${array[0]}"
DAY="${array[1]}"

mkdir -p $YEAR/$DAY/01
mkdir -p $YEAR/$DAY/02

cp template/* $YEAR/$DAY/01/
