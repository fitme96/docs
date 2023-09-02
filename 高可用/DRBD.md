---
title: DRBD
type: 高可用
 ---




## BRBD安装

```bash
apt install -y drbd-utils
modprobe drbd
cp global_common.conf{,.bak}

```
### 如下需要在两台机器操作
- 编辑global_common.conf
```
# DRBD is the result of over a decade of development by LINBIT.
# In case you need professional services for DRBD or have
# feature requests visit http://www.linbit.com

global {
	usage-count no;

	# Decide what kind of udev symlinks you want for "implicit" volumes
	# (those without explicit volume <vnr> {} block, implied vnr=0):
	# /dev/drbd/by-resource/<resource>/<vnr>   (explicit volumes)
	# /dev/drbd/by-resource/<resource>         (default for implict)
	udev-always-use-vnr; # treat implicit the same as explicit volumes

	# minor-count dialog-refresh disable-ip-verification
	# cmd-timeout-short 5; cmd-timeout-medium 121; cmd-timeout-long 600;
}

common {
        protocol C;
	handlers {
		# These are EXAMPLE handlers only.
		# They may have severe implications,
		# like hard resetting the node under certain circumstances.
		# Be careful when choosing your poison.

		# pri-on-incon-degr "/usr/lib/drbd/notify-pri-on-incon-degr.sh; /usr/lib/drbd/notify-emergency-reboot.sh; echo b > /proc/sysrq-trigger ; reboot -f";
		# pri-lost-after-sb "/usr/lib/drbd/notify-pri-lost-after-sb.sh; /usr/lib/drbd/notify-emergency-reboot.sh; echo b > /proc/sysrq-trigger ; reboot -f";
		# local-io-error "/usr/lib/drbd/notify-io-error.sh; /usr/lib/drbd/notify-emergency-shutdown.sh; echo o > /proc/sysrq-trigger ; halt -f";
		# fence-peer "/usr/lib/drbd/crm-fence-peer.sh";
		# split-brain "/usr/lib/drbd/notify-split-brain.sh root";
		# out-of-sync "/usr/lib/drbd/notify-out-of-sync.sh root";
		# before-resync-target "/usr/lib/drbd/snapshot-resync-target-lvm.sh -p 15 -- -c 16k";
		# after-resync-target /usr/lib/drbd/unsnapshot-resync-target-lvm.sh;
		# quorum-lost "/usr/lib/drbd/notify-quorum-lost.sh root";
	}

	startup {
		# wfc-timeout degr-wfc-timeout outdated-wfc-timeout wait-after-sb
	        wfc-timeout          240;
       		degr-wfc-timeout     240;
        	outdated-wfc-timeout 240;
	}

	options {
		# cpu-mask on-no-data-accessible

		# RECOMMENDED for three or more storage nodes with DRBD 9:
		# quorum majority;
		# on-no-quorum suspend-io | io-error;
	}

	disk {
		on-io-error detach;
		# size on-io-error fencing disk-barrier disk-flushes
		# disk-drain md-flushes resync-rate resync-after al-extents
                # c-plan-ahead c-delay-target c-fill-target c-max-rate
                # c-min-rate disk-timeout
	}

	net {
		# protocol timeout max-epoch-size max-buffers
		# connect-int ping-int sndbuf-size rcvbuf-size ko-count
		# allow-two-primaries cram-hmac-alg shared-secret after-sb-0pri
		# after-sb-1pri after-sb-2pri always-asbp rr-conflict
		# ping-timeout data-integrity-alg tcp-cork on-congestion
		# congestion-fill congestion-extents csums-alg verify-alg
		# use-rle
	        cram-hmac-alg md5; ##设置加密算法md5
        	shared-secret "testdrbd"; ##加密密码
	}
}


```

- 新建/etc/drbd.d/r0.res

```
## sec主机名
resource r0 {
on sec {
  device     /dev/drbd0;
  disk       /dev/mapper/ubuntu--vg-ubuntu--lv;
  address    192.168.180.22:9527;
  meta-disk  internal;
 }
on sec2 {
  device     /dev/drbd0;
  disk       /dev/mapper/ubuntu--vg-ubuntu--lv;
  address    192.168.180.23:9527;
  meta-disk  internal;
 }
}


```
```bash
mknod /dev/drbd0 b 147 0
drbdadm create-md r0
systemctl start drbd && systemctl enable drbd
```
### 以下在主节点操作
```bash
# 强制为主
drbdsetup /dev/drbd0 primary --force  
mkfs.ext4 /dev/drbd0 
mount /dev/drbd0  /var/lib/docker/volumes/

### 增加数据
umount /var/lib/docker/volumes
drbdsetup /dev/drbd0 secondary

```

### 故障切换如下在从节点

```bash
drbdsetup /dev/drbd0 primary
docker compose -f host.yaml up -d


```

















