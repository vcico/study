# sysstet 统计系统的各种资源占用情况。 包括 CPU 内存 IP 网络

yum install sysstat
yum list sysstat


由cron生成统计文件 再有sysstat分析
/etc/cron.d/sysstat

# Run system activity accounting tool every 10 minutes
*/10 * * * * root /usr/lib64/sa/sa1 1 1   cpu占用的统计？？？
*/1 * * * * root /usr/lib64/sa/sa1 1 1   开发时改成 每一分钟获取一次
    cat /var/log/sa/sa29  文件存在这里 但乱码的 二进制的
# 0 * * * * root /usr/lib64/sa/sa1 600 6 &
# Generate a daily summary of process accounting at 23:53
53 23 * * * root /usr/lib64/sa/sa2 -A   每天一个总结报告


## sysstat 命令

### CPU监控
`sar -q -f /var/log/sa/sa08` 占用情况 sa08是/var/log/sa/下文件名 

    10:04:54 AM   runq-sz  plist-sz   ldavg-1   ldavg-5  ldavg-15
    10:04:56 AM         0       110      0.00      0.00      0.00
    10:10:01 AM         0       111      0.00      0.00      0.00
    Average:            0       110      0.00      0.00      0.00

    runq-sz  等待执行的队列任务 越长阻塞越严重
    plist-sz   任务队列中的任务总数
    ldavg-1/5/15 load average 1分钟 5分钟 15分钟 内系统负载的描述
        值是通过执行中的任务和等待执行的任务的个数的平均值得到的
        如果值 == cpu个数 等于满载执行。
        2个cpu 值是6 相当于每人能做一个 却给了3个任务 等待队列会变长
        为什么 1-15分钟进行统计？
            1： 6 ； 5：2 ；15：0.5 这也就意味着一分钟内忽然涌入大量请求 然后忽然降下来 可以用top找谁占用多
            1:0.5 ； 5： 3 ；15:  5  反之
    Average
    
`sar -p -f /var/log/sa/sa08` 看占用百分比
                                                                                
    10:04:54 AM     CPU     %user     %nice   %system   %iowait    %steal     %idle 
    10:04:56 AM     all      0.45      0.00      1.34      0.45      0.00     97.77
    10:10:01 AM     all      0.02      0.00      0.15      0.01      0.00     99.82
    Average:        all      0.02      0.00      0.16      0.01      0.00     99.81

    额外概念： PRI 进程优先权 值越小 优先级越高越早被执行
    
    nice    改过优先级的进程CPU占用率 值小级高
    idle 空闲
    steal  在等CPU。越高说明CPU任务越繁重。（ 管理程序(hypervisor)为另一个虚拟进程提供服务而等待虚拟cpu的百分比）
    IOwait 等待磁盘读写 越高磁盘任务繁重
    
## 内存监控

sar -r -f /var/log/sa/sa29

    10:04:54 AM kbmemfree kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit
    10:04:56 AM    491124    512988     51.09     68912    235944    393256     12.95
    10:10:01 AM    490892    513220     51.11     68984    235948    394976     13.01
    10:20:01 AM    490628    513484     51.14     69120    235952    394976     13.01
    10:30:01 AM    490544    513568     51.15     69232    235956    394976     13.01
    Average:       490797    513315     51.12     69062    235950    394546     13.00

    kbmemfree 空闲 物理内存
    kbmemused 已使用 物理内存
    memused 内存使用率
    kbbuffers 缓存使用量
    kbcached 缓存使用量
    kbcommit 为了保证程序正常运行需要的内存
    commit  当前已申请的内容还剩余多少 
    如果 commit + memused 超出 100% 就会再申请内存 频繁换页 会用到虚拟内容临时存放
    
    buffer 和 cached 都是为了提高文件读取性能的磁盘缓存
    page cache 针对文件系统的 是文件的缓存 
    buffer cache 针对磁盘快(文件磁盘块)的缓存
    
sar -B -f /var/log/sa/sa29

    10:04:54 AM  pgpgin/s pgpgout/s   fault/s  majflt/s  pgfree/s pgscank/s pgscand/s pgsteal/s    %vmeff
    10:04:56 AM      0.00     21.43    645.09      0.00    208.93      0.00      0.00      0.00      0.00
    10:10:01 AM      0.00      0.67     10.06      0.00      2.97      0.00      0.00      0.00      0.00
    10:20:01 AM      0.00      0.58     14.90      0.00      4.86      0.00      0.00      0.00      0.00
    10:30:01 AM      0.00      0.42      4.89      0.00      1.71      0.00      0.00      0.00      0.00
    Average:         0.00      0.57     10.87      0.00      3.53      0.00      0.00      0.00      0.00

    
    pgpgin  换页 换入   换页用到虚拟内存 用到磁盘读写 会冲击IO
    pgpgout 换页 换出
    fault 每秒钟系统产生的缺页数 
    majflt 每秒钟产生的主缺页数
    主要分析以上4个
    pgfree 每秒被放入空闲队列的页个数
    pgscank 每秒被 kswapd 扫描的页个数
    pgscand 每秒直接被扫描的页个数
    pgsteal 每秒钟从cache中被清除来满足内?要的页个数
    vmeff 每秒清除的页 pgsteal 站总扫描页(pgscank+pgscand)的百分比
    
    缺页： 内容放不下 放到了磁盘虚拟内存  读取内存找不到 
        major 主要内容缺页 必须进行换页
        minor 次要内容缺页 内容机制可以解决掉
        
 sar -W -f /var/log/sa/sa29
 
    10:04:54 AM  pswpin/s pswpout/s
    10:04:56 AM      0.00      0.00
    10:10:01 AM      0.00      0.00
    10:20:01 AM      0.00      0.00
    10:30:01 AM      0.00      0.00
    10:40:01 AM      0.00      0.00
    10:50:01 AM      0.00      0.00
    Average:         0.00      0.00
    
    page swap in  用交互分区 倒腾数据的个数
    page swap out 

    
## IO监控

sar -b -f /var/log/sa/sa29

    10:04:54 AM       tps      rtps      wtps   bread/s   bwrtn/s
    10:04:56 AM      6.70      0.00      6.70      0.00     85.71
    10:10:01 AM      0.30      0.00      0.30      0.00      2.68
    10:20:01 AM      0.26      0.00      0.26      0.00      2.32
    10:30:01 AM      0.19      0.00      0.19      0.00      1.68
    10:40:01 AM      0.19      0.00      0.19      0.00      1.66
    10:50:01 AM      0.22      0.00      0.22      0.00      1.92
    Average:         0.23      0.00      0.23      0.00      2.05

    tps 每秒物理设备的IP请求次数
    rtps 每秒从物理设备读取的请求次数 
    wtps 每秒向物理设备写入的请求次数 次数多 可以缓存内容 一起写入
    bread/s 每秒从物理设备读取的数据量 单位块/s
    bwrtn/s 每秒向物理设备写入的数据量 单位块/s
    
sar -d -f /var/log/sa/sa29


    10:50:01 AM       DEV       tps  rd_sec/s  wr_sec/s  avgrq-sz  avgqu-sz     await     svctm     %util
    11:00:01 AM   dev11-0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    11:00:01 AM    dev8-0      0.09      0.00      0.83      9.36      0.00      1.34      1.06      0.01
    11:00:01 AM  dev253-0      0.10      0.00      0.83      8.00      0.00      1.29      0.90      0.01
    11:00:01 AM  dev253-1      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    Average:      dev11-0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    Average:       dev8-0      0.10      0.00      0.99     10.07      0.00      1.39      1.06      0.01
    Average:     dev253-0      0.12      0.00      0.99      8.00      0.00      1.58      0.85      0.01
    Average:     dev253-1      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

    DEV       
    tps   每秒从物理磁盘IO的次数 多个逻辑请求会被合并为一个IO请求 一次传输的数据大小不确定
    rd_sec/s   每秒读扇区的次数
    wr_sec/s   每秒写扇区的次数
    avgrq-sz  平均每次设备IO的数据大小(扇区)
    avgqu-sz    磁盘请求队列的平均长度
    await     从请求磁盘操作到系统完成处理每次平均的消耗时间。包括请求队列的等待时间 的那位毫秒
    svctm     系统处理每次请求的平均时间 不包括请求队列中消耗的时间
    %util   IO请求站CPU的百分比 比例越大 越饱和 或者 负担越大
    
## newwork 监控


DEV 显示网络接口信息
EDEV 显示关于网络错误的统计数据
NFS 统计活动的NFS客户端的信息
NFSD 统计NFS服务器的信息
sock 显示套接字信息
ALL 显示所有5个开关 
他们可以单独或一起使用

sar -n DEV -f /var/log/sa/sa29  

    
    10:04:54 AM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
    10:04:56 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:04:56 AM      eth0      5.36      2.68      0.61      0.34      0.00      0.00      0.00
    10:10:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:10:01 AM      eth0      0.54      0.22      0.05      0.04      0.00      0.00      0.00
    10:20:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:20:01 AM      eth0      0.37      0.12      0.04      0.01      0.00      0.00      0.00
    10:30:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:30:01 AM      eth0      0.34      0.07      0.03      0.03      0.00      0.00      0.00
    10:40:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:40:01 AM      eth0      0.26      0.08      0.03      0.01      0.00      0.00      0.00
    10:50:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    10:50:01 AM      eth0      0.29      0.04      0.03      0.00      0.00      0.00      0.00
    11:00:01 AM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    11:00:01 AM      eth0      0.48      0.12      0.05      0.01      0.00      0.00      0.00
    Average:           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00
    Average:         eth0      0.37      0.10      0.04      0.02      0.00      0.00      0.00

    IFACE   eth0 第一块网卡 
    rxpck/s    每秒接受的数据包
    txpck/s    每秒钟发送的数据包
    rxkB/s    每秒接受的字节数
    txkB/s   每秒发送的字节数
    rxcmp/s   每秒接受的压缩数据包
    txcmp/s  每秒发送的压缩数据包
    rxmcst/s 每秒接受的多播数据包
    
sar -n NFS -f /var/log/sa/sa29

    [为什么和我执行出来的不一样](ser-nNFS-f.png)
    sysstat 版本问题？  系统版本问题 centos6.9
    
    10:04:54 AM    call/s retrans/s    read/s   write/s  access/s  getatt/s
    11:10:01 AM      0.00      0.00      0.00      0.00      0.00      0.00
    Average:         0.00      0.00      0.00      0.00      0.00      0.00

    call/s  
    retrans/s    
    read/s   
    write/s  
    access/s  
    getatt/s

## 评估磁盘读写性能极限

    fio -filename=/root/test -direct=1 -iodepth 1 -thread -rw=randrw -ioengine=psync   -bs=16k -size 2G -numjobs=10 -runtime=30 -group_reporting -name=mytest13
    
    
    
    mytest13: (g=0): rw=randrw, bs=16K-16K/16K-16K/16K-16K, ioengine=psync, iodepth=1
    ...
    mytest13: (g=0): rw=randrw, bs=16K-16K/16K-16K/16K-16K, ioengine=psync, iodepth=1
    fio-2.0.13
    Starting 10 threads
    mytest13: Laying out IO file(s) (1 file(s) / 2048MB)
    Jobs: 10 (f=10): [mmmmmmmmmm] [100.0% done] [1023K/868K/0K /s] [63 /54 /0  iops] [eta 00m:00s]
    mytest13: (groupid=0, jobs=10): err= 0: pid=3537: Sat Dec 29 11:20:49 2018
      read : io=30464KB, bw=1011.9KB/s, iops=63 固态硬盘 1181 , runt= 30107msec
        clat (usec): min=28 , max=1701.3K, avg=56375.37, stdev=151031.17
         lat (usec): min=28 , max=1701.4K, avg=56375.72, stdev=151031.12
        clat percentiles (usec):
         |  1.00th=[   30],  5.00th=[   41], 10.00th=[   50], 20.00th=[   70],
         | 30.00th=[   84], 40.00th=[  107], 50.00th=[  161], 60.00th=[11328],
         | 70.00th=[55552], 80.00th=[90624], 90.00th=[144384], 95.00th=[201728],
         | 99.00th=[774144], 99.50th=[1368064], 99.90th=[1564672], 99.95th=[1695744],
         | 99.99th=[1695744]
        bw (KB/s)  : min=    9, max=  819, per=11.61%, avg=117.42, stdev=101.55
      write: io=32784KB, bw=1088.1KB/s, iops=68 固态硬盘 1180 普通 200-300, runt= 30107msec
        clat (usec): min=490 , max=573412 , avg=94037.37, stdev=83996.57
         lat (usec): min=491 , max=573415 , avg=94052.24, stdev=84000.60
        clat percentiles (usec):
         |  1.00th=[  700],  5.00th=[ 2064], 10.00th=[ 5984], 20.00th=[17280],
         | 30.00th=[33024], 40.00th=[52992], 50.00th=[74240], 60.00th=[93696],
         | 70.00th=[123392], 80.00th=[158720], 90.00th=[216064], 95.00th=[264192],
         | 99.00th=[342016], 99.50th=[374784], 99.90th=[423936], 99.95th=[497664],
         | 99.99th=[569344]
        bw (KB/s)  : min=    9, max=  637, per=10.98%, avg=119.45, stdev=75.27
        lat (usec) : 50=4.68%, 100=13.13%, 250=8.58%, 500=0.99%, 750=0.73%
        lat (usec) : 1000=0.35%
        lat (msec) : 2=1.67%, 4=1.87%, 10=3.82%, 20=5.62%, 50=11.48%
        lat (msec) : 100=18.90%, 250=23.35%, 500=3.97%, 750=0.35%, 1000=0.13%
        lat (msec) : 2000=0.38%
      cpu          : usr=0.00%, sys=1.63%, ctx=18446744073709546677, majf=18446744073709551598, minf=18446744073709422918
      IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
         submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
         complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
         issued    : total=r=1904/w=2049/d=0, short=r=0/w=0/d=0

    Run status group 0 (all jobs):
       READ: io=30464KB, aggrb=1011KB/s, minb=1011KB/s, maxb=1011KB/s, mint=30107msec, maxt=30107msec
      WRITE: io=32784KB, aggrb=1088KB/s, minb=1088KB/s, maxb=1088KB/s, mint=30107msec, maxt=30107msec

    Disk stats (read/write):
        dm-0: ios=1903/2060, merge=0/0, ticks=37728/29685, in_queue=67414, util=99.48%, aggrios=1908/2064, aggrmerge=0/0, aggrticks=37849/29721, aggrin_queue=67568, aggrutil=99.39%
      sda: ios=1908/2064, merge=0/0, ticks=37849/29721, in_queue=67568, util=99.39%
