#!/usr/bin/env bash

# brew install wrk for MacOS

#wrk -t6 -c10 -d30s http://127.0.0.1:6666/tf/version
wrk -t6 -c100 -d60s http://127.0.0.1:6666/tf/version
