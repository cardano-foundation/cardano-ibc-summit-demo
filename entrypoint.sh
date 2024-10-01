#!/bin/bash

source $HOME/.bashrc
cd /root/vesseloracle/workspace/vesseloracle

DO_NOT_TRACK=1 GOFLAGS='-buildvcs=false' ignite chain serve -y -v -r

exec "$@"