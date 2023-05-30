<script>
    export default {
        components: {
            'attributeseditor': () => import("/vue/products/attributeseditor.js"),
            'variationeditor': () => import("/vue/products/variationeditor.js"),
        },
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Variation: "",
                        CategoryID: 0,
                        SubcategoryID: 0,
                        ChildcategoryID: 0,
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
                CategoryID: 0,
                SubcategoryID: 0,
                ChildcategoryID: 0,
                Category: null,
                Subcategory: null,
                Childcategory: null,
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.SetData()
                }
            },
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
            CategoryID: function (newValue, oldValue) {
                if (!this.loading) {
                    this.SubcategoryID = 0
                }
            },
            SubcategoryID: function (newValue, oldValue) {
                if (!this.loading) {
                    this.ChildcategoryID = 0
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            prefix: function () {
                return 'categories'
            },
            CategoryIDError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.CategoryID) {
                    return "please select category"
                }
            },
            SubcategoryIDError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.SubcategoryID) {
                    return "please select subcategory"
                }
            },
        },
        methods: {
            SetData() {
                if (this.value.ID > 0) {
                    this.loading = true
                    this.CategoryID = this.value.CategoryID
                    this.SubcategoryID = this.value.SubcategoryID
                    this.ChildcategoryID = this.value.ChildcategoryID
                    setTimeout(() => {
                        this.loading = false
                    }, 100);
                }
            },
            submit() {
                this.submitted = true
                if (this.CategoryIDError || this.SubcategoryIDError) { return false }
                var product = _.clone(this.value)
                product.CategoryID = this.CategoryID
                product.SubcategoryID = this.SubcategoryID
                product.ChildcategoryID = this.ChildcategoryID
                if (!product.Logo) {
                    if (this.Childcategory && this.Childcategory.Logo) {
                        product.Logo = this.Childcategory.Logo
                    } else if (this.Subcategory && this.Subcategory.Logo) {
                        product.Logo = this.Subcategory.Logo
                    } else if (this.Category && this.Category.Logo) {
                        product.Logo = this.Category.Logo
                    }
                }
                return {
                    product: product,
                }
            }
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload')
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="d-flex flex-column text-center">
        <p class="display-6 m-5">Select in which category you want the item to appear on website</p>
        <formitem name="inputCategoryID" label="Category" :error="CategoryIDError" v-model="CategoryID" type="select"
            @onselect="Category = $event" service="allcategories" displayby="Name" selectby="ID" :customLayout="true" />
        <formitem ref="SubcategoryInput" name="inputSubcategoryID" label="Subcategory" :error="SubcategoryIDError"
            v-model="SubcategoryID" type="select" service="allsubcategories" displayby="Name" selectby="ID"
            :filter="{ CategoryID: CategoryID }" @onselect="Subcategory = $event" :customLayout="true" />
        <formitem ref="ChildcategoryInput" name="inputChildcategoryID" label="Childcategory" v-model="ChildcategoryID"
            type="select" service="allchildcategories" displayby="Name" selectby="ID"
            :filter="{ SubcategoryID: SubcategoryID }" @onselect="Childcategory = $event" :customLayout="true" />
    </div>
</template>