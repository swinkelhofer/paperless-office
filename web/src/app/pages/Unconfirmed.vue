<template>
<div>
    <v-row>
        <v-col v-if="$root.viewMode == 'list' && unconfirmedItemsList.length > 0" cols="12">
            <v-list class="py-0">
                <ListItem v-for="(item, index) in unconfirmedItems[page - 1]" :key="index" v-model="unconfirmedItems[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></ListItem>
            </v-list>
        </v-col>
        <v-col v-if="$root.viewMode == 'grid' && unconfirmedItemsList.length > 0" cols="12" md="6" lg="3" v-for="(item, index) in unconfirmedItems[page - 1]" :key="index">
            <CardItem v-model="unconfirmedItems[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></CardItem>
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
</div>
</template>

<script>
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
            unconfirmedItems: [],
            unconfirmedItemsList: [],
            menu: [],
            deleteDialog: false,
            deleteIndex: -1,
            deleteItem: {},
            maxPage: 0,
            page: 1,
            emptyState: {
                enabled: false,
                icon: "",
                title: "",
                subtitle: "",
            },
            itemsConfig: {
                date: false,
                tags: false,
            }
        }
    },
    watch: {
        'page': 'storeSearchParams',
    },
    mounted() {
        this.loadSearchParams()
        this.$root.actionBar.toggleView.enabled = true
        this.$http.post("api/search", {
                unconfirmed: true,
                inTrash: false,
            })
            .then(resp => {
                this.unconfirmedItemsList = resp.data
                this.menu = resp.data.map(x => false)
                this.applyModel()
            })
            .finally(() => {
                this.$root.pageTitle = "Unconfirmed"
            })
    },
    created() {
        this.$root.pageTitle = "Unconfirmed"
        this.$root.actionBar.toggleView.enabled = true
    },
    beforeDestroy() {
        this.$root.pageTitle = ""
        this.$root.actionBar.toggleView.enabled = false
    },
    methods: {
        applyModel: function() {
            this.unconfirmedItems = []
            var items = [...this.unconfirmedItemsList]
            while (items.length) {
                this.unconfirmedItems.push(items.splice(0, 20))
            }
            this.maxPage = Math.ceil(this.unconfirmedItemsList.length / 20)
                if(this.maxPage < 2)
                    this.page = 1
            this.storeSearchParams()
            this.emptyStateCalc()
        },
        loadSearchParams: function() {
            this.searchParams = this.$route.query
            if(this.searchParams["p"])
                this.page = JSON.parse(this.searchParams["p"])
        },
        storeSearchParams: function() {
            this.searchParams = {}
            this.searchParams["p"] = JSON.stringify(this.page)
            var route = {
                        name: this.$route.name,
                        params: this.$route.params,
                        query: this.searchParams
                    }
            this.$router.push(route)
        },
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
        emptyStateCalc: function() {
            if(this.unconfirmedItemsList.length == 0) {
                this.emptyState = {
                    enabled: true,
                    icon: "mdi-check-circle",
                    title: "Nothing Left To Do",
                    subtitle: "It seems you confirmed all documents",
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
                    this.unconfirmedItemsList.splice(index, 1)
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
                    this.applyModel()
                })
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
