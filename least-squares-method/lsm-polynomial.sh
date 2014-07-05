#!/bin/bash

for num in `seq 2 31`
do
  go run lsm-polynomial.go $num
done
