哨兵error定义好
dao wrap层抛出
service 无法避免error则不处理直接透传
service 可降级处理则处理完 err清空

降级err处理  日志打印在service层最近api
透传err  打印在处理req的service顶层
api-service这么处理挺好的
但是功能类的包反而非透明的error比较好