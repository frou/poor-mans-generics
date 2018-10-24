#!/bin/sh

set -e

find ./* -maxdepth 0 -type d | while read -r DIR
do
	case "$DIR" in
		(./internal)
			continue
		;;
		(*)
			pushd "$DIR" >/dev/null
			go generate
			go test
			popd >/dev/null
		;;
	esac
done
