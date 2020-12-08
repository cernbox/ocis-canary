"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _canary = require("./client/canary");

var _axios = _interopRequireDefault(require("axios"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

// eslint-disable-next-line camelcase
var state = {
  config: null,
  error: ''
};
var getters = {
  config: function config(state) {
    return state.config;
  },
  error: function error(state) {
    return state.error;
  }
};
var actions = {
  // Used by ocis-web.
  loadConfig: function loadConfig(_ref, config) {
    var commit = _ref.commit;
    console.log("load config", config);
    commit('LOAD_CONFIG', config);
  },
  setCanary: function setCanary(_ref2, version) {
    var commit = _ref2.commit,
        dispatch = _ref2.dispatch,
        getters = _ref2.getters,
        rootGetters = _ref2.rootGetters;
    injectAuthToken(rootGetters);
    (0, _canary.Canary_SetCanary)({
      $domain: rootGetters.configuration.server,
      body: {
        version: version
      }
    }).then(function (response) {
      if (response.status === 200 || response.status === 201) {
        document.cookie = response.data.cookie + "=" + version + "; max-age=" + response.data.ttl + "; path=/";
        window.location.href = "/";
      } else {
        dispatch('showMessage', {
          title: 'Failed to set the version',
          desc: response.statusText,
          status: 'danger'
        }, {
          root: true
        });
      }
    })["catch"](function (error) {
      console.error(error);
      dispatch('showMessage', {
        title: 'Failed to set version',
        desc: error.message,
        status: 'danger'
      }, {
        root: true
      });
    });
  }
};
var mutations = {
  SET_MESSAGE: function SET_MESSAGE(state, payload) {
    state.error = payload;
  },
  LOAD_CONFIG: function LOAD_CONFIG(state, config) {
    state.config = config;
  }
};
var _default = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};
exports["default"] = _default;

function injectAuthToken(rootGetters) {
  _axios["default"].interceptors.request.use(function (config) {
    if (typeof config.headers.Authorization === 'undefined') {
      var token = rootGetters.user.token;

      if (token) {
        config.headers.Authorization = "Bearer ".concat(token);
      }
    }

    return config;
  });
}