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
                HaveVariation: false,
                Option: {
                    Text: "Option",
                    Color: "#ffffff",
                    Image: ""
                },
                Verities: [],
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
            HaveVariation: function (newValue, oldValue) {
                if (newValue && !this.Verities.length) {
                    this.pushVerity()
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            HaveVariationError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.HaveVariation) {
                    return false
                } else if (this.HaveVariation) {
                    if (!this.Verities.length) {
                        return "Please add verity or say no to above question"
                    }
                } else {
                    return "Please select Yes or No for above question"
                }
            },
        },
        methods: {
            SetData() {
                if (this.value) {
                    this.loading = true
                    this.HaveVariation = this.value.HaveVariation
                    if (this.value.Variation) {
                        try {
                            this.Verities = JSON.parse(this.value.Variation)
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
                    value.HaveVariation = this.HaveVariation
                    value.Variation = JSON.stringify(this.Verities)
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
                if(!this.HaveVariation){
                    return false
                }
                if(!this.Verities.length){
                    return true
                }
                for (let index = 0; index < this.Verities.length; index++) {
                    const verity = this.Verities[index];
                    if (!verity.Name) {
                        return true
                    }
                    if(!verity.Options.length){
                        return true
                    }
                    for (let index = 0; index < verity.Options.length; index++) {
                        const option = verity.Options[index];
                        if(!option.Name){
                            return true
                        }
                        if(verity.Color && !option.Color){
                            return false
                        }
                        if(verity.Image && !option.Image){
                            return false
                        }
                    }
                }
                return false
            },
            pushVerity() {
                var Name = "verity type " + (this.Verities.length + 1)
                var verity = {
                    Name: Name,
                    Options: []
                }
                if(!this.Verities.length){
                    verity.Name = "Color"
                    verity.Color = true
                }
                if(this.Verities.length == 1){
                    verity.Name = "Size"
                }
                this.Verities.push(verity);
            },
            setVerityType(verity, typename) {
                if (!typename) {
                    verity.Error = "please enter name for type of Verity"
                } else {
                    verity.Name = typename
                    delete verity.Error
                }
            },
            addOption(verity){
                verity.Options.push({
                    Text: "Option",
                    Color: "#ffffff",
                    Image: ""
                })
            },
            prefix(){
                return "product/"+ this.value.ID + "/options"
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
            <formitem name="inputHaveVariation" label="" :error="HaveVariationError"
                v-model="HaveVariation" :customLayout="true">
                <div class="d-flex flex-row">
                    <label class="m-1">Does this product have verity ? </label>
                    <div class="form_control m-1">
                        <input type="radio" id="inputVerityYes" name="radioVerity" :value="true"
                            v-model="HaveVariation" />
                        <label for="inputVerityYes">Yes</label>
                    </div>
                    <div class="form_control m-1">
                        <input type="radio" id="inputVerityNo" name="radioVerity" :value="false" v-model="HaveVariation" />
                        <label for="inputVerityNo">No</label>
                    </div>
                </div>
            </formitem>
            <div v-if="HaveVariation" v-for="(verity, index) in Verities" class="row">
                <div class="col-md-4 col-lg-3">
                    <label class="form-label"> Type of Verity : </label>
                    <input type="text" :name="'inputVerity'+index" v-model="verity.Name" class="form-control"/>
                    <div v-if="!verity.Name" class="invalid-feedback" style="display: block;">Please enter label for type of verity</div>
                    <div class="d-flex">
                        <div class="form_control m-1">
                            <input type="checkbox" :id="'inputVerityColor'+index" :name="'inputVerityColor'+index" v-model="verity.Color" />
                            <label :for="'inputVerityColor'+index">Have Color</label>
                        </div>
                        <div class="form_control m-1 flex-1">
                            <input type="checkbox" :id="'inputVerityImage'+index" :name="'inputVerityImage'+index" v-model="verity.Image" />
                            <label :for="'inputVerityImage'+index">Have Image</label>
                        </div>
                        <button type="button" class="btn btn-sm btn-danger" @click="Verities.splice(index, 1);"><i class="fas fa-trash-alt mr-1"></i></button>
                    </div>
                </div>
                <div class="col-md-8 col-lg-9 ">
                    <label class="form-label"> Options : </label>
                    <div class="d-flex flex-wrap align-items-start">
                        <div v-for="(option, option_index) in verity.Options" :key="'option_'+index+'_'+option_index"
                            class="d-flex flex-column form-control w-auto mb-1 me-1">
                                <div class="d-flex">
                                    <input type="text" :name="'inputoption_'+index+'_'+option_index" v-model="option.Name" style="width: auto;" class="form-control"/>
                                    <button type="button" class="btn btn-sm btn-danger" @click="verity.Options.splice(option_index, 1);"><i class="fas fa-trash-alt mr-1"></i></button>
                                </div>
                                <input v-if="verity.Color" type="color" :name="'inputoption_'+index+'_'+option_index" v-model="option.Color" style="width: auto;"/>
                                <ImageFile v-if="verity.Image" :prefix="prefix()" maxWidth="120px" maxHeight="120px" v-model="option.Image"/>
                                <div v-if="!option.Name" class="invalid-feedback" style="display: block;">Enter option name</div>
                                <div v-if="verity.Color && !option.Color" class="invalid-feedback" style="display: block;">Enter option color</div>
                                <div v-if="verity.Image && !option.Image" class="invalid-feedback" style="display: block;">Enter option image</div>
                        </div>
                        <button type="button" class="btn btn-sm btn-success mb-1 me-1" @click="addOption(verity)"> + </button>
                    </div>
                    <div v-if="2 > verity.Options.length" class="invalid-feedback" style="display: block;">Enter atlist two options for {{verity.Name}} variation</div>
                </div>
            </div>
            <div v-if="HaveVariation" class="row">
                <div class="col-md-4 col-lg-3">
                    <button type="button" class="btn btn-sm btn-success mb-1 me-1" @click="pushVerity()"> Add type of variation </button>
                </div>
                <div class="col-md-8 col-lg-9">

                </div>
            </div>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="SetData"> Cancel </button>
            </div>
        </form>
    </divloading>
</template>