```bash

wget https://download.libvirt.org/libvirt-9.1.0.tar.xz

tar xf libvirt-9.1.0.tar.xz
cd libvirt-9.1.0/
apt install python3-pip -y
pip install meson docutils ninja
apt install -y build-essential libxml2-utils xsltproc libxml2-dev libglib2.0-dev libgnutls28-dev libpciaccess-dev libtirpc-dev libssh-dev libyajl-dev  qemu-utils
meson setup build -Dlibssh=enabled -Ddriver_remote=enabled -Ddriver_qemu=enabled

ninja -C build install

virsh list

apt install -y qemu-kvm

useradd libvirt-qemu

systemctl start libvirtd.socket



```