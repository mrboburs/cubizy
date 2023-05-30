<script>
    export default {
        props: {
            value: {},
            error: {
                type: Boolean,
                required: false
            },
            hide_lable: {
                type: Boolean,
                required: false
            },
            compact: {
                type: Boolean,
                required: false
            },
        },
        data: () => {
            return {}
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div v-if="value" class="position-relative d-flex align-items-center alert-dismissible fade show"
        :class="{'alert-success': !error, 'alert-danger': error, 'alert': !compact }" role="alert">
        <strong v-if="!error && !hide_lable">Success: </strong>
        <strong v-if="error && !hide_lable">Error: </strong>
        <span class="ms-1" v-html="value"></span>
        <button type="button" class="btn-close" @click.prevent="$emit('input', false);" aria-label="Close"></button>
    </div>
</template>