window.AppConfig = {
    "backend_api": "http://127.0.0.1:18000",
    "default_swagger_json_url": "https://raw.githubusercontent.com/haozzzzzzzz/api_hub/master/backend/swagger.json",
    "swagger_request_plugins": [
        "./plugins/swagger_request_plugin.js"
    ],
    "swagger_request_proxy": {
        "localhost:18000": "127.0.0.1:18000/api/api_hub/v1/reverse_proxy/api_hub"
    }
};