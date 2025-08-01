#!/usr/bin/env bash
## #ddev-generated
set -euo pipefail

NAME=$(ddev debug configyaml | grep '^name:' | cut -d' ' -f2)
TLD=$(ddev debug configyaml | grep '^project_tld:' | cut -d' ' -f2)
TLD=${TLD:-ddev.site}
HOST="kafka-ui.$NAME.$TLD"

cd kafka-go
docker build -t kafka-hostname-updater .
cp ../config.yaml .
docker run --rm -v $(pwd):/data -u $(id -u):$(id -g) kafka-hostname-updater "kafka-ui.$NAME"
docker rmi kafka-hostname-updater
cp config.yaml ..
cd ..

sed "s/NAME_REPLACE_THIS/$NAME/" traefik/config/kafka.yaml > traefik/config/kafka.tmp
sed "s/HOST_REPLACE_THIS/$HOST/" traefik/config/kafka.tmp > traefik/config/kafka.yaml
rm traefik/config/kafka.tmp
