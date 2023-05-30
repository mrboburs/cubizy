<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                AppName : window.application.AppName
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
            AddressProofError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.AddressProof) {
                    return "AddressProof can not  be empty"
                }
            },
            RegistretionProofError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.RegistretionProof) {
                    return "RegistretionProof can not  be empty"
                }
            },
        },
        methods: {
            submit() {
                this.submitted = true
                this.loading = true
                this.$store.dispatch('call', {
                    api: "me",
                    data: {
                        user: {
                            Joined: true
                        }
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
            }
        },
        mounted: function () {
            if(this.user.Joined && this.account.Status > 0){
                this.$router.push('/')
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <divloading :fullpage="false" :loading="loading" class="row">
        <v-alert v-model="message" :error="error" />
        <div v-if="user.ID == account.CreatedBy && account.Status == 0">
            <div class="mt-5 text-center m-auto w-50">
                <h1 class="m-auto">THANK YOU FOR CHOOSING {{AppName}}</h1>
                <p class="mt-5 text-left">
                    Admin will review your documents and activate your account fully within 2
                    business days, will contact you if having any problem.
                </p>
                <p class="mt-5 text-left">
                    Meanwhile please complete your business and personal profile details.
                </p>
                <v-alert v-model="message" :error="error" />
            </div>
            <div class="d-flex justify-content-between">
                <router-link to="/setup/documents" class="btn btn-primary">
                    Back
                </router-link>
                <button type="submit" class="btn btn-success" @click="submit">
                    Finish
                </button>
            </div>
        </div>
        <div v-else-if="account.Status == 0">
            <div class="mt-5 text-center m-auto w-50">
                <h2 class="m-auto">Welcome {{user.Name}}</h2>
                <h4 class="mt-5">Please contact "{{account.Title}}" account owner and ask to complete the setup process.</h4>
                <button type="submit" class="btn btn-success" @click="submit">
                    Finish
                </button>
            </div>
        </div>
        <div v-else-if="!user.Joined">
            <div class="mt-5 text-center m-auto w-50">
                <h2 class="m-auto">Welcome {{user.Name}}</h2>
                <h4 class="mt-5">You have successfully joined "{{account.Title}}" account. </h4>
                <button type="submit" class="btn btn-success" @click="submit">
                    Finish
                </button>
            </div>
        </div>
    </divloading>
</template>