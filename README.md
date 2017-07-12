# What
collectd plugin for zfs disk stats using go and the go-libzfs C libs that talk directly to zfs apis, this should work well on zfsonlinux.
things like :
```
used
available
usedbydataset
usedbychildren
usedbysnapshots
written
logicalused
volsize
reservatio
```

# Why
Because before I was using the "exec" plugin in collectd to run a bash script with sudo every x minutes and it was noticeble affecting performance when using AWS t2.micro instances.

This is an attempt to make getting metrics to collectd fast and NOT kill small machines.


# How

### Build

go build

### Put the produced binary in collect d plugins directory..
```
/usr/share/collectd
```

### Change the permissions so we run this as a root always:

```
chown root:root /usr/share/collectd/gollectz
chmod u+s /usr/share/collectd/gollectz
```

### Add a user to run the plugin as:
```
useradd mrzufse
usermod mrzufse -G root
```

### Create a plugin config file:
```
/etc/collectd.d/gollectz.conf
```

```
LoadPlugin exec
<Plugin "exec">
  Exec "mrzufse" "/usr/share/collectd/gollectz"
</Plugin>
``` 


