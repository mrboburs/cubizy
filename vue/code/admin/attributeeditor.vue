<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                    }
                }
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                FieldTypes: [ 'Text', 'Number', 'Dropdown' ],
                Name: "",
                FieldType: "",
                Options: "",
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.SetData()
                    this.$emit('onset', this.value)
                }
            },
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            NameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Name.trim()) {
                    return "Name can not  be empty"
                }
            },
            FieldTypeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.FieldType.trim()) {
                    return "Please select type of field"
                }
            },
            OptionsLabel: function(){
                if(this.FieldType == "Dropdown"){
                    return "Options"
                }else{
                    return "Default"
                }
            },
            OptionsError: function(){
                if (!this.submitted) {
                    return false
                }
                if(this.FieldType == "Dropdown"){
                    if (!this.Options.trim()) {
                        return "Please give options to show in dropdown"
                    }
                }
            },
            OptionsType: function(){
                switch (this.FieldType) {
                    case "Dropdown":
                        return "tags"
                    case "Number":
                        return "number"
                    default:
                    return "text"
                }
            },
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    if (this.value.Name) {
                        this.Name = this.value.Name
                        this.FieldType = this.value.FieldType
                        this.Options = this.value.Options
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.FieldTypeError || this.OptionsError) { return }
                this.value.Name = this.Name
                this.value.FieldType = this.FieldType
                this.value.Options = this.Options
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <form @submit.prevent="submit">
        <div class="row">
            <div class="col">
                <formitem name="inputName" label="Title" :error="NameError" v-model="Name" />
                <formitem name="inputFieldType" label="Field Type" :error="FieldTypeError" v-model="FieldType">
                    <div class="d-flex align-items-center justify-content-evenly m-2">
                        <div v-for="(_FieldType, index) in FieldTypes">
                            <input  type="radio" :id="'radio_fieldtype' + index" name="radio_fieldtype" :value="_FieldType" v-model="FieldType"/>
                            <label :for="'radio_fieldtype' + index">{{_FieldType}}</label>
                        </div>
                    </div>
                </formitem>
                <formitem name="inputOptions" :label="OptionsLabel" :error="OptionsError" v-model="Options" :type="OptionsType" />
                <div class="d-flex align-items-center end m-2">
                    <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                    <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
                </div>
            </div>
        </div>
    </form>
</template>