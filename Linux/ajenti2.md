FROM ubuntu:bionic
ARG HTTP_RPOXY
ARG HTTPS_PROXY
RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    export http_proxy=$HTTP_RPOXY && \
    export https_proxy=$HTTPS_PROXY && \
    apt update && \
    DEBIAN_FRONTEND="noninteractive" apt-get install -y build-essential python3-pip python3-dev python3-lxml libssl-dev python3-dbus python3-augeas python3-apt ntpdate python3-testresources tzdata netplan.io && \
    pip3 install setuptools pip wheel -U && \
    pip3 install ajenti-panel ajenti.plugin.ace ajenti.plugin.augeas ajenti.plugin.auth-users ajenti.plugin.core ajenti.plugin.dashboard ajenti.plugin.datetime ajenti.plugin.filemanager ajenti.plugin.filesystem ajenti.plugin.network ajenti.plugin.notepad ajenti.plugin.packages ajenti.plugin.passwd ajenti.plugin.plugins ajenti.plugin.power ajenti.plugin.services ajenti.plugin.settings ajenti.plugin.terminal && \
    echo "root:Sec1024" | chpasswd
CMD ["ajenti-panel"]

#### build

docker build --build-arg HTTP_RPOXY=http://192.168.60.68:8001 --build-arg HTTPS_PROXY=http://192.168.60.68:8001 -t hub.bugfeel.net:8443/ck-test/ajenti2:latest . 

#### run

 docker run -d --network host -v /etc/netplan:/etc/netplan  -v /root/images/ajenti/config.yml:/etc/ajenti/config.yml -v /root/images/ajenti/users.yml:/etc/ajenti/users.yml hub.bugfeel.net:8443/ck-test/ajenti2:latest