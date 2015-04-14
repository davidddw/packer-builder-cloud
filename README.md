this is packer build for kvm and xenserver

xe sr-create name-label=ISOSR type=iso device-config:location=/iso_sr \
    device-config:legacy_mode=true content-type=iso
