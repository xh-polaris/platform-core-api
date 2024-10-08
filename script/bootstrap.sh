#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=platform.core-api
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}