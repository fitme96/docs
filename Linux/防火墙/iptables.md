### Netfilter/Iptables 信息包过滤系统可以当成一个整体，netfilter是内核的模块实现，iptables是对上层操作工具

#### 表
- raw
- mangle
- filter
- nat

####  链
每个链都有多条规则，且有默认规则，规则从第一条开始匹配，--line-numbers参数可以查看规则号，匹配未中时执行默认规则
- INPUT
- OUTPUT
- OUPUT
- PREROUTING
- POSTROUTING
#### 规则


