# LOSGoI

**LOSGoI**, short for **L**ineage**OS** **Go**es to **I**PFS, is a Go daemon that hashes [LineageOS](https://download.lineageos.org/) builds using [URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore).

It provides a way to access LineageOS' builds after they're removed from the website, assuming there are people who pinned the hashes. As such, it is an unreliable service to be used only as a proof-of-concept.

## URLStore

[URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore) is an experimental feature. If it becomes unsupported from future IPFS versions, then this repository will be archived.

This program will invoke the command `ipfs urlstore add <url>`. As such, you must have [IPFS](https://github.com/ipfs/go-ipfs) installed.

## IPFS Cluster

This program will attempt to pin the hashes it created to the local IPFS Cluster with the command `ipfs-cluster-ctl pin add <hash> --name <build> --replication-min <min> --replication-max <max>`.

Replication factors are calculated as following : 

```Go
size := buildHash.Build.Size
speed := 10 * 1024 * 1024
seconds := 60

min := strconv.Itoa(1)
max := strconv.Itoa(size/(seconds*speed) + 1)
```

As such, replication factors will grow according to the following table :

| Size     | Min | Max |
| -------- | --- | --- |
|  500 MiB |   1 |   1 |
|  600 MiB |   1 |   2 |
| 1200 MiB |   1 |   3 |
|    2 GiB |   1 |   4 |
