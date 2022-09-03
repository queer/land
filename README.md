# land

docker image -> firecracker vm!

## go

```bash
# build the vm
./land <docker image[:tag]>
# launch the vm!
./boot
# TODO: networkng goes here
```

### deps

- golang ;-;
- git
- jq
- e2fsprogs

## how does it work?

- extract layers + metadata from docker
- rebuild layers into one rootfs
- convert the rootfs tarball into an ext4 disk image
- mount the rootfs image as the vm rootfs
- :tada: