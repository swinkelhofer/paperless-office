import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

const opts = {
    theme: {
        themes: {
            light: {
                primary: '#4C7337',
                secondary: '#02394A',
                accent: '#B2B1CF',
                error: '#885053',
            },
        },
    },
}

export default new Vuetify(opts)