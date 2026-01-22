
# what done

ipfs-cid print ipfs cid info string to stdout by read any file you give it

# how to use

## use by build

```
go install github.com/lrobot/ipfs-cid@latest
export PATH=$PATH:~/go/bin
ipfs-cid <your_file_name>
cat <your_file_name> | ipfs-cid -stdin
```

# how to build 

```
go build .
```

