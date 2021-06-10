package znet

import "awesomeProject1/src/zinxtcpsocket/zinxface"

//实现Router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类进行重写就好
type BaseRouter struct {


}
//这里之所以BaseRouter的方法都为空
//是因为有的Router不希望有PreHandle,PostHandle这两个业务
//所以Router全部继承BaseRouter的好处就是不需要实现PreHandle,PostHandle

//在处理conn业务之前的钩子方法Hook
func (br*BaseRouter) PreHandle(request zinxface.IRequest){}
//在处理conn业务的主钩子方法Hook
func (br*BaseRouter)Handler(request zinxface.IRequest){}
//在处理conn业务之后的主钩子方法Hook
func (br*BaseRouter)PostHandle(request zinxface.IRequest){}