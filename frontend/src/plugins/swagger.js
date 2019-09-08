import proxy from './proxy'

let SwaggerRequestProxyPlugin = {
    requestInterceptor(req) {
        req.url = proxy.getProxyUrl(req.url);
        return req;
    },
    responseInterceptor(resp) {
        return resp;
    }
};

let SwaggerRequestPlugins = {
    plugins: [SwaggerRequestProxyPlugin],
    requestInterceptor(req) {
        for (let plugin of this.plugins) {
            req = plugin.requestInterceptor(req);
        }
        return req;
    },

    responseInterceptor(resp) {
        for (let plugin of this.plugins) {
            resp = plugin.responseInterceptor(resp);
        }
        return resp;
    }
};

export default {
    SwaggerRequestPlugins
}