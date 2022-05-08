
### asynq命令

1. 安装
   > go get -u github.com/hibiken/asynq/tools/asynq
2. 查看cron任务
   > asynq cron ls -p "3hcml2gaYn&Q93hM"
3. 查看cron历史任务
   > asynq cron history b9749c03-8c30-4dff-a692-1f146fcdedd1 -p "3hcml2gaYn&Q93hM"
   > history后面是entry id
4. webUI
   > 直接下载asynqmon然后
   > asynqmon -redis-password="3hcml2gaYn&Q93hM"即可