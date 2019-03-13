#!/bin/sh

ipfs-cluster-ctl add -w --chunker=rabin --rmin 1 --rmax 1 

ipfs-cluster-ctl add -w --chunker=rabin --rmin 1 --rmax 1 https://mirrorbits.lineageos.org/full/hammerhead/20190126/lineage-14.1-20190126-nightly-hammerhead-signed.zip
