<template>
<v-app>
    <v-navigation-drawer v-model="drawerLeft" width="500px" temporary app fixed>
        <div class="d-flex d-lg-none justify-end mr-2 mt-2">
            <v-btn fab color="secondary" small @click="drawerLeft=false">
                <v-icon>mdi-arrow-left</v-icon>
            </v-btn>
        </div>
        <MainNavigation v-model="navItems" :unconfirmedNum="unconfirmedNum"></MainNavigation>
    </v-navigation-drawer>
    <v-app-bar color="secondary" app dark>
        <v-app-bar-nav-icon @click.stop="drawerLeft = !drawerLeft"></v-app-bar-nav-icon>

        <v-toolbar-title>
            <router-link to="/" style="color: white" class="text-decoration-none">{{ ($root.pageTitle.length > 0 ? $root.pageTitle : "Paperless Office") }}</router-link>
        </v-toolbar-title>
        <v-progress-linear :active="$root.loading" :indeterminate="true" absolute bottom color="white"></v-progress-linear>
        <v-spacer></v-spacer>

        <div v-for="(item, name) in $root.actionBar" v-if="item.enabled" :key="name">
            <v-badge v-if="item.badge" top color="red darken-2" dot overlap offset-x="16" offset-y="16">
                <v-btn :title="item.title" icon @click.stop="item.click">
                    <v-icon>{{ (typeof item.icon == "function" ? item.icon() : item.icon) }}</v-icon>
                </v-btn>
            </v-badge>
            <v-btn :title="item.title" v-if="!item.badge" icon @click.stop="item.click">
                <v-icon>{{ (typeof item.icon == "function" ? item.icon() : item.icon) }}</v-icon>
            </v-btn>

        </div>
        <v-btn title="Go back" icon @click="$router.go(-1)">
            <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
        <v-btn title="Go forward" icon @click="$router.go(1)">
            <v-icon>mdi-arrow-right</v-icon>
        </v-btn>

    </v-app-bar>

    <v-navigation-drawer v-model="$root.drawerRight" width="500px" temporary right app fixed>
        <div class="d-flex d-lg-none justify-start ml-2 mt-2">
            <v-btn fab color="secondary" small @click="$root.drawerRight=false">
                <v-icon>mdi-arrow-right</v-icon>
            </v-btn>
        </div>
        <FilterOptions :filterTagsItems="filterTagsItems"></FilterOptions>
    </v-navigation-drawer>

    <v-main style="margin-top: 64px" class="pt-12 pb-12">
        <v-container>
            <router-view></router-view>
        </v-container>
    </v-main>

    <Footer></Footer>
    <div class="text-center ma-2">
        <v-snackbar v-model="$root.snackBar.open" timeout="3000">
            {{ $root.snackBar.text }}

            <template v-slot:action="{ attrs }">
                <v-btn :color="$root.snackBar.color" text v-bind="attrs" @click="$root.snackBar.open = false">
                    Close
                </v-btn>
            </template>
        </v-snackbar>
    </div>
</v-app>
</template>

<script>
import FilterOptions from "../components/FilterOptions"
import MainNavigation from "../components/MainNavigation"
import Footer from "../components/Footer"
import initSqlJs from "sql.js"

export default {
    components: {
        FilterOptions: FilterOptions,
        MainNavigation: MainNavigation,
        Footer: Footer,
    },
    data: () => ({
        drawerLeft: false,
        filterTagsItems: [],
        unconfirmedNum: 0,
        navItems: {
            navigation: [
                {
                    to: "/",
                    icon: "mdi-home",
                    title: "Home",
                },
                {
                    to: "/unconfirmed",
                    icon: "mdi-marker-check",
                    title: "Unconfirmed",
                },
            ],
            controls: [
                {
                    to: "/trash",
                    icon: "mdi-trash-can-outline",
                    title: "Trash"
                },
            ]
        },
    }),
    watch: {
        "$root.DB": "init"
    },
    mounted() {
        document.addEventListener("keydown", this.doOpen)
    },
    beforeDestroy() {
        document.removeEventListener("keydown", this.doOpen)
    },
    created() {
        initSqlJs({
            locateFile: () => `https://cdnjs.cloudflare.com/ajax/libs/sql.js/1.5.0/sql-wasm.wasm`
        })
        .then(SQL => {
            this.$http.get("db.sqlite", {
                    responseType: 'arraybuffer'
                })
                .then(resp => {
                    this.$root.DB = new SQL.Database(new Uint8Array(resp.data))
                    if(this.$root.DB != null) {
                        var res = this.$root.DB.exec("SELECT * FROM pdf_entries WHERE confirmed = false AND in_trash = false ORDER BY date DESC")
                        this.unconfirmedNum = this.$root.parseResult(res).length
                    }
                })
        })
        .catch(err => {
            console.log(err)
        })
    },
    methods: {
        init() {
            if(this.$root.DB != null) {
                var res = this.$root.DB.exec("SELECT * FROM tags")
                this.filterTagsItems = this.$root.parseResult(res).map(x => x.title)
            }
        },
        doOpen: function(e) {
            if (!(e.keyCode === 79 && (e.ctrlKey || e.metaKey))) {
                return
            }
            e.preventDefault()
            this.$router.push({ name: 'Upload', params: { triggered: true }})
        },
    }
}
</script>
