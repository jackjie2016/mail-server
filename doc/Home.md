简单的搭建一个邮件服务，是我的一个愿望。欢迎大家一起交流。

### 下载安装

- 二进制安装

```
* 检查环境要求是否已满足
* 解压压缩包。
* 使用命令 cd 进入到刚刚创建的目录。
* 执行命令 ./mail-server service。
* mail-server 默认会在端口 1080 启动 HTTP 服务）。
```

安装完成后可继续参照配置与运行。

- 源码安装

mail-server要求至少使用 Go 1.15 或更高的版本进行编译，具体安装步骤请参考[官方文档](https://golang.org/doc/install)。

```
# 克隆仓库到 "mail-server" 子目录
git clone --depth 1 https://github.com/phper95/mail-server.git
# 修改工作目录
cd mail-server
# 编译主程序，这个步骤会下载所有依赖
go build main.go
```

### 邮件域名配置
- [域名配置](https://github.com/phper95/mail-server/wiki/%E9%82%AE%E4%BB%B6%E5%9F%9F%E5%90%8D%E9%85%8D%E7%BD%AE)


## 快速安装

```
curl -fsSL  https://raw.githubusercontent.com/phper95/mail-server/master/scripts/install.sh | sh
```

## 快速开发
```
curl -fsSL  https://raw.githubusercontent.com/phper95/mail-server/master/scripts/install_dev.sh | sh
```
