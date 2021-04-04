<template>
    <v-list-item link class="elevation-1 mb-2 py-2">
        <v-list-item-action>
            <v-btn fab color="primary" dark>
                <v-icon>mdi-file-pdf</v-icon>
            </v-btn>
        </v-list-item-action>
        <v-list-item-content>
            <v-list-item-title class="text-h6" v-if="value.title.length > 0" v-text="value.title"></v-list-item-title>
            <v-list-item-title class="text-h6" v-if="! value.title || value.title.length == 0" v-text="value.filename"></v-list-item-title>

            <v-list-item-subtitle>
                <v-chip v-for="(tag, index) in value.tags" :key="tag.title + index" v-if="index < 2" class="mt-2 mr-2" color="green" outlined :color="tag.color">
                    {{ tag.title }}
                </v-chip>
                <v-chip outlined v-if="value.tags.length > 2" class="mt-2 mr-2">
                    + {{ value.tags.length - 2 }} more tags...
                </v-chip>
                <v-chip outlined v-if="value.tags.length == 0" class="mt-2 mr-2">
                    No tags set
                </v-chip>
            </v-list-item-subtitle>
        </v-list-item-content>
        <v-chip color="primary" outlined class="mx-4">
            <v-icon class="mr-2">mdi-calendar-month</v-icon>{{ date(value.date) }}
        </v-chip>

        <v-list-item-action v-for="button in actionButtons" :key="button.icon">
            <v-btn fab small color="primary" dark @click="button.func">
                <v-icon v-text="button.icon"></v-icon>
            </v-btn>
        </v-list-item-action>

<!--
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
//-->
    </v-list-item>
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
<style lang="scss" scoped>
.v-list-item.v-list-item--link {
    border: 1px solid rgba(0,0,0,.1);
}
</style>