import config from '@/config/config.js'

let SwaggerRequestPlugins = {
    plugins: [],
    loadPlugins(){
        let swaggerPluginPaths = config.AppConfig.swagger_request_plugins;
        if (swaggerPluginPaths===undefined || swaggerPluginPaths.length === 0) {
            return
        }

        for (let path of swaggerPluginPaths) {
            import(/* webpackIgnore: true */`${path}`).then(({default: plugin}) => {
                this.plugins.push(plugin);
            })
        }

    },

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

SwaggerRequestPlugins.loadPlugins();

export default {
    SwaggerRequestPlugins
}