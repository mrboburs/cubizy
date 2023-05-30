
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                password_set : false,
                resetcode_on: "Email",
                resetcode_sent: 0,
                username: "",
                resetcode: "",
                password: "",
                conform_password: "",
                showpassword: false,
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
            canSignUp: function () {
                if(window.location.host == "seller."+  window.application.BaseDomin || window.location.host == "student."+  window.application.BaseDomin){
                    return true
                }
            },
            usernameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.username.trim()) {
                    return "Please provide a valid " + this.resetcode_on
                }
            },
            resetcodeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.resetcode.trim()) {
                    return "Please provide a valid reset code"
                }
            },
            passwordError: function () {
                if (!this.submitted) {
                    return false
                }
                var password = this.password.trim()
                if (!password) {
                    return "Please provide a valid password"
                }
                if (!/[a-z]/.test(password)) {
                    return "Password must have atlist one small character"
                }
                if (!/[A-Z]/.test(password)) {
                    return "Password must have atlist one capital character"
                }
                if (!/[0-9]/.test(password)) {
                    return "Password must have atlist one number"
                }
            },
            conform_passwordError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.conform_password.trim()) {
                    return "Please enter same password again in conform password"
                }
                if (this.conform_password != this.password) {
                    return "Conform password and password should be exactly same"
                }
            },
        },
        methods: {
            submit() {
                this.submitted = true
                if (!this.resetcode_sent) {
                    this.send_resetcode()
                } else {
                    this.reset_password()
                }
            },
            can_login() {
                if (_.isEmpty(this.username.trim()) || _.isEmpty(this.password.trim())) {
                    return false
                }
                return true
            },
            send_resetcode() {
                if (this.usernameError) {
                    return
                }
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "resetcode",
                    data: {
                        resetcode_on: this.resetcode_on,
                        username: this.username,
                    }
                }).then(function (data) {
                    component.message = data.Message;
                    if (data.Status == 2) {
                        component.resetcode_sent++
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
            reset_password() {
                if (this.usernameError || this.resetcodeError || this.passwordError || this.conform_passwordError) {
                    return
                }
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "resetpassword",
                    data: {
                        resetcode_on: this.resetcode_on,
                        username: this.username,
                        resetcode: this.resetcode,
                        password: this.password,
                    }
                }).then(function (data) {
                    component.message = data.Message;
                    if (data.Status == 2) {
                        component.password_set = true
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
            if (this.$route.query.email) {
                this.username = this.$route.query.email
                this.resetcode_on = "Email"
            }
            if (this.$route.query.mobile) {
                this.username = this.$route.query.mobile
                this.resetcode_on = "Mobile"
            }
            if (this.$route.query.code) {
                this.resetcode = this.$route.query.code
                this.resetcode_sent = 1
            }
        },
        template: `
    <div>
        <!-- title-->
        <h4 class="mt-0">Reset Password</h4>
        <p v-if="!resetcode_sent" class="text-muted mb-4">Please enter your email or mobile address and get reset code
            to reset password.</p>
        <p v-if="resetcode_sent" class="text-muted mb-4">Please enter your reset code , new password and conform
            password.</p>
        <!-- form -->
        <div v-if="message" class="alert alert-dismissible fade show"
            :class="{'alert-success': !error, 'alert-danger': error }">
            <strong v-if="!error">Success!</strong>
            <strong v-if="error">Eror!</strong>
            {{message}}
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
                @click="message = false"></button>
        </div>
        <divloading :fullpage="false" :loading="loading" class="modal-body">
            <form v-if="!password_set" action="#" @submit.prevent="submit">
                <div v-if="!resetcode_sent" class="mb-2">
                    <label class="form-label">Get Reset code on... </label>
                    <div class="form-check form-check-inline">
                        <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadioEmail"
                            value="Email" v-model="resetcode_on">
                        <label class="form-check-label" for="inlineRadioEmail">Email</label>
                    </div>
                    <div class="form-check form-check-inline">
                        <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadioMobile"
                            value="Mobile" v-model="resetcode_on">
                        <label class="form-check-label" for="inlineRadioMobile">Mobile</label>
                    </div>
                </div>
                <div class="mb-2">
                    <label for="inputusername" class="form-label">{{resetcode_on}}</label>
                    <input class="form-control" :class="{'is-invalid': usernameError }" id="inputusername"
                        :placeholder="'Enter your '+ resetcode_on" max="30" v-model="username">
                    <div v-if="usernameError" class="invalid-feedback">
                        {{usernameError}}
                    </div>
                </div>
                <div class="d-grid text-center">
                    <button v-if="!resetcode_sent" class="btn btn-primary" type="submit">Get reset code </button>
                </div>
                <div v-if="resetcode_sent" class="mb-2">
                    <label for="inputresetcode" class="form-label">Reset Code</label>
                    <input class="form-control" :class="{'is-invalid': resetcodeError }" id="inputresetcode"
                        :placeholder="'Enter reset code recived on '+ resetcode_on" max="30" v-model="resetcode">
                    <div v-if="resetcodeError" class="invalid-feedback">
                        {{resetcodeError}}
                    </div>
                    <a v-if="resetcode_sent" href="#" class="text-muted float-end" @click.prevent="resend">Resend reset
                        code</a>
                </div>
                <div v-if="resetcode_sent" class="mb-2">
                    <label for="password" class="form-label">Password</label>
                    <input type="password" id="password" class="form-control" placeholder="Enter new password"
                        maxlength="30" v-model="password" :class="{'is-invalid': passwordError }">
                    <div v-if="passwordError" class="invalid-feedback">
                        {{passwordError}}
                    </div>
                </div>
                <div v-if="resetcode_sent" class="mb-2">
                    <label for="inputconform_password" class="form-label">Conform Password</label>
                    <div class="input-group input-group-merge">
                        <input class="form-control" :class="{'is-invalid': conform_passwordError }"
                            :type="showpassword?'':'password'" id="inputconform_password"
                            placeholder="Enter password once again same as above" maxlength="30"
                            v-model="conform_password">
                        <div class="input-group-text" :class="{'show-password': showpassword}" data-password="false"
                            @click="showpassword = !showpassword">
                            <span class="password-eye"></span>
                        </div>
                        <div v-if="conform_passwordError" class="invalid-feedback">
                            {{conform_passwordError}}
                        </div>
                    </div>
                </div>
                <div v-if="resetcode_sent" class="d-grid text-center">
                    <button class="btn btn-primary" type="submit">Set New Password </button>
                </div>
            </form>
            <router-link v-else to="/auth/login" class="btn btn-primary w-50 m-auto d-block" type="submit">Login </router-link>
        </divloading>
        <!-- end form-->
        <!-- Footer-->
        <footer>
            <p class="text-muted">
                Remember password ?
                <router-link to="/auth/login" class="text-muted float-end">
                    <small> Login </small>
                </router-link>
            </p>
            <p class="text-muted" v-if="canSignUp">
                Don't have an account?
                <router-link to="/auth/register" class="text-muted float-end">
                    <small> Sign Up </small>
                </router-link>
            </p>
        </footer>
    </div>
`
    }
