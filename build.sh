#!/bin/bash

HERE=$(cd $(dirname $BASH_SOURCE) && pwd)

if [ ! -e $HERE/src/github.com/kellegous/godirs ]; then
  mkdir -p $HERE/src/github.com/kellegous
  ln -s ../../.. $HERE/src/github.com/kellegous/godirs
fi

GOPATH=$HERE go install \
  github.com/kellegous/godirs/gaan \
  github.com/kellegous/godirs/gaandoc

STATUS=$?

rm -rf src