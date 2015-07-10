this is packer build for kvm and xenserver

xe sr-create name-label=ISOSR type=iso device-config:location=/iso_sr \
    device-config:legacy_mode=true content-type=iso


yum install golang golang-pkg-linux-amd64 golang-pkg-linux-386 golang-pkg-linux-arm git \
golang-pkg-windows-amd64 golang-pkg-windows-386 golang-pkg-darwin-amd64 golang-pkg-darwin-386