### 下载

#### 支持的操作系统

* Windows
* Linux
* macOS
* FreeBSD
* OpenBSD

#### 使用二进制发行版

1. 下载 `pmsg` [最新版本](https://github.com/BleethNie/pmsg/releases)

1. 开始运行它:

   linux

   ```shell
   $ ./pmsg --help
   
   Usage:
     pmsg [command]
   
   Available Commands:
     dingtalk    ding talk
     feishu      fei shu
     help        Help about any command
     slack       slack   
     weixin      weixin        
     workweixin  work weixin
   
   Flags:
     -h, --help      help for pmsg
     -v, --version   version for pmsg
   
   Use "pmsg [command] --help" for more information about a command.
   ```

#### linux容器映像

1. `Packages` https://github.com/BleethNie/pmsg/pkgs/container/pmsg

1. 拉取容器映像
   ```shell
   $ docker pull ghcr.io/BleethNie/pmsg
   ```

1. 开始运行它
   ```shell
   $ docker run --rm ghcr.io/BleethNie/pmsg --help
   ```

#### 源代码

```shell
$ git clone https://github.com/BleethNie/pmsg.git
```