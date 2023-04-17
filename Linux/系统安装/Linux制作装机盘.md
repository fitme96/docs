syslinux包中isohybird可以为iso添加MBR/EFI

sudo isohybrid CLS-YLW-server-V4.1-dev-1026_amd64.iso 默认MBR

EFI

sudo isohybird --uefi CLS-YLW-server-V4.1-dev-1026_amd64.iso

sudo dd if=CLS-YLW-server-V4.1-dev-1026_amd64.iso of=/dev/sda bs=4M status=progress