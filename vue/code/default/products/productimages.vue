<script>
    export default {
        components: {
            'verityeditor': () => import("/vue/products/verityeditor.js"),
        },
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
                Logo: "",
                Images: [],
                Verities: [],
                ImageBy : false
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
                    this.Logo = this.value.Logo
                    if (this.value.Variation) {
                        try {
                            this.Verities = JSON.parse(this.value.Variation)
                        } catch (error) {
                            console.log(error)
                        }
                    }
                    if(this.value.ImageBy){
                        this.Verities.forEach(element => {
                            if(element.Name == this.value.ImageBy ){
                                this.ImageBy = element
                            }
                        });
                    }
                    this.ImageBy
                    if (this.value.Images) {
                        try {
                            this.Images = JSON.parse(this.value.Images)
                        } catch (error) {
                            console.log(error)
                        }
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
                    value.Logo = this.Logo
                    value.ImageBy = this.ImageBy.Name
                    value.Images = JSON.stringify(this.Images)
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
                if (!this.HaveVariation) {
                    return false
                }
                if (!this.Verities.length) {
                    return true
                }
                for (let index = 0; index < this.Verities.length; index++) {
                    const verity = this.Verities[index];
                    if (!verity.Name) {
                        return true
                    }
                    if (!verity.Options.length) {
                        return true
                    }
                    for (let index = 0; index < verity.Options.length; index++) {
                        const option = verity.Options[index];
                        if (!option.Name) {
                            return true
                        }
                        if (verity.Color && !option.Color) {
                            return false
                        }
                        if (verity.Image && !option.Image) {
                            return false
                        }
                    }
                }
                return false
            },
            addImage(verity) {
                this.Images.push({
                    Image: "",
                    Variant: "",
                })
            },
            prefix() {
                return "product/" + this.value.ID + "/images"
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
            <div class="form-group  m-3" v-if="value.HaveVariation">
                <label>Dose have different Image by <span v-if="value.Service">Service</span> <span
                        v-else>Product</span> variety !</label>
                <select class="form-select" v-model="ImageBy">
                    <option value="">No</option>
                    <option v-for="(verity, index) in Verities" :key="index" :value="verity">
                        {{verity.Name}}
                    </option>
                </select>
            </div>
            <div class="d-flex flex-wrap align-items-start">
                <div class="form-control w-auto mb-1 me-1">
                    <ImageFile v-model="Logo" :prefix="prefix()" maxHeight="260px" maxWidth="260px"></ImageFile>
                    <label class="form-control">Logo(Default Image)</label>
                </div>
                <div v-for="(item, index) in Images" :key="'item_'+index" class="d-flex flex-column form-control w-auto mb-1 me-1">
                    <ImageFile v-model="item.Image" :prefix="prefix()" maxHeight="260px" maxWidth="260px"></ImageFile>
                    <div class="d-flex">
                        <select v-if="ImageBy" class="form-select" v-model="item.Variant">
                            <option value="">in all</option>
                            <option v-for="variant in ImageBy.Options" :value="variant.Name">{{variant.Name}}</option>
                        </select>
                        <button type="button" class="btn btn-sm btn-danger" @click="Images.splice(index, 1);"><i class="fas fa-trash-alt mr-1"></i></button>
                    </div>
                </div>
                <button type="button" class="btn btn-sm btn-success mb-1 me-1" @click="addImage()"> + </button>
            </div>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="SetData"> Cancel </button>
            </div>
        </form>
    </divloading>
</template>