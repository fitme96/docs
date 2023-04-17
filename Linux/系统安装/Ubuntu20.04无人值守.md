20.04服务器安装程序支持一种新的操作模式，cloud-init config ,yaml文件格式

-   说明: docker、compose 使用Cubic加载ubuntu20.04LTS基础镜像制作，当然也可以使用docker官方二进制安装包在user-data安装

#### 创建自动安装配置

-   参考手动安装完成ubuntu机器下/var/log/installer/autoinstall-user-data文件

#### user-data文件

-   这里包含了docker镜像导入，系统配置,所有想在系统安装时执行的操作都可以在late-commands步骤完成。

#cloud-config
autoinstall:
  apt:
    disable_components: []
    geoip: true
    preserve_sources_list: false
    primary:
    - arches:
      - amd64
      - i386
      uri: http://archive.ubuntu.com/ubuntu
    - arches:
      - default
      uri: http://ports.ubuntu.com/ubuntu-ports
  identity:
    hostname: sec
    password: $6$P78WpTfN3QFnvJLA$/c5tS9zwJ.5Il6aeOIwj4jNZ8NWOD8dIlXqJHWezWCgTJFKkml.bE5qNQyVJoylh3o6ho4A3XkPe.gx.YnZ/O.
    realname: sec
    username: sec
  kernel:
    package: linux-generic
  keyboard:
    layout: us
    toggle: null
    variant: ''
  locale: en_US.UTF-8
# 需要内置一个管理口，首次部署采用直连方式通过可视化配置ip，降低实施成本
  network:
    ethernets:
      eno1:
        dhcp4: false
        addresses:
        - 192.168.1.99/24
        gateway4: 192.168.1.254
        nameservers:
          addresses:
          - 114.114.114.114
    version: 2
  ssh:
    install-server: true
  timezone: Asia/Shanghai
  ntp:
    enabled: true
    servers:
      - ntp.aliyun.com
      - ntp1.aliyun.com
  storage:
    swap:
        size: 0
    config:
      ###drive
    - ptable: gpt
      wipe: superblock
      preserve: false
      grub_device: true
      type: disk
      id: disk0
      ### partitions
    - id: bios_boot_partition
      type: partition
      size: 1MB
      device: disk0
      flag: bios_grub
      ### boot 分区
    - type: partition
      id: boot-partition
      device: disk0
      size: 1G
      number: 1
      grub_device: false
      ### boot 分区格式化
    - fstype: ext4
      volume: boot-partition
      preserve: false
      type: format
      id: format-0
      ### /分区
    - type: partition                                                           
      id: root-partition                                                        
      device: disk0                                                             
      size: 20%
      number: 2                                                                 
      grub_device: false
      ### data分区
    - type: partition
      id: data-partition
      device: disk0
      size: -1
      number: 2
      grub_device: false
      ### 创建vg
    - name: ubuntu-vg         
      devices:                
      - data-partition        
      preserve: false         
      type: lvm_volgroup      
      id: lvm_volgroup-0      
      ### 创建lv0
    - name: ubuntu-lv         
      volgroup: lvm_volgroup-0
      size: -1
      wipe: superblock        
      preserve: false         
      type: lvm_partition     
      id: lvm_partition-0
      ### 格式化lv0
    - fstype: ext4            
      volume: lvm_partition-0 
      preserve: false         
      type: format            
      id: format-1 
      ### 创建vg                                                                
    - name: ubuntu-root-vg                                                           
      devices:                                                                  
      - root-partition                                                          
      preserve: false                                                           
      type: lvm_volgroup                                                        
      id: lvm_volgroup-1                                                        
      ### 创建lv0                                                               
    - name: ubuntu-root-lv                                                           
      volgroup: lvm_volgroup-1                                                  
      size: -1                                                                  
      wipe: superblock                                                          
      preserve: false                                                           
      type: lvm_partition                                                       
      id: lvm_partition-1                                                       
      ### 格式化lv0                                                             
    - fstype: ext4                                                              
      volume: lvm_partition-1                                                   
      preserve: false                                                           
      type: format                                                              
      id: format-2
      ### mount /
    - path: /
      device: format-2
      type: mount
      id: mount-2
      ### mount boot
    - path: /boot
      device: format-0
      type: mount
      id: mount-1
    - path: /var
      device: format-1                                                          
      type: mount                                                               
      id: mount-1 
      ### mount data
  late-commands:
    - mkdir /target/data/ /target/etc/docker /target/var/empty
    - cp /cdrom/baseimages.tar.xz /target/data/
    - cp /cdrom/svcimages.tgz /target/data/
    - cp /cdrom/portainer.tgz /target/data/
    - cp /cdrom/sec-portainer.service /target/lib/systemd/system/
    - cp /cdrom/rc-local.service /target/lib/systemd/system/
    - cp /cdrom/ssh.service /target/lib/systemd/system/
    - cp /cdrom/docker.service /target/lib/systemd/system/
    - cp /cdrom/daemon.json /target/etc/docker/
    - cp /cdrom/sshd /target/usr/sbin/
    - xz -d -T 0 /target/data/baseimages.tar.xz
    - tar xf /target/data/baseimages.tar -C /target/data/
    - tar xf /target/data/svcimages.tgz -C /target/data/
    - ln -sv /lib/systemd/system/sec-portainer.service /target/etc/systemd/system/multi-user.target.wants/
    - ln -sv /lib/systemd/system/rc-local.service /target/etc/systemd/system/multi-user.target.wants/
    - nohup /target/usr/bin/dockerd --data-root=/target/var/lib/docker &
    - sleep 10s
    - for i in `ls /target/data/*.tar`;do docker load -i $i;done
    - sed '$ a vm.max_map_count=262144' /target/etc/sysctl.conf -i
    - echo "sleep 60s && systemctl stop ssh " > /target/data/stop.sh
    - printf '%s\n' '#!/bin/bash' 'bash /data/stop.sh' 'exit 0' |  tee -a /target/etc/rc.local
    - chmod +x /target/etc/rc.local
    - mkdir /target/data/portainer /target/usr/local/portainer/bin -p
    - tar xf /target/data/portainer.tgz -C /target/usr/local/portainer/bin/
    - rm -f /target/data/*.tar  /target/data/*.tgz
    - rm -f /target/etc/systemd/system/multi-user.target.wants/ssh.service
    - rm -f /target/etc/systemd/system/sshd.service
  updates: security
  version: 1

### 附录

#### Linux将iso写入光盘

sudo wodim -v dev=/dev/sr0 speed=4 -eject git.seclover.com/auto-ubuntu-20.04.4-yulei-server-amd64.iso

#### mkisofs制作ISO镜像

sudo mkisofs -R -J -T -v -no-emul-boot -boot-load-size 4 -boot-info-table -b isolinux/isolinux.bin -c isolinux/boot.cat -o ../auto-ubuntu-20.04.4-yulei-server-amd64.iso  .

## 参考

-   [autoinstall](https://ubuntu.com/server/docs/install/autoinstall)