#!/usr/bin/env bash

# brew install wrk for MacOS

readonly CURRENT_PATH=$(cd $(dirname "$0"); pwd)
wrk -t6 -c10 -d30s -s ${CURRENT_PATH}/tf_serving.lua http://localhost:8501/v1/models/half_plus_two:predict

