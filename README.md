# Behu
> 使用GO语言实现的知乎后端

## 简介
Behu是一个基于 [go-zero](https://github.com/zeromicro/go-zero)微服务框架实现的一个简单知乎网站后端

## 功能
### 认证模块
- 使用[casdoor](https://casdoor.org/zh/)作为身份认证平台，完成用户登录、注册等与用户鉴权相关操作
- 支持**第三方登录(OAuth2.0)**，目前支持Github登录
- 支持**SSO单点登录**
- [TODO] 黑名单机制，使得用户退出登录后，token失效