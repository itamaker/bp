#!/bin/bash

# build from source
echo "Building..."
go build

# install or update
if [ -f "/usr/local/bin/bp" ]
	then
	echo "Updating..."	
else
	echo "Installing..."
fi
mv bp /usr/local/bin
echo "Bp updated."