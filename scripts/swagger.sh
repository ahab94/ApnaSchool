#!/bin/bash

mkdir gen gen/models gen/restapi
swagger generate server -f ./swagger.yaml -m gen/models -s gen/restapi -A ApnaSchool
