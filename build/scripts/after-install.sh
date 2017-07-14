#!/bin/bash
set -ex
chown root:root /usr/share/collectd/gollectz
chmod u+s /usr/share/collectd/gollectz

cat <<HERETHIS >> /usr/share/collectd/types.db
zfs-allocated     value:GAUGE:0:U
zfs-size     value:GAUGE:0:U
zfs-free     value:GAUGE:0:U
HERETHIS

systemctl restart collectd.service


