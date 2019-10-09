#!/bin/sh
set -e

wget https://dl.google.com/go/go1.12.10.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.12.10.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
