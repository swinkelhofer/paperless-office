<template>
    <v-responsive class="empty-state">
        <div class="text-center mb-5">
            <v-icon color="grey" class="text-h2">{{ value.icon }}</v-icon>
        </div>
        <div class="text-h4 text-uppercase text-center mb-1 transition-swing grey--text">{{ value.title }}</div>
        <div class="text-h5 text-center grey--text">{{ value.subtitle }}</div>
    </v-responsive>
</template>

<script>
export default {
    props: ["value"]
}
</script>
