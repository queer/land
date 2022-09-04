#!/usr/bin/env python3

import json
import sys


def spacer(times=2):
    for _ in range(times):
        print()


def boolify(s):
    if s in ["True", "true", "1", "y", "t"]:
        return True
    elif s in ["False", "false", "0", "n", "f"]:
        return False
    else:
        return False


def die(check, msg):
    if check:
        print(msg)
        sys.exit(1)


print("hello! let's configure your new firecracker vm together :D")
spacer(3)

print("first, we need to know how many vcpus you want give the vm")
print("this is just a number (1, 7, 4, etc)")
vcpus = input("vcpus = ")
die(not vcpus.isdigit(), "vcpus must be a number")
spacer()

print("next, we need to know how much memory you want to give the vm")
print("this is the number of mib you want to assign (256, 512, 2048, etc)")
memory = input("memory = ")
die(not memory.isdigit(), "memory must be a number")
spacer()

print("do you want to enable smt? (y/n) (default: n)")
smt = boolify(input("smt = "))
spacer()

print("where is your kernel located? (default: ./kernel.bin)")
kernel = input("kernel = ")
if kernel == "":
    print("using default kernel location!")
    kernel = "./kernel.bin"
spacer()

print("where is your rootfs located? (default: ./rootfs.ext4)")
rootfs = input("rootfs = ")
if rootfs == "":
    print("using default rootfs location!")
    rootfs = "./rootfs.ext4"
spacer()

print("do you want to pass a custom kernel cmdline?")
default_cmdline = "console=ttyS0 reboot=k panic=1 pci=off"
print(f"default: {default_cmdline}")
cmdline = input("cmdline = ")
if cmdline == "":
    print("using default kernel cmdline!")
    cmdline = default_cmdline
spacer()

print("good job! :D")

config = {
    "boot-source": {
        "kernel_image_path": kernel,
        "boot_args": cmdline,
        "initrd_path": None,
    },
    "drives": [
        {
            "drive_id": "rootfs",
            "path_on_host": rootfs,
            "is_root_device": True,
            "is_read_only": False,
        }
    ],
    "machine-config": {
        "vcpu_count": int(vcpus),
        "mem_size_mib": int(memory),
        "smt": smt,
        "track_dirty_pages": False,
    },
}

with open("./vm_config.json", "w") as f:
    json.dump(config, f, indent=2)

print("your vm config has been written to ./vm_config.json")
