
1. 打开多个文件分屏
```bash
vim -O 1.txt 2.txt 横向打开
vim -o 1.txt 2.txt 纵向打开

切换窗口
ctrl + w, jk控制上下
ctrl + w 按住ctrl 不放 继续按w跳到下一个窗口

```
2. 打开两个文件比对
```bash
vimdiff 1.txt 2.txt
```
3. 粘贴不自动缩进
```shel
:set paste
:set nopaste

```



:h 

dw  删除单词
x     一个字符
dap 删除全部


[vim 实用技巧](https://agou-images.oss-cn-qingdao.aliyuncs.com/pdfs/Vim%E5%AE%9E%E7%94%A8%E6%8A%80%E5%B7%A7%EF%BC%88%E7%AC%AC2%E7%89%88%EF%BC%89.pdf) 