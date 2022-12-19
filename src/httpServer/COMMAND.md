## 运行命令

### 构建镜像、查看配置
```sh
//构建镜像
docker build . -t httpserver:0.0.4
//启动镜像
docker run -d httpserver:0.0.4
//查看进程
docker ps
//查看细节拿到pid
docker inspect <containerId>
//查看容器配置
nsenter -t <pid> -n ip a
````

### 推送到dockerhub
```sh
//注册hub.docker.com
//创建仓库ellyshenxinyu/httpserver
//登陆
sudo docker login
//修改镜像名称
docker images
docker tag httpserver:0.0.4 ellyshenxinyu/httpserver:v1
//推送到dockerhub
docker images
docker push ellyshenxinyu/httpserver:v1
````