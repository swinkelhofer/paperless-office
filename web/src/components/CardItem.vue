<template>
    <v-card rounded>
        <v-toolbar color="primary" dark>
            <v-btn icon>
                <v-icon>mdi-file-pdf</v-icon>
            </v-btn>
            <v-toolbar-title v-if="value.title.length > 0" v-text="value.title"></v-toolbar-title>
            <v-toolbar-title v-if="! value.title || value.title.length == 0" v-text="value.filename"></v-toolbar-title>

            <v-spacer></v-spacer>
            <v-btn icon @click="button.func" v-for="button in actionButtons" :key="button.icon">
                <v-icon v-text="button.icon"></v-icon>
            </v-btn>
        </v-toolbar>
        <v-card-text>
            <v-img :aspect-ratio="1/1" :src="$root.staticFilesPathPrefix + value.filename.replace('.pdf', '.png')"></v-img>
        </v-card-text>
        <v-divider class="mx-4" v-if="config.date || config.tags"></v-divider>
        <v-card-actions v-if="config.tags">
            <v-chip v-for="(tag, index) in value.tags" :key="tag.title + index" v-if="index < 2" class="mx-2 mt-2" color="green" outlined :color="tag.color">
                {{ tag.title }}
            </v-chip>
            <v-chip outlined v-if="value.tags.length > 2" class="mx-2 mt-2">
                + {{ value.tags.length - 2 }} more tags...
            </v-chip>
            <v-chip color="#cccccc" outlined v-if="value.tags.length == 0" class="mx-2 mt-2">
                No tags set
            </v-chip>
        </v-card-actions>

        <v-card-actions class="pt-0 pb-3" v-if="config.date">
            <v-chip color="primary" class="mx-2">
                <v-icon class="mr-2">mdi-calendar-month</v-icon>{{ date(value.date) }}
            </v-chip>
        </v-card-actions>
    </v-card>
</template>
<script>
const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
export default {
    props: ["value", "actionButtons", "config"],
    methods: {
        date: function(date) {
            if(date) {
                var d = new Date(date)
                return [monthNames[d.getMonth()], d.getFullYear()].join(' ')
            }
            return ""
        },
    },
}
</script>