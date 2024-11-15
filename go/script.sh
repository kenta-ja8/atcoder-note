#!/bin/bash

set -euxo pipefail
shopt -s expand_aliases

source ../activate.sh

if [ "$#" -eq 0 ]; then
  echo "Usage: $0 <function_name>"
  exit 1
fi

function submit() {
  contest=$(cat ./contest.txt)
  task=$2
  dir="archive/$contest/$task"

  go run github.com/mrsombre/codingame-golang-merger/cmd/cgmerge -o "$dir/merged.go"
  cp ./input.txt "$dir"
  cd "$dir"
  acc submit --contest "$contest" merged.go
  cd -
}

function new() {
  contest="$1"
  dir="archive"

  cd "$dir"
  acc new "$contest"
  cd -
  cat "$1" > contest.txt
}

# function add() {
#   dir="archive/$1"
#   cd "$dir"
#   acc add
#   cd -
# }

function test() {
  task=$1
  contest=$(cat ./contest.txt)

  oj test -c "go run ." -d "./archive/${contest}/${task}/tests/"
}

case "$1" in
submit)
  submit "$2"
  ;;
new)
  new "$2"
  ;;
# add)
#   add "$2"
#   ;;
test)
  test "$2"
  ;;
*)
  echo "Unknown function: $1"
  exit 1
  ;;
esac
