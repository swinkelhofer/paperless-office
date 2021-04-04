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
        '$root.DB': 'init',
    },
    mounted() {
        this.init()
        this.$root.actionBar.toggleView.enabled = true
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
        init: function() {
            this.loadSearchParams()
            if(this.$root.DB != null) {
                var result = this.$root.DB.exec("SELECT * FROM pdf_entries WHERE confirmed = false AND in_trash = false ORDER BY date DESC")
                this.unconfirmedItemsList = this.$root.parseResult(result).map(entry => {
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
                    }
                })
                this.$root.pageTitle = "Unconfirmed"
                this.menu = this.unconfirmedItems.map(x => false)
                this.applyModel()
            }
        },
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
