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
            user: function (newValue, oldValue) {
                if (newValue) {
                    this.init()
                }else{
                    this.$router.push('/auth/login')
                }
            },
        },
        computed: {
            ...Vuex.mapState(["user", "account", "last_message", "last_message_error"]),
            year : function (){
                return new Date().getFullYear()
            }
        },
        methods: {
            logout() {
                if (!this.user) { return }
                this.$store.dispatch('call', { api: "logout", data: {} })
            },
            init(){
                if(!this.user.EmailVerified){

                }
            },
        },
        mounted: function () {
            //this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <section class="">
        <div class="account-pages mt-5 mb-5">
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-xl-9">
                        <div class="card">
                            <div class="card-body p-4">
                                <div class="text-center mb-4">
                                    <div class="auth-logo">
                                        <a href="index.html" class="logo logo-dark text-center">
                                            <span class="logo-lg">
                                                <img src="/assets/images/logo-dark.png" alt="" height="22">
                                            </span>
                                        </a>
                                    </div>
                                    <h1>Account setup</h1>
                                </div>
                                <transition name="custom-classes-transition" mode="out-in"
                                    enter-active-class="animate__animated animate__slideInRight animate__faster">
                                    <router-view></router-view>
                                </transition>
                                <!-- end row-->
                            </div> <!-- end card-body -->
                        </div>
                        <!-- end card -->
                        <div class="row mt-3">
                            <div class="col-12 text-center">
                                <p class="text-muted">You can <a href="#" @click.prevent="logout" > click here to logout</a>  and login again later to continue setup any time....</a></p>
                            </div> <!-- end col -->
                        </div>
                    </div> <!-- end col -->
                </div>
                <!-- end row -->
            </div>
            <!-- end container -->
        </div>
        <!-- end page -->
        <footer class="footer footer-alt">
            {{year}} © Powered by <a href=""
                class="text-dark">{{AppName}}</a>
        </footer>
    </section>
</template>