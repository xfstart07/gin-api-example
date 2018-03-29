例子： Gin 做 API 支持版本

## API 版本

在 Gin 中实现 API 版本的思路，通过中间件获取 Header Accept 中的版本，然后在控制器中根据版本再做一次路由分发。

### 资源

[设计一套良好的 REST API](https://zhuanlan.zhihu.com/p/34289466?utm_source=wechat_timeline&utm_medium=social&from=timeline)
