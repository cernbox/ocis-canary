"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

require("regenerator-runtime/runtime");

var _App = _interopRequireDefault(require("./components/App.vue"));

var _store = _interopRequireDefault(require("./store"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var appInfo = {
  name: 'Canary',
  id: 'canary',
  icon: 'info',
  isFileEditor: false,
  extensions: []
};
var routes = [{
  name: 'canary',
  path: '/',
  components: {
    app: _App["default"]
  }
}];
var navItems = [{
  name: 'Canary',
  iconMaterial: appInfo.icon,
  route: {
    name: 'canary',
    path: "/".concat(appInfo.id, "/")
  },
  menu: 'user'
}];
var _default = {
  appInfo: appInfo,
  store: _store["default"],
  routes: routes,
  navItems: navItems
};
exports["default"] = _default;