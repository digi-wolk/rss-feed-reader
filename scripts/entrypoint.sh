#!/usr/bin/env bash
# This file was inspired by gosec project

# Binary name based on the value specified in Dockerfile
BIN_NAME=readrss

# Expand the arguments into an array of strings. This is required because the GitHub action
# provides all arguments concatenated as a single string.
ARGS=("$@")

/bin/${BIN_NAME} "${ARGS[*]}"