

<必须输入项> [可选输入项] {最起码输入一个}

nmap [扫描类型] [设置] {设备地址}

	地址类型：主机名、IP、网段等
	-iL <文件名>：通过文件输入地址
	-iR <IP地址数目：20 / 0 无限>  随机生成地址然后扫描
	--exclude <host1,[host2],..> 排除主机或网段
	--exclude <exclude_file> 排除文件中的地址


输出格式参数
	命名中使用时间 %H，%M，%S，%m,%d,%y
Nmapde 的输出格式	

	-oN filename
	-oX filename
	-oS filename
	-oG filename
	-oA fileBaseName

输出详情和调试设置

	-V      可以多个V
	-d[0-9]  debug
	--reason   发的包和收的包都抓出来
	--packet-trace
	--open
	--iflist
	--log-errors


其他输出设置

	--append-output
	--resume <文件名>
	--stylesheet <路径或URL地址>
	--webxml (导入从nmap.org下载的stylesheet)