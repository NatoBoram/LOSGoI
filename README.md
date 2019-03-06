# LOSGoI

[![Pipeline Status](https://gitlab.com/NatoBoram/LOSGoI/badges/master/pipeline.svg)](https://gitlab.com/NatoBoram/LOSGoI/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/NatoBoram/LOSGoI)](https://goreportcard.com/report/gitlab.com/NatoBoram/LOSGoI)
[![GoDoc](https://godoc.org/gitlab.com/NatoBoram/LOSGoI?status.svg)](https://godoc.org/gitlab.com/NatoBoram/LOSGoI)

**LOSGoI**, short for **L**ineage**OS** **Go**es to **I**PFS, is a Go daemon that hashes [LineageOS](https://download.lineageos.org/) builds using [URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore).

It provides a way to access LineageOS' builds after they're removed from the website, assuming there are people who pinned the hashes. **As such, it is an unreliable service to be used only as a proof-of-concept.**

## URLStore

[URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore) is an experimental feature. If it becomes unsupported from future IPFS versions, then this repository will be archived.

This program will invoke the command `ipfs urlstore add <url>`. As such, you must have [IPFS](https://github.com/ipfs/go-ipfs) installed.

## IPFS Cluster

This program will attempt to pin the hashes it created to the local [IPFS Cluster](https://github.com/ipfs/ipfs-cluster) with the command `ipfs-cluster-ctl pin add <hash> --name <build> --replication-min <min> --replication-max <max>`.

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

The goal of this replication factor is to be able to download a build within **60 seconds** at **10 mbps**.

## Getting started

### Dependencies

```bash
sudo apt install mariadb-server
sudo snap install ipfs --edge
sudo snap install ipfs-cluster --edge
```

### Installation

```bash
go get -u -v -fix gitlab.com/NatoBoram/LOSGoI
```

You also need to import [`Configs/database.sql`](https://gitlab.com/NatoBoram/LOSGoI/blob/master/Configs/database.sql). You can optionally get inspiration from the `systemd` services in [`Configs`](https://gitlab.com/NatoBoram/LOSGoI/tree/master/Configs).

### Configuration

MariaDB is required. It will establish a MySQL connection using the configuration file located at `./LOSGoI/database.json`. Here's a template.

```json
{
	"User": "LOSGoI",
	"Password": "",
	"Address": "localhost",
	"Port": 3306,
	"Database": "LOSGoI"
}
```

## Usage

Only one instance of LOSGoI should run on any given IPFS Cluster. However, every single IPFS Cluster running LOSGoI will improve the network, even if they do not share the same cluster.

Here's how to get the best results.

1. Setup the master node, preferably the one which will also host the associated website. You'll need to install **IPFS**, **IPFS Cluster**, **MariaDB**, and the necessary for the public-facing website. You can find a functioning website template at [gitlab.com/NatoBoram/lineageos-on-ipfs.com](https://gitlab.com/NatoBoram/lineageos-on-ipfs.com).
2. Run LOSGoI on the master node
3. **Wait for LOSGoI to complete its first pinning round.** This can take a day or two, as it needs to hash approximatively 700 LineageOS builds.
4. Only then you can bootstrap other IPFS Cluster nodes. It's not necessary if you have enough disk space on the original node, but it can help serve larger builds and to pin them faster.