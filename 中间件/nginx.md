nginx: [emerg] module "/usr/lib64/nginx/modules/ngx_http_image_filter_module.so" version 1016001 instead of 1020001 in /usr/share/nginx/modules/mod-http-image-filter.conf:1


```
 yum remove nginx-mod*
```

```
 yum install nginx-mod*
systemctl restart nginx

```


limit_conn_zone

limit_zone 该指令在 1.1.8 版中已过时，并在 1.7.6 版中删除。应使用语法有所改变的等效 limit_conn_zone 指令,[官网说明](https://nginx.org/en/docs/http/ngx_http_limit_conn_module.html#limit_zone)


防盗链
```
location ~* ^.+\.(gif|jpg|png|swf|flv|rar|zip)$ {  
valid_referers none blocked www.baidu.com;  
	if ($invalid_referer) {  
		return 403;  
	}  
}
```


隐藏nginx版本
```
自定义
修改nginx解压路径(eg:/usr/local/nginx-1.5.6/src/http/ngx_http_header_filter_module.c)文件的第48和49行内容，自定义头信息：  
static char ngx_http_server_string[] = “Server:XXXXX.com” CRLF;  
static char ngx_http_server_full_string[] = “Server:XXXXX.com” CRLF;  

隐藏版本号http区域增加如下  
server_tokens off;
```

限制ip访问

```
deny 192.168.1.1; #拒绝IP  
allow 192.168.1.0/24; #允许IP  
allow 10.1.1.0/16; #允许IP  
deny all; #拒绝其他所有IP
```



过去3个月中