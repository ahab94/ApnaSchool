#!/bin/bash

 grep --color=always -rni -E "TODO|FIXME" **/*.go > TODO.md
 sed -i '1s/^/# TODOs\n /' TODO.md
