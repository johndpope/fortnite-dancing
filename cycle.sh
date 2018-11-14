#!/bin/bash

for filename in frames/*.txt;
do
    echo $filename
    cat $filename
done


# This stuff was for testing
#value=`cat frames/0.txt`
#echo "$value"
