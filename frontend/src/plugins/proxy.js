import config from '@/config/config.js'
let proxyMap = config.AppConfig.plugin_config.proxy;

/**
 * 获取代理的url
 * @param reqUrl
 * @returns {*}
 */
function getProxyUrl(reqUrl) {
    for (let host in proxyMap) {
        let idx = reqUrl.indexOf(host);
        if ( idx >= 0){
            let proxyUrl = reqUrl.replace(host, proxyMap[host]);
            // eslint-disable-next-line no-console
            console.log(`proxy: ${reqUrl} => ${proxyUrl}`);
            return proxyUrl
        }
    }

    return reqUrl
}

export default {
    getProxyUrl
}