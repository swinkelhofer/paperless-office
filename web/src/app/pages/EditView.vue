<template>
<div>
    <v-row>
        <v-col cols="12">
            <v-card>
                <v-toolbar color="primary" dark>
                    <v-btn icon>
                        <v-icon>
                            mdi-information-outline
                        </v-icon>
                    </v-btn>
                    <v-toolbar-title class="pl-0">
                        Meta
                    </v-toolbar-title>
                </v-toolbar>
                <v-card-text>
                    <v-row>
                        <v-col cols="12" md="6">
                            <v-text-field v-model="value.title" prepend-icon="mdi-format-title" label="Title" color="primary"></v-text-field>

                            <v-autocomplete hide-selected cache-items hide-no-data hide-details v-model="value.tags" :items="tagItems" :item-text="tagText" :item-value="tagValue" label="Tags" item-color="primary" color="primary" prepend-icon="mdi-tag-multiple" chips small-chips multiple>
                                <template v-slot:selection="{ item, index, select, selected }">
                                    <v-chip :color="item.color" close @click:close="value.tags.splice(index, 1)" outlined>
                                        <span v-text="item.title"></span>
                                    </v-chip>
                                </template>
                                <template v-slot:append-outer>
                                    <v-btn icon @click="createTagDialog=true">
                                        <v-icon>mdi-tag-plus</v-icon>
                                    </v-btn>
                                </template>
                            </v-autocomplete>

                            <v-menu ref="menu" v-model="menu" :close-on-content-click="false" offset-y transition="scale-transition" offset-y :nudge-right="40" min-width="auto">
                                <template v-slot:activator="{ on, attrs }">
                                    <v-text-field v-model="computedDate" label="Date" persistent-hint prepend-icon="mdi-calendar" v-bind="attrs" v-on="on" color="primary"></v-text-field>
                                </template>
                                <v-date-picker v-model="value.date" color="primary" :title-date-format="formatDate" scrollable @input="menu = false"></v-date-picker>
                            </v-menu>

                        </v-col>
                        <v-col cols="12" md="6">
                            <v-expansion-panels popout>
                                <v-expansion-panel>
                                    <v-expansion-panel-header ripple>
                                        Detected Mail Adresses
                                    </v-expansion-panel-header>
                                    <v-expansion-panel-content>
                                        <v-text-field dense v-model="value.emailAdresses[index]" prepend-icon="mdi-email-outline" v-for="(mail, index) in value.emailAdresses" :key="'mail' + index" color="primary">
                                            <v-btn slot="append" icon @click="value.emailAdresses.splice(index, 1)">
                                                <v-icon>
                                                    mdi-trash-can-outline
                                                </v-icon>
                                            </v-btn>
                                        </v-text-field>
                                        <v-btn color="primary" dark small @click="value.emailAdresses.push('')">Add Mail Address</v-btn>
                                    </v-expansion-panel-content>
                                </v-expansion-panel>

                                <v-expansion-panel>
                                    <v-expansion-panel-header ripple>
                                        Detected URLs
                                    </v-expansion-panel-header>
                                    <v-expansion-panel-content>
                                        <v-text-field dense v-model="value.urls[index]" prepend-icon="mdi-web" v-for="(url, index) in value.urls" :key="'url' + index" color="primary">
                                            <v-btn slot="append" icon @click="value.urls.splice(index, 1)">
                                                <v-icon>
                                                    mdi-trash-can-outline
                                                </v-icon>
                                            </v-btn>
                                        </v-text-field>
                                        <v-btn color="primary" dark small @click="value.urls.push('')">Add URL</v-btn>

                                    </v-expansion-panel-content>
                                </v-expansion-panel>

                                <v-expansion-panel>
                                    <v-expansion-panel-header ripple>
                                        Detected Phone Numbers
                                    </v-expansion-panel-header>
                                    <v-expansion-panel-content>
                                        <v-text-field dense v-model="value.phoneNumbers[index]" prepend-icon="mdi-web" v-for="(phoneNumber, index) in value.phoneNumbers" :key="'phonenumber' + index" color="primary">
                                            <v-btn slot="append" icon @click="value.phoneNumbers.splice(index, 1)">
                                                <v-icon>
                                                    mdi-trash-can-outline
                                                </v-icon>
                                            </v-btn>
                                        </v-text-field>
                                        <v-btn color="primary" dark small @click="value.phoneNumbers.push('')">Add Phone Number</v-btn>
                                    </v-expansion-panel-content>
                                </v-expansion-panel>

                            </v-expansion-panels>
                        </v-col>

                    </v-row>
                </v-card-text>
            </v-card>

        </v-col>

        <v-col cols="12" md="6" v-if="! hasRescanned">
            <v-card class="pdfeditor">
                <v-responsive :aspect-ratio="3/4">
                    <v-toolbar color="primary" dark>
                        <v-text-field v-model="value.title" label="Title" color="white" prepend-icon="mdi-format-title" hide-details="auto"></v-text-field>
                    </v-toolbar>
                    <v-card-text>
                        <v-textarea v-model="value.content" filled prepend-icon="mdi-comment" color="primary"></v-textarea>
                    </v-card-text>
                </v-responsive>
            </v-card>
        </v-col>

        <v-col cols="12" md="6">
            <v-card>
                <v-toolbar color="primary" dark v-if="hasRescanned">
                    Original Scan
                </v-toolbar>
                <v-responsive :aspect-ratio="3/4">
                    <div id="pdf"></div>
                </v-responsive>
            </v-card>
        </v-col>

        <v-col cols="12" md="6" v-if="hasRescanned">
            <v-card>
                <v-toolbar color="primary" dark>
                    Rescanned Document
                    <v-spacer></v-spacer>

                    <v-btn icon @click="abortRescan">
                        <v-icon>mdi-trash-can-outline</v-icon>
                    </v-btn>
                    <v-btn icon @click="finalizeRescan">
                        <v-icon>mdi-check</v-icon>
                    </v-btn>
                </v-toolbar>
                <v-responsive :aspect-ratio="3/4">
                    <div id="rescanned_pdf"></div>
                </v-responsive>
            </v-card>
        </v-col>
    </v-row>

    <v-overlay :value="overlay">
        <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>

    <v-dialog v-model="createTagDialog" max-width="400">
        <v-card>
            <v-card-title class="headline">
                Create new tag
            </v-card-title>

            <v-card-text>
                <v-text-field v-model="newTag.title" label="Tag Name" color="primary"></v-text-field>
                <v-color-picker dot-size="15" hide-mode-switch mode="hexa" show-swatches swatches-max-height="100" v-model="newTag.color"></v-color-picker>
            </v-card-text>
            <v-card-actions fixed bottom>
                <v-spacer></v-spacer>

                <v-btn text @click="createTagDialog = $root.loading = false">
                    Cancel
                </v-btn>

                <v-btn color="primary" text @click="createTag" :disabled="newTag.title.length == 0">
                    Create
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>

</div>
</template>

<script>
import PDFObject from "pdfobject"
export default {
    props: ["value"],
    data() {
        return {
            menu: false,
            tagItems: [],
            createTagDialog: false,
            overlay: false,
            newTag: {
                title: "",
                color: "",
            },
        }
    },
    mounted() {
        document.addEventListener("keydown", this.doSave)
        this.$root.actionBar.save = {
            enabled: true,
            icon: 'mdi-content-save-outline',
            click: this.save,
        }

        this.$root.actionBar.rescan = {
            enabled: true,
            icon: 'mdi-ocr',
            click: this.triggerRescan,
        }

        this.$http.get("api/tags")
            .then(resp => {
                this.tagItems = resp.data
            })
    },
    beforeDestroy() {
        document.removeEventListener("keydown", this.doSave)
        this.$root.actionBar.rescan.enabled = false
        this.$root.actionBar.save.enabled = false
    },
    watch: {
        "value.filename": "initPDFViewer",
        "value.rescanFilename": "initRescanViewer",
    },
    computed: {
        computedDate: function () {
            if (this.value.date)
                return new Date(this.value.date).toISOString().substr(0, 10)
            else
                return null
        },
        hasRescanned: function () {
            return typeof this.value.rescanFilename !== typeof undefined && this.value.rescanFilename.length > 0
        }
    },
    methods: {
        doSave(e) {
            if (!(e.keyCode === 83 && (e.ctrlKey || e.metaKey))) {
                return
            }
            e.preventDefault()
            this.save()
        },
        initPDFViewer: function () {
            PDFObject.embed(`data/processed/${this.value.filename}`, "#pdf", {
                pdfOpenParams: {
                    view: 'FitV'
                },
            })
        },
        initRescanViewer: function () {
            var ctx = this
            setTimeout(function () {
                if (ctx.value.rescanFilename) {
                    PDFObject.embed(`data/processed/${ctx.value.rescanFilename}`, "#rescanned_pdf", {
                        pdfOpenParams: {
                            view: 'FitV'
                        },
                    })
                }
            }, 200)
        },
        triggerRescan: function () {
            this.$root.loading = true
            this.overlay = true
            var post = JSON.parse(JSON.stringify(this.value))
            post.date = new Date(post.date).toGMTString()
            post.emailAdresses = JSON.stringify(post.emailAdresses)
            post.urls = JSON.stringify(post.urls)
            post.phoneNumbers = JSON.stringify(post.phoneNumbers)
            this.$http.post('api/rescan', post)
                .then(resp => {
                    this.$set(this.value, 'rescanFilename', resp.data.rescanFilename)
                })
                .finally(() => {
                    this.$root.loading = false
                    this.overlay = false
                })
        },
        abortRescan: function () {
            this.$root.loading = true
            var post = JSON.parse(JSON.stringify(this.value))
            post.date = new Date(post.date).toGMTString()
            post.emailAdresses = JSON.stringify(post.emailAdresses)
            post.urls = JSON.stringify(post.urls)
            post.phoneNumbers = JSON.stringify(post.phoneNumbers)
            this.$http.post('api/rescan/abort', post)
                .then(resp => {
                    this.$set(this.value, 'rescanFilename', "")
                })
                .finally(() => {
                    this.$root.loading = false
                })
        },
        finalizeRescan: function () {
            this.$root.loading = true
            var post = JSON.parse(JSON.stringify(this.value))
            post.date = new Date(post.date).toGMTString()
            post.emailAdresses = JSON.stringify(post.emailAdresses)
            post.urls = JSON.stringify(post.urls)
            post.phoneNumbers = JSON.stringify(post.phoneNumbers)
            this.$http.post('api/rescan/finalize', post)
                .then(resp => {
                    this.$set(this.value, 'rescanFilename', "")
                    this.initPDFViewer()
                })
                .finally(() => {
                    this.$root.loading = false
                })
        },
        save: function () {
            var post = JSON.parse(JSON.stringify(this.value))
            post.date = new Date(post.date).toGMTString()
            post.emailAdresses = JSON.stringify(post.emailAdresses)
            post.urls = JSON.stringify(post.urls)
            post.phoneNumbers = JSON.stringify(post.phoneNumbers)
            post.confirmed = true
            this.$root.loading = true
            this.$http.post('api/documents', post)
                .then(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Successfully saved document",
                        color: "success"
                    }
                    this.$router.push({
                        name: "View",
                        params: {
                            id: this.$route.params.id
                        }
                    })
                })
                .catch(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Failed saving document",
                        color: "red"
                    }
                })
                .finally(() => {
                    this.$root.loading = false
                })
        },
        formatDate: function (dateString) {
            return new Date(dateString).toISOString().substr(0, 10)
        },
        tagText: function (tag) {
            return tag.title
        },
        tagValue: function (tag) {
            return tag
        },
        createTag: function () {
            this.$root.loading = true
            this.$http.put("api/tags", this.newTag)
                .then(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Successfully created tag",
                        color: "success"
                    }
                    this.tagItems.push(resp.data)
                })
                .catch(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Failed creating tag",
                        color: "red"
                    }
                })
                .finally(() => {
                    this.createTagDialog = false
                    this.$root.loading = false
                    this.newTag = {
                        title: "",
                        color: "",
                    }
                })
        }
    },
}
</script>

<style lang="scss">
.pdfobject-container {
    width: 100%;
    height: 100%;
}

.pdfeditor {
    .v-card__text {
        height: calc(100% - 64px);

        .v-input.v-textarea.theme--light.v-text-field.v-text-field--filled.v-text-field--is-booted.v-text-field--enclosed {
            height: 100%;

            .v-input__control {
                height: 100%;

                .v-input__slot {
                    height: 100%;
                    overflow: hidden;

                    .v-text-field__slot {
                        height: 100%;

                        textarea {
                            height: 100%;
                            font-size: 0.5rem;
                            font-family: "Roboto Mono";
                        }
                    }

                }
            }

        }
    }
}
</style>
