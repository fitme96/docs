
# Windows

#### 安装OpenSSH
- [安装ssh](https://learn.microsoft.com/zh-cn/windows-server/administration/openssh/openssh_install_firstuse)
```shell

# 设置ssh连接后默认shell为powershell
New-ItemProperty -Path "HKLM:\SOFTWARE\OpenSSH" -Name DefaultShell -Value "C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe" -PropertyType String -Force
```



# Linux


# MacOS

# 其他

### KVM-NAT网络模式
- 默认子网192.168.122.0/24，可通过virsh net-edit --network defalut 修改子网，需要重启网络，DHCP是由dnsmasq实现
- 关于DHCP模式下固定虚拟机IP地址，可通过绑定MAC地址与IP地址实现
```xml
<ip address='192.168.123.1' netmask='255.255.255.0'> 
  <dhcp> 
    <range start='192.168.123.2' end='192.168.123.254'/> 
    <host mac='52:54:00:aa:55:39' name='win2k19' ip='192.168.123.200'/> 
  </dhcp> 
</ip>

```

# 问题记录

### win2k19安装ssh失败
```powershell
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0 + CategoryInfo : NotSpecified: (:) [Add-WindowsCapability], COMException + FullyQualifiedErrorId : Microsoft.Dism.Commands.AddWindowsCapabilityCommand
```
[打补丁](https://github.com/MicrosoftDocs/windowsserverdocs/issues/2074)