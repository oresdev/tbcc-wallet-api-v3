#!/bin/bash

# Add Golang to path
echo "Configuring environment variables..."
GOROOT="/usr/local/go"
GOPATH="/mnt/d/Workspace/Goland"
if [ -d "$GOROOT" ]; then
    if ! grep -q 'GOPATH' $HOME/.bashrc ; then
        touch "$HOME/.bashrc"
        {
            echo ''
            echo '# GOLANG'
            echo 'export GOROOT='$GOROOT
            echo 'export GOPATH='$GOPATH
            echo 'export GOBIN=$GOPATH/bin'
            echo 'export PATH=$PATH:$GOROOT/bin:$GOBIN'
            echo ''
        } >> "$HOME/.bashrc"
        source "$HOME/.bashrc"
        echo "GOROOT set to $GOROOT"
        echo "GOPATH set to $GOPATH"
    fi
fi