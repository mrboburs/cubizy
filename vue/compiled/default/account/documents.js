
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                IDProof: "",
                AddressProof: "",
                RegistretionProof: "",
                OtherDocument: "",
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
            IDProofError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.IDProof) {
                    return "IDProof can not  be empty"
                }
            },
            AddressProofError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.AddressProof) {
                    return "AddressProof can not  be empty"
                }
            },
            RegistretionProofError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.RegistretionProof) {
                    return "RegistretionProof can not  be empty"
                }
            },
        },
        methods: {
            Reset() {
                this.SetData()
            },
            SetData() {
                this.submitted = false
                if (this.account && this.account.ID > 0) {
                    this.submitted = false
                    this.IDProof = this.account.IDProof
                    this.AddressProof = this.account.AddressProof
                    this.RegistretionProof = this.account.RegistretionProof
                    this.OtherDocument = this.account.OtherDocument
                }
            },
            submit() {
                this.submitted = true
                if (this.IDProofError || this.AddressProofError || this.RegistretionProofError) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        account: {
                            IDProof: this.IDProof.toString(),
                            AddressProof: this.AddressProof.toString(),
                            RegistretionProof: this.RegistretionProof.toString(),
                            OtherDocument: this.OtherDocument.toString()
                        }
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if (this.message.trim() && this.$route.name == 'setdocuments') {
                            this.$router.push('/setup/welcome')
                        }
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                })
            }
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload')
        },
        template: `
    <form @submit.prevent="submit" class="card container-fluid">
        <h4>Documents for account verification</h4>
        <p class="text-muted mb-4" v-if="$route.name == 'setdocuments'">Please enter valid account detials that will be visible on websit listing .</p>
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="row">
            <div class="col-6 col-md-4 col-lg-3">
                <formitem name="inputIDProof" type="file" label="IDProof *" :error="IDProofError">
                    <v-files-static @oncount="IDProof = $event" :only_image="false"
                        :prefix="'account/idproof'" />
                </formitem>
            </div>
            <div class="col-6 col-md-4 col-lg-3">
                <formitem name="inputAddressProof" type="file" label="AddressProof *" :error="AddressProofError">
                    <v-files-static @oncount="AddressProof = $event" :only_image="false"
                        :prefix="'account/addressproof'" />
                </formitem>
            </div>
            <div class="col-6 col-md-4 col-lg-3">
                <formitem name="inputRegistretionProof" type="file" label="RegistretionProof *"
                    :error="RegistretionProofError">
                    <v-files-static @oncount="RegistretionProof = $event" :only_image="false"
                        :prefix="'account/registretionproof'" />
                </formitem>
            </div>
            <div class="col-6 col-md-4 col-lg-3">
                <formitem name="inputOtherDocument" type="file" label="OtherDocument *">
                    <v-files-static @oncount="OtherDocument = $event" :only_image="false"
                        :prefix="'account/otherdocuments'" />
                </formitem>
            </div>
            <div v-if="$route.name == 'setdocuments'" class="d-flex justify-content-between">
                <router-link to="/setup/location" class="btn btn-primary">
                    Back
                </router-link>
                <button type="submit" class="btn btn-success" :disabled="loading">
                    Next
                </button>
            </div>
            <div v-else class="d-flex justify-content-end">
                <button type="submit" class="btn btn-success m-1" :disabled="loading">
                    <b-spinner small v-if="loading"></b-spinner>
                    Save
                </button>
            </div>
        </divloading>
    </form>
`
    }
