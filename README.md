# land

<sup><sup>*(5 hours and 165 loc to get the mvp!)*</sup></sup>

docker image -> firecracker vm!

## go

```bash
# build the vm
./land <docker image[:tag]>
# create your config file
./configure_vm.py # or just copy + edit vm_config.example.json
# launch the vm!
./boot [my_vm_config.json] # ./vm_config.json is the default
# TODO: networkng goes here
```

### deps

- golang ;-;
  - required for building tar2ext4
- git
  - hopefully obvious, but it's for submodules
- jq
  - used for parsing docker manifests
- e2fsprogs
  - used for building the final rootfs image
- firecracker
  - run the vm! :D

## how does it work?

- extract layers + metadata from docker
- rebuild layers into one rootfs
- convert the rootfs tarball into an ext4 disk image
- mount the rootfs image as the vm rootfs
- :tada:
