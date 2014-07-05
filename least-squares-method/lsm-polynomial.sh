#!/bin/bash

for num in `seq 2 11`
do
  go run lsm-polynomial.go $num
done
