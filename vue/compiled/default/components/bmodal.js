
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
            size:{
                default: "lg"
            }
        },
        data: () => {
            return {}
        },
        methods: {
            tryclose() {
                this.$emit('input', false)
            },
        },
        template: `
    <div v-if="value" class="modal fade show"  tabindex="-1" aria-labelledby="Label"
        aria-hidden="true" :data-bs-backdrop="canclose? '': 'static'" :data-bs-keyboard="canclose"
        style="display: block;">
        <div class="modal-dialog modal-dialog-scrollable" :class="size?'modal-'+ size:''">
            <div class="modal-content shadow-lg">
                <div v-if="header" class="modal-header">
                    <slot name="header">
                        <h5 class="modal-title">{{title}}</h5>
                    </slot>
                    <button v-if="canclose" type="button" class="btn-close" @click.prevent="tryclose"
                        aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <slot></slot>
                </div>
                <div v-if="footer" class="modal-footer">
                    <slot name="footer">
                        TEst
                    </slot>
                </div>
            </div>
        </div>
    </div>
`
    }
