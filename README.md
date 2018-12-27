demo用于演示这个[base业务框架](https://github.com/microsvs/base)的使用，方便大家对它更加熟悉

demo由四个微服务构成：gateway、token、address和user

**注意：如果大家想直接利用这个环境，我可以提供相关分支，测试没有问题再合并tag。但是不要在配置中心和mysql上删除数据, 同时如果大家愿意让这台ECS和弹性公网IP一直工作，可以资助下，保证该服务能够一直运行，或者免费提供ECS等**

**2018.12.26上线的demo，我看到有很多用户都在访问这个服务。**

```shell
# TODO list
1. 提供自动提交代码，并自动构建部署
2. Dockfile容器化部署
```
1. gateway作为网关，对外提供web服务
2. token对已登录用户进行身份验证
3. user获取用户相关信息
4. address提供地址服务。


## DEMO截图

### 配置中心

![查询地址接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG48.jpeg)

### 查询接口

1. 查询me接口

![查询me接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG44.jpeg)

2. 查询错误码接口

![查询错误码接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG45.jpeg)

3. 查询地址接口

![查询地址接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG46.jpeg)

4. 日志目录结构与日志信息截图

![查询地址接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG47.jpeg)

5. mysql截图

![查询地址接口](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG49.jpeg)

## 演示环境


### 服务列表

| 服务名 | 服务地址 | 用户名 | 密码 |
|---|---|---|---|
| zkui | [zkui](http://39.96.95.220:9090/login) | admin | manager |
| gateway | [demo-dev](http://39.96.95.220:8081?token=e3215ffa-8bd6-4010-aafb-d7817f3103dc) | - | - |
| token | 192.168.0.79:8084 | - | - |
| user | 192.168.0.79:8085 | - | -|
| mysql | 39.96.95.220:3306 | demo_dev | UcvIZn8QKAs7| 

### 捐赠

保证ECS和弹性公网IP，能够一直对外提供服务。自己在七牛云上购买了一台2核2G Ubuntu14的ECS和流量计费的1M带宽弹性公网IP

![支付宝](https://gewuwei.oss-cn-shanghai.aliyuncs.com/tracelearning/WechatIMG50.jpeg)

捐赠人名单：


