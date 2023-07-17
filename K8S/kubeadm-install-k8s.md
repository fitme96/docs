## 安装containerd

-   注意:Containerd与kubelet CgroupsDriver需要一致)

# 下载containerd containerd.service runc cni-plugins 
wget https://github.com/containerd/containerd/releases/download/v1.5.11/containerd-1.5.11-linux-amd64.tar.gz
wget https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz
wget https://github.com/opencontainers/runc/releases/download/v1.1.4/runc.amd64
wget https://raw.githubusercontent.com/containerd/containerd/main/containerd.service

# containerd
tar Cxzvf /usr/local/ containerd-1.5.11-linux-amd64.tar.gz
mv containerd.service  /etc/systemd/system/
systemctl daemon-reload 
systemctl enable --now containerd
mkdir /etc/containerd/
# 生成默认配置文件
containerd config default > /etc/containerd/config.toml

# runc
install -m 755 runc.amd64 /usr/local/sbin/runc
# cni-plugins
mkdir /opt/cni/bin -p
tar Cxzvf /opt/cni/bin/ cni-plugins-linux-amd64-v1.1.1.tgz

## 安装kubectl kubeadm kubelet

# 安装kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
#  安装crictl
DOWNLOAD_DIR="/usr/local/bin"
mkdir -p "$DOWNLOAD_DIR"
CRICTL_VERSION="v1.25.0"
ARCH="amd64"
curl -L "https://github.com/kubernetes-sigs/cri-tools/releases/download/${CRICTL_VERSION}/crictl-${CRICTL_VERSION}-linux-${ARCH}.tar.gz" | sudo tar -C $DOWNLOAD_DIR -xz

cat << EOF > /etc/crictl.yaml
runtime-endpoint: unix:///run/containerd/containerd.sock
image-endpoint: unix:///run/containerd/containerd.sock
timeout: 10
debug: true
EOF


## 下载安装kubeadm kubelet
RELEASE="$(curl -sSL https://dl.k8s.io/release/stable.txt)"
ARCH="amd64"
cd $DOWNLOAD_DIR
sudo curl -L --remote-name-all https://dl.k8s.io/release/${RELEASE}/bin/linux/${ARCH}/{kubeadm,kubelet}
sudo chmod +x {kubeadm,kubelet}
## 下载kubelet.service
RELEASE_VERSION="v0.4.0"
curl -sSL "https://raw.githubusercontent.com/kubernetes/release/${RELEASE_VERSION}/cmd/kubepkg/templates/latest/deb/kubelet/lib/systemd/system/kubelet.service" | sed "s:/usr/bin:${DOWNLOAD_DIR}:g" | sudo tee /etc/systemd/system/kubelet.service
sudo mkdir -p /etc/systemd/system/kubelet.service.d
## kubeadm.conf
curl -sSL "https://raw.githubusercontent.com/kubernetes/release/${RELEASE_VERSION}/cmd/kubepkg/templates/latest/deb/kubeadm/10-kubeadm.conf" | sed "s:/usr/bin:${DOWNLOAD_DIR}:g" | sudo tee /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
systemctl enable --now kubelet

crictl completion bash > /etc/bash_completion.d/crictl
kubectl completion bash > /etc/bash_completion.d/kubectl
kubeadm completion bash > /etc/bash_completion.d/kubeadm
source /etc/bash_completion.d/crictl 
source /etc/bash_completion.d/kubect
source /etc/bash_completion.d/kubeadm

## kubeadm init

## 
apt install -y socat conntrack

# 激活 br_netfilter 模块
modprobe br_netfilter
cat << EOF > /etc/modules-load.d/k8s.conf
br_netfilter
EOF

# 内核参数设置：开启IP转发，允许iptables对bridge的数据进行处理
cat << EOF > /etc/sysctl.d/k8s.conf 
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

# 立即生效
sysctl --system


# 列出images
kubeadm config images list

## 下载k8s组件镜像(配置国内仓库)
kubeadm config images pull --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers

## 修改containerd config.toml文件中pause镜像
sandbox_image = "registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.8"
## 初始化(创建控制面组件pod）
kubeadm init --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers  --kubernetes-version=v1.25.4 --service-cidr=10.1.0.0/16 --pod-network-cidr=10.244.0.0/16


export KUBECONFIG=/etc/kubernetes/admin.conf

## kubeadm join

# Master生成
kubeadm token create --print-join-command

# node执行（需要安装crictl socat conntrack 加载内核模块
kubeadm join 192.168.190.223:6443 --token 9x67ma.kpeeld3wenrqzwsi --discovery-token-ca-cert-hash sha256:ff8de2cb35695b9d3435bcf8d184b317964fda766c840cb42fb21a3f271518fc

## 安装flannel

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml