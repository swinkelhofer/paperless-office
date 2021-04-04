import Vue from 'vue'
import Router from 'vue-router'
import Home from '../../pages/Home'
import Unconfirmed from '../../pages/Unconfirmed'
import Confirm from '../../pages/Confirm'
import View from '../../pages/View'
import Trash from '../../pages/Trash'
import Upload from '../../pages/Upload'

const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home,
        },
        {
            path: '/unconfirmed',
            name: 'Unconfirmed',
            component: Unconfirmed,
        },
        {
            path: '/trash',
            name: 'Trash',
            component: Trash,
        },
        {
            path: '/upload',
            name: 'Upload',
            component: Upload,
        },
        {
            path: '/confirm/:id',
            name: 'Confirm',
            component: Confirm,
        },
        {
            path: '/view/:id',
            name: 'View',
            component: View,
        },
    ]
})
