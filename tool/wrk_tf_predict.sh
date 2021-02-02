#!/usr/bin/env bash

# brew install wrk for MacOS

readonly CURRENT_PATH=$(cd $(dirname "$0"); pwd)
wrk -t6 -c10 -d30s -s ${CURRENT_PATH}/tf_predict.lua http://localhost:6666/tf/predict
