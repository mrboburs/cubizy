<script>
    export default {
        props: {
            label: {
                type: String,
                required: true
            },
            value: {},
            help: {},
            error: {},
            message: {},
            customLayout: {
                type: Boolean,
                default: false,
            },
            inputgroup: {
                type: Boolean,
                default: false,
            },
            suffix: {
                type: String,
                default: ""
            },
            labelclass: {
                type: String,
                default: "col-md-4 col-lg-3 text-md-end"
            },
            inputclass: {
                type: String,
                default: "col-md-8 col-lg-9"
            },
            name: {
                type: String,
                default: ""
            },
            type: {
                type: String,
                default: "",
            },
            service: {
                type: String,
                default: "",
            },
            placeholder: {
                type: String,
                default: "",
            },
            autocomplete: {
                type: Boolean,
                default: true,
            },
            submitted: {
                type: Boolean,
                default: false,
            },
            selectby: {
                type: String,
                default: "ID",
            },
            displayby: {
                type: String,
                default: "Title",
            },
            true_value: {
                default: true
            },
            false_value: {
                default: false
            },
            values: {
                type: Array,
                default: null,
            },
            prefix: {
                type: String,
                default: "account",
            },
            filter: {},
            nothing_selected_text: {
                type: String,
                default: "Nothing selected",
            },
            select_out_of_single: {
                type: Boolean,
                default: true,
            },
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue != this.internal_value) {
                    this.internal_value = newValue
                }
            },
            internal_value: function (newValue, oldValue) {
                if (newValue != this.value) {
                    this.$emit('input', newValue)
                }
            },
        },
        computed: {
            _inputclass() {
                var inputclass = ""
                if (this.customLayout) {
                    inputclass += this.inputclass
                }
                // if (this.inputgroup) {
                //     inputclass += ' input-group'
                // }
                return inputclass
            },
        },
        mounted() {
            this.internal_value = this.value
        },
        methods: {
            oninput(value) {
                this.$emit('input', value)
            },
        },
        data: () => {
            return {
                internal_value: false,
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="mb-2" :class="{ row:customLayout }">
        <label v-if="type != 'checkbox' && type != 'switch' && label" :for="name" class="form-label"
            :class="customLayout?labelclass:''">
            <slot name="label">
                {{label}}
            </slot>
        </label>
        <div v-if="(type == 'checkbox' || type == 'switch' || !label) && customLayout" class="form-label"
            :class="customLayout?labelclass:''"></div>
        <div :class="_inputclass">
            <slot>
                <div :class="{'input-group' : inputgroup, 'w-100' : !inputgroup}">
                    <input v-if="type =='readonly'" readonly :value="value" class="form-control"
                        :class="{ 'is-invalid': submitted && $v.location.Country.$error }" />
                    <div v-else-if="type == 'checkbox'" class="form-check">
                        <input class="form-check-input" type="checkbox" :name="name" :id="name" :true-value="true_value"
                            :false-value="false_value" v-model="internal_value">
                        <label class="form-check-label text-nowrap" :for="name">
                            {{label}}
                        </label>
                    </div>
                    <div v-else-if="type == 'switch'" class="form-check form-switch ml-2">
                        <input class="form-check-input" type="checkbox" :name="name" :id="name" :true-value="true_value"
                            :false-value="false_value" v-model="internal_value">
                        <label class="form-check-label text-nowrap" :for="name">
                            {{label}}
                        </label>
                    </div>
                    <textarea v-else-if="type == 'textarea'" class="form-control" rows="3" :placeholder="placeholder"
                        :autocomplete="autocomplete" @input="$emit('input', $event.target.value)"
                        style="margin-top: 0px;margin-bottom: 0px;min-height: 116px;"
                        :class="{ 'is-invalid': submitted && $v.location.Country.$error }" :value="value"> </textarea>
                    <v-taginput v-else-if="type =='tags'" :placeholder="placeholder" :value="value" @input="oninput"
                        :class="{ 'is-invalid': submitted && $v.location.Country.$error }" />
                    <v-select v-else-if="type =='select'" :service="service" :value="value"
                        @input="$emit('input', $event)" :selectby="selectby" :displayby="displayby" :filter="filter"
                        :values="values" :nothing_selected_text="nothing_selected_text"
                        :select_out_of_single="select_out_of_single" @onselect="$emit('onselect', $event)" />
                    <ImageFile v-else-if="type =='image'" :value="value" @input="$emit('input', $event)"
                        :prefix="prefix" maxWidth="100%" maxHeight="auto">
                    </ImageFile>
                    <anyfile v-else-if="type =='file'" :value="value" @input="$emit('input', $event)"
                        :prefix="prefix" maxWidth="100%" maxHeight="auto">
                    </anyfile>
                    <input v-else :id="name" :name="name" :value="value" :type="type" class="form-control"
                        :placeholder="placeholder" :autocomplete="autocomplete" step="any"
                        @input="$emit('input', $event.target.value)"
                        :class="{ 'is-invalid': submitted && $v.location.Country.$error }" />
                    <span v-if="suffix" class="input-group-text">{{suffix}}</span>
                </div>
            </slot>
            <div v-if="error" class="invalid-feedback" style="display: block;">{{error}}</div>
            <div v-if="message" class="valid-feedback" style="display: block;">{{message}}</div>
            <div v-if="help" class="form-text">{{help}}</div>
        </div>
    </div>
</template>