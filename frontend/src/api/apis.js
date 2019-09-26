import axios from "axios"
import config from "@/config/config.js"

const client = axios.create({
    baseURL: config.AppConfig.backend_api,
    timeout: 10000,
});

function handleResponse(comp, resp, err, callback) {
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
        callback(null, err);
        return
    }

    if (!resp) {
        comp.$notify({
            type: 'error',
            title: '请求出错',
            message: err.toString()
        });
        callback(null, err);
        return
    }

    let respData = resp.data;
    if (respData.ret !== 0) {
        // eslint-disable-next-line no-console
        console.log(respData);
        comp.$notify({
           type: 'error',
           title: '接口返回错误',
           message: 'ret: '  + respData.ret + ', msg: ' + respData.msg
        });
        callback(respData, err);
        return
    }

    callback(respData, err);
}

let handlingRequest = {};
export default {
    /**
     * api document list
     * @param comp component
     * @param pageId page id
     * @param limit page limit size
     * @param search search word
     * @param callback callback function
     */
    docList(comp, pageId, limit, search, callback) {
        if (handlingRequest[this.docList]) {
            callback(null, "last doc list request is not finish");
            return
        }

        handlingRequest[this.docList] = true;
        client({
            method: 'get',
            url: '/api/api_hub/v1/doc/doc/list',
            params: {
                page: pageId,
                limit: limit,
                search: search,
            }
        }).then((response) => {
            handleResponse(comp, response, null, callback);
            handlingRequest[this.docList] = false;
        }).catch((err) => {
            handleResponse(comp, null, err, callback);
            handlingRequest[this.docList] = false;
        });
    },

    docDetailSpecUrl(docId) {
        return config.AppConfig.backend_api + '/api/api_hub/v1/doc/detail/spec/' + docId
    }
}