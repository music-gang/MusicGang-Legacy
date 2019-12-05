import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import MusicModule from "@/store/modules/music"

export default new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    modules: {
        MusicModule
    }
})
