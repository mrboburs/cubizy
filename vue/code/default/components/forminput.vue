<script>
    export default {
        props: {
            value : {},
            help: {},
            error : {},
            message: {},
            customLayout: {
                type: Boolean,
                default: false,
            },
            inputgroup: {
                type: Boolean,
                default: false,
            },
            labelclass: {
                type: String,
                default: "col-sm-2 col-lg-4"
            },
            inputclass: {
                type: String,
                default: "col-sm-10 col-lg-8"
            },
            name: {
                type: String,
                default: ""
            },
            type: {
                type: String,
                default : "",
            },
            placeholder:{
                type: String,
                default : "",
            },
            autocomplete:{
                type: Boolean,
                default: true,
            },
            submitted:{
                type: Boolean,
                default: false,
            }
        },
        computed: {
            _inputclass() {
                var inputclass = ""
                if(this.customLayout){
                    inputclass += this.inputclass
                }
                if(this.inputgroup){
                    inputclass += ' input-group'
                }
                return inputclass
            },
        },
        data: () => {
            return {
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
        <div :class="_inputclass">
            <slot>
                <input :id="name" :name="name" :value="value" :type="type" class="form-control"
                            :placeholder="placeholder" :autocomplete="autocomplete" step="any"
                            @input="$emit('input', $event.target.value)"
                            :class="{ 'is-invalid': submitted && $v.location.Country.$error }" />
            </slot>
            <div v-if="error" class="invalid-feedback" style="display: block;">{{error}}</div>
            <div v-if="message" class="valid-feedback" style="display: block;">{{message}}</div>
            <div v-if="help" class="form-text">{{help}}</div>
        </div>
</template>