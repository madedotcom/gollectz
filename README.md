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

This was tested and used on:
* `UBUNTU Linux 4.4.0-77-generic #98-Ubuntu SMP Wed Apr 26 08:34:02 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux, dkms/xenial-updates,xenial-updates,now 2.2.0.3-2ubuntu11.3, zfsutils-linux/now 0.6.5.6-0ubuntu16`
* `CENTOS 3.10.0-514.16.1.el7.x86_64 #1 SMP Wed Apr 12 15:04:24 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux, zfs-0.6.5.9-1.el7.centos.x86_64, zfs-release-1-3.el7.centos.noarch`

 

# Why
Because before I was using the "exec" plugin in collectd to run a bash script with sudo every x minutes and it was noticeble affecting performance when using AWS t2.micro instances.

This is an attempt to make getting metrics to collectd fast and NOT kill small machines.

# How

This uses `libzfslinux-dev`, you should have this installed on your build machine.

### Build
I am using golang version 1.8.3
I am assuming you have your go environment setup correctly, for example:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Then build:
```
go get github.com/madedotcom/gollectz
go build github.com/madedotcom/gollectz
```

You could also use something like this:
https://github.com/egidijus/go-builder

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

# More

If you would like to report other things, you can add more properties to monitor in this list:
```
var properties = []zfs.Prop{
    zfs.PoolPropSize,
    zfs.PoolPropFree,
    zfs.PoolPropAllocated,
}

```
Be aware that not all values are parsed correctly or at all by this plugin.
