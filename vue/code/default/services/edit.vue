<script>
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
            'serviceseditor': () => import("/vue/services/serviceseditor.js"),
            'serviceimages': () => import("/vue/services/serviceimages.js"),
            'pricingeditor': () => import("/vue/services/pricingeditor.js"),
        },
        data: () => {
            return {
                loading: false,
                error: false,
                message: "",
                tab: 'serviceseditor',
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
        template: `{{{template}}}`,
    };
</script>
<template>
    <div class="d-flex flex-column">
        <v-alert v-model="message" :error="error" />
        <ul class="nav nav-tabs">
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'serviceseditor'}" @click.prevent="tab = 'serviceseditor'"
                    href="#">Basic Details</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'pricingeditor'}"
                    @click.prevent="tab = 'pricingeditor'" href="#">Variation & Price</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" :class="{'active': tab == 'serviceimages'}" @click.prevent="tab = 'serviceimages'"
                    href="#">Images</a>
            </li>
        </ul>
        <div class="mt-1">
            <serviceseditor v-if="tab == 'serviceseditor'" ref="serviceseditor" v-model="product" ></serviceseditor>
            <pricingeditor v-else-if="tab == 'pricingeditor'" ref="pricingeditor" v-model="product" />
            <serviceimages v-else-if="tab == 'serviceimages'" ref="serviceimages" v-model="product" />
        </div>
    </div>
</template>