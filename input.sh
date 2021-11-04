#!/bin/bash

if [ -z "$1" ]
then
  echo "Usage: $(basename $0) <day>" >&2
  exit 1
fi

day="$(printf "day%02d\n" $1)"
dir="inputs/${day}"

name="sample.txt"

inputs="$(ls ${dir})"
if [ -n "${inputs}" ]
then
  for f in ${inputs}
  do
    fbase=$(basename $f .txt)
    [ "${fbase}" = "sample" ] && continue
    num=$(echo "${fbase}" | cut -d- -f1)
  done
  let 'num = num + 1' 
  name="$(printf "%02d-input.txt" ${num})"
fi
pbpaste > inputs/${day}/${name}
echo >> inputs/${day}/${name}
echo "Created inputs/${day}/${name}."
