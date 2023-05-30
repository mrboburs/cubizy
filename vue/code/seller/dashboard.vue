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
            ...Vuex.mapState(['user', 'account', 'last_message', 'last_message_error']),
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
                        this.SessionActive = data.Result.SessionActive
                        this.StudentActive = data.Result.StudentActive
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
    <!-- Start Content-->
    <div class="container-fluid">
        <v-alert v-model="message" :error="error" />
        <div class="row">
            <div class="col">
                <div class="card">
                    <div class="card-body">
                        <h3>{{account.Title}}</h3>
                        <div class="summary">
                            <dt class="">AccountType : </dt>
                            <dd class=" text-capitalize"> {{account.AccountType}} </dd>
                            <span class="clearfix"></span>
                            <dt class="">Status : </dt>
                            <dd class=" text-capitalize">
                                <span class="text-capitalize text-warning" v-if="account.Status == 1">Under
                                    Review</span>
                                <span class="text-capitalize text-danger" v-if="account.Status == 2">Rejected</span>
                                <span class="text-capitalize text-warning" v-if="account.Status == 3">OnHold</span>
                                <span class="text-capitalize text-danger" v-if="account.Status == 4">Expired</span>
                                <span class="text-capitalize text-success" v-if="account.Status == 10">Active</span>

                                <div v-if="10 > account.Status" class="d-flex flex-column">
                                    <span v-if="account.IDStatus" class="text-success"> Id Verified <i
                                            class="far fa-check-circle"></i> </span>
                                    <span v-if="!account.IDStatus" class="text-danger"> Id not Verified <i
                                            class="far fa-times-circle"></i></span>
                                    <span v-if="account.AddressStatus" class="text-success"> Address Verified <i
                                            class="far fa-check-circle"></i></span>
                                    <span v-if="!account.AddressStatus" class="text-danger"> Address not Verified <i
                                            class="far fa-times-circle"></i></span>
                                    <span v-if="account.RegistretionStatus" class="text-success"> Registretion Verified
                                        <i class="far fa-check-circle"></i></span>
                                    <span v-if="!account.RegistretionStatus" class="text-danger"> Registretion not
                                        Verified <i class="far fa-times-circle"></i></span>
                                </div>
                            </dd>
                            <span class="clearfix"></span>
                            <dt class="" v-if="account.Status == 10">Website </dt>
                            <dd class=" text-capitalize">
                                <span v-if="account.CanActive" class="text-success">Available</span>
                                <span v-if="!account.CanActive" class="text-warning">Not Available</span>
                            </dd>
                        </div>
                        <span class="clearfix"></span>
                        <span class="text-capitalize" v-if="account.StatusComment">{{account.StatusComment}}</span>
                    </div>
                </div>
            </div><!-- end col -->

            <div v-if="account.Sessions" class="col-auto col-md-3">
                <divloading :fullpage="false" :loading="loading" class="card">
                    <div class="card-body">
                        <div class="d-flex justify-content-center align-items-center flex-wrap">
                            <div class="knob-chart" dir="ltr">
                                <v-knobcontrol :value="SessionActive" :max="account.Sessions"></v-knobcontrol>
                            </div>
                            <div class="text-end">
                                <h3 class="mb-1 mt-0"> <span data-plugin="counterup">{{account.Sessions}}</span> </h3>
                                <p class="text-muted mb-0">Cources</p>
                            </div>
                        </div>
                    </div>
                </divloading>
            </div>
            <!-- end col -->

            <div v-if="account.Students" class="col-auto col-md-3">
                <divloading :fullpage="false" :loading="loading" class="card">
                    <div class="card-body">
                        <div class="d-flex justify-content-center align-items-center flex-wrap">
                            <div class="knob-chart" dir="ltr">
                                <v-knobcontrol :value="StudentActive" :max="account.Students"></v-knobcontrol>
                            </div>
                            <div class="text-end">
                                <h3 class="mb-1 mt-0"><span data-plugin="counterup">{{account.Students}}</span> </h3>
                                <p class="text-muted mb-1">Students</p>
                            </div>
                        </div>
                    </div>
                </divloading>
            </div>
            <!-- end col -->

            <div class="col">
                <div class="card">
                    <div class="card-body">
                        <h3>Checklist !</h3>
                        <div class="d-flex flex-column">
                            <span v-if="account.Keywords" class="text-success" v-tooltip:bottom="'SEO, Helps user to find your seller more faster and easily.'">
                                <i class="far fa-check-circle"></i>
                                SEO Keywords added
                            </span>
                            <router-link v-if="!account.Keywords" class="text-danger" to="/account"
                                v-tooltip:bottom="'SEO, Helps user to find your seller more faster and easily.'">
                                <i class="far fa-times-circle"></i>
                                Add SEO Keywords
                            </router-link>
                            <span v-if="account.WideLogo" class="text-success" v-tooltip:bottom="'Logo with name in styled font.'">
                                <i class="far fa-check-circle"></i>
                                Widelogo added
                            </span>
                            <router-link v-if="!account.WideLogo" class="text-danger" to="/account"
                                v-tooltip:bottom="'Logo with name in styled font.'">
                                <i class="far fa-times-circle"></i>
                                Add Widelogo
                            </router-link>
                            <span v-if="account.Banner" class="text-success" v-tooltip:bottom="'Banner image used on Details card on Social Media, Seller listing etc.'">
                                <i class="far fa-check-circle"></i>
                                Banner added
                            </span>
                            <router-link v-if="!account.Banner" class="text-danger" to="/account"
                                v-tooltip:bottom="'Banner image used on Details card on Social Media, Seller listing etc.  '">
                                <i class="far fa-times-circle"></i>
                                Add Banner
                            </router-link>
                            <span
                                v-if="account.Youtube || account.Facebook || account.Instagram || account.Pinterest || account.WhatsApp"
                                class="text-success">
                                <i class="far fa-check-circle"></i>
                                <span v-tooltip:bottom="'FB link, Youtube link, Instagram link etc '">Social media links
                                    added</span>
                            </span>
                            <router-link to="/account"
                                v-if="!account.Youtube && !account.Facebook && !account.Instagram && !account.Pinterest && !account.WhatsApp"
                                class="text-danger">
                                <i class="far fa-times-circle"></i>
                                <span v-tooltip:bottom="'FB link, Youtube link, Instagram link etc '">Add social media
                                    links</span>
                            </router-link>
                        </div>
                        <div v-if="account.Status > 9" class="d-flex flex-column">
                            <span v-if="account.Subjects" class="text-success">
                                <i class="far fa-check-circle"></i>
                                Subjected Selected
                            </span>
                            <router-link to="/subjects" v-if="!account.Subjects" class="text-danger">
                                <i class="far fa-times-circle"></i>
                                Tell us what subjects you teach 
                            </router-link>
                            <span v-if="account.Sessions" class="text-success">
                                <i class="far fa-check-circle"></i>
                                Cources Added
                            </span>
                            <router-link to="/sessions" v-if="!account.Sessions" class="text-danger">
                                <i class="far fa-times-circle"></i>
                                Add your first cource
                            </router-link>
                            <span v-if="account.Students" class="text-success">
                                <i class="far fa-check-circle"></i>
                                Student Added
                            </span>
                            <router-link to="/sessions" v-if="!account.Students" class="text-danger">
                                <i class="far fa-times-circle"></i>
                                Add your first student
                            </router-link>
                            <span v-if="account.Notes" class="text-success">
                                <i class="far fa-check-circle"></i>
                                Noted Added
                            </span>
                            <router-link to="/notes" v-if="!account.Notes" class="text-danger">
                                <i class="far fa-times-circle"></i>
                                Add your first note
                            </router-link>
                            <span v-if="account.Subdomain && account.Active" class="text-success">
                                <i class="far fa-check-circle"></i>
                                Website Activeted
                            </span>
                            <router-link to="/notes" v-if="!account.Subdomain || !account.Active" class="text-danger">
                                <i class="far fa-times-circle"></i>
                                Activate Website
                            </router-link>
                        </div>
                    </div>
                </div>
            </div><!-- end col -->
        </div>
        <!-- end row -->
    </div> <!-- container -->
</template>