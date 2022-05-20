package request

type RequestInterface interface {
	MethodName() string
	WithData(data interface{})
	Data() interface{}
}
