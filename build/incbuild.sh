#!/usr/bin/env sh
if test -f "$1"; then
  oldnum=`cat $1`
  echo `expr $oldnum + 1` > $1
else
  echo 1 > $1
fi
