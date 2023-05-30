
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Email: "",
                EmailChanged: false,
                EmailCode: "",
                SentOnInterval: 60,
                LastSentOn: JSON.parse(localStorage.getItem('LastSentOn')),
                SentOnTimes: JSON.parse(localStorage.getItem('SentOnTimes')),
                canSendCodeIn: 0,
                regEmailExp: /\S+@\S+\.\S+/,
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
            user: function (newValue, oldValue) {
                if (newValue) {
                    this.init()
                }
            },
            SentOnTimes: function (newValue, oldValue) {
                localStorage.setItem("SentOnTimes", this.SentOnTimes)
            },
            Email: function (newValue, oldValue) {
                if (newValue != this.user.Email) {
                    this.EmailChanged = true
                } else {
                    this.EmailChanged = false
                }
            }
        },
        computed: {
            ...Vuex.mapState(["user"]),
            EmailError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Email.trim()) {
                    return "Please provide a valid Email"
                } else if (!this.regEmailExp.test(this.Email)) {
                    return "Please enter valid email"
                } else {
                    return false
                }
            },
            EmailCodeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.EmailCode.trim()) {
                    return "Please enter email virification code"
                }
            },
        },
        methods: {
            init() {
                if (this.user) {
                    if (this.user.EmailVerified) {
                        if (!this.message) {
                            this.message = "Email verified"
                        }
                    } else if (this.Email != this.user.Email) {
                        this.Email = this.user.Email
                        this.EmailCode = ""
                        this.EmailChanged = false
                        if (this.user.EmailCodeSet) {
                            this.canSendCodeIn = this.SentOnInterval
                        }
                    }
                }
                setInterval(() => {
                    if (this.canSendCodeIn > 0) {
                        this.canSendCodeIn--
                    }
                }, 1000);
            },
            sendVirificationEmail() {
                this.submitted = true
                if (this.canSendCodeIn || this.EmailError) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "sendvirificationemail",
                    data: {
                        Email: this.Email
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
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
            verifyEmail() {
                this.submitted = true
                if (this.EmailCodeError || this.EmailError) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "verifyemail",
                    data: {
                        Email: this.Email,
                        EmailCode: this.EmailCode
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
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
            this.init()
        },
        template: `
    <div class="m-auto w-50">
        <!-- title-->
        <h4 class="mt-0">Email virification</h4>
        <p class="text-muted mb-4">Please verify your email to use this application.</p>
        <div v-if="message" class="alert d-flex align-items-center alert-dismissible fade show"
            :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
            <strong v-if="!error">Success: </strong>
            <strong v-if="error">Error: </strong>
            <span class="ms-1" v-html="message"></span>
            <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
        </div>
        <divloading :fullpage="false" :loading="loading" class="container" v-if="!user.EmailVerified">
            <formitem name="inputEmail" label="Email" :error="EmailError" v-model="Email" />
            <div class="d-grid text-center mb-3">
                <button type="button" class="btn btn-outline-primary" @click="sendVirificationEmail"
                    :disabled="!!canSendCodeIn">
                    <span v-if="EmailChanged">Update Email & Get verification code on email</span>
                    <span v-else-if="!user.EmailCodeSet">Get verification code on email</span>
                    <span v-else-if="canSendCodeIn">Can resend code in {{canSendCodeIn}} seconds</span>
                    <span v-else>Resend code</span>
                </button>
            </div>
            <formitem v-if="user.EmailCodeSet" name="inputEmailCode" label="Email Verification Code"
                :error="EmailCodeError" v-model="EmailCode" />
            <div v-if="user.EmailCodeSet" class="d-grid text-center mb-3">
                <button type="button" class="btn btn-primary" @click="verifyEmail">
                    Verify your Email
                </button>
            </div>
        </divloading>
        <div v-else>
            <div class="d-grid text-center mb-3">
                <router-link to="setup/verifymobile" class="btn btn-primary">
                    Next
                </router-link>
            </div>
        </div>
    </div>
`
    }
