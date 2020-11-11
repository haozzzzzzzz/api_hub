window.AppConfig = {
    "backend_api": "http://127.0.0.1:18000",
    "default_swagger_json_url": "https://raw.githubusercontent.com/haozzzzzzzz/api_hub/master/backend/swagger.json",
    "plugins": [ // plugin
        "./swagger_request_plugin.js",
        "./private/api_hub_plugins/nosign_request.js"
    ],

    // 插件配置
    "plugin_config": {
        // 签名跳过
        "nosign_request": {
            "domains": [
                "test-m.videobuddy.vid007.com",
                "test-sg.videobuddy.vid007.com",
                "test-vn.videobuddy.vid007.com",
            ]
        },

        // 代理
        "proxy": {
            "test-m.videobuddy.vid007.com": "127.0.0.1:18000/api/api_hub/v1/reverse_proxy/test_videobuddy"
        },
    },

};