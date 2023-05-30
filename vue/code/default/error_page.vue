<template>
    <h1>Somthing went wrong.</h1>
</template>
<script>
    export default {
        data() {
            return {
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