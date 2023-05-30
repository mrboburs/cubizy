
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Mobile: "",
                MobileChanged: false,
                MobileCode: "",
                SentOnInterval: 60,
                LastSentOn: JSON.parse(localStorage.getItem('LastSentOn')),
                SentOnTimes: JSON.parse(localStorage.getItem('SentOnTimes')),
                canSendCodeIn: 0,
                regMobileExp: /^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\s\./0-9]*$/,
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
            Mobile: function (newValue, oldValue) {
                if (newValue != this.user.Mobile) {
                    this.MobileChanged = true
                } else {
                    this.MobileChanged = false
                }
            }
        },
        computed: {
            ...Vuex.mapState(["user"]),
            MobileError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Mobile.trim()) {
                    return "Please provide a valid Mobile"
                } else if (!this.regMobileExp.test(this.Mobile)) {
                    return "Please enter valid mobile"
                } else {
                    return false
                }
            },
            MobileCodeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.MobileCode.trim()) {
                    return "Please enter mobile virification code"
                }
            },
        },
        methods: {
            init() {
                if (this.user) {
                    if (this.user.MobileVerified) {
                        if (!this.message) {
                            this.message = "Mobile verified"
                        }
                    } else if (this.Mobile != this.user.Mobile) {
                        this.Mobile = this.user.Mobile
                        this.MobileCode = ""
                        this.MobileChanged = false
                        if (this.user.MobileCodeSet) {
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
            sendVirificationMobile() {
                this.submitted = true
                if (this.canSendCodeIn || this.MobileError) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "sendvirificationmobile",
                    data: {
                        Mobile: this.Mobile
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
            verifyMobile() {
                this.submitted = true
                if (this.MobileCodeError || this.MobileError) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "verifymobile",
                    data: {
                        Mobile: this.Mobile,
                        MobileCode: this.MobileCode
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
        <h4 class="mt-0">Mobile virification</h4>
        <p class="text-muted mb-4">Please verify your mobile.</p>
        <div v-if="message" class="alert d-flex align-items-center alert-dismissible fade show"
            :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
            <strong v-if="!error">Success: </strong>
            <strong v-if="error">Error: </strong>
            <span class="ms-1" v-html="message"></span>
            <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
        </div>
        <divloading :fullpage="false" :loading="loading" class="container" v-if="!user.MobileVerified">
            <formitem name="inputMobile" label="Mobile" :error="MobileError" v-model="Mobile" />
            <div class="d-grid text-center mb-3">
                <button type="button" class="btn btn-outline-primary" @click="sendVirificationMobile"
                    :disabled="!!canSendCodeIn">
                    <span v-if="MobileChanged">Update Mobile & Get verification code on mobile</span>
                    <span v-else-if="!user.MobileCodeSet">Get verification code on mobile</span>
                    <span v-else-if="canSendCodeIn">Can resend code in {{canSendCodeIn}} seconds</span>
                    <span v-else>Resend code</span>
                </button>
            </div>
            <formitem v-if="user.MobileCodeSet" name="inputMobileCode" label="Mobile Verification Code"
                :error="MobileCodeError" v-model="MobileCode" />
            <div v-if="user.MobileCodeSet" class="d-grid text-center mb-3">
                <button type="button" class="btn btn-primary" @click="verifyMobile">
                    Verify your Mobile
                </button>
            </div>
        </divloading>
        <div v-else>
            <div class="d-grid text-center mb-3">
                <router-link to="setup/account" class="btn btn-primary">
                    Next
                </router-link>
            </div>
        </div>
    </div>
`
    }
