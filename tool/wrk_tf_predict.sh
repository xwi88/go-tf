#!/usr/bin/env bash

# brew install wrk for MacOS

#wrk -t12 -c200 -d30s http://127.0.0.1:6666/tf/predict
wrk -t12 -c200 -d60s http://127.0.0.1:6666/tf/predict
