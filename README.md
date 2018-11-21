# LOSGoI

**LOSGoI**, short for **L**ineage**OS** **Go**es to **I**PFS, is a Go daemon that hashes [LineageOS](https://download.lineageos.org/) builds using [URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore).

It provides a way to access LineageOS' builds after they're removed from the website, assuming there are people who pinned the hashes. As such, it is an unreliable service to be used only as a proof-of-concept.

## URLStore

[URLStore](https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore) is an experimental feature. If it becomes unsupported from future IPFS versions, then this repository will be archived.

This program will invoke the command `ipfs urlstore add -t <url>`. As such, you must have [IPFS](https://github.com/ipfs/go-ipfs) installed.

```markdown
USAGE
  ipfs urlstore add <url> - Add URL via urlstore.

SYNOPSIS
  ipfs urlstore add [--trickle | -t] [--] <url>

ARGUMENTS

  <url> - URL to add to IPFS

OPTIONS

  -t, --trickle bool - Use trickle-dag format for dag generation.

DESCRIPTION

  Add URLs to ipfs without storing the data locally.

  The URL provided must be stable and ideally on a web server under your
  control.

  The file is added using raw-leaves but otherwise using the default
  settings for 'ipfs add'.

  The file is not pinned, so this command should be followed by an 'ipfs
  pin add'.

  This command is considered temporary until a better solution can be
  found.  It may disappear or the semantics can change at any
  time.
```

This program will not attempt to pin the hashes it created. It will store them in a database so they don't get re-hashed again.
