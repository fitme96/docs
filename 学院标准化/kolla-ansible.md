

```shell
export https_proxy=http://192.168.60.63:7890
yum install epel-release  python3-pip  ansible -y
#pip3.9 install -U pip

curl -sSL https://github.com/openstack/kolla-ansible/archive/refs/tags/queens-eol.tar.gz -o queens-eol.tar.gz

gzip -d queens-eol.tar.gz 
cd kolla-ansible-queens-eol/
pip3.9 install -r requirements.txt 
git init
python3.9 setup.py install


cp -r /usr/local/python3.7/share/kolla-ansible/etc_examples/kolla/* /etc/kolla/

```


```shell

    1  https://github.com/openstack/kolla-ansible/archive/refs/tags/9.3.2.tar.gz
    2  export https_proxy=http://192.168.60.63:7890
    3  curl https://github.com/openstack/kolla-ansible/archive/refs/tags/9.3.2.tar.gz -o 9.3.2.tar.gz
    4  ls
    5  curl https://github.com/openstack/kolla-ansible/archive/refs/tags/9.3.2.tar.gz -o 9.3.2.tar.gz
    6  ls
    7  curl -sSL https://github.com/openstack/kolla-ansible/archive/refs/tags/9.3.2.tar.gz -o 9.3.2.tar.gz
    8  ls
    9  gzip -d 9.3.2.tar.gz 
   10  ls
   11  tar xf 9.3.2.tar 
   12  ls
   13  cd kolla-ansible-9.3.2/
   14  ls
   15  pip install -r requirements.txt 
   16  yum install epel-release  python-pip  python-devel libffi-devel gcc openssl-devel libselinux-python ansible -y
   17  unset https
   18  unset https_proxy
   19  yum install epel-release  python-pip  python-devel libffi-devel gcc openssl-devel libselinux-python ansible -y
   20  export https_proxy=http://192.168.60.63:7890
   21  yum install -y python-pip
   22  yum list|grep pip
   23  yum install -y python3-pip
   24  unset https_proxy
   25  yum install -y python3-pip
   26  ls
   27  export https_proxy=http://192.168.60.63:7890
   28  pip install -r requirements.txt 
   29  pip3 install -r requirements.txt 
   30  pip3 install pip
   31  pip3 install -U pip
   32  pip install -r requirements.txt 
   33  python3 setup.py install
   34  git init
   35  yum install -y git
   36  unset https_proxy
   37  yum install -y git
   38  export https_proxy=http://192.168.60.63:7890
   39  git init
   40  python3 setup.py install
   41  pip install ansible
   42  host_key_checking=False
   43  pipelining=True
   44  cp -r /usr/local/share/kolla-ansible/etc_examples/kolla/* /etc/kolla
   45  sudo mkdir -p /etc/kolla
   46  sudo chown $USER:$USER /etc/kolla
   47  cp -r /usr/local/share/kolla-ansible/etc_examples/kolla/* /etc/kolla
   48  cd
   49  cp -r /usr/local/share/kolla-ansible/ansible/inventory/* . 
   50  kolla-ansible install-deps
   51  pip uninstall ansbile
   52  pip uninstall ansible
   53  kolla-ansible install-deps
   54  ansible -v
   55  ansible -V
   56  ansible --version
   57  kolla-ansible
   58  curl https://github.com/openstack/kolla-ansible/archive/refs/tags/12.1.0.tar.gz -o 12.1.0.tar.gz
   59  ls
   60  curl https://github.com/openstack/kolla-ansible/archive/refs/tags/12.1.0.tar.gz -o 12.1.0.tar.gz
   61  ls
   62  export https_proxy=http://192.168.60.63:7890
   63  ls
   64  curl https://github.com/openstack/kolla-ansible/archive/refs/tags/12.1.0.tar.gz -o 12.1.0.tar.gz
   65  curl -L https://github.com/openstack/kolla-ansible/archive/refs/tags/12.1.0.tar.gz -o 12.1.0.tar.gz
   66  ls
   67  gzip -d 12.1.0.tar.gz 
   68  tar xf 12.1.0.tar 
   69  cd kolla-ansible-12.1.0/
   70  ls
   71  pip install -r requirements.txt 
   72  python3 setup.py install
   73  git init
   74  python3 setup.py install
   75  kolla-ansible install-deps
   76  cd
   77  ls
   78  curl -L https://github.com/openstack/kolla-ansible/archive/refs/tags/queens-eol.tar.gz -o queens-eol.tar.gz
   79  gzip -d queens-eol.tar.gz 
   80  tar queens-eol.tar 
   81  tar xf queens-eol.tar 
   82  cd kolla-ansible-queens-eol/
   83  ls
   84  pip install -r requirements.txt 
   85  python3 setup.py install
   86  git innt
   87  git inint
   88  git init
   89  python3 setup.py install
   90  kolla-ansible install-deps
   91  pip install 'ansible>=2.9'
   92  pip install 'ansible>=2.9<2.10'
   93* pip install 
   94  pip install 'ansible==2.9'
   95  kolla-ansible install-deps
   96  ansible -h
   97  ansible --version
   98  pip install 'ansible==2.8'
   99  /usr/local/lib/python3.6/site-packages/ansible --version
  100  ansible --version
  101  python3 ansible --version
  102  ls
  103  ansible
  104  cd
  105  ls
  106  curl -L https://github.com/openstack/kolla-ansible/archive/refs/tags/10.3.0.tar.gz -o 10.3.0.tar.gz
  107  ls
  108  gzip -d 10.3.0.tar.gz 
  109  tar xf 10.3.0.tar 
  110  cd kolla-ansible-10.3.0/
  111  ls
  112  pip install -r requirements.txt 
  113  python3 setup.py install
  114  git init
  115  python3 setup.py install
  116  kolla-ansible
  117  kolla-ansiblekolla-ansible install-deps
  118  kolla-ansible install-deps
  119  cd
  120  vim /etc/kolla/globals.yml 
  121  kolla-genpwd
  122  kolla-ansible -i all-in-one bootstrap-servers
  123  pip install 'ansible==2.9'
  124  kolla-ansible -i all-in-one bootstrap-servers
  125  pip install AnsibleCollectionLoader
  126  history 

```