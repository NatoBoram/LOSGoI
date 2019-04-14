@Echo Off

Rem Make sure you're not in $GOPATH
CD %UserProfile%

Rem Enable Go 1.11 Module
Set GO111MODULE=on

Rem Go Get - IPFS
Go get -u -v -fix github.com/ipfs/go-ipfs

Rem Go Get - IPFS Cluster
Go get -u -v -fix github.com/ipfs/ipfs-cluster

Rem Make Install - IPFS
CD %UserProfile%/go/src/github.com/ipfs/go-ipfs
Make install

Rem Make Install - IPFS Cluster
CD %UserProfile%/go/src/github.com/ipfs/ipfs-cluster
Make install

CD %UserProfile%
