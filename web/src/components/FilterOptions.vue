<template>
<v-container fluid>
    <v-subheader>FILTER</v-subheader>
    <v-col>
        <v-text-field v-model="$root.filter.title" label="Title" color="primary"></v-text-field>
        <v-text-field v-model="$root.filter.content" label="Volltext" color="primary"></v-text-field>
        <v-text-field v-model="$root.filter.mail" label="E-Mail-Adresse" color="primary"></v-text-field>
        <v-text-field v-model="$root.filter.url" label="URL" color="primary"></v-text-field>
        <v-menu v-model="fromDateToggle" :close-on-content-click="false" :nudge-right="40" transition="scale-transition" offset-y min-width="auto">
            <template v-slot:activator="{ on, attrs }">
                <v-text-field v-model="$root.filter.fromDate" label="From Date" readonly v-bind="attrs" v-on="on"></v-text-field>
            </template>
            <v-date-picker color="primary" v-model="$root.filter.fromDate" @input="fromDateToggle = false"></v-date-picker>
        </v-menu>

        <v-menu v-model="toDateToggle" :close-on-content-click="false" :nudge-right="40" transition="scale-transition" offset-y min-width="auto">
            <template v-slot:activator="{ on, attrs }">
                <v-text-field v-model="$root.filter.toDate" label="To Date" readonly v-bind="attrs" v-on="on"></v-text-field>
            </template>
            <v-date-picker color="primary" v-model="$root.filter.toDate" @input="toDateToggle = false"></v-date-picker>
        </v-menu>

        <v-autocomplete class="mb-5" hide-selected cache-items hide-no-data hide-details v-model="$root.filter.tags" :items="filterTagsItems" label="Tags" item-color="primary" color="primary" chips small-chips multiple>
            <template v-slot:selection="{ item, index, select, selected }">
                <v-chip color="primary" outlined close @click:close="$root.filter.tags.splice(index, 1)">
                    <span v-text="item"></span>
                </v-chip>
            </template>
        </v-autocomplete>

        <v-btn-toggle>
            <v-btn tile large color="primary" @click="$root.filter.apply" active-class="white--text">
                <v-icon left color="white">
                    mdi-filter
                </v-icon>
                Filter
            </v-btn>

            <v-btn tile large @click="$root.filter.clear" active-class="dark--text">
                <v-icon left>
                    mdi-close
                </v-icon>
                Clear
            </v-btn>
        </v-btn-toggle>
    </v-col>
</v-container>
</template>

<script>
export default {
    props: ["filterTagsItems"],
    data() {
        return {
            fromDateToggle: false,
            toDateToggle: false,
        }
    },
}
</script>
