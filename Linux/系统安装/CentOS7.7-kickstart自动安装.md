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
### 复制U盘数据到目标系统，U盘启动盘默认挂载至/run/install/repo下，目标系统/目录为/mnt/sysimage/
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

传统BIOS isolinux/isolinux.cfg文件
```yaml
default vesamenu.c32
timeout 10

display boot.msg

# Clear the screen when exiting the menu, instead of leaving the menu displayed.
# For vesamenu, this means the graphical background is still displayed without
# the menu itself for as long as the screen remains in graphics mode.
menu clear
menu background splash.png
menu title CentOS 7
menu vshift 8
menu rows 18
menu margin 8
#menu hidden
menu helpmsgrow 15
menu tabmsgrow 13

# Border Area
menu color border * #00000000 #00000000 none

# Selected item
menu color sel 0 #ffffffff #00000000 none

# Title bar
menu color title 0 #ff7ba3d0 #00000000 none

# Press [Tab] message
menu color tabmsg 0 #ff3a6496 #00000000 none

# Unselected menu item
menu color unsel 0 #84b8ffff #00000000 none

# Selected hotkey
menu color hotsel 0 #84b8ffff #00000000 none

# Unselected hotkey
menu color hotkey 0 #ffffffff #00000000 none

# Help text
menu color help 0 #ffffffff #00000000 none

# A scrollbar of some type? Not sure.
menu color scrollbar 0 #ffffffff #ff355594 none

# Timeout msg
menu color timeout 0 #ffffffff #00000000 none
menu color timeout_msg 0 #ffffffff #00000000 none

# Command prompt text
menu color cmdmark 0 #84b8ffff #00000000 none
menu color cmdline 0 #ffffffff #00000000 none

# Do not display the actual menu unless the user presses a key. All that is displayed is a timeout message.

menu tabmsg Press Tab for full configuration options on menu items.

menu separator # insert an empty line
menu separator # insert an empty line

label linux
  menu label ^Install CentOS 7 USB
  kernel vmlinuz
  append initrd=initrd.img inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 inst.ks=hd:LABEL=CentOS\x207\x20x86_64:/ks.cfg quiet

label linux
  menu label ^Install CentOS 7 CDROM
  kernel vmlinuz
  append initrd=initrd.img inst.ks=cdrom:/ks.cfg quiet

label check
  menu label Test this ^media & install CentOS 7
  menu default
  kernel vmlinuz
  append initrd=initrd.img inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 rd.live.check quiet

menu separator # insert an empty line

# utilities submenu
menu begin ^Troubleshooting
  menu title Troubleshooting

label vesa
  menu indent count 5
  menu label Install CentOS 7 in ^basic graphics mode
  text help
	Try this option out if you're having trouble installing
	CentOS 7.
  endtext
  kernel vmlinuz
  append initrd=initrd.img inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 xdriver=vesa nomodeset quiet

label rescue
  menu indent count 5
  menu label ^Rescue a CentOS system
  text help
	If the system will not boot, this lets you access files
	and edit config files to try to get it booting again.
  endtext
  kernel vmlinuz
  append initrd=initrd.img inst.stage2=hd:LABEL=CentOS\x207\x20x86_64 rescue quiet

label memtest
  menu label Run a ^memory test
  text help
	If your system is having issues, a problem with your
	system's memory may be the cause. Use this utility to
	see if the memory is working correctly.
  endtext
  kernel memtest

menu separator # insert an empty line

label local
  menu label Boot from ^local drive
  localboot 0xffff

menu separator # insert an empty line
menu separator # insert an empty line

label returntomain
  menu label Return to ^main menu
  menu exit

menu end


```
```


##### 制作 UEFI iso及U盘启动盘
***必须修改以下三个文件
-   `isolinux/isolinux.cfg`
-   `EFI/BOOT/grub.cfg`
-   `images/efiboot.img`
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



[参考官方kickstart](https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/7/html/installation_guide/chap-kickstart-installations)
[关于KVM测试成功，物理机random: crng init done的问题](https://access.redhat.com/discussions/5555121)
[关于同时支持传统BIOS与UEFI](https://unix.stackexchange.com/questions/418974/red-hat-7-4-how-to-inject-kickstart-file-into-usb-media-for-uefi-only-system)
[制作UEFI kickstart ISO](https://unix.stackexchange.com/questions/418974/red-hat-7-4-how-to-inject-kickstart-file-into-usb-media-for-uefi-only-system)