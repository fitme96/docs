```
root@kvm:~# systemctl cat iptables.service 
# /lib/systemd/system/iptables.service
[Unit]
Description=netfilter persistent configuration
DefaultDependencies=no
Wants=network-pre.target systemd-modules-load.service local-fs.target
Before=network-pre.target shutdown.target libvirtd.target
After=systemd-modules-load.service local-fs.target
Conflicts=shutdown.target
Documentation=man:netfilter-persistent(8)

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/sbin/netfilter-persistent start
ExecStop=/usr/sbin/netfilter-persistent stop

```
关键字
- "Requires"关键字表示一个服务依赖于另一个服务，如果被依赖的服务没有运行或者失败，那么依赖它的服务也无法运行。
- "Wants"关键字表示一个服务希望依赖于另一个服务，如果被依赖的服务没有运行或者失败，不会影响依赖它的服务的正常运行。
- "Conflicts"关键字表示两个服务之间存在冲突，只能同时运行一个。
- "After"关键字表示一个服务应该在另一个服务之后启动。
- "Before"关键字表示一个服务应该在另一个服务之前启动，与"After"关键字相反。