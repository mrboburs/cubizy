
    Vue.use(VueFormWizard)
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return null
                }
            },
        },
        components: {
            'producteditor': () => import("/vue/products/producteditor.js"),
            'variationeditor': () => import("/vue/products/variationeditor.js"),
            'productimages': () => import("/vue/products/productimages.js"),
            'shippingeditor': () => import("/vue/products/shippingeditor.js"),
            'stockeditor': () => import("/vue/products/stockeditor.js"),
            'descriptioneditor': () => import("/vue/products/descriptioneditor.js"),
            'reviews': () => import("/vue/products/reviews.js"),
        },
        data: () => {
            return {
                loading: false,
                error: false,
                message: "",
                tab: 'producteditor',
                product: {},
                attributes: [],
                extra_details: {},
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
            value: function (newValue) {
                this.product = newValue
            },
        },
        mounted: function () {
            if (this.value) {
                this.product = this.value
            }
        },
        template: `
    <div class="d-flex flex-column">
        <v-alert v-model="message" :error="error" />
        <ul class="nav nav-tabs">
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'producteditor'}" @click.prevent="tab = 'producteditor'"
                    href="#">Basic Details</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'variationeditor'}"
                    @click.prevent="tab = 'variationeditor'" href="#">Variation</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'stockeditor'}"
                    @click.prevent="tab = 'stockeditor'" href="#">Stock & Price</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'productimages'}" @click.prevent="tab = 'productimages'"
                    href="#">Images</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'shippingeditor'}" @click.prevent="tab = 'shippingeditor'"
                    href="#">Shipping</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'descriptioneditor'}" @click.prevent="tab = 'descriptioneditor'"
                    href="#">Description</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'reviews'}" @click.prevent="tab = 'reviews'"
                    href="#">Reviews</a>
            </li>
        </ul>
        <div class="mt-1">
            <producteditor v-if="tab == 'producteditor'" ref="producteditor" v-model="product" ></producteditor>
            <variationeditor v-else-if="tab == 'variationeditor'" ref="variationeditor" v-model="product" />
            <stockeditor v-else-if="tab == 'stockeditor'" ref="stockeditor" v-model="product" />
            <productimages v-else-if="tab == 'productimages'" ref="productimages" v-model="product" />
            <shippingeditor v-else-if="tab == 'shippingeditor'" ref="shippingeditor" v-model="product" />
            <descriptioneditor v-else-if="tab == 'descriptioneditor'" ref="descriptioneditor" v-model="product" />
            <reviews v-else-if="tab == 'reviews'" ref="reviews" v-model="product" />
        </div>
    </div>
`,
    };
