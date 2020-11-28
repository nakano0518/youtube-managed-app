//Nuxtはログイン時に返却されたJWTトークンをブラウザのクッキーに保存する
//下記で(ページ読み込み時に)そのJWTトークンをVuexのストアに保存する処理を行う
export default ({app, store}) => {
    const token = app.$cookies.get('jwt_token')
    if (token) {
        store.dispatch('setToken', token)
    }
}