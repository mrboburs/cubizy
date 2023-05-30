
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                name: "",
                email: "",
                mobile: "",
                password: "",
                showpassword: false,
                regMobileExp: /^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\s\./0-9]*$/,
                regEmailExp: /\S+@\S+\.\S+/,
            }
        },
        watch: { 
        },
        computed: {
            nameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.name.trim()) {
                    return "Name can not be empty"
                }
            },
            emailError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.email.trim()) {
                    return "Email can not be empty"
                } else if (!this.regEmailExp.test(this.email)) {
                    return "Please enter valid email"
                } else {
                    return ""
                }
            },
            mobileError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.mobile.trim()) {
                    return "Mobile can not be empty"
                } else if (!this.regMobileExp.test(this.mobile)) {
                    return "Please enter valid mobile number"
                } else {
                    return ""
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
        },
        methods: {
            register() {
                this.submitted = true
                if (this.nameError || this.emailError || this.mobileError || this.passwordError) {
                    return
                }
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "register",
                    data: {
                        Name: this.name,
                        Email: this.email,
                        Mobile: this.mobile,
                        Password: this.password,
                    }
                }).then(function (data) {
                    component.message = data.Message;
                    if (data.Status == 2) {
                        component.error = false
                        component.$router.push('/')
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
    <div>
        <!-- title-->
        <h4 class="mt-0">Sign Up</h4>
        <p class="text-muted mb-4">Enter your name, email address and mobile number to start signup.</p>
        <div v-if="message" class="alert alert-dismissible fade show" role="alert"
            :class="{'alert-danger': error, 'alert-success': !error }">
            <strong v-if="!error">Success!</strong>
            <strong v-if="error">Eror!</strong>
            {{message}}
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
                @click="message = false"></button>
        </div>
        <!-- form -->
        <form action="#" @submit.prevent="register">
            <div class="form-floating mb-2">
                <input id="inputname" placeholder="name" class="form-control" :class="{'is-invalid': nameError }"
                    max="50" v-model="name">
                <label for="inputname">Name</label>
                <div v-if="nameError" class="invalid-feedback">
                    {{nameError}}
                </div>
            </div>
            <div class="form-floating mb-2">
                <input type="email" id="inputemailaddress" placeholder="name@example.com" class="form-control"
                    :class="{'is-invalid': emailError }" max="50" v-model="email">
                <label for="inputemailaddress">Email address</label>
                <div v-if="emailError" class="invalid-feedback">
                    {{emailError}}
                </div>
            </div>
            <div class="form-floating mb-2">
                <input id="inputmobile" placeholder="(+0x) xxxxx xxxxx" class="form-control"
                    :class="{'is-invalid': mobileError }" max="50" v-model="mobile">
                <label for="inputmobile">Mobile number</label>
                <div v-if="mobileError" class="invalid-feedback">
                    {{mobileError}}
                </div>
            </div>
            <div class="form-floating mb-2">
                <input :type="showpassword?'':'password'" id="inputpassword" class="form-control"
                    placeholder="Enter your password" type="password" maxlength="30" v-model="password"
                    :class="{'is-invalid':passwordError }">
                <div class="input-group-text" :class="{'show-password': showpassword}" data-password="false"
                    @click="showpassword = !showpassword">
                    <span class="password-eye"></span>
                </div>
                <label for="inputpassword">Password</label>
                <div v-if="passwordError" class="invalid-feedback">
                    {{passwordError}}
                </div>
            </div>
            <div class="d-grid text-center">
                <button class="btn btn-primary" type="submit">Create Account</button>
            </div>
        </form>
        <!-- end form-->

        <!-- Footer-->
        <footer class="m-3">
            <p class="text-muted">
                Have an account?
                <router-link to="/auth/login" class="text-muted float-end">
                    <small> Login </small>
                </router-link>
            </p>
        </footer>
    </div>
`
    }
