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
            category_id: {
                type: Number,
                default: 0
            },
            basecategory_id: {
                type: Number,
                default: 0
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
                Active : false,
                CategoryID : 0,
                BaseSubcategoryID : 0,
                BaseChildcategoryID : 0,
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
            CategoryIDError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.CategoryID) {
                    return "Category can not  be empty"
                }
            },
            BaseSubcategoryError: function () {
                if (!this.submitted || this.account.AccountType == 'admin') {
                    return false
                }
                if (!this.BaseSubcategoryID) {
                    return "Base Subcategory can not  be empty"
                }
            },
            BaseChildcategoryError: function () {
                if (!this.submitted || this.account.AccountType == 'admin') {
                    return false
                }
                if (!this.BaseChildcategoryID) {
                    return "Base Childcategory can not  be empty"
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
                    this.Active = this.value.Active
                    this.BaseSubcategoryID = this.value.BaseSubcategoryID
                    this.BaseChildcategoryID = this.value.BaseChildcategoryID
                    this.CategoryID = this.value.CategoryID
                }
                if(this.value.ID == 0){
                    this.Active = true
                    if(this.category_id){
                        this.CategoryID = this.category_id
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.CategoryIDError || this.BaseSubcategoryError || this.BaseChildcategoryError) { return }
                this.value.Name = this.Name
                this.value.Description = this.Description
                this.value.Logo = this.Logo
                this.value.Active = this.Active
                this.value.BaseSubcategoryID = this.BaseSubcategoryID
                this.value.BaseChildcategoryID = this.BaseChildcategoryID
                this.value.CategoryID = this.CategoryID
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
            <formitem name="inputCategoryID" label="Category" :error="CategoryIDError" v-model="CategoryID" type="select" 
            service="allcategories" displayby="Name" selectby="ID"  v-if="!category_id"  />
            <formitem v-if="account.AccountType != 'admin'" name="inputBaseSubcategoryID" label="Cubizy Subcategory" v-model="BaseSubcategoryID" type="select" 
            service="allsubcategories" :filter="{ CategoryID: basecategory_id }" displayby="Name" :error="BaseSubcategoryError"
            selectby="ID" />
            <formitem v-if="account.AccountType != 'admin'" name="inputBaseChildcategoryID" label="Cubizy Childcategory" v-model="BaseChildcategoryID" type="select" 
            service="allchildcategories" :filter="{ SubcategoryID: BaseSubcategoryID }" displayby="Name" :error="BaseChildcategoryError"
            selectby="ID" />
        </div>
        <div class="d-flex centered">
            <button type="submit" class="btn btn-success m-1" :disabled="loading">
                Save
            </button>
            <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
        </div>
    </form>
</template>