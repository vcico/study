crontab 是什么？

如同win里面的定时任务

配置文件  文件方式设置定时任务
crontab 配置工具 用于调整定时任务
系统服务 crond  每分钟从配置内容刷新定时任务/执行计划任务


定时任务列表 ： cron table
单个定时任务： cron job

yum install -y vixie-cron 安装 
yum install -y crontabs 安装 
crontab -l  查看定时任务列表
service crond status 查看运行状态


crontab -e 进入编辑页面
*/1 * * * * [user-name] date >> /tmp/date.txt 输入此内容 每分钟输入一次时间到该文件
crontab -l
tail -f /tmp/date.txt  查看文件最后几行的内容  -f 动态

tail -f /var/log/cron 查看执行日志


crontab的配置文件格式

* * * * * COMMAND
| 分钟 0-59
  | 小时 0-23
    | 日期 1-31
      |月份 1-12
         |日期 0-7 (0/7代表星期天)
         
 * 表示任何时候都匹配
 A,B,C 表示A或B或C时执行
 A-B 表示A-B之间执行
 */A 表示每A 分/时/.. 执行一次 
 A-B/A 可以结合使用

例子：

    30 21 * * * service httpd restart
    每晚21:30 重启apache

    45 4 1,10,22 * * service httpd restart 
    每月1、10、22 日的4:45 重启 apache
    
    
    45 4 1-10 * * service httpd restart 
    每月1-10日4:45重启apache
    
    */2 * * * * service httpd restart
    每隔两分钟重启httpd
    
    1-59/2 * * * * command
    1-59分钟区间(奇数分钟) 每隔两分钟执行一次
    
    0 23-7/1 * * * command
    晚11 到早 7点之间每隔1小时执行 (分钟是0 如果是* 23-7之间任何一分钟都执行)
    
    0,30 18-23 * * * command  或者 0-59/30 18-23 * * * command
    18-23点之间每隔30分钟执行
    
    
    
crontab 配置文件

全局 配置文件 （只有root用户 才能增加计划任务）
    系统级的配置  crontab -l -u 无法显示系统级的配置
        /etc/crontab  在这里面设置的应该是系统计划任务
        /etc/cron.d/* 系统配置文件的补充 可以自己添加文件进去 里面的任务一般加 用户名
        /etc/cron.d/sysstat 系统计划任务 载入为root用户任务计划 crontab sysstat  导入该文件的内容到计划任务里(原本有的会消失) -- crontab -l 查看
    用户级配置文件
    crontab -l 和  cat /var/spool/cron/root 文件内容 一模一样
    crontab -e 实际修改的就是 /var/spool/cron/[用户名]




/var/log/cron 今日日志 或 /var/log/cron-date 每一天日志(简化的)
    
/etc/crontab 配置里有一个MAILTO 配置 吧所有（执行详情)发送mail给某用户

/var/spool/mail 下面有 各个用户名的文件  用以接手 所有用户执行命令情况的 mail  


常见错误

环境变量：
    修改 ~/.bash_profile 
    作用 source .bash_profile 
    这里面的用户环境变量 不起作用 
    加入到全局环境变量可以？
    
    
双引号中使用% 未加 \ 
    4月第一个星期日早晨1时59分运行a.sh
        错误
            59 1 1-7 4 0 /root/a.sh
            = 59分   一时  第一星期  月 周日   /root/a.sh
            1-7 和0 是或的关系 也就是第一星期 或者 之后的周日都执行  执行了 10 次
        正确  
            59 1 1-7 4 * test `date + \%w` -eq 0 && /root/a.sh
                         测试如果当前的日期的星期 如果是0 就执行       
        命令 date +%w 当前日期 星期几
            test `date +%w` -eq 6
            echo $?   // 0是成功的
            test 1=1 / 1=0 都是 0 test 认为 1=1 就是一字符串 用 -eq  或者 1 = 1  加空格

            
分钟设置误用
    两小时运行一次
    错误 * 0,2,4,6,8,10,12,14,16,18,20,22 * * * date
    是所有分钟都执行了 0点的所有分钟 2点的所有分钟等
    
    0 */2 * * * date 正确的设置
    只有在每两个小时 0 分钟的时候执行
    
    每2分钟执行一次
    正确 1-59/2 * * * * date
    先满足/前面的条件 再满足后面的条件
    

补充：
    半分钟执行一个任务
    shell 脚本配合 
        date ; sleep 0.5s; date
    */1 * * * * date
    */1 * * * * sleep 0.5s; date
    两个同一时间执行 后一个有休眠
    
    
    