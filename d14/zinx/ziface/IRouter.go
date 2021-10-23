package ziface

// 路由的抽象接口,
type IRouter interface{
	//处理之前的方法
	BeforeHandle(request IRequest)
	//处理时的方法
	Handle(request IRequest)
	//处理后的方法
    AfterHandle(request IRequest)
	//
}