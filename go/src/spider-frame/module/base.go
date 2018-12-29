package module


/* 
值一般由 3 部分组成: 
	标识组件类型的字母
	代表生成顺序的序列号
	<ip>:<port> 127.0.0.1:8020 用于定位组件的网络地址(可选 组件可以和爬虫主程序处于同一进程)
*/
type MID string

// 组件ID的模板
var midTemplate = "%s%d|%s" 
// 组件的类型
type Type String
//  当前认可的组件类型的常量
const (
	TYPE_DOWNLOADER Type = "downloader"
	TYPE_ANALYZER Type = "analyzer"
	TYPE_PIPELINE Type = "pipeline"
)
// 合法组件类型 - 字母 的映射
var legalTypeLetterMap = map[Type]string{
	TYPE_DOWNLOADER:"D",
	TYPE_ANALYZER:"A",
	TYPE_PIPELINE:"P",
}
// 序列号生成器的接口类型
type SNGenertor interface{
	//获取预设的最小序列号
	Start() uint64
	// 获取预设的最大序列号 使用方 初始化生成器时 给定 (上同)
	Max() uint64
	// 获取下一个序列号
	Next() uint64
	// 获取循环计数  
	CycleCount() uint64
	// 获得一个序列号并准备下一个序列号
	Get() uint64
}
// 组件注册器接口
type Registrar interface{
	//  用于注册组件实例
	Register(module Module)(bool,error)
	// 用于注销组件实例
	Unregister(mid MID)(bool,error)
	// 基于负载均衡 返回一个指定类型的实例
	Get(moduleType Type)(Module,error)
	// 获取指定类型的所有组件实例
	GetAllByType(moduleType Type)(map[MID]Module,error)
	// 用于获取所有组件实例
	GetAll() map[MID]Module
	// 清除所有的组件注册记录
	Clear()
}




// 下载器接口 （下载器功能： 发送请求 接受响应。从参数签名可以体现）
Download(req *Request)(*Response,error)


// 用于计算组件评分的函数类型
type CalculateScore func(counts Counts) uint64

// 所有处理模块(包括下载器)都应具备一些方法 进行统计 描述 需要更加抽象的接口类型
// module代表组件的基础接口类型
// 改接口的实现类型必须的并发安全的
type Module interface{
	// 用于获取当前组件的ID
	ID() MID
	// 用于获取当前组件的网络地址的字符串形式 <ip>:<port>  类型：远程模块 远程加载 并非爬取页面的网址
	Addr() string
	// 用于获取当前组件的评分
	Score() uint64
	// 设置当前组件的评分
	SetScore(score uint64)
	// 用于获取评分计算器
	ScoreCalculator() CalculateScore
	// 用于获取当前组件被调用的计数
	CalledCount() uint64
	// 用于获取当前组件接收的调用的计数（有可能由于超负荷或参数有误拒绝调用）
	AcceptedCount() uint64
	// 用于获取当前组件已成功完成的调用的计数
	CompletedCount() uint64
	// 用于获取当前组件正在处理的调用的数量
	HandlingNumber() uint64
	// 用于一次性获取所有计数
	Counts() Counts
	// 用于获取组件的摘要
	Summary() SummaryStruct
}