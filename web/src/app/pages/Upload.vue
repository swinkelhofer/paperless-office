<template>
<v-card rounded>
    <v-toolbar color="primary" dark>
        <v-btn icon>
            <v-icon>mdi-upload</v-icon>
        </v-btn>
        <v-toolbar-title>Upload Documents</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
        <v-file-input ref="upload" v-model="files" color="primary" counter truncate-length="30" label="Select Documents" multiple prepend-icon="mdi-upload" :show-size="1000">
            <template v-slot:selection="{ index, text }">
                <v-chip color="secondary" dark>
                    {{ text }}
                </v-chip>
            </template>
        </v-file-input>
        <div class="pa-4"></div>
        <div class="d-flex justify-end">
            <v-btn ref="submit" color="primary" large dark @click="upload">
                Submit
            </v-btn>
        </div>
    </v-card-text>
</v-card>
</template>

<script>
export default {
    data: () => ({
        files: [],
    }),
    created() {
        this.$root.pageTitle = "Upload Documents"
    },
    mounted() {
        document.addEventListener("keydown", this.keydownCallback)
        if(this.$route.params.triggered)
            this.open()
    },
    beforeDestroy() {
        this.$root.pageTitle = ""
        document.removeEventListener("keydown", this.keydownCallback)
    },
    watch: {
        '$route.params.triggered': 'open',
    },
    methods: {
        keydownCallback: function(e) {
            if (e.keyCode === 79 && (e.ctrlKey || e.metaKey)) {
                e.preventDefault()
                this.open()
            } else if (e.keyCode === 13) {
                e.preventDefault()
                this.$refs.submit.$refs.link.click()
            }
        },
        open: function() {
            this.$refs.upload.$refs.input.click()
        },
        upload: function() {
            var uploadFile = new Object()
            var ctx = this

            for(var i = this.files.length - 1; i >= 0; i--) {
                var reader = new FileReader()
                reader.onload = function(fileData) {
                    console.log(fileData)
                    uploadFile.fileData = fileData.target.result

                    ctx.$http.post("/api/upload", uploadFile)
                    .then(resp => {
                        ctx.files.splice(i, 1)
                        ctx.$root.snackBar = {
                            open: true,
                            text: "Successfully uploaded " + fileData.target.name,
                            color: "success"
                        }
                    })
                    .catch(resp => {
                        ctx.$root.snackBar = {
                            open: true,
                            text: "Failed to upload " + fileData.target.name + ". " + resp,
                            color: "red"
                        }
                    })
                }

                uploadFile.name = ctx.files[i].name
                uploadFile.date = ctx.files[i].lastModified
                uploadFile.size = ctx.files[i].size
                uploadFile.type = ctx.files[i].type
                reader.readAsDataURL(ctx.files[i])
            }
        }
    },
}
</script>
