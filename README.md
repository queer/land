# land

docker image -> firecracker vm!

## go

```bash
git submodule init && git submodule update --recursive
cd hcsshim
go build cmd/tar2ext4
cp tar2ext4 ../
cd ..
./land <docker image>
```