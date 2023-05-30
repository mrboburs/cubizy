<script>
    export default {
        props: {
            value: {},
        },
        data: () => {
            return {
                submitted: false,
                ThemeSettings: {
                    key: {
                        title: "property",
                        value: "value",
                        type: "text",
                    }
                }
            }
        },
        watch: {
            value : function (newValue, oldValue) {
                this.load()
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        methods: {
            load(){
                try {
                    this.ThemeSettings = JSON.parse(this.value)
                } catch (error) {
                    console.log("error while loading theme setting json")
                    console.log(error)
                }
            },
            submit(key, value) {
                try {
                    this.ThemeSettings[key].value = value
                    var ThemeSettingsJSON = JSON.stringify(this.ThemeSettings)
                    this.$emit('input', ThemeSettingsJSON)
                } catch (error) {
                    console.log("error while submiting theme setting json")
                    console.log(error)
                }
                
                
            },
        },
        mounted: function () {
            this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>
        <div class="card-body">
            <div class="container">
                <div  class="row" >
                    <div v-for="key in Object.keys(ThemeSettings)" class="col-12 col-md-6">
                        <formitem :customLayout="true" :name="'input'+key" :label="ThemeSettings[key].title" 
                        :value="ThemeSettings[key].value" @input="submit(key, $event)"
                        :type="ThemeSettings[key].type" prefix="website"/>
                    </div>
                </div>
            </div>
        </div>
</template>