#!/usr/bin/env bash
main () {
  local count
  for (( i=0; i<${#1}; i++ )); do
    (( count += ${1:$i:1} ** ${#1} ))
  done
  [ $count -eq $1 ] && echo true || echo false
}
main "$@"
exit 0