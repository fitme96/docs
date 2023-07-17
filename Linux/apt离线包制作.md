
```shell
## 制作
mkdir local && cd local
apt-get download $(apt-cache depends --recurse --no-recommends --no-suggests --no-conflicts --no-breaks --no-replaces --no-enhances libvirt-clients gnutls-bin qemu-kvm libvirt-daemon-system dnsmasq | grep "^\w" | sort -u)

dpkg-scanpackages . /dev/null | gzip -9c > Packages.gz

cd ..
tar zcf offline-packages.tgz local/
## 使用

tar xf offline-packages.tgz -C /var/lib/
echo "deb [trusted=yes] file:///var/lib/local /" > etc/apt/sources.list.d/local.list
```