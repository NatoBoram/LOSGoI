@Echo Off

Rem Make sure you're not in $GOPATH
CD %UserProfile%

Rem Go Get
Go get -u -v github.com/ipfs/go-ipfs
Go get -u -v github.com/ipfs/ipfs-cluster

Rem Go Install
Go install github.com/ipfs/go-ipfs/cmd/ipfs
Go install github.com/ipfs/ipfs-cluster/cmd/ipfs-cluster-service
Go install github.com/ipfs/ipfs-cluster/cmd/ipfs-cluster-ctl

CD %UserProfile%
