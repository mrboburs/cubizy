<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Name: "",
                        Status: true
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

                Name : "",
                Description: "",
                Logo : "",
                Icon : "",
                Active : false,
                Top : 0,
                Featured : 0,
                BaseCategoryID : 0,
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
            prefix: function () {
                return 'categories'
            },
            NameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Name.trim()) {
                    return "Name can not  be empty"
                }
            },
            BaseCategoryError: function () {
                if (!this.submitted || this.account.AccountType == 'admin') {
                    return false
                }
                if (!this.BaseCategoryID) {
                    return "BaseCategory can not  be empty"
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
                    this.Name = this.value.Name
                    this.Description = this.value.Description
                    this.Logo = this.value.Logo
                    this.Icon = this.value.Icon
                    this.Active = this.value.Active
                    this.Top = this.value.Top
                    this.Featured = this.value.Featured
                    this.BaseCategoryID = this.value.BaseCategoryID
                }
                if(this.value.ID == 0){
                    this.Active = true
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.BaseCategoryError) { return }
                this.value.Name = this.Name
                this.value.Description = this.Description
                this.value.Logo = this.Logo
                this.value.Icon = this.Icon
                this.value.Active = this.Active
                this.value.Top = this.Top
                this.value.Featured = this.Featured
                this.value.BaseCategoryID = this.BaseCategoryID
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload')
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <form @submit.prevent="submit" class="row">
        <div class="col-auto">
            <formitem name="inputLogo" label="Logo" type="image" :prefix="prefix" v-model="Logo" style="max-width: 300px;"/>
            <formitem v-if="account.AccountType == 'admin'" name="inputIcon" label="Icon" v-model="Icon"/>
            <formitem name="inputActive" label="Status" v-model="Active">
                <div class="form-check form-switch ml-2">
                    <input class="form-check-input" type="checkbox" id="inputIsSuperAdmin" v-model="Active">
                    <label class="form-check-label" for="inputIsSuperAdmin">
                        <span v-if="Active">Active</span>
                        <span v-if="!Active">Not Active</span>
                    </label>
                </div>
            </formitem>
        </div>
        <div class="col">
            <formitem name="inputName" label="Name" :error="NameError" v-model="Name"/>
            <formitem name="inputDescription" label="Description" v-model="Description" type="textarea"/>
            <formitem v-if="account.AccountType != 'admin'" name="inputBaseCategoryID" label="Base Category" v-model="BaseCategoryID" type="select" 
            service="allcategories" :filter="{ BaseCategoryID: 0 }" displayby="Name" :error="BaseCategoryError" 
            selectby="ID" />
            <formitem name="inputTop" label="Top" type="number" v-model="Top" :customLayout="true"/>
            <formitem name="inputFeatured" label="Featured" type="number" v-model="Featured" :customLayout="true"/>
        </div>
        <div class="d-flex centered">
            <button type="submit" class="btn btn-success m-1" :disabled="loading">
                <b-spinner small v-if="loading"></b-spinner>
                Save
            </button>
            <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
        </div>
    </form>
</template>