namespace go hello.demo

/**
 * 结构体定义
 */
struct HelloReply{
    1: string Message,
}
struct  HelloRequest {
	1: string Name
}

service helloThrift {
        HelloReply SayHello(1:HelloRequest helloReq),
}
