**常用命令**

1. 从数据库表生成model:
> goctl model mysql datasource -url="root:3hcml2gaYn&Q93hM@tcp(127.0.0.1:3306)/gozerosingle" -table="*"  -dir="./model"

2. api文件生成项目代码:
> goctl api go -api entry.api -dir . -style gozero

3. 启动命令
> 异步任务(包括延迟任务): go run mq/mq.go -f mq/etc/qiyu-mq.yaml
> api: go run qiyu.go
