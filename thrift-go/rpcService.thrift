namespace go test.rpc
namespace java test.rpc
 
// 测试服务
service RpcService {
 
    // 发起远程调用
    list<string> funCall(1:i64 callTime, 2:map<string, string> paramMap),
 
}
