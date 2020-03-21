#!/bin/bash

dirs=`find . -name "*.go"`

shape=()
for dir in $dirs;
do
  shape+=(`echo $dir | awk -F'.' '{print $2}' | awk -F '/' '{print $2}' | awk -F '_' '{print $1}'`)
done
# mkdir `printf "%s\n" "${shape[@]}" | sort -u`
uniq_shape=`printf "%s\n" "${shape[@]}" | sort -u`
for us in $uniq_shape;
do
  mv $(ls . | grep ${us}_) ${us}
done
# num=151
# echo `find -E . -type f -regex "028_\w+.go"`
