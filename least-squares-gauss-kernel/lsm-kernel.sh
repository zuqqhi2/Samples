#!/bin/bash

for num in `seq 1 50`
do
  LEARN=`expr $num \* 50`
  go run lsm-kernel.go $LEARN
done
