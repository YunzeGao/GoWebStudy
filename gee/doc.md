## 设计Context
### 必要性
对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。  
但这两个对象提供的接口粒度太细，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。  
Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。  
Context 就像一次会话的百宝箱，可以找到任何东西。