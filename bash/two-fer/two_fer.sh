#!/usr/bin/env bash

main() {
    local NAME="${1:-you}"
    printf "One for %s, one for me." "$NAME"
}
main "$@"