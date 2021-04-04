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
        '$root.DB': 'init',
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
        this.init()
        this.$root.actionBar.toggleView.enabled = true
    },
    methods: {
        init: function() {
            this.loadSearchParams()
            if(this.$root.DB != null) {
                var result = this.$root.DB.exec("SELECT * FROM pdf_entries WHERE in_trash = true ORDER BY date DESC")
                this.trashItemsList = this.$root.parseResult(result).map(entry => {
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
                        tags: []
                    }
                })
                this.$root.pageTitle = "Trash"
                this.applyModel()
            }
        },
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
        actionButtons: function(item, index) {
            var ctx = this
            return [
                {
                    icon: "mdi-eye",
                    func: function() {
                        ctx.$router.push({ name: 'View', params: { 'id': `${item.ID}` }})
                    }
                },
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
