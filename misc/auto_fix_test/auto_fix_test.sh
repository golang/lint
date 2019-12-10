#!/bin/sh

set -ue

cp ./misc/auto_fix_test/auto_fix_given.go.orig ./misc/auto_fix_test/auto_fix_given.go
go run ./golint/* --fix -- ./misc/auto_fix_test/auto_fix_given.go

if ! diff ./misc/auto_fix_test/auto_fix_given.go ./misc/auto_fix_test/auto_fix_expected; then
  echo not ok
  exit 1
fi

echo ok

