nginx: [emerg] module "/usr/lib64/nginx/modules/ngx_http_image_filter_module.so" version 1016001 instead of 1020001 in /usr/share/nginx/modules/mod-http-image-filter.conf:1


```
 yum remove nginx-mod*
```

```
 yum install nginx-mod*
systemctl restart nginx

```
