//アクションの定義
import {createRequestClient} from '~/store/request-client';
import firebase from '~/plugins/firebase';

export const state = () => ({
    items: [],
    relatedItems: [],
    searchItems: [],
    item: {},
    meta: {},
    searchMeta: {},
    token: '',
})

export const actions = {
    async fetchPopularVideos({commit}, payload) {
        const client = createRequestClient(this.$axios)
        const res = await client.get(payload.uri, payload.params)
        commit("mutatePopularVideos", res)
    },
    async findVideo({commit}, payload) {
        const client = createRequestClient(this.$axios)
        const res = await client.get(payload.uri)
        const params = {
            ...res.video_list,
        }
        commit('mutateVideo', params)
    },
    async fetchRelatedVideos({commit}, payload) {
        const client = createRequestClient(this.$axios)
        const res = await client.get(payload.uri)
        commit('mutateRelatedVideos', res)
    },
    async searchVideos({commit}, payload) {
        const client = createRequestClient(this.$axios, this.$cookies, this)
        const res = await client.get(payload.uri, payload.params)
        commit('mutateSearchVideos', res)
    },
    async signUp({commit, dispatch}, payload) {
        console.log(process.env.API_KEY)
        await firebase.auth().createUserWithEmailAndPassword(payload.email, payload.password)
        const res = await firebase.auth().signInWithEmailAndPassword(payload.email, payload.password)
        const token = await res.user.getIdToken()
        this.$cookies.set('jwt_token', token)
        commit('mutateToken', token)
        this.app.router.push('/')
    },
    async setToken({commit}, payload) {
        commit('mutateToken', payload)
    },
}

export const mutations = {
    mutatePopularVideos(state, payload) {
        state.items = payload.items ? state.items.concat(payload.items) : []
        state.meta = payload
    },
    mutateVideo(state, payload) {
        const params = (payload.items && payload.items.length > 0) ? payload.items[0] : {}
        state.item = params
    },
    mutateRelatedVideos(state, payload) {
        state.relatedItems = payload.items || []
    },
    mutateSearchVideos(state, payload) {
        state.searchItems = payload.items ? state.searchItems.concat(payload.items) : []
        state.searchMeta = payload
    },
    mutateToken(state, payload) {
        state.token = payload
    },
}

export const getters = { //Vueコンポーネントでstateを参照するためのgetter
    getPopularVideos(state) {
        return state.items 
    },
    getMeta(state) {
        return state.meta
    },
    getVideo(state) {
        return state.item
    },
    getRelatedVideos(state) {
        return state.relatedItems
    },
    getSearchVideos(state) {
        return state.searchItems
    },
    getSearchMeta(state) {
        return state.searchMeta
    },
}