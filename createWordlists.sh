#!/bin/bash

if [ $# -ne 2 ]; then
    echo "usage: $0 <keyword> <wordlist>"
    exit 1;
fi

sed -e "s/^/$1-/" $2 >> $1-buckets.txt
sed -e "s/^/$1/" $2 >> $1buckets.txt

# WHY YOU NOT WORK SED

#sed -e "s/$/-$1/" $2 >> buckets-$1.txt
ex +"%s/$/-demo/g" -cwq buckets-$1.txt

#sed -e "s/$/$1/" $2 >> buckets$1.txt
ex +"%s/$/-demo/g" -cwq buckets$1.txt
