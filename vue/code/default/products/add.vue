<script>
    Vue.use(VueFormWizard)
    export default {
        components: {
            'productcategory': () => import("/vue/products/productcategory.js"),
            'producteditor': () => import("/vue/products/producteditor.js"),
            'variationeditor': () => import("/vue/products/variationeditor.js"),
            'productimages': () => import("/vue/products/productimages.js"),
            'shippingeditor': () => import("/vue/products/shippingeditor.js"),
            'stockeditor': () => import("/vue/products/stockeditor.js"),
        },
        data: () => {
            return {
                loading: false,
                error: false,
                message: "",
                product: {
                    ID: 0,
                    Name: "",
                    Summary: "",
                    Condition: "",
                    QuickPoints: "",
                    Comment: "",
                    EVRating: 0,
                    CategoryID: 0,
                    SubcategoryID: 0,
                    ChildcategoryID: 0,
                    Variation: "",
                },
            }
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        methods: {
            onComplete(arg) {
                this.$router.push('/products')
            },
            checkDetails: function () {
                var flag = Math.floor(Math.random() * 10)
                if (flag % 2) {
                    return true
                } else {
                    return false
                }
            },
            validateStep: function (name) {
                var refToValidate = this.$refs[name];
                if (!refToValidate) return false
                return Promise.resolve(refToValidate.submit()).then((data) => {
                    if (data.product) {
                        this.product = data.product
                    }
                    if (!data) {
                        return false
                    } else {
                        return true
                    }
                });
            },
        },
        template: `{{{template}}}`,
    };
</script>
<template>
    <div class="row add_product_page">
        <div class="col card">
            <form-wizard @on-complete="onComplete" color="#45afda" title="" subtitle="">
                <tab-content title="Category" :before-change="()=>validateStep('productcategory')">
                    <productcategory ref="productcategory" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Basic Details" :before-change="()=>validateStep('producteditor')">
                    <producteditor ref="producteditor" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Variation" :before-change="()=>validateStep('variationeditor')">
                    <variationeditor ref="variationeditor" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Stock & Price" :before-change="()=>validateStep('stockeditor')">
                    <stockeditor ref="stockeditor" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Images" :before-change="()=>validateStep('productimages')">
                    <productimages ref="productimages" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Shipping" :before-change="()=>validateStep('shippingeditor')">
                    <shippingeditor ref="shippingeditor" mode="wizard" v-model="product" />
                </tab-content>
            </form-wizard>
        </div>
    </div>
</template>