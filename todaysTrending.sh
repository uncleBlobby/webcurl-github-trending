#!/bin/sh

curl https://github.com/trending > results.txt 2>&1

go run . 