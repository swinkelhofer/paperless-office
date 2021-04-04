<template>
<div>
    <v-row>
        <v-col v-if="$root.viewMode == 'list' && trashItemsList.length > 0" cols="12">
            <v-list class="py-0">
                <ListItem v-for="(item, index) in trashItems[page - 1]" :key="index" v-model="trashItems[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></ListItem>
            </v-list>
        </v-col>
        <v-col v-if="$root.viewMode == 'grid' && trashItems.length > 0" cols="12" md="6" lg="3" v-for="(item, index) in trashItems[page - 1]" :key="index">
            <CardItem v-model="trashItems[page - 1][index]" :actionButtons="actionButtons(item, index)" :config="itemsConfig"></CardItem>
        </v-col>

        <v-col cols="12" v-if="emptyState.enabled">
            <EmptyState v-model="emptyState"></EmptyState>
        </v-col>

        <v-col cols="12" v-if="! emptyState.enabled">
            <div class="text-center">
                <v-pagination v-model="page" :length="maxPage" :total-visible="7"></v-pagination>
            </div>
        </v-col>
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
            trashItems: [],
            trashItemsList: [],
            menu: [],
            maxPage: 0,
            page: 1,
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
    watch: {
        'page': 'storeSearchParams',
    },
    created() {
        this.$root.actionBar.emptyTrash = {
            enabled: true,
            icon: 'mdi-trash-can-outline',
            click: this.emptyTrash,
            badge: false,
            title: "Empty trash"
        }
        this.$root.pageTitle = "Trash"
        this.$root.actionBar.toggleView.enabled = true
    },
    beforeDestroy() {
        this.$root.actionBar.emptyTrash.enabled = false
        this.$root.actionBar.toggleView.enabled = false
        this.$root.pageTitle = ""
    },
    mounted() {
        this.loadSearchParams()
        this.$root.actionBar.toggleView.enabled = true
        this.$http.post("api/search", {
                inTrash: true
            })
            .then(resp => {
                this.trashItemsList = resp.data
                this.menu = resp.data.map(x => false)
                this.applyModel()
            })
            .finally(() => {
                this.$root.pageTitle = "Trash"
                this.emptyStateCalc()
            })
    },
    methods: {
        applyModel: function() {
            this.trashItems = []
            var items = this.trashItemsList.map(x => x)
            while (items.length) {
                this.trashItems.push(items.splice(0, 20))
            }
            this.maxPage = Math.ceil(this.trashItemsList.length / 20)
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
        emptyTrash: function() {
            for(var index = this.trashItems.length - 1; index >= 0; --index) {
                this.deleteDocument(this.trashItems[index], index)
            }
            this.applyModel()
        },
        actionButtons: function(item, index) {
            var ctx = this
            return [
                {
                    icon: "mdi-trash-can-outline",
                    func: function() {
                        ctx.deleteDocument(item, index)
                    }
                },
                {
                    icon: "mdi-restore",
                    func: function() {
                        ctx.restoreDocument(item, index)
                    }
                }
            ]
        },
        emptyStateCalc: function() {
            if(this.trashItemsList.length == 0) {
                this.emptyState = {
                    enabled: true,
                    icon: "mdi-trash-can-outline",
                    title: "Nothing In Trash",
                    subtitle: "It seems you didn't delete anything",
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
        deleteDocument: function (item, index) {
            this.$root.loading = true
            this.$http.delete("api/documents", {
                    data: item
                })
                .then(resp => {
                    this.trashItemsList.splice(index, 1)
                    this.$root.snackBar = {
                        open: true,
                        text: "Successfully deleted document permanently",
                        color: "success"
                    }
                })
                .catch(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Failed deleting document",
                        color: "red"
                    }
                })
                .finally(() => {
                    this.$root.loading = false
                    this.applyModel()
                })
        },
        restoreDocument: function (item, index) {
            this.$root.loading = true
            item.inTrash = false
            this.$http.post("api/documents", item)
                .then(resp => {
                    this.trashItemsList.splice(index, 1)
                    this.$root.snackBar = {
                        open: true,
                        text: "Successfully restored document",
                        color: "success"
                    }
                })
                .catch(resp => {
                    this.$root.snackBar = {
                        open: true,
                        text: "Failed restoring document",
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
