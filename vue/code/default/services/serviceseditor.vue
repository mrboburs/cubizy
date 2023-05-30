<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Variation: "",
                        Summary: "",
                        Condition: "",
                        QuickPoints: "",
                        Comment: "",
                        ExtraDetails: {}
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

                Conditions: ['At service center', 'Onsite', 'Online'],
                NamePlaceholder: "e.g : Apple iMac",
                SummaryPlaceholder: "Please enter short discription of your product/service",

                Name: "",
                Condition: "",
                Summary: "",
                Service: false,
                QuickPoints: "",
                Comment: "",
                Status: false,
                Keywords: [],

                attributes: [],
                ExtraDetails: {},
                Category: null,
                Subcategory: null,
                Childcategory: null,
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
            NameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Name.trim()) {
                    return "Name can not  be empty"
                }
            },
            SummaryError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Summary.trim()) {
                    return "Summary can not  be empty"
                }
            },
            ConditionError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Condition.trim()) {
                    return "please select condition"
                }
            },
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.loading = true

                    this.Name = this.value.Name
                    this.Summary = this.value.Summary
                    this.Condition = this.value.Condition
                    this.QuickPoints = this.value.QuickPoints
                    this.Keywords = this.value.Keywords
                    this.Comment = this.value.Comment
                    return this.$store.dispatch('call', {
                        api: "attributes",
                        data: {
                            SubcategoryID: this.value.SubcategoryID,
                            ChildcategoryID: this.value.ChildcategoryID,
                            ProductID: this.value.ID,
                        },
                    }).then((data) => {
                        if (Array.isArray(data.data)) {
                            this.attributes = data.data
                        }
                        if (data.Result.ProductDetails) {
                            this.ExtraDetails = data.Result.ProductDetails
                        }
                        this.attributes.forEach(attribute => {
                            if (!this.ExtraDetails[attribute.ProductColumn]) {
                                if (attribute.FieldType == "Dropdown") {
                                    this.ExtraDetails[attribute.ProductColumn] = ""
                                } else {
                                    this.ExtraDetails[attribute.ProductColumn] = attribute.Options ? attribute.Options : ""
                                }
                            }
                        });
                    }).finally(() => {
                        this.loading = false
                        if (!this.error) {
                            this.message = ""
                        }
                    });
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.SummaryError || this.ConditionError || this.is_exter_details_invalid()) { return false }
                var value = {}
                if (this.value) {
                    value = _.clone(this.value)
                }
                value.Service = true
                value.Name = this.Name
                value.Summary = this.Summary
                value.Condition = this.Condition
                value.QuickPoints = this.QuickPoints
                value.Keywords = this.Keywords
                value.Comment = this.Comment

                var details = {}
                this.attributes.forEach(attribute => {
                    details[attribute.ProductColumn] = this.ExtraDetails[attribute.ProductColumn]
                });
                value.ExtraDetails = details

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
            },
            is_exter_details_invalid() {
                if (!this.submitted) {
                    return false
                }
                for (let index = 0; index < this.attributes.length; index++) {
                    const attribute = this.attributes[index];
                    if (this.error_in(attribute)) {
                        return true
                    }
                }
                return false
            },
            error_in(attribute) {
                if (!this.submitted) {
                    return false
                }
                if (!this.ExtraDetails[attribute.ProductColumn]) {
                    return "Please enter the value for " + attribute.Name
                }
            },
            field_type(FieldType) {
                switch (FieldType) {
                    case "Dropdown":
                        return "select"
                    default:
                        return FieldType.toLowerCase()
                }
            },
            dropdown_options(attribute) {
                if (attribute.FieldType == "Dropdown") {
                    return attribute.Options.split(",")
                } else {
                    return []
                }
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
            <formitem name="inputName" label="Name" :error="NameError" v-model="Name" :customLayout="true" />
            <formitem name="inputCondition" label="Condition" :error="ConditionError" v-model="Condition"
                :customLayout="true">
                <div class="d-flex flex-row form-control">
                    <div v-for="(_Condition, index) in Conditions" class="mx-2">
                        <input type="radio" :id="'radio_condition' + index" name="radio_condition" :value="_Condition"
                            v-model="Condition" />
                        <label :for="'radio_condition' + index">{{_Condition}}</label>
                    </div>
                </div>
            </formitem>
            <formitem name="inputSummary" label="Summary" :error="SummaryError" v-model="Summary" type="textarea"
                :customLayout="true" />
            <formitem name="inputQuickPoints" label="QuickPoints" v-model="QuickPoints" type="tags" :customLayout="true" class="quick_points" />
            <formitem name="inputKeywords" label="More Keywords" v-model="Keywords" type="tags" :customLayout="true" />
            <formitem v-for="attribute in attributes" :key="attribute.ProductColumn"
                :name="'input'+attribute.ProductColumn" :label="attribute.Name" :error="error_in(attribute)"
                v-model="ExtraDetails[attribute.ProductColumn]" :type="field_type(attribute.FieldType)" displayby="Name"
                selectby="ID" :customLayout="true" :values="dropdown_options(attribute)" />
            <formitem name="inputStatus" label="Status" v-model="Status" v-if="mode == 'form'" :customLayout="true">
                <div class="form-check form-switch ml-2">
                    <input class="form-check-input" type="checkbox" id="inputIsSuperAdmin" v-model="Status">
                    <label class="form-check-label" for="inputIsSuperAdmin">
                        <span v-if="Status">Active</span>
                        <span v-if="!Status">Not Active</span>
                    </label>
                </div>
            </formitem>
            <div class="d-flex centered" v-if="mode == 'form'">
                <button type="submit" class="btn btn-success m-1" :disabled="loading"> Save </button>
                <button class="btn btn-danger m-1" @click="Reset"> Cancel </button>
            </div>
        </form>
    </divloading>
</template>