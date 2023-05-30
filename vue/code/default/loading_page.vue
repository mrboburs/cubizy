<template>
    <b-loading :is-full-page="isFullPage" v-model="isLoading" :can-cancel="true">
        <b-icon
            pack="fas"
            icon="sync-alt"
            size="is-large"
            custom-class="fa-spin">
        </b-icon>
    </b-loading>
</template>
<script>
    export default {
        data() {
            return {
                isLoading: true,
                isFullPage: true
            }
        },
        methods: {
            openLoading() {
                this.isLoading = true
                setTimeout(() => {
                    this.isLoading = false
                }, 10 * 1000)
            }
        }
    }
</script>