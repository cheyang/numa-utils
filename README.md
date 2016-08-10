# numa-utils
go bindings for numa

## godep and dependency management

Mindfly uses [godep](https://github.com/tools/godep) to manage dependencies. It is not strictly required for building Mindfly but it is required when managing dependencies under the Godeps/ tree, and is required by a number of the build and test scripts. Please make sure that ``godep`` is installed and in your ``$PATH``.

### Installing godep

There are many ways to build and host go binaries. Here is an easy way to get utilities like `godep` installed:


1) Create a new GOPATH for your tools and install godep:

```sh
export GOPATH=$HOME/go-tools
mkdir -p $GOPATH
go get github.com/tools/godep
```

2) Add $GOPATH/bin to your path. Typically you'd add this to your ~/.profile:

```sh
export GOPATH=$HOME/go-tools
export PATH=$PATH:$GOPATH/bin
```

### Using godep

Here's a quick walkthrough of one way to use godeps to add or update a Mindfly dependency into Godeps/_workspace. For more details, please see the instructions in [godep's documentation](https://github.com/tools/godep).



```
cd $GOPATH/src/github.com/cheyang/numa-utils
git pull https://github.com/cheyang/numa-utils.git
godep restore -v
# To add a new dependency, do:
go get path/to/dependency
godep save ./...
```

## Build

```sh
cd $GOPATH/src/github.com/cheyang/numa-utils
make
```

## Make sure `libnuma.so.1` is in you LD_LIBRARY_PATH or default lib path like `/usr/lib`

you can download `ftp://oss.sgi.com/www/projects/libnuma/download/numactl-2.0.11.tar.gz`


## Run

```sh
cd $GOPATH/src/github.com/cheyang/numa-utils
./gonumactl
available: 2 nodes
node 0 cpus: [0 1 2 3 8 9 10 11]
node 0 size: 147446 MB
node 0 free: 9293 MB
node 1 cpus: [4 5 6 7 12 13 14 15]
node 1 size: 147456 MB
node 1 free: 3470 MB
node distances:
node   0   1
  0:  10  21
  1:  21  10
```



