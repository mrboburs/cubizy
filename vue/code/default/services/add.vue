<script>
    Vue.use(VueFormWizard)
    export default {
        components: {
            'productcategory': () => import("/vue/products/productcategory.js"),
            'serviceseditor': () => import("/vue/services/serviceseditor.js"),
            'serviceimages': () => import("/vue/services/serviceimages.js"),
            'pricingeditor': () => import("/vue/services/pricingeditor.js"),
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
                <tab-content title="Basic Details" :before-change="()=>validateStep('serviceseditor')">
                    <serviceseditor ref="serviceseditor" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Variation & Price" :before-change="()=>validateStep('pricingeditor')">
                    <pricingeditor ref="pricingeditor" mode="wizard" v-model="product" />
                </tab-content>
                <tab-content title="Images" :before-change="()=>validateStep('serviceimages')">
                    <serviceimages ref="serviceimages" mode="wizard" v-model="product" />
                </tab-content>
            </form-wizard>
        </div>
    </div>
</template>