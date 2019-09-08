import config from '@/config/config.js'

let proxyMap = config.AppConfig.proxy;

function getProxyUrl(reqUrl) {
    for (let urlPrefix in proxyMap) {
        let idx = reqUrl.indexOf(urlPrefix);
        if ( idx >= 0){
            let proxyUrl = reqUrl.replace(urlPrefix, proxyMap[urlPrefix]);
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