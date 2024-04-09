要在 Linux 中打开 qcow2 磁盘镜像并修改其中的文件，您可以按照以下步骤进行操作：

1. 安装 qemu-utils 工具，这个工具包含了 qemu-img 工具，可用于管理虚拟磁盘镜像。您可以在终端中运行以下命令来安装它：

   
   sudo apt install qemu-utils
   

2. 挂载 qcow2 磁盘镜像文件。在终端中运行以下命令：

   
   sudo modprobe nbd max_part=16
   sudo qemu-nbd -c /dev/nbd0 /path/to/your/qcow2/image
   sudo fdisk -l /dev/nbd0
   

   第一个命令用于加载 nbd 内核模块；第二个命令用于将 qcow2 磁盘镜像挂载到 /dev/nbd0 上；第三个命令用于列出已挂载的磁盘分区。

3. 挂载 qcow2 磁盘镜像中所需的分区。在终端中运行以下命令：

   
   sudo mount /dev/nbd0pX /mnt
   

   其中，X 为挂载分区的编号，/mnt 是挂载的挂载点。

4. 进入挂载的目录 /mnt，修改其中的文件。

5. 关闭挂载的文件系统。在终端中运行以下命令：

   
   sudo umount /mnt
   

6. 关闭 qcow2 磁盘镜像。在终端中运行以下命令：

   
   sudo qemu-nbd -d /dev/nbd0
   

   这将卸载映像并释放与之关联的资源。
