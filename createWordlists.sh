#!/bin/bash

if [ $# -ne 2 ]; then
    echo "usage: $0 <keyword> <wordlist>"
    exit 1;
fi

cp common.txt buckets-$1.txt
cp common.txt buckets$1.txt

sed -e "s/^/$1-/" $2 >> $1-buckets.txt
sed -e "s/^/$1/" $2 >> $1buckets.txt

#sed -e "s/$/-$1/" $2 >> buckets-$1.txt
ex +"%s/$/-demo/g" -cwq buckets-$1.txt

#sed -e "s/$/$1/" $2 >> buckets$1.txt
ex +"%s/$/-demo/g" -cwq buckets$1.txt
