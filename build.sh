#!/bin/sh

output_name='logpasta-'$GOOS'-'$GOARCH
if [ $GOOS = "windows" ]; then
    output_name+='.exe'
fi

GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name .