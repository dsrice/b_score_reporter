#!/bin/bash

if [ $# -ne 3 ]; then
  exit 1
fi

go run infra/generate/generate.go $1 $2 $3