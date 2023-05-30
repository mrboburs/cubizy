
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                account: {
                    Title: "",
                    Content: "",
                    Image: "",
                },
                sellersubjects: [],
                sessions: [],
                sellers: [],
                previous: null,
                next: null,
                address: null,
                oldid: 0
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
            account: function (newValue) {
                scroll(0, 0)
            },
        },
        computed: {
            ...Vuex.mapGetters(['getAccountUrl', 'getFullDate', 'getMonth', 'getDuration', 'getDate', 'getFBLink', 'getTwitterLink', 'getGooglePlusLink', 'getPinterestLink']),
        },
        methods: {
            load() {
                if (this.loading) { return }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        id: this.account.ID
                    }
                }).then((data) => {
                    this.content = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.account = data.Result.account
                        this.sellers = data.Result.sellers
                        this.sessions = data.Result.sessions
                        this.sellersubjects = data.Result.sellersubjects
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
            if (window.application.Account.AccountType != 'admin') {
                this.account = window.application.Account
            }
            this.load()
        },
        template: `
    <div class="panel panel-default">
        <div class="panel-heading">
            <div class="member-single-post">
                <div class="row">
                    <div class="col-md-4">
                        <div class="item-thumbnail">
                            <img :src="account.Logo" alt="image"></a>
                        </div>
                    </div>

                    <div class="col-md-8">
                        <div class="content-pad">
                            <div class="item-content">
                                <h3 class="item-title">
                                    {{account.Title}}
                                </h3>
                                <h4 class="small-text">
                                    {{account.Keywords}}
                                </h4>
                                <div class="row">
                                    <div class="col-sm-12 col-md-6">
                                        <div class="member-tax small-text">
                                            <h4 class="h4">Experience : </h4>
                                            <a v-for="subject in sellersubjects" href="#"
                                                class="cat-link">{{subject.SubjectName}}({{subject.SubLevelName}}-
                                                {{subject.LevelName}})</a>
                                        </div>
                                    </div>
                                    <div class="col-sm-12 col-md-6">
                                        <div class="member-tax small-text">
                                            <h4 class="h4">Seller : </h4>
                                            <div class="media" v-for="seller in sellers">
                                                <div class="media-left">
                                                    <a href="#" style="width: 64px; height: 64px;">
                                                        <img class="media-object" :src="seller.Photo" alt="...">
                                                    </a>
                                                </div>
                                                <div class="media-body">
                                                    <h4 class="media-heading">{{seller.Name}}</h4>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <!--/content-pad-->
                    </div>
                    <!--/col-md-8-->
                </div>
                <!--/row-->
            </div>
        </div>
        <div class="panel-body">

            <ul class="list-inline social-light">
                <li v-if="account.Mobile">
                    <a class="btn btn-default social-icon" :href="'tel:'+account.Mobile">
                        <i class="fa fa-mobile"></i>
                    </a>
                </li>
                <li v-if="account.Email">
                    <a class="btn btn-default social-icon" :href="'mailto:'+account.Email">
                        <i class="fa fa-at"></i>
                    </a>
                </li>
                <li v-if="account.Facebook">
                    <a class="btn btn-default social-icon" :href="account.Facebook" target="_blank">
                        <i class="fab fa-facebook"></i>
                    </a>
                </li>
                <li v-if="account.Youtube">
                    <a class="btn btn-default social-icon" :href="account.Youtube" target="_blank">
                        <i class="fab fa-youtube"></i>
                    </a>
                </li>
                <li v-if="account.Instagram">
                    <a class="btn btn-default social-icon" :href="account.Instagram" target="_blank">
                        <i class="fab fa-instagram"></i>
                    </a>
                </li>
                <li v-if="account.Pinterest">
                    <a class="btn btn-default social-icon" :href="account.Pinterest" target="_blank">
                        <i class="fab fa-pinterest"></i>
                    </a>
                </li>
                <li v-if="account.WhatsApp">
                    <a class="btn btn-default social-icon" :href="account.WhatsApp" target="_blank">
                        <i class="fab fa-whatsapp"></i>
                    </a>
                </li>
            </ul>
            <div class="pure-content">
                <div class="content-pad">
                    <h4 class="h4">About Me : </h4>
                    <p>{{account.Description}}</p>
                </div>
            </div>
        </div>
    </div>
`
    }
