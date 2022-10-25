# CONFIG MEMCACHED

```shell
 sudo dockerd
 docker run -d --restart always -m 256M -p 127.0.0.1:11211:11211 --name memcached -it memcached:alpine

 docker stop memcached
```
