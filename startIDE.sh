#!/bin/bash
set -e
mkdir -p chedata
docker run --name ga-ide -it --rm -v /var/run/docker.sock:/var/run/docker.sock -v chedata:/data -v `pwd`:/chedir eclipse/che dir up || docker start ga-ide
