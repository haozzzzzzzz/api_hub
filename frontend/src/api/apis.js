import axios from "axios"
import config from "@/config/config"

const client = axios.create({
    baseURL: config.backendApi,
    timeout: 5000,
});

function handleResponse(comp, respData, err, callback) {
    if (err !== null) {
        let errMsg = '';
        if (err.response) {
            errMsg = 'status: ' + err.response.status + ', status_text: ' + err.response.statusText
        } else if (err.request) {
            errMsg = err.message
        } else {
            errMsg = err.message;
        }

        comp.$notify({
            type: 'error',
            title: '请求出错',
            message: errMsg
        });
        callback(respData, err);
        return
    }

    if (!respData) {
        comp.$notify({
            type: 'error',
            title: '请求出错',
            message: err.toString()
        });
        callback(respData, err);
        return
    }

    if (respData.ret !== 0) {
        comp.$notify({
           type: 'error',
           title: '接口返回错误',
           message: 'ret: '  + respData.ret + ', msg: ' + respData.message
        });
        callback(respData, err);
        return
    }

    callback(respData, err);
}

export default {
    docList(comp, pageId, limit, callback) {
        client({
            url: '/api/api_hub/v1/doc/doc/list'
        }).then(function (response) {
            handleResponse(comp, response, null, callback);
        }).catch(function (err) {
            handleResponse(comp, null, err, callback);
        });
    }
}