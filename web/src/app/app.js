import Vue from 'vue'
import VueCookies from 'vue-cookies'
import vuetify from './plugins/vuetify' // path to vuetify export
import App from './App'
import router from './plugins/router'
import axios from 'axios'

Vue.prototype.$http = axios;
Vue.use(VueCookies)

var ctx = new Vue({
    el: '#app',
    router: router,
    template: '<App/>',
    vuetify: vuetify,
    data: {
        actionBar: {
            save: {},
            filter: {},
            emptyTrash: {},
            toggleView: {
                enabled: false,
                icon: function() {
                    if(Vue.$cookies.get('viewMode') == "list") {
                        return "mdi-format-list-bulleted"
                    } else {
                        return "mdi-view-grid"
                    }
                },
                click: function() {
                    if(Vue.$cookies.get('viewMode') == "list") {
                        Vue.$cookies.set('viewMode','grid')
                        ctx.viewMode = "grid"
                        this.icon = "mdi-view-grid"
                    } else {
                        Vue.$cookies.set('viewMode','list')
                        ctx.viewMode = "list"
                        this.icon = "mdi-format-list-bulleted"
                    }
                },
                badge: false,
                title: "Toggle View"
            },
        },
        snackBar: {
            open: false,
            text: "",
            color: "danger"
        },
        pageTitle: "",
        staticFilesPathPrefix: "data/processed/",
        loading: false,
        drawerRight: false,
        viewMode: function() {
            return (Vue.$cookies.get('viewMode') ? Vue.$cookies.get('viewMode') : "grid")
        }(),
        filter: {
            tags: [],
            mail: "",
            url: "",
            content: "",
            title: "",
            fromDate: "",
            toDate: "",
            apply: () => {},
            clear: () => {},
            active: false,
        },
    },
    components: {
        "App": App,
    }
})
