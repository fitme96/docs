```shell
# 新建文件
echo "this is test file" > test.txt
# inode 等于1069637
root@ck-test-65:~/ck/inode# ls -i test.txt 
1069637 test.txt
# 重定向inode不变
root@ck-test-65:~/ck/inode# echo "dd"> test.txt 
root@ck-test-65:~/ck/inode# ls -i test.txt 
1069637 test.txt
# vim sed 改变inode
root@ck-test-65:~/ck/inode# vim test.txt 
root@ck-test-65:~/ck/inode# ls -i test.txt 
1069693 test.txt
# 容器inode还是1069637
root@ck-test-65:~/ck/inode# docker exec -ti 8a4a bash
root@8a4a167bc951:/# ls -i test.txt 
1069637 test.txt
# 重启,inode改变
root@ck-test-65:~/ck/inode# docker restart 8a4a
8a4a
root@ck-test-65:~/ck/inode# docker exec -ti 8a4a bash
root@8a4a167bc951:/# ls -i test.txt 
1069693 test.txt
# 文件内容同步
root@8a4a167bc951:/# cat test.txt 
dd
ddd
```