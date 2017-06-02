# What
collectd plugin for zfs disk stats using go c lib, this should work well on zfsonlinux.
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

