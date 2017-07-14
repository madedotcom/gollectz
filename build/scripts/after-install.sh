#!/bin/bash
set -ex
chown root:root /usr/share/collectd/gollectz
chmod u+s /usr/share/collectd/gollectz

cat <<HERETHIS >> /usr/share/collectd/types.db
allocated     value:GAUGE:0:U
size     value:GAUGE:0:U
free     value:GAUGE:0:U
HERETHIS

systemctl restart collectd.service
