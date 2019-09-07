let SwaggerRequestPlugin = {
    requestInterceptor(req){
        console.log(req);
        return req
    },
    responseInterceptor(resp){
        console.log(resp);
        return resp
    }
};

export default SwaggerRequestPlugin