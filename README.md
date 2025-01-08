# goiam
An identity and access management service built with the Go language.

由Go语言构建的身份和访问管理服务。本项目根据[marmotedu/iam](https://github.com/marmotedu/iam)项目进行重构，
功能上将基本保持一致，但在实现上会有一些差异。

> 这是一个 **学习项目**，内部实现上可能会出现一些`不太必要的代码`
>
> 例如，本项目的应用构建是使用`chhz0/gosimplecobra`，这是自己实现的一个简易命令行工具，功能上比`marmotedu/iam`下的`pkg/app`应用构建工具要简单。建议参考使用`marmotedu/iam`下的`pkg/app`应用构建工具。
>
> 此外，还会由不同的其它实现，例如`log`,`errcode`等，但大部分代码将参考`marmotedu/iam`下的实现。

## 组件
核心组件：
- `iam-api-server`: 认证和授权服务
- `iam-authz-server`: 授权服务
- `iam-pump`: 数据清洗服务
- `iam-watcher`: 分布式作业服务
- `iam-sdk-go`: SDK
- `iamctl`: 命令行工具

旁路组件：
- `iam-app`: 应用服务
- `iam-webconsole`: Web控制台
- `iam-operating-system`: 运营系统
- `iam-loadbalancer`: 负载均衡器