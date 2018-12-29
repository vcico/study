

## TOP获取进程的资源占用信息

    $top
    
    
    top - 09:48:38 up 2 days, 23:34,  2 users,  load average: 0.01, 0.01, 0.00
    Tasks:  90 total,   1 running,  89 sleeping,   0 stopped,   0 zombie
    Cpu(s):  0.0%us,  0.3%sy,  0.0%ni, 99.3%id,  0.0%wa,  0.0%hi,  0.3%si,  0.0%st
    Mem:   1004112k total,   509192k used,   494920k free,    68396k buffers
    Swap:  2031612k total,        0k used,  2031612k free,   233436k cached
    进程ID 用户 - - 虚拟内容 物理内存 共享内存  - - 运行时间长度  运行的命令
     PID USER      PR  NI  VIRT  RES  SHR S %CPU %MEM    TIME+  COMMAND                                                                                       
     1601 mysql     20   0  476m  41m 4360 S  0.3  4.2   2:06.55 mysqld                                                                                         
     3366 root      20   0 15024 1316 1016 R  0.3  0.1   0:00.28 top                                                                                            
     1 root      20   0 19232 1496 1232 S  0.0  0.1   0:01.27 init                                                                                           
     2 root      20   0     0    0    0 S  0.0  0.0   0:00.00 kthreadd   
     .
     额外概念： PRI 进程优先权 值越小 优先级越高越早被执行
     top中的 NI  进程可被执行的优先级的修正数值
     PRI(NEW) = PRI（old） + nice(sysstat中有)
     .
     .

    在top内按h 进入help (esc返回 回车亦可)
    
    Help for Interactive Commands - procps version 3.2.8
    Window 1:Def: Cumulative mode Off.  System: Delay 3.0 secs; Secure mode Off.

      Z,B       Global: 'Z' change color mappings; 'B' disable/enable bold
      l,t,m     Toggle Summaries: 'l' load avg; 't' task/cpu stats; 'm' mem info
      1,I       Toggle SMP view: '1' single/separate states; 'I' Irix/Solaris mode

      f,o     . Fields/Columns: 'f' add or remove; 'o' change display order 排序显示的指定列 每个列都有解释
      F or O  . Select sort field  根据某列排序
      <,>     . Move sort field: '<' next col left; '>' next col right
      R,H     . Toggle: 'R' normal/reverse sort; 'H' show threads
      c,i,S   . Toggle: 'c' cmd name/line; 'i' idle tasks; 'S' cumulative time
      x,y     . Toggle highlights: 'x' sort field; 'y' running tasks
      z,b     . Toggle: 'z' color/mono; 'b' bold/reverse (only if 'x' or 'y')
      u       . Show specific user only
      n or #  . Set maximum tasks displayed

      k,r       Manipulate tasks: 'k' kill; 'r' renice
      d or s    Set update interval
      W         Write configuration file
      q         Quit
              ( commands shown with '.' require a visible task display window ) 
    Press 'h' or '?' for help with Windows,
    any other key to continue 

    要显示列 排序 按 o  (esc返回)
    
    Current Fields:  AEHIOQTWKNMbcdfgjplrsuvyzX  for window 1:Def
    Upper case letter moves field left, lower case right  
    按选项前的字母 大写向前移动 小写向后移动 
    * 表示已经显示的列

    * A: PID        = Process Id                                                    0x001D0000  special states (2.5)
    * E: USER       = User Name                                                     0x00100000  PF_USEDFPU (thru 2.4)
    * H: PR         = Priority
    * I: NI         = Nice value
    * O: VIRT       = Virtual Image (kb)
    * Q: RES        = Resident size (kb)
    * T: SHR        = Shared Mem size (kb)
    * W: S          = Process Status
    * K: %CPU       = CPU usage
    * N: %MEM       = Memory usage (RES)
    * M: TIME+      = CPU Time, hundredths
      b: PPID       = Parent Process Pid
      c: RUSER      = Real user name
      d: UID        = User Id
      f: GROUP      = Group Name
      g: TTY        = Controlling Tty
      j: P          = Last used cpu (SMP)
      p: SWAP       = Swapped size (kb)
      l: TIME       = CPU Time
      r: CODE       = Code size (kb)
      s: DATA       = Data+Stack size (kb)
      u: nFLT       = Page Fault count
      v: nDRT       = Dirty Pages count
      y: WCHAN      = Sleeping in Function
      z: Flags      = Task Flags <sched.h>
    * X: COMMAND    = Command name/line

    Flags field:
      0x00000001  PF_ALIGNWARN
      0x00000002  PF_STARTING
      0x00000004  PF_EXITING
      0x00000040  PF_FORKNOEXEC
      0x00000100  PF_SUPERPRIV
      0x00000200  PF_DUMPCORE
      0x00000400  PF_SIGNALED
      0x00000800  PF_MEMALLOC
      0x00002000  PF_FREE_PAGES (2.5)
      0x00008000  debug flag (2.5)
      0x00024000  special threads (2.5)
    Upper case letter moves field left, lower case right 

    * A: PID        = Process Id                                                    0x001D0000  special states (2.5)
    * E: USER       = User Name                                                     0x00100000  PF_USEDFPU (thru 2.4)
    * H: PR         = Priority
    * I: NI         = Nice value
    * O: VIRT       = Virtual Image (kb)
    * Q: RES        = Resident size (kb)
    * T: SHR        = Shared Mem size (kb)
    * W: S          = Process Status
    * K: %CPU       = CPU usage
    * N: %MEM       = Memory usage (RES)
    * M: TIME+      = CPU Time, hundredths
      b: PPID       = Parent Process Pid
      c: RUSER      = Real user name
      d: UID        = User Id
      f: GROUP      = Group Name
      g: TTY        = Controlling Tty
      j: P          = Last used cpu (SMP)
      p: SWAP       = Swapped size (kb)
      l: TIME       = CPU Time
      r: CODE       = Code size (kb)
      s: DATA       = Data+Stack size (kb)
      u: nFLT       = Page Fault count
      v: nDRT       = Dirty Pages count
      y: WCHAN      = Sleeping in Function
      z: Flags      = Task Flags <sched.h>
    * X: COMMAND    = Command name/line

    Flags field:
      0x00000001  PF_ALIGNWARN
      0x00000002  PF_STARTING
      0x00000004  PF_EXITING
      0x00000040  PF_FORKNOEXEC
      0x00000100  PF_SUPERPRIV
      0x00000200  PF_DUMPCORE
      0x00000400  PF_SIGNALED
      0x00000800  PF_MEMALLOC
      0x00002000  PF_FREE_PAGES (2.5)
      0x00008000  debug flag (2.5)
      0x00024000  special threads (2.5)

命令例子

    top -ab -n 1 | grep mysql
    a 根据内存倒序排序
    b 所有列全部列出
    -n 1 只执行一次