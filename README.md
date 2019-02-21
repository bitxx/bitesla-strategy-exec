# bitesla-strategy-exec
用于为[bitesla](https://github.com/jason-wj/bitesla)编写策略，目前仅支持golang编写

该项目使用方法后面等项目完善后再编辑

简单说明一下:
1. 用户只要参考根目录的`strategy.go`编写策略即可
2. Dockerfile文件是用来给[bitesla](https://github.com/jason-wj/bitesla)构建镜像的
3. /example中的策略是已经写好的，可直接使用，就是说，将其中内容直接复制到接口中调用