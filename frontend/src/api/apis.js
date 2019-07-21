import axios from "axios"
import config from "@/config/config"

const client = axios.create({
    baseURL: config.backendApi,
    timeout: 5000,
});

function parseRespData(respData) {
    if (respData.ret !== 0) {
        return [null, respData.msg]
    } else {
        return [respData.data, null]
    }
}

export default {
    docList(pageId, limit, callback) {
        client({
            url: '/api/api_hub/v1/doc/doc/list'
        }).then(function (response) {
            let [data, err] = parseRespData(response);
            callback(data, err);
        }).catch(function (err) {
            callback(null, err.toString());
        });

        console.log('hello');
    }
}