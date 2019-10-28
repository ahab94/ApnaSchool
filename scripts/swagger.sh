#!/bin/bash

if [ ! -d "gen" ]; then
  mkdir gen gen/models gen/restapi
fi
swagger generate server -t gen --exclude-main -A ApnaSchool
swagger generate client -t gen -A ApnaSchool
