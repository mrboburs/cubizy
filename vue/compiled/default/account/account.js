
    export default {
        components: {
            'AccountEditor': () => import("/vue/account/accounteditor.js"),
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
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
            ...Vuex.mapState(['user', 'account']),
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
            },
        },
        methods: {
            can_load() {
                if (this.loading) {
                    return false
                }
                return true
            },
            load(data) {
                if (!data) {
                    data = {}
                }
                data.fix_condition = {
                    created_by: this.user.ID
                }

            },
            submit(record) {
                if (!record) {
                    return
                }
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        account: record
                    }
                }).then((data)=>{
                    component.message = data.Message;
                    if (data.Status == 2) {
                        if(this.$route.name == 'creataccount'){
                            this.$router.push('/setup/location')
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
        },
        template: `
    <section class="section is-medium card">
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="card-body container-fluid">
            <h4 v-if="$route.name == 'creataccount'">Account details for listing</h4>
            <p v-if="$route.name == 'creataccount'" class="text-muted mb-4">Please enter valid account detials; that will be visible on websit listing .</p>
            <AccountEditor :value="account" @input="submit">
            </AccountEditor>
        </divloading>
    </section>
`
    }
