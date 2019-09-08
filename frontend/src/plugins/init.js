import config from '@/config/config.js'

let pluginPaths = config.AppConfig.plugins;

if (pluginPaths !== undefined && pluginPaths.length > 0 ) {
    for (let path of pluginPaths) {
        import(`${path}`)
    }
}