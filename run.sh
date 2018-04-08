#!/bin/sh

# install glide
PKG_OK=$(command -v glide)
if [ "" = "$PKG_OK" ]; then
    echo "Setting up glide ..."
    curl https://glide.sh/get | sh
fi

glide init