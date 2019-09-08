import swaggerPlugins from './swagger'

let SwaggerRequestPlugin = {
    requestInterceptor(req){
        // eslint-disable-next-line no-console
        console.log("swagger_request_plugin.js request:", req);
        return req
    },
    responseInterceptor(resp){
        // eslint-disable-next-line no-console
        console.log("swagger_request_plugin.js response:", resp);
        return resp
    }
};

swaggerPlugins.SwaggerRequestPlugins.plugins.push(SwaggerRequestPlugin);

export default SwaggerRequestPlugin