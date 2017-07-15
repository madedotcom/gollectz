#!/bin/bash
set -ex

THEVERSION=$(git describe --tags --long)

fpm -C build/therpm \
    -s dir \
    -t rpm \
    -n collectd-gollectz-zfs \
    --description "Collectd plugin, writtend in golang, uses exec plugin to read disk stats from your ZFS pools." \
    -v ${THEVERSION} \
    --url "https://github.com/madedotcom/gollectz" \
    --license "MIT" \
    -m "Egidijus Ligeika" \
    --architecture noarch \
    -x "*/.git" \
    --verbose \
    --after-install build/scripts/after-install.sh \
    --after-remove build/scripts/after-remove.sh \
    --after-upgrade build/scripts/after_upgrade.sh
