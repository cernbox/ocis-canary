// eslint-disable-next-line camelcase
import { Canary_SetCanary } from './client/canary'
import axios from 'axios'

const state = {
  config: null,
  error: ''
}

const getters = {
  config: state => state.config,
  error: state => state.error
}

const actions = {
  // Used by ocis-web.
  loadConfig ({ commit }, config) {
    console.log("load config", config)
    commit('LOAD_CONFIG', config)
  },

  setCanary ({ commit, dispatch, getters, rootGetters }, version) {
    injectAuthToken(rootGetters)
    Canary_SetCanary({
      $domain: rootGetters.configuration.server,
      body: { version: version }
    })
      .then(response => {
        if (response.status === 200 || response.status === 201) {
          document.cookie = response.data.cookie+"="+version+"; max-age="+response.data.ttl+"; path=/";
          window.location.href= "/";
        } else {
          dispatch('showMessage', {
            title: 'Failed to set the version',
            desc: response.statusText,
            status: 'danger'
          }, { root: true })
        }
      })
      .catch(error => {
        console.error(error)

        dispatch('showMessage', {
          title: 'Failed to set version',
          desc: error.message,
          status: 'danger'
        }, { root: true })
      })
  }
}

const mutations = {
  SET_MESSAGE (state, payload) {
    state.error = payload
  },

  LOAD_CONFIG (state, config) {
    state.config = config
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}

function injectAuthToken (rootGetters) {
  axios.interceptors.request.use(config => {
    if (typeof config.headers.Authorization === 'undefined') {
      const token = rootGetters.user.token
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
    }
    return config
  })
}
