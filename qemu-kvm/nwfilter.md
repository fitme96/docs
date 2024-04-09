
通过定义network filter XML实现网络过滤器，filter被修改时，所有引用此filter的虚拟机都会自动更新规则

filter 支持网络类型：
- network(NAT)
- bridge

KVM 自带许多filter

virsh nwfilter-list 查看 