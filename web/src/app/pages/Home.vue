<template>
    <v-row>
        <v-col v-if="$root.viewMode == 'list' && documents.length > 0" cols="12">
            <v-list class="py-0">
                <ListItem v-for="(item, index) in filteredDocuments[page - 1]" :key="index" v-model="filteredDocuments[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></ListItem>
            </v-list>
        </v-col>
        <v-col v-if="$root.viewMode == 'grid' && documents.length > 0" cols="12" md="6" lg="3" v-for="(item, index) in filteredDocuments[page - 1]" :key="index">
            <CardItem v-model="filteredDocuments[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></CardItem>
        </v-col>

        <v-col cols="12" v-if="emptyState.enabled">
            <EmptyState v-model="emptyState"></EmptyState>
        </v-col>

        <v-col cols="12" v-if="! emptyState.enabled">
            <div class="text-center">
                <v-pagination v-model="page" :length="maxPage" :total-visible="7"></v-pagination>
            </div>
        </v-col>

        <v-dialog v-model="deleteDialog" max-width="400">
            <v-card>
                <v-card-title class="headline">
                    Move to trash
                </v-card-title>
                <v-card-text>
                    Do you really want to move this document to trash?
                </v-card-text>
                <v-card-actions fixed bottom>
                    <v-spacer></v-spacer>

                    <v-btn text @click="deleteDialog = $root.loading = false">
                        No
                    </v-btn>

                    <v-btn color="primary" text @click="deleteDocument(); deleteDialog = false">
                        Yes
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-row>
</template>

<script>

import Router from 'vue-router'
import CardItem from '../../components/CardItem'
import ListItem from '../../components/ListItem'
import EmptyState from '../../components/EmptyState'
export default {
    components: {
        CardItem: CardItem,
        ListItem: ListItem,
        EmptyState: EmptyState,
    },
    data() {
        return {
            documents: [],
            filteredDocuments: [],
            menu: [],
            deleteDialog: false,
            deleteIndex: -1,
            deleteItem: {},
            page: 1,
            maxPage: 0,
            searchParams: null,
            emptyState: {
                enabled: false,
                icon: "",
                title: "",
                subtitle: "",
            },
            itemsConfig: {
                date: true,
                tags: true,
            }
        }
    },
    created() {
        this.$root.actionBar.filter = {
            enabled: true,
            icon: 'mdi-filter',
            click: () => {
                this.$root.drawerRight = !this.$root.drawerRight
            },
            badge: false,
            title: "Filter"
        }
        this.$root.filter.apply = this.applyFilter
        this.$root.filter.clear = this.clearFilter

        this.$root.actionBar.toggleView.enabled = true
    },
    beforeDestroy() {
        this.$root.actionBar.filter.enabled = false
        this.$root.actionBar.toggleView.enabled = false
    },
    mounted() {
        this.loadSearchParams()
        this.$root.actionBar.toggleView.enabled = true
        this.$root.loading = true
        this.$http.get("api/documents")
            .then(resp => {
                for (var index = 0; index < resp.data.length; ++index)
                    try {
                        resp.data[index].tags = JSON.parse(resp.data[index].tags)
                    } catch (e) {}
                var docs = resp.data
                this.menu = docs.map(x => false)
                this.documents = docs
                this.applyFilter()
            })
            .finally(() => {
                this.$root.loading = false
            })
    },
    watch: {
        '$root.filter': 'applyFilter',
        'page': 'storeSearchParams',
    },
    methods: {
        actionButtons: function(item, index) {
            var ctx = this
            return [
                {
                    icon: "mdi-trash-can-outline",
                    func: function() {
                        ctx.deleteItem = item
                        ctx.deleteIndex = index
                        ctx.deleteDialog = true
                    }
                },
                {
                    icon: "mdi-eye",
                    func: function() {
                        ctx.$router.push({ name: 'View', params: { 'id': `${item.ID}` }})
                    }
                }
            ]
        },
        loadSearchParams: function() {
            this.searchParams = this.$route.query
            for(var key in this.$root.filter) {
                if(this.searchParams[key]) {
                    this.$root.filter[key] = JSON.parse(this.searchParams[key])
                }
            }
            if(this.searchParams["p"])
                this.page = JSON.parse(this.searchParams["p"])
        },
        storeSearchParams: function() {
            this.searchParams = {}
            for(var key in this.$root.filter) {
                if(typeof this.$root.filter[key] != "function" && this.$root.filter[key].length > 0) {
                    this.searchParams[key] = JSON.stringify(this.$root.filter[key])
                }
            }
            this.searchParams["p"] = JSON.stringify(this.page)
            var route = {
                        name: this.$route.name,
                        params: this.$route.params,
                        query: this.searchParams
                    }
            this.$router.push(route)
        },
        emptyStateCalc: function() {
            if(this.documents.length == 0) {
                this.emptyState = {
                    enabled: true,
                    icon: "mdi-bed-empty",
                    title: "NO DOCUMENTS FOUND",
                    subtitle: "It seems you didn't start your paperless office",
                }
            } else if (this.filteredDocuments.length == 0) {
                this.emptyState = {
                    enabled: true,
                    icon: "mdi-feature-search",
                    title: "NO DOCUMENTS FOUND",
                    subtitle: "Your search filters don't match any documents",
                }
            } else {
                this.emptyState = {
                    enabled: false,
                    icon: "",
                    title: "",
                    subtitle: "",
                }
            }
        },
        deleteDocument: function () {
            var item = this.deleteItem
            item.inTrash = true
            var index = this.deleteIndex
            this.$root.loading = true
            this.$http.post("api/documents", item)
                .then(resp => {
                    this.documents.splice(index, 1)
                    this.$root.snackBar = {
                        open: true,
                        text: "Successfully moved document to trash",
                        color: "success"
                    }
                })
                .catch(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Failed moving document to trash",
                        color: "red"
                    }
                })
                .finally(() => {
                    this.$root.loading = false
                    this.applyFilter()
                })
        },
        clearFilter: function() {
            this.$root.filter.content = ""
            this.$root.filter.title = ""
            this.$root.filter.mail = ""
            this.$root.filter.url = ""
            this.$root.filter.fromDate = ""
            this.$root.filter.toDate = ""
            this.$root.filter.tags = []
            this.storeSearchParams()
            this.applyFilter()
        },
        applyFilter: function (docs) {
            this.$root.loading = true
            var filteredDocs = JSON.parse(JSON.stringify(this.documents))

            this.$root.actionBar.filter.badge = false

            if (this.$root.filter.title.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => document.title.toLowerCase().includes(this.$root.filter.title.toLowerCase()))
            }

            if (this.$root.filter.content.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => document.content.toLowerCase().includes(this.$root.filter.content.toLowerCase()))
            }

            if (this.$root.filter.mail.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => document.emailAdresses.toLowerCase().includes(this.$root.filter.mail.toLowerCase()))
            }

            if (this.$root.filter.url.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => document.urls.toLowerCase().includes(this.$root.filter.url.toLowerCase()))
            }

            if (this.$root.filter.fromDate.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => (new Date(document.date)).getTime() >= (new Date(this.$root.filter.fromDate)).getTime())
            }

            if (this.$root.filter.toDate.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => (new Date(document.date)).getTime() <= (new Date(this.$root.filter.toDate)).getTime())
            }

            if (this.$root.filter.tags.length > 0) {
                this.$root.actionBar.filter.badge = true
                filteredDocs = filteredDocs.filter(document => {
                    var tags = document.tags.map(x => x.title)
                    return this.$root.filter.tags.every(i => tags.includes(i))
                })
            }
            this.$root.loading = false
            this.filteredDocuments = []
            this.maxPage = Math.ceil(filteredDocs.length / 20)
            if(this.maxPage < 2)
                this.page = 1
            this.storeSearchParams()
            while (filteredDocs.length) {
                this.filteredDocuments.push(filteredDocs.splice(0, 20));
            }
            this.emptyStateCalc()
        }
    },
}
</script>

<style lang="scss">
.empty-state {
    height: calc(100vh - 118px);
    margin-top: -82px;
    justify-content: center;
    align-items: center;
}
</style>
