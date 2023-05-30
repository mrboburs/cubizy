<script>
    export default {
        props: {
            is_model: {
                default: false,
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                validation: {
                    username: "",
                    password: "",
                },
                username: "",
                password: "",
                rememberme: "",
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
            username: function (newValue, oldValue) {
                if (_.isEmpty(newValue.trim())) {
                    this.validation.username = "Email/Mobile can not be empty"
                } else {
                    this.validation.username = ""
                }
            },
            password: function (newValue, oldValue) {
                if (_.isEmpty(newValue.trim())) {
                    this.validation.password = "Password can not be empty"
                } else {
                    this.validation.password = ""
                }
            }
        },
        computed: {
            canSignUp: function () {
                if(window.location.host == "seller."+  window.application.BaseDomin || window.location.host == "student."+  window.application.BaseDomin){
                    return true
                }
            }
        },
        methods: {
            can_login() {
                if (_.isEmpty(this.username.trim()) || _.isEmpty(this.password.trim())) {
                    return false
                }
                return true
            },
            login() {
                if (!this.can_login()) { return }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "login",
                    data: {
                        username: this.username,
                        password: this.password,
                        rememberme: this.rememberme,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if(!this.is_model){
                            this.$router.push('/')
                            this.tab_index = 0
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
            if (this.$route.query.email) {
                this.username = this.$route.query.email
            }
            if (this.$route.query.mobile) {
                this.username = this.$route.query.mobile
            }
            if (this.$route.query.code) {
                this.password = this.$route.query.code
                if(this.password && this.username){
                    this.rememberme = false
                    this.login()
                }
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div>
        <v-alert v-model="message" :error="error" />
        <!-- title-->
        <h4 class="mt-0">Sign In</h4>
        <p class="text-muted mb-4">Enter your email address and password to access admin panel.</p>
        <!-- form -->
        <form action="#" @submit.prevent="login">

            <div class="form-floating mb-2">
                <input type="email" id="inputusername" placeholder="name@example.com" class="form-control"
                    :class="{'is-invalid': validation.username }" max="50" v-model="username">
                <label for="inputusername">Email / Mobile</label>
                <div v-if="validation.username" class="invalid-feedback">
                    {{validation.username}}
                </div>
            </div>
            <div class="form-floating mb-2">
                <input :type="showpassword?'':'password'" id="inputpassword" class="form-control"
                    placeholder="Enter your password" type="password" maxlength="30" v-model="password"
                    :class="{'is-invalid': validation.password }">
                <div class="input-group-text" :class="{'show-password': showpassword}" data-password="false"
                    @click="showpassword = !showpassword">
                    <span class="password-eye"></span>
                </div>
                <label for="inputpassword">Password</label>
                <div v-if="validation.password" class="invalid-feedback">
                    {{validation.password}}
                </div>
            </div>
            <div class="mb-2">
                <router-link to="/auth/recoverpw" class="text-muted float-end"><small> Forgot your password? </small>
                </router-link>
            </div>

            <div class="mb-3">
                <div class="form-check">
                    <input class="form-check-input" type="checkbox" id="checkbox-signin" v-model="rememberme">
                    <label class="form-check-label" for="checkbox-signin">
                        Remember me
                    </label>
                </div>
            </div>
            <div class="d-grid text-center">
                <button class="btn btn-primary" type="submit">Log In </button>
            </div>
        </form>
        <!-- end form-->

        <!-- Footer-->
        <footer class="m-3">
            <p class="text-muted" v-if="canSignUp">
                Don't have an account?
                <router-link to="/auth/register" class="text-muted float-end">
                    <small> Sign Up </small>
                </router-link>
            </p>
        </footer>
    </div>
</template>