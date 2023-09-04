#!/bin/bash


for dir in `ls -l .|grep "^d"|awk '{print $NF}'`; ## 获取项目根目录所有一级目录开启循环
do
        echo $dir|awk -F "/" '{print "- ""["$NF"]("$dir"/)"}' ## 先打印一级目录
        dirsum=`ls -l $dir |grep "^d"|wc -l` 
        if [ $dirsum -ne 0 ];then  ## 判断是否有二级目录
                subdirs=`ls -l $dir |grep "^d"|awk '{print $NF}'` 
		for subdir in `ls -l $dir |grep "^d"|awk '{print $NF}'`;
		do
                	echo "  - [$subdir]($dir/$subdir)"  ## 先打印二级目录
                	find  $dir/$subdir  -maxdepth 1 -name "*.md" |awk -F "/" '{print "    - ["$NF"]""("$i")"}' ## 再打印二级目录所有markdown
		done
        fi
        find  $dir -maxdepth 1 -name "*.md" |awk -F "/" '{print "  - ["$NF"]""("$i")"}' ## 最后打印一级目录markdown
done
