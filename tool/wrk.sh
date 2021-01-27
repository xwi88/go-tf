#!/usr/bin/env bash

# brew install wrk for MacOS

#wrk -t6 -c100 -d30s http://127.0.0.1:6666/api/list
#wrk -t6 -c100 -d60s http://127.0.0.1:6666/api/list
#wrk -t12 -c100 -d30s http://127.0.0.1:6666/api/list
wrk -t12 -c200 -d30s http://127.0.0.1:6666/api/list
wrk -t12 -c200 -d60s http://127.0.0.1:6666/api/list
