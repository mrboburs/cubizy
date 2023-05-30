<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Password: "",
                ConformPassword: "",
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
        },
        computed: {
            ...Vuex.mapState(["user"]),
            passwordError: function () {
                if (!this.submitted) {
                    return false
                }
                var password = this.Password.trim()
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
            ConformPasswordError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.ConformPassword.trim()) {
                    return "Please enter same password again in conform password"
                }
                if (this.ConformPassword != this.Password) {
                    return "Conform password and password should be exactly same"
                }
            },
        },
        methods: {
            init() {
                if (this.user) {
                }
            },
            submit() {
                this.submitted = true
                if (this.passwordError || this.ConformPasswordError) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "setpassword",
                    data: {
                        Password: this.Password,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$router.push('/')
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
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="m-auto w-50">
        <!-- title-->
        <h4 class="mt-0">Set Password</h4>
        <p class="text-muted mb-4">Please set password so you can login from next time. 
            If password not set right now, you will need to user forgot password next time to login again  </p>
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="container">
            <formitem name="inputPassword" label="Password" :error="passwordError" v-model="Password" type="password" />
            <formitem name="inputconformPassword" label="Conform Password" :error="ConformPasswordError" v-model="ConformPassword" />
            <div class="d-grid text-center mb-3">
                <button type="button" class="btn btn-primary" @click="submit">
                    Set Password
                </button>
            </div>
        </divloading>
    </div>
</template>