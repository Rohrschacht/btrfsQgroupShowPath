# btrfsQgroupShowPath

## About

btrfsQgroupShowPath shows the same output as regular `btrfs qgroup show`, but
with an additional column showing the path of the corresponding subvolume.

## Usage

Simply run btrfsQgroupShowPath as root and give it a mountpoint to the btrfs
filesystem:

```
$ sudo btrfsQgroupShowPath /mnt/btrfs

qgroupid         rfer         excl  path
--------         ----         ----  ----
0/5          16.00KiB     16.00KiB  /
0/1098       41.80GiB      4.81GiB  @rootfs
0/1100       16.00KiB     16.00KiB  @snapshots
0/1521       85.95GiB     34.85MiB  @home
0/1575       42.16GiB      2.64GiB  @snapshots/@rootfs.20210624T0026
0/1592        1.25GiB      1.25GiB  @libvirtimages
0/1642       43.08GiB      1.66GiB  @snapshots/@rootfs.20210630T1840
0/1724       39.70GiB    975.56MiB  @snapshots/@rootfs.20210706T1527
0/1745       40.61GiB    746.57MiB  @snapshots/@rootfs.20210708T1555
0/1750       16.00KiB     16.00KiB  @remotesnapshots
0/1759       39.92GiB     94.48MiB  @remotesnapshots/@rootfs.20210709T2019
0/1760       81.50GiB      2.33GiB  @remotesnapshots/@home.20210709T2019
0/1770       39.59GiB     84.02MiB  @snapshots/@rootfs.20210711T1717
0/1824       40.12GiB    258.02MiB  @snapshots/@rootfs.20210715T1644
0/1873       40.57GiB    361.75MiB  @snapshots/@rootfs.20210720T2355
0/1874       84.61GiB    236.67MiB  @snapshots/@home.20210721T0001
0/1879       85.95GiB     40.02MiB  @snapshots/@home.20210721T0101
0/1880       85.95GiB     25.87MiB  @snapshots/@home.20210721T0201
```

# Install

Install the tool using go get:

```
$ go get github.com/rohrschacht/btrfsQgroupShowPath
```
