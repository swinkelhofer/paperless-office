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
                                    <v-chip :color="item.color"  outlined>
                                        <span v-text="item.title"></span>
                                    </v-chip>
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

        <v-col cols="12" md="6">
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
                <v-responsive :aspect-ratio="3/4">
                    <div id="pdf"></div>
                </v-responsive>
            </v-card>
        </v-col>

    </v-row>
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
        this.init()
    },
    watch: {
        "value.filename": "initPDFViewer",
        "value.rescanFilename": "initRescanViewer",
        "$root.DB": "init"
    },
    computed: {
        computedDate: function () {
            if (this.value.date)
                return new Date(this.value.date).toISOString().substr(0, 10)
            else
                return null
        },
    },
    methods: {
        init() {
            if(this.$root.DB != null) {
                var res = this.$root.DB.exec("SELECT * FROM tags")
                this.tagItems = this.$root.parseResult(res)
            }
        },
        initPDFViewer: function () {
            PDFObject.embed(`${this.value.filename}`, "#pdf", {
                pdfOpenParams: {
                    view: 'FitV'
                },
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
