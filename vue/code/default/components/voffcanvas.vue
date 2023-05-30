<script>
    export default {
        props: {
            value: {
                type: Boolean,
                default: true
            },
            title: {
                type: String,
                required: false
            },
            fullpage: {
                type: Boolean,
                required: false
            },
            canclose: {
                type: Boolean,
                default: true
            },
            header: {
                type: Boolean,
                default: true
            },
            footer: {
                type: Boolean,
                required: false
            },
            placement:{
                default: "end"
            },
        },
        data: () => {
            return {}
        },
        methods: {
            tryclose() {
                this.$emit('input', false)
            },
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div>
        <div v-if="value" class="modal-backdrop fade show" :class="{'show': value}" @click.prevent="tryclose"></div>
        <div class="offcanvas offcanvas-auto" :class="(placement?'offcanvas-'+ placement:'')+ (value?' show':'') " tabindex="-1" aria-labelledby="offcanvasLabel">
            <div class="offcanvas-header">
                <slot name="header">
                    <h5 id="offcanvasLabel"> {{title}} </h5>
                </slot>                
                <button v-if="canclose" type="button" class="btn-close text-reset" @click.prevent="tryclose"
                    aria-label="Close"></button>
            </div>
            <div class="offcanvas-body">
                <slot></slot>
            </div>
            <div v-if="footer" class="offcanvas-footer">
                <slot name="footer">
                    TEst
                </slot>
            </div>
        </div>
    </div>
</template>