<script>
    export default {
        props: { // 
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Variation: "",
                    }
                }
            },
            mode: {
                type: String,
                default: "form"
            }
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                ShippingServices: [],
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
            ShippingServices: function (newValue, oldValue) {
                if (!this.ShippingServices.length) {
                    this.addShippingService()
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        methods: {
            SetData() {
                if (this.value) {
                    this.loading = true
                    
                    if(this.value.Shipping){
                        try {
                            this.ShippingServices = JSON.parse(this.value.Shipping)
                        } catch (error) {
                            console.log(error)
                        }
                    }
                        if(!this.ShippingServices.length){
                            this.addShippingService()
                        }
                    setTimeout(() => {
                        this.loading = false
                    }, 100);
                }
            },
            submit() {
                this.submitted = true
                if (this.isShippingInvalid()) { return false }
                try {
                    var value = {}
                    if (this.value) {
                        value.ID = this.value.ID
                    }
                    value.Shipping = JSON.stringify(this.ShippingServices)
                    this.loading = true
                    return this.$store.dispatch('call', {
                        api: "product",
                        data: {
                            product: value
                        },
                    }).then((data) => {
                        this.message = data.Message;
                        if (data.Status == 2) {
                            if(this.mode == "form"){
                                this.$emit('input', data.Result.product)
                            }else{
                                return {
                                    product : data.Result.product
                                }
                            }
                        } else {
                            this.error = true
                            return false
                        }

                    }).catch((error) => {
                        console.error('Error:', error);
                        this.error = true
                        this.message = error
                    }).finally(() => {
                        this.loading = false
                        if (!this.error) {
                            this.message = ""
                        }
                    })
                } catch (error) {
                    console.log(error)
                    this.message = error
                    this.error = true
                }
            },
            isShippingInvalid() {
                if (!this.submitted) {
                    return false
                }
                var ShippingServiceNames = []
                for (let index = 0; index < this.ShippingServices.length; index++) {
                    const ShippingService = this.ShippingServices[index];
                    ShippingService.Name = ShippingService.Name.trim()
                    if (!ShippingService.Name) {
                        return true
                    }
                    if(ShippingServiceNames.includes(ShippingService.Name)){
                        ShippingService.NameRepeted = true
                        return true
                    }else{
                        ShippingService.NameRepeted = false
                    }
                    ShippingServiceNames.push(ShippingService.Name)
                }
                return false
            },
            addShippingService() {
                this.ShippingServices.push({
                    Image : "",
                    Name: "Shipping Free",
                    NameRepeted : false,
                    Price: 0,
                    EDTMin: 7,
                    EDTMax: 14
                })
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
    <divloading :fullpage="false" :loading="loading" class="mt-1">
        <form @submit.prevent="submit" class="d-flex flex-column">
            <v-alert v-model="message" :error="error" />
            <div class="d-flex flex-column align-items-start">
                <div v-for="(ShippingService, index) in ShippingServices" :key="'item_'+index" class="form-control mb-1">
                    <div class="d-flex">
                        <formitem :name="'input_shipping_service_name'+index" label="Shipping Service Name" v-model="ShippingService.Name" class="mb-1 me-1 flex-1"/>
                        <formitem :name="'input_shipping_service_price'+index" label="Price" v-model="ShippingService.Price" type="number" class="mb-1 me-1"/>
                        <formitem :name="'input_shipping_service_edmin'+index" label="Minimum days" v-model="ShippingService.EDTMin" type="number" class="mb-1 me-1"/>
                        <formitem :name="'input_shipping_service_edmax'+index" label="Maximum days" v-model="ShippingService.EDTMax" type="number" class="mb-1 me-1"/>
                        <button type="button" class="btn btn-sm btn-danger m-1" @click="ShippingServices.splice(index, 1);"><i class="fas fa-trash-alt mr-1"></i></button>
                    </div>
                    <label> Available In...</label>
                    <div class="d-flex">
                        <formitem name="inputCountry" label="Country" v-model="ShippingService.Country" type="select" service="countries" displayby="Country" selectby="Country" class="mb-1 me-1" nothing_selected_text="All Countries"  :select_out_of_single="false"/>
                        <formitem name="inputDistrict" label="District" v-model="ShippingService.District" type="select" class="mb-1 me-1" service="districts" :filter="{ country: ShippingService.Country }" displayby="District" selectby="District" nothing_selected_text="All Districts" :select_out_of_single="false"/>
                        <formitem name="inputLocality" label="Locality" v-model="ShippingService.Locality" type="select" class="mb-1 me-1" service="localities" :filter="{ district: ShippingService.District }" displayby="Locality" selectby="Locality"  nothing_selected_text="All Localities" :select_out_of_single="false"/>
                        <formitem name="inputSubLocality" label="SubLocality" v-model="ShippingService.SubLocality" type="select" class="mb-1 me-1 flex-1" service="sublocalities" :filter="{ locality: ShippingService.Locality }" displayby="SubLocality" selectby="SubLocality" nothing_selected_text="All SubLocality" :select_out_of_single="false"/>
                    </div>
                    <div v-if="!ShippingService.Name" class="invalid-feedback" style="display: block;">Please enter shipping service name</div>
                    <div v-if="submitted && ShippingService.NameRepeted" class="invalid-feedback" style="display: block;">Please enter unique shipping service name</div>
                </div>
                <button type="button" class="btn btn-sm btn-success m-1" @click="addShippingService()"> Add new shipping service </button>
            </div>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="SetData"> Cancel </button>
            </div>
        </form>
    </divloading>
</template>