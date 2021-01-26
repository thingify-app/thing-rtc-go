#!/bin/bash

set -e

# Trap signals so that we kill child processes before exiting.
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

# Start up server in background.
npm --prefix server run serve &

# Live-compile peer library in background.
npm --prefix peer run buildWatch &

# Serve example in background.
npm --prefix peer/example run serve &

# Keep running until we kill the script and all child processes.
wait
