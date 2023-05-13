import HttpRequest from '../request/index'

// 获取服务信息
export const AddFile = (url: string)=> {
    return HttpRequest.post('/gRPCTool/serverInfo', {url: url})
}

// 调用接口
export const Query = (serviceName: string, methodName: string, url: string, data: string)=> {
    return HttpRequest.post('/gRPCTool/call', {url: url, serviceName: serviceName, methodName: methodName ,data: data})
}

export const GetMethodParam = (serviceName, methodName,url)=> {
    return HttpRequest.post('/gRPCTool/methodParam', {url: url, serviceName: serviceName, methodName: methodName})
}

// 保存接口和方法
export const SetFile = (data)=>{
    return HttpRequest.post('/gRPCTool/set',{data: data})
}

// 获取接口和方法
export const GetFile = ()=>{
    return HttpRequest.post('get')
}