
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                popupitem : false
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
            showpopup: {
                // getter
                get: function () {
                    if (this.popupitem) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.popupitem = false
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'danger'
                } else {
                    return 'success'
                }
            }
        },
        methods: {
            load() {
                if(this.loading){
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "test",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.popupitem = data.Result.content
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
            },
        },
        mounted: function () {
            this.load()
            for (let index = 0; index < array.length; index++) {
                const element = array[index];
                
            }
        },
        template: `
    <section class="section is-medium">
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="container">
            <div class="columns is-vcentered">
                <div class="column has-text-centered">
                    <h1 class="title">Hi , this is test page</h1>
                    <p class="subtitle">Server works fine .</p>
                    <a class="button">Home</a>
                    <a class="button">Contact</a>
                </div>
            </div>
        </divloading>
        <v-modal  v-model="showpopup" title="Popup Message"
            header-close-variant="light" title-class="font-18" hide-footer>
            {{popupitem}}
        </v-modal>
    </section>
`
    }
