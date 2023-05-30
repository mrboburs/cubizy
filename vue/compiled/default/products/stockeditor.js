
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
                HaveVariation: false,
                Verities: [],
                Stock: [],
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
                    this.Verities = []
                    this.Stock = []
                    this.HaveVariation = this.value.HaveVariation
                    if (this.value.Variation) {
                        try {
                            this.Verities = JSON.parse(this.value.Variation)
                        } catch (error) {
                            console.log(error)
                        }
                    }
                    if (this.value.Stock) {
                        try {
                            this.Stock = JSON.parse(this.value.Stock)
                        } catch (error) {
                            console.log(error)
                        }
                    }
                    if (!this.Stock.length) {
                        this.AddMissingStockRecords()
                    }
                    setTimeout(() => {
                        this.loading = false
                    }, 100);
                }
            },
            AddMissingStockRecords() {
                var baseArray = [{
                    SKU: "",
                    Variation: {},
                    Active: true,
                    Quantity: 1,
                    Sold: 0,
                    Price: 0,
                    Discount: 0,
                    Cost: 0,
                    LowStockLimit: 1,
                    AddToCart: false,
                }]
                if (this.HaveVariation) {
                    this.Verities.forEach(variant => {
                        var outputArray = []
                        baseArray.forEach(item => {
                            variant.Options.forEach(Option => {
                                var new_item = _.cloneDeep(item)
                                new_item.Variation[variant.Name] = Option.Name
                                outputArray.push(new_item)
                            });
                        });
                        baseArray = outputArray
                    });
                    if (this.Stock.length) {
                        baseArray.forEach(element => {
                            this.addStockRecord(element.Variation)
                        });
                    } else {
                        this.Stock = baseArray
                    }
                } else {
                    this.Stock = baseArray
                }
            },
            addStockRecord(Variation) {
                if (this.Stock.length) {
                    var duplicateStock = this.Stock.filter(variant => this.shallowEqual(variant.Variation, Variation));
                    if (duplicateStock.length) {
                        this.message = "Allready present"
                        this.error = true
                        return
                    }
                }
                this.Stock.push({
                    SKU: "",
                    Variation: {},
                    Active: true,
                    Quantity: 1,
                    Sold: 0,
                    Price: 0,
                    Discount: 0,
                    Cost: 0,
                    LowStockLimit: 1,
                    AddToCart: false,
                })
            },
            shallowEqual(object1, object2) {
                const keys1 = Object.keys(object1);
                const keys2 = Object.keys(object2);
                if (keys1.length !== keys2.length) {
                    return false;
                }
                for (let key of keys1) {
                    if (object1[key] !== object2[key]) {
                        return false;
                    }
                }
                return true;
            },
            submit() {
                this.submitted = true
                this.message = ""
                this.error = false
                if (this.isStockInvalid()) {
                    if(this.message == ""){
                        this.message = "Please enter valid details in all fields"
                    }
                    this.error = true
                    return false
                }
                try {
                    var value = {}
                    if (this.value) {
                        value.ID = this.value.ID
                    }
                    if (this.Stock.length > 1) {
                        value.SKU = ""
                    } else if (this.Stock.length) {
                        value.SKU = this.Stock[0].SKU
                    }
                    value.Quantity = 0
                    for (let index = 0; index < this.Stock.length; index++) {
                        const record = this.Stock[index];
                        value.Quantity += record.Quantity
                        if (value.Price > record.Price|| !value.Price) {
                            value.Price = record.Price
                        }
                        if (value.Discount < record.Discount || !value.Discount) {
                            value.Discount = record.Discount
                        }
                        if (value.Cost > record.Cost || !value.Cost) {
                            value.Cost = record.Cost
                        }
                        if (!value.LowStockLimit || value.LowStockLimit < record.LowStockLimit) {
                            value.LowStockLimit = record.LowStockLimit
                        }
                        if (!value.MaxPrice || value.MaxPrice < record.Price) {
                            value.MaxPrice = record.Price
                        }
                        if (!value.MaxCost || value.MaxCost < record.Cost) {
                            value.MaxCost = record.Cost
                        }
                    }
                    value.Stock = JSON.stringify(this.Stock)
                    this.loading = true
                    return this.$store.dispatch('call', {
                        api: "product",
                        data: {
                            product: value
                        },
                    }).then((data) => {
                        this.message = data.Message;
                        if (data.Status == 2) {
                            if (this.mode == "form") {
                                this.$emit('input', data.Result.product)
                            } else {
                                return {
                                    product: data.Result.product
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
            isStockInvalid() {
                if (!this.submitted) {
                    return false
                }
                var SKUs = []
                for (let index = 0; index < this.Stock.length; index++) {
                    const record = this.Stock[index];
                    if (!record.SKU) {
                        return true
                    }else if(SKUs.includes(record.SKU)){
                        this.message = "SKU can not be repeated"
                        this.error = true
                        return true
                    }else{
                        SKUs.push(record.SKU)
                    }
                    if (record.Price < 0) {
                        return true
                    }
                    if (record.Discount < 0) {
                        return true
                    }
                    if (record.Cost < 0) {
                        return true
                    }
                    if (record.LowStockLimit < 1) {
                        return true
                    }
                }
                return false
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
        template: `
    <divloading :fullpage="false" :loading="loading" class="mt-1">
        <form @submit.prevent="submit" class="d-flex flex-column">
            <v-alert v-model="message" :error="error" />
            <div class="d-flex centered">
                <button type="button" class="btn btn-sm btn-success flex-1 m-1" @click="AddMissingStockRecords()">
                    Add All Missing Stock Records by Variation
                </button>
                <button type="button" class="btn btn-sm btn-success m-1" @click="addStock()"> Add Stock Record
                </button>
            </div>
            <div class="d-flex flex-column">
                <div v-for="(record, index) in Stock" :key="'item_'+index" class="form-control mb-1">
                    <div class="d-flex align-items-start">
                        <formitem v-for="(key, variation_index) in Object.keys(record.Variation)"
                            :key="'input_stock_variation'+index+'_'+variation_index"
                            :name="'input_stock_variation'+index+'_'+variation_index" :label="key"
                            :value="record.Variation[key]" :readonly="true" class="mb-1 me-1 flex-1" />
                        <formitem :name="'input_stock_sku'+index" label="SKU" v-model="record.SKU"
                            class="mb-1 me-1 flex-grow-1" />
                        <formitem :name="'input_stock_quantity'+index" label="Quantity" v-model="record.Quantity"
                            type="number" class="mb-1 me-1 flex-1" />
                        <formitem :name="'input_stock_lowquantity'+index" label="Low Stock Limit"
                            v-model="record.LowStockLimit" type="number" class="mb-1 me-1 flex-1" />
                        <formitem :name="'input_stock_price'+index" label="Base Price" :value="record.Price"
                            type="number" class="mb-1 me-1 flex-1" @input="setPrice(record, $event)" />
                        <formitem :name="'input_stock_discount'+index" label="Discount" v-model="record.Discount"
                            type="number" class="mb-1 me-1 flex-1" @input="setDiscount(record, $event)" :inputgroup="true" suffix="%" />
                        <formitem :name="'input_stock_cost'+index" label="Final Cost" v-model="record.Cost"
                            type="number" class="mb-1 me-1 flex-1" />
                        <button type="button" class="btn btn-danger" style="margin-top: 2em;"
                            @click="Stock.splice(index, 1);">
                            <i class="fas fa-trash-alt mr-1"></i>
                        </button>
                    </div>
                    <div v-if="!record.SKU" class="invalid-feedback" style="display: block;">
                        Please enter stock SKU
                    </div>
                    <div v-if="0 > record.Quantity" class="invalid-feedback" style="display: block;">
                        Please enter valid quantity
                    </div>
                    <div v-if="0 > record.Price" class="invalid-feedback" style="display: block;">
                        Please enter valid Price
                    </div>
                    <div v-if="0 > record.Discount" class="invalid-feedback" style="display: block;">
                        Please enter valid Discount
                    </div>
                    <div v-if="1 > record.LowStockLimit" class="invalid-feedback" style="display: block;">
                        Please enter Low Stock Limit more then 1
                    </div>
                </div>
            </div>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="SetData"> Cancel </button>
            </div>
        </form>
    </divloading>
`
    }
