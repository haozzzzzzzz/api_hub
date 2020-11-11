import proxy from './proxy'
import Url from 'url-parse'
let SwaggerRequestProxyPlugin = {
    requestInterceptor(req) {
        req.url = proxy.getProxyUrl(req.url);
        return req;
    },
    responseInterceptor(resp) {
        return resp;
    }
};

// swagger请求中间件
let SwaggerRequestPlugins = {
    plugins: [SwaggerRequestProxyPlugin],
    requestInterceptor(req) {
        let url = new Url(req.url, true);
        req.headers["Host"] = url.host;
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