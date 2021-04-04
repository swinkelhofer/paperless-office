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
        '$route': 'init',
        '$root.DB': 'init',
    },
    methods: {
        init: function () {
            if(this.$root.DB != null) {
                var result = this.$root.DB.exec("SELECT * FROM pdf_tags")
                var pdf_tag_preload = this.$root.parseResult(result)

                var result = this.$root.DB.exec(`SELECT * FROM pdf_entries WHERE id = ${this.$route.params.id}`)
                var documents = this.$root.parseResult(result).map(entry => {
                    var matching_tags = pdf_tag_preload.filter(x => x.pdf_entry_id == entry.id)

                    var tags = []
                    for(var matching_tag of matching_tags) {
                        var tag_result = this.$root.DB.exec(`SELECT * FROM tags WHERE id = ${matching_tag.tag_id}`)
                        var tag = this.$root.parseResult(tag_result)
                        if(tag.length > 0)
                            tags.push(tag[0])
                    }

                    return {
                        ID: entry.id,
                        content: entry.content,
                        title: entry.title,
                        confirmed: entry.confirmed,
                        filename: entry.filename,
                        date: entry.date,
                        emailAdresses: entry.e_mail_adresses,
                        inTrash: entry.in_trash,
                        phoneNumbers: entry.phoneNumbers,
                        urls: entry.urls,
                        tags: tags
                    }
                })
                if(documents.length > 0) {
                    this.model = documents[0]
                    try {
                        this.model.emailAdresses = JSON.parse(this.model.emailAdresses)
                    } catch(e) {}
                    try {
                        this.model.urls = JSON.parse(this.model.urls)
                    } catch(e) {}
                    try {
                        this.model.phoneNumbers = JSON.parse(this.model.phoneNumbers)
                    } catch(e) {}
                }
            }
        }
    },
}
</script>
