<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                SellerTotal: 0,
                SellerActive: 0,

                SessionTotal: 0,
                SessionActive: 0,

                StudentTotal: 0,
                StudentActive: 0,

                SMSBalance: 0,
                SMSSent: 0,

                EmailBalance: 0,
                EmailSent: 0,
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
            last_message: function (newValue, oldValue) {
                this.set_last_message()
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'last_message', 'last_message_error']),
            ...Vuex.mapGetters(['getFullDateTime']),
        },
        methods: {
            set_last_message() {
                if (this.last_message) {
                    this.message = this.last_message
                    this.error = this.last_message_error
                    this.$store.commit('set_last_message', false)
                }
            },
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "dashbord",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if(data.Result.SellerTotal){
                            this.SellerTotal = data.Result.SellerTotal
                            this.SellerActive = data.Result.SellerActive

                            this.SMSBalance = data.Result.SMSBalance
                            this.SMSSent = data.Result.SMSSent

                            this.EmailBalance = data.Result.EmailBalance
                            this.EmailSent = data.Result.EmailSent
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
            },
        },
        mounted: function () {
            this.set_last_message()
            this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>

    <div class="col-12">
        <v-alert v-model="message" :error="error" />
        <div class="row">
            <div class="col">
                <div class="card">
                    <div class="card-body">
                        <h3>Welcome admin</h3>
                        <p class="fs-4">{{getFullDateTime(new Date()/1000)}}</p>
                    </div>
                </div>
            </div><!-- end col -->
            <div class="col-auto col-md-3">
                <divloading :fullpage="false" :loading="loading" class="card">
                    <div class="card-body">
                        <div class="d-flex justify-content-center align-items-center flex-wrap">
                            <div class="knob-chart" dir="ltr">
                                <v-knobcontrol :value="SellerActive" :max="SellerTotal"></v-knobcontrol>
                            </div>
                            <div class="text-end">
                                <h3 class="mb-1 mt-0"> <span data-plugin="counterup">{{SellerTotal}}</span> </h3>
                                <p class="text-muted mb-0">Seller Accounts</p>
                            </div>
                        </div>
                    </div>
                </divloading>
            </div><!-- end col -->
        </div>
        <!-- end row -->
    </div>
</template>