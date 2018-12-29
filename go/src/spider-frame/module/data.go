package module

import (
	"net/http"
)

// 数据的接口类型
// reqeust response item 都是此类型
type Data interface {

	// 用于判断数据是否有效
	Valid() bool
}

// 条目处理器由框架使用者提供
type Item map[string]interface{}

// 数据请求的类型
type Request struct {
	// http请求
	httpReq *http.Request
	// 请求的深度
	depth uint32
}

// 数据响应的类型
type Response struct {
	// 结果
	httpResp *http.Response
	// 深度
	depth uint32
}

// 创建一个新的响应实例
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp, depth}
}

func (resp *Response) HttpResp() *http.Response {
	return resp.httpResp
}

func (resp *Response) Depth() *uint32 {
	return resp.depth
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.body != nil
}

func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request(httpReq, depth)
}

// 获取http请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// 获取采集深度
func (req *Request) Depth() uint32 {
	return req.depth
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (item Item) Valid() bool {
	return item != nil
}
