//axiosをラップした自作ライブラリ
export class RequestClient {
    constructor(axios) {
        this.axios = axios
    }
    async get(uri, params = {}) {
        const queryString = Object.keys(params).map(key => key + '=' + params[key]).join('&');
        const query = queryString.length > 0 ? `${uri}?${queryString}` : uri
        return await this.axios.$get(query)
    }
    async post(uri) {
        return await this.axios.$post(uri)
          .catch(err => {
            return this.retry(err)
          })
    }
}
export function createRequestClient(axios) {
    return new RequestClient(axios)       
}