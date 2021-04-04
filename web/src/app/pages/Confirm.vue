<template>
<div>
    <EditView v-model="model"></EditView>
</div>
</template>

<script>
import EditView from './EditView'
export default {
    components: {
        EditView: EditView,
    },
    data() {
        return {
            model: {},
        }
    },
    mounted() {
        this.init()
    },
    watch: {
        '$route': 'init'
    },
    methods: {
        init: function () {
            this.$http.post("api/search", {
                    ID: this.$route.params.id
                })
                .then(resp => {
                    if (resp.data.length > 0) {
                        this.model = resp.data[0]
                        this.model.emailAdresses = JSON.parse(this.model.emailAdresses)
                        this.model.urls = JSON.parse(this.model.urls)
                        this.model.phoneNumbers = JSON.parse(this.model.phoneNumbers)
                    }
                })
        }
    },
}
</script>
