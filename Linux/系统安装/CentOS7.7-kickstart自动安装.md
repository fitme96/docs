kickstart文件
```yaml
## cat ks.cfg

#version=DEVEL
# System authorization information
auth --enableshadow --passalgo=sha512
# Use CDROM installation media
cdrom
# Use text install
text
### SELinux configuration ###
selinux --disabled
### Firewall configuration ###
firewall --disabled
### Reboot after installation ###
reboot --eject

# Run the Setup Agent on first boot
firstboot --enable
ignoredisk --only-use=sda
# Keyboard layouts
keyboard --vckeymap=us --xlayouts='us'
# System language
lang en_US.UTF-8

# Network information
network  --bootproto=static --device=eth0 --onboot=on --ip=192.168.1.98 --netmask=255.255.255.0 --ipv6=auto --activate
network  --hostname=seclover.com

# Partition clearing information
clearpart --all --initlabel
# Root password
rootpw --iscrypted $6$.FCgbdSV4UFcxm3u$JhP3TnM43W2kDTp36XUApp0r6he/3.wJge8c8QiEiXtHpoPNN5Imn.wOaatD3dYP5Zsf6ZN7wVA71y.9IR1IL1
# System services
services --disabled="chronyd"
# System timezone
timezone Asia/Shanghai --isUtc --nontp
# System bootloader configuration
bootloader --location=mbr --boot-drive=sda
# Disk partitioning information
#autopart --type=lvm
part pv.315 --fstype="lvmpv" --ondisk=sda --size=101375 --grow
part /boot --fstype="xfs" --ondisk=sda --size=1024
part /boot/efi --fstype="efi" --size=200 --asprimary
volgroup centos --pesize=4096 pv.315
logvol swap  --fstype="swap" --size=4096 --name=swap --vgname=centos
logvol /  --fstype="xfs" --grow --size=2048 --name=root --vgname=centos


%packages
@^minimal
@core

%end

%addon com_redhat_kdump --disable --reserve-mb='auto'

%end
### 对于sda外的其他磁盘做LVM 挂载至/data 用于数据盘,由于不同设备名称可能不同，所以在post阶段通过脚本完成LVM创建,grep -v 排除sdb是因为U盘启动盘也会作为一个磁盘，这里没有做判断，如果必要通过磁盘大小排除即可，这里偷懒了。
%post --interpreter=/bin/bash
num_disks=`lsblk -d -n |grep -Ev "loop|sr0|sda|sdb"|wc -l`  # 获取硬盘数量
if [ $num_disks -gt 0 ]; then
        lvm_disks=`lsblk -d -n |grep -Ev "loop|sr0|sda|sdb"|awk '{printf("/dev/%s ", $1); } END { printf("\n"); }'`
	pvcreate $lvm_disks
        vgcreate vg01 $lvm_disks
        lvcreate --name lv_data -l 100%FREE vg01
        mkfs.xfs /dev/vg01/lv_data
	echo "/dev/vg01/lv_data /data xfs defaults	0 0" >> /etc/fstab
	mkdir /data
fi

%end
%post --nochroot --log=/mnt/sysimage/root/ks-post.log
cp -a /run/install/repo/test /mnt/sysimage/root/
%end


```
EFI grub.cfg
```yaml
set default="0"

function load_video {
  insmod efi_gop
  insmod efi_uga
  insmod video_bochs
  insmod video_cirrus
  insmod all_video
}

load_video
set gfxpayload=keep
insmod gzio
insmod part_gpt
insmod ext2

set timeout=10
### END /etc/grub.d/00_header ###

search --no-floppy --set=root -l 'CentOS 7 x86_64'

### BEGIN /etc/grub.d/10_linux ###
menuentry 'Install CentOS 7 usb' --class fedora --class gnu-linux --class gnu --class os {
	linuxefi /images/pxeboot/vmlinuz inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 inst.ks=hd:LABEL=CentOS\x207\x20x86_64:/ks.cfg quiet
	initrdefi /images/pxeboot/initrd.img
}
menuentry 'Install CentOS 7 cdrom' --class fedora --class gnu-linux --class gnu --class os {
        linuxefi /images/pxeboot/vmlinuz inst.ks=cdrom:/ks.cfg quiet
        initrdefi /images/pxeboot/initrd.img
}

menuentry 'Test this media & install CentOS 7' --class fedora --class gnu-linux --class gnu --class os {
	linuxefi /images/pxeboot/vmlinuz inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 rd.live.check quiet
	initrdefi /images/pxeboot/initrd.img
}
submenu 'Troubleshooting -->' {
	menuentry 'Install CentOS 7 in basic graphics mode' --class fedora --class gnu-linux --class gnu --class os {
		linuxefi /images/pxeboot/vmlinuz inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 xdriver=vesa nomodeset quiet
		initrdefi /images/pxeboot/initrd.img
	}
	menuentry 'Rescue a CentOS system' --class fedora --class gnu-linux --class gnu --class os {
		linuxefi /images/pxeboot/vmlinuz inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 rescue quiet
		initrdefi /images/pxeboot/initrd.img
	}
}

```



制作 UEFI iso及U盘启动盘
```shell
## 复制官方ISO文件到制作自定义ISO工作目录，注意cp源目录加. 复制隐藏文件
mkdir centos make_point
mount -o loop CentOS-7-x86_64-Minimal-1908.iso centos/
cp -v -r centos/. make_point/
## 复制文档开头制作的ks.cfg文件
cp ks.cfg make_point
## 修改EFI grub.cfg ，详细参考文档开头grub.cfg文件
### 注意:以下两处grub.cfg都要更改
cd make_point
mount images/efiboot.img /mnt/
cp grub.cfg /mnt/EFI/BOOT/
umount /mnt
## 同时修改EFI/BOOT 目录grub.cfg 
cp grub.cfg EFI/BOOT/
## 生成iso
mkisofs -o auto.iso -b isolinux/isolinux.bin -J -R -l -c isolinux/boot.cat -no-emul-boot -boot-load-size 4 -boot-info-table -eltorito-alt-boot -e images/efiboot.img -no-emul-boot -graft-points -V "CentOS 7 x86_64" .
## 增加EFI分区
isohybrid --uefi auto.iso
## 刻录到U盘
sudo dd if=auto.iso of=/dev/sda conv=fsync
```


[参考kickstart](https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/7/html/installation_guide/chap-kickstart-installations)