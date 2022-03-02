package codec

import "io"

type Header struct {
	ServiceMethod string //服务名和方法(通常与go中的结构体和方法相映射)
	Seq           uint64 //客户端请求ID
	Error         string //错误信息
}

//抽象出对消息体进行编解码的接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

//抽象Codec的构造器
type NewCodecFunc func(io.ReadWriteCloser) Codec

//客户端根据Type来得到构造器
type Type string

//定义支持的序列类型
const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
