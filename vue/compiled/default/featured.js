
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                accounts: [],
                sort_by: "updated_at",
                sortdesc: true,
                limit: 5,
                counter: 0,
                timer: false,
                showBanner: true,
                showInfo: true,
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
            ...Vuex.mapGetters(['getAccountUrl', 'getFullDate', 'getMonth', 'getDuration', 'getDate', 'getFBLink', 'getTwitterLink', 'getGooglePlusLink', 'getPinterestLink']),
        },
        methods: {
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "accounts",
                    data: {
                        sort: this.sort_by,
                        sortdesc: this.sortdesc,
                        limit: this.limit,
                        page: 0,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.accounts = data.data
                        if (this.accounts.length) {
                            setTimeout(() => {
                                this.glide = new Glide('.glide', {
                                    autoplay: 5000,
                                    hoverpause: true,
                                    perView: 1,
                                    type: 'carousel' ,
                                }).mount()
                                this.glide.on('run.before', (data) => {
                                    if (this.$refs['featured_intro' + this.glide.index][0]) {
                                        this.$refs['featured_intro' + this.glide.index][0].classList.toggle('animate__animated')
                                        this.$refs['featured_intro' + this.glide.index][0].classList.toggle('animate__bounceInUp')
                                    }
                                })
                                this.glide.on('move.after', (data) => {
                                    if(this.$refs['featured_intro' + this.glide.index][0]){
                                        this.$refs['featured_intro' + this.glide.index][0].classList.toggle('animate__animated')
                                        this.$refs['featured_intro' + this.glide.index][0].classList.toggle('animate__bounceInUp')
                                    }
                                })
                            }, 300);
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
            this.load()
        },
        template: `
    <div style="overflow: hidden;">
        <div class="glide">
            <div class="glide__track" data-glide-el="track">
                <ul class="glide__slides">
                    <li v-for="(account,index) in accounts" class="glide__slide">
                        <div :src="account.Banner" class="featured_banner image_holder"
                            :style="{ 'background-image' : 'url('+ account.Banner +')' }">
                            <div :ref="'featured_intro'+ index" class="featured_intro">
                                <div class="title main-color-1 font-2 h2">{{account.Title}}</div>
                                <div class="content"> {{account.Description}} </div>
                                <div class="content"> {{account.Keywords}} </div>
                                <a class="btn btn-primary pull-right" target="_blank" :href="getAccountUrl(account)">
                                    View More...
                                </a>
                                <span class="clearfix"></span>
                            </div>
                        </div>
                    </li>
                </ul>
            </div>
            <div class="glide__arrows" data-glide-el="controls">
                <button class="glide__arrow glide__arrow--left" data-glide-dir="<"><i
                        class="fas fa-chevron-left"></i></button>
                <button class="glide__arrow glide__arrow--right" data-glide-dir=">"><i
                        class="fas fa-chevron-right"></i></button>
            </div>
        </div>
    </div>
`
    }
