# 分布式实时日志

## 三种模式
### client
接收日志实时显示
### agent
负责采集日志向server上报
### server
负责收集agent上报日志向client下发

## 使用
### 启动server
```shell
./logd -server
```
默认`9876`端口监听`agent`，`6789`端口监听`client`

可使用`-config`指定配置文件
```yaml
# 配置文件模版
listen:
    agent: 0.0.0.0:9876
    client: 0.0.0.0:6789
```
### 启动agent
```shell
./logd -agent -config logd.yaml
```
`agent`模式启动必须指定`-config`参数
```yaml
# 配置文件模版
server: app-server # client端显示server名称
upstream: 1.2.3.4:9876 # 收集server地址
logs:
  - file: /app.log # 日志路径
    name: app # client端显示log名称
```
### 启动client
```shell
# -upstream 指定server
./logd_darwin_amd64 -client -upstream 1.2.3.4:6789
```
