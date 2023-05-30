
    export default {
        components: {
            AddressEditor: () => import("/vue/addresseditor.js"),
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                accountAddress: {
                    ID: 0,
                    Title: "Account Address",
                    Mobile: "",
                    AddressLine1: "",
                    AddressLine2: "",
                    AddressLine3: "",
                    Longitude: "",
                    Latitude: "",
                    Code: "",
                    SubLocality: false,
                    Locality: false,
                    District: false,
                    Country: false,
                }
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
        },
        computed: {
            showmessage: {
                // getter
                get: function () {
                    if (this.message) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.message = ""
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'alert-danger'
                } else {
                    return 'alert-success'
                }
            }
        },
        methods: {
            submit(address) {
                if(address){
                    this.load({
                        address : address
                    })
                }
            },
            load(data) {
                if (!data) {
                    data = {}
                }
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "accountaddress",
                    data: data
                }).then((data) => {
                    component.message = data.Message;
                    if (data.Status == 2) {
                        component.error = false
                        if (data.Result.address) {
                            component.accountAddress = data.Result.address
                        }
                        if(component.message.trim() && this.$route.name == 'setlocation'){
                            this.$router.push('/setup/documents')
                        }
                    } else {
                        component.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    component.error = true
                    component.message = error
                }).finally(() => {
                    component.loading = false
                })
            },
        },
        mounted: function () {
            this.load()
        },
        template: `
    <section class="card section is-medium">
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="card-body">
            <p>Set your account's main location.</p>
            <AddressEditor :value="accountAddress" @input="submit">
            </AddressEditor>
        </divloading>
    </section>
`
    }
