#!/bin/bash
set -ex
export BUILD_NUMBER=10

cd build
fpm -s dir \
    -t rpm \
    -n collectd-gollectz-zfs \
    --description "golang exec plugin to read disk stats from your ZFS pools" \
    -v 0.0.${BUILD_NUMBER} \
    --url "https://github.com/madedotcom/gollectz" \
    --license "MIT" \
    -m "Egidijus Ligeika" \
    --architecture noarch \
    -x "*/.git" \
    --verbose \
    --after-install scripts/after-install.sh \
    therpm/=/




