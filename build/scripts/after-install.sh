#!/bin/bash
set -ex
chown root:root /usr/share/collectd/gollectz
chmod u+s /usr/share/collectd/gollectz

cat <<HERETHIS >> /usr/share/collectd/types.db
zfs_allocated     value:GAUGE:0:U
zfs_size     value:GAUGE:0:U
zfs_free     value:GAUGE:0:U
HERETHIS

systemctl restart collectd.service


