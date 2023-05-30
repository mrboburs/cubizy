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
                Variation: [],
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
        },
        methods: {
            SetData() {
                if (this.value) {
                    this.loading = true
                    
                    if(this.value.Variation){
                        try {
                            this.Variation = JSON.parse(this.value.Variation)
                        } catch (error) {
                            console.log(error)
                        }
                    }
                    if(!this.Variation.length){
                        this.addVerity()
                    }
                    setTimeout(() => {
                        this.loading = false
                    }, 100);
                }
            },
            submit() {
                this.submitted = true
                if (this.isVariationInvalid()) { return false }
                try {
                    var value = {}
                    if (this.value) {
                        value.ID = this.value.ID
                    }
                    if(this.Variation.length > 1){
                        value.HaveVariation = true
                    }else{
                        value.HaveVariation = false
                    }

                    if (this.Variation.length > 1) {
                        value.SKU = ""
                    } else if (this.Variation.length) {
                        value.SKU = this.Variation[0].SKU
                    }
                    for (let index = 0; index < this.Variation.length; index++) {
                        const record = this.Variation[index];
                        if (value.Price > record.Price || !value.Price) {
                            value.Price = record.Price
                        }
                        if (value.Discount < record.Discount || !value.Discount) {
                            value.Discount = record.Discount
                        }
                        if (value.Cost > record.Cost || !value.Cost ) {
                            value.Cost = record.Cost
                        }
                        if (value.MaxPrice < record.Price || !value.MaxPrice) {
                            value.MaxPrice = record.Price
                        }
                        if (value.MaxCost < record.Cost || !value.MaxCost ) {
                            value.MaxCost = record.Cost
                        }
                    }

                    value.Variation = JSON.stringify(this.Variation)
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
            isVariationInvalid() {
                if (!this.submitted) {
                    return false
                }
                if (!this.Variation.length) {
                    return true
                }
                for (let index = 0; index < this.Variation.length; index++) {
                    const Verity = this.Variation[index];
                    if (!Verity.SKU) {
                        return true
                    }
                    if (Verity.Price < 0) {
                        return true
                    }
                    if (Verity.Discount < 0) {
                        return true
                    }
                    if (Verity.Cost < 0) {
                        return true
                    }
                }
                return false
            },
            addVerity() {
                this.Variation.push({
                    Name: "",
                    SKU: "",
                    Price: 0,
                    Discount: 0,
                    Cost: 0,
                })
            },
            setPrice(record, value) {
                record.Price = value
                if (record.Discount) {
                    record.Cost = record.Price - (record.Price * (record.Discount / 100))
                } else {
                    record.Cost = record.Price
                }
            },
            setDiscount(record, value) {
                record.Discount = value
                if (record.Discount) {
                    record.Cost = record.Price - (record.Price * (record.Discount / 100))
                } else {
                    record.Cost = record.Price
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
    <divloading :fullpage="false" :loading="loading" class="mt-1">
        <form @submit.prevent="submit" class="d-flex flex-column">
            <v-alert v-model="message" :error="error" />
            <div class="d-flex flex-column align-items-start">
                <div v-for="(Verity, index) in Variation" :key="'item_'+index" class="form-control mb-1">
                    <div class="d-flex">
                        <formitem v-if="Variation.length > 1" :name="'input_verity_name'+index" label="Verity Name" v-model="Verity.Name" class="mb-1 me-1 flex-1"/>
                        <formitem :name="'input_verity_sku'+index" label="Service Code" v-model="Verity.SKU"
                            class="mb-1 me-1 flex-grow-1" />
                        <formitem :name="'input_verity_price'+index" label="Base Price" :value="Verity.Price"
                            type="number" class="mb-1 me-1 flex-1" @input="setPrice(Verity, $event)" />
                        <formitem :name="'input_verity_discount'+index" label="Discount" v-model="Verity.Discount"
                            type="number" class="mb-1 me-1 flex-1" @input="setDiscount(Verity, $event)" :inputgroup="true" suffix="%" />
                        <formitem :name="'input_verity_cost'+index" label="Final Cost" v-model="Verity.Cost"
                            type="number" class="mb-1 me-1 flex-1" />
                        <button type="button" class="btn btn-sm btn-danger m-1" @click="Variation.splice(index, 1);"><i class="fas fa-trash-alt mr-1"></i></button>
                    </div>
                    <div v-if="!Verity.Name && Variation.length > 1" class="invalid-feedback" style="display: block;">Please enter verity name</div>
                    <div v-if="!Verity.SKU" class="invalid-feedback" style="display: block;">Please enter service code</div>
                </div>
                <button type="button" class="btn btn-sm btn-success m-1" @click="addVerity()"> Add new verity of service </button>
            </div>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="SetData"> Cancel </button>
            </div>
        </form>
    </divloading>
</template>