import 'regenerator-runtime/runtime'
import App from './components/App.vue'
import store from './store'

const appInfo = {
  name: 'Canary',
  id: 'canary',
  icon: 'info',
  isFileEditor: false,
  extensions: []
}

const routes = [
  {
    name: 'canary',
    path: '/',
    components: {
      app: App
    }
  }
]

const navItems = [
  {
    name: 'Canary',
    iconMaterial: appInfo.icon,
    route: {
      name: 'canary',
      path: `/${appInfo.id}/`
    },
    menu: 'user'
  }
]

export default {
  appInfo,
  store,
  routes,
  navItems
}
