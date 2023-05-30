
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Code: "",
                        SubLocality: "",
                        Locality: "",
                        District: "",
                        Country: "Mauritius",
                        Latitude: 0,
                        Longitude: 0,
                    }
                }
            },
            title: {
                type: String,
                required: false
            },
            fullpage: {
                type: Boolean,
                required: false
            },
            canclose: {
                type: Boolean,
                default: true
            }
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Code: "",
                SubLocality: "",
                Locality: "",
                District: "",
                Country: "",
                Longitude: 0.0,
                Latitude: 0.0,
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
            CodeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Code.trim()) {
                    return "Please provide a valid Area Code/ Pin Code"
                }
            },
            SubLocalityError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.SubLocality.trim()) {
                    return "Please provide a valid SubLocality"
                }
            },
            LocalityError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Locality.trim()) {
                    return "Please provide a valid Locality"
                }
            },
            DistrictError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.District.trim()) {
                    return "Please provide a valid District"
                }
            },
            CountryError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Country.trim()) {
                    return "Please select a valid Country"
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
                    this.submitted = false
                    this.Code = this.value.Code
                    this.SubLocality = this.value.SubLocality
                    this.Locality = this.value.Locality
                    this.District = this.value.District
                    this.Country = this.value.Country
                }
            },
            submit() {
                this.submitted = true
                if (this.CodeError || this.SubLocalityError || this.LocalityError || this.DistrictError || this.CountryError) { return }
                this.value.Code = this.Code
                this.value.SubLocality = this.SubLocality
                this.value.Locality = this.Locality
                this.value.District = this.District
                this.value.Country = this.Country
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload')
        },
        template: `
    <form @submit.prevent="submit">
        <formitem :customLayout="true" name="inputCountry" label="Country" :error="CountryError" v-model="Country"/>
        <formitem :customLayout="true" name="inputDistrict" label="District" :error="DistrictError" v-model="District"/>
        <formitem :customLayout="true" name="inputLocality" label="Locality" :error="LocalityError" v-model="Locality"/>
        <formitem :customLayout="true" name="inputSubLocality" label="SubLocality" :error="SubLocalityError" v-model="SubLocality"/>
        <formitem :customLayout="true" name="inputCode" label="Pin Code" :error="CodeError" v-model="Code"/>
        <formitem :customLayout="true" name="inputLongitude" label="Longitude" v-model="Longitude" type="number"/>
        <formitem :customLayout="true" name="inputLatitude" label="Latitude" v-model="Latitude" type="number"/>
        <formitem :customLayout="true" label="">
            <button type="submit" class="btn btn-success" :disabled="loading">
                <b-spinner small v-if="loading"></b-spinner>
                Save
            </button>
            <button class="btn btn-danger ml-1" @click="Reset">Cancel</button>
        </formitem>
    </form>
`
    }
