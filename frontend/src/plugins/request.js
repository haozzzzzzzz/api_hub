import config from '@/config/config.js'

let SwaggerRequestProxyPlugin = {
    proxyMap: config.AppConfig.swagger_request_proxy,
    requestInterceptor(req) {
        let reqUrl = req.url;
        for (let urlPrefix in this.proxyMap) {
            let idx = reqUrl.indexOf(urlPrefix);
            if ( idx >= 0){
                let proxyUrl = reqUrl.replace(urlPrefix, this.proxyMap[urlPrefix]);
                // eslint-disable-next-line no-console
                console.log(`proxy: ${reqUrl} => ${proxyUrl}`);
                req.url = proxyUrl;
            }
        }
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