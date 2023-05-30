<script>
    export default {
        data: () => {
            return {
                loading: false,
                loading_subjects: false,
                loading_recent_posts : false,
                loading_recent_events : false,
                submitted: false,
                error: false,
                message: "",
                AppName : window.application.AppName,
                AccountType : window.application.Account.AccountType,
                Title : window.application.Account.Title,
                Logo : window.application.Account.Logo,
                Description : window.application.Account.Description,

                Youtube : window.application.Account.Youtube,
                Facebook : window.application.Account.Facebook,
                Instagram : window.application.Account.Instagram,
                Pinterest : window.application.Account.Pinterest,
                WhatsApp : window.application.Account.WhatsApp,

                Address : window.application.Address,
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
        components: {
            'address_view': () => import("/vue/address.js"),
        },
        computed: {
            ...Vuex.mapState(['pages', 'sublevels','subjects','recent_posts', 'recent_events', 'recent_sessions']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getDate']),
            showmessage: {
                // getter
                get: function () {
                    if (this.message) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.message = ""
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'alert-danger'
                } else {
                    return 'alert-success'
                }
            },
            year : function (){
                return new Date().getFullYear()
            }
        },
        methods: {
            can_load() {
                if(this.loading){
                    return false
                }
                return true
            },
            
            get_subjects() {
                if (this.loading_subjects) {
                    return
                }
                this.loading_subjects = true
                this.$store.dispatch('call', {
                    api: "allsubjects",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_subjects', data.data)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading_subjects = false
                })
            },
            
            get_recent_posts() {
                if (this.loading_recent_posts) {
                    return
                }
                this.loading_recent_posts = true
                this.$store.dispatch('call', {
                    api: "blogs",
                    data: {
                        sort: "updated_at",
                        sortdesc: true,
                        limit: 7,
                        page: 0,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_recent_posts', data.data)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading_recent_posts = false
                })
            },
            
            get_recent_events() {
                if (this.loading_recent_events) {
                    return
                }
                this.loading_recent_events = true
                this.$store.dispatch('call', {
                    api: "events",
                    data: {
                        sort: "updated_at",
                        sortdesc: true,
                        limit: 7,
                        page: 0,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_recent_events', data.data)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading_events_posts = false
                })
            },
        },
        mounted: function () {
            this.get_subjects()
            if(this.AccountType == 'admin'){
                this.get_recent_posts()
                this.get_recent_events()
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <footer class="footer full-color">
        <section id="bottom">
            <div class="section-inner">
                <div class="container">
                    <div class="row normal-sidebar">
                        <div class="widget widget-text">
                            <div class="row">
                                <div class="col-md-3">
                                    <div class=" widget-inner">
                                        <div class="textwidget">
                                            <div id="un-icon-box-1" class="media un-icon-box" data-delay="0">
                                                <div class="pull-left">
                                                    <div class="un-icon">
                                                        <i class="fa fa-cube"></i>
                                                    </div>
                                                </div>
                                                <div class="media-body">
                                                    <h4 class="media-heading">Research</h4>
                                                    <p>Atero voluptatum ptatum eos et accusamus et iusto</p>    
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                                <div class=" col-md-3">
                                    <div class=" widget-inner">
                                        <div class="textwidget">
                                            <div id="un-icon-box-2" class="media un-icon-box  " data-delay="0">
                                                <div class="pull-left">
                                                    <div class="un-icon">
                                                        <i class="fa fa-leaf"></i>
                                                    </div>
                                                </div>
                                                <div class="media-body ">
                                                    <h4 class="media-heading">Engage</h4>
                                                    <p>At vero eos et accusamus et iusto odio dignissimos ducimus</p>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                                <div class=" col-md-3">
                                    <div class=" widget-inner">
                                        <div class="textwidget">
                                            <div id="un-icon-box-3" class="media un-icon-box  " data-delay="0">
                                                <div class="pull-left">
                                                    <div class="un-icon">
                                                        <i class="fa fa-thumbs-up"></i>
                                                    </div>
                                                </div>
                                                <div class="media-body ">
                                                    <h4 class="media-heading">Commitment</h4>
                                                    <p>Ptatum eos et accusamus et iusto odio dignissimos ducimus</p>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                                <div class=" col-md-3">
                                    <div class=" widget-inner">
                                        <div class="textwidget">
                                            <div id="un-icon-box-4" class="media un-icon-box  " data-delay="0">
                                                <div class="pull-left">
                                                    <div class="un-icon">
                                                        <i class="fa fa-rocket"></i>
                                                    </div>
                                                </div>
                                                <div class="media-body ">
                                                    <h4 class="media-heading">Innovation</h4>
                                                    <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem</p>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class=" widget divider-widget">
                            <div class=" widget-inner">
                                <div class="un-heading un-separator">
                                    <div class="un-heading-wrap">
                                        <span class="un-heading-line un-heading-before">
                                            <span></span>
                                        </span>
                                        <button class="flat-button style1">ENROLL TODAY <i class="fa fa-angle-right"></i>
                                        </button>
                                        <span class="un-heading-line un-heading-after">
                                            <span></span>
                                        </span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
                            </div>
                        </div>

                        <div v-for="sublevel in sublevels.slice(0, 3)" class="col-md-3  widget widget-nav-menu">
                            <div class=" widget-inner">
                                <h2 class="widget-title maincolor1">{{sublevel.Name}} <span v-if="sublevel.SessionCount">({{sublevel.SessionCount}})</span></h2>
                                <div class="menu-law-business-container">
                                    <ul id="menu-law-business" class="menu">
                                        <li v-for="subject in sublevel.Subjects" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-1280">
                                            <router-link :to="'/courses/?level='+subject.LevelName+'&sub_level='+subject.SubLevelName+'&subject='+subject.Name">
                                                {{subject.Name}} <span v-if="subject.SessionCount">({{subject.SessionCount}})</span>
                                            </router-link>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </div>

                        <div class=" col-md-3  widget divider-widget">
                            <div class=" widget-inner">
                                <div class="un-heading un-separator">
                                    <div class="un-heading-wrap">
                                        <span class="un-heading-line un-heading-before"><span>
                                            </span></span>
                                        <span class="un-heading-line un-heading-after">
                                            <span></span>
                                        </span>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
                            </div>
                        </div>

                        <div v-if="AccountType != 'admin'" class=" col-md-3  widget widget-recent-entries">
                            <div class=" widget-inner">
                                <img :src="Logo" alt="image" style="width: 100%; height: auto;">
                            </div>
                        </div>

                        <div class=" col-md-3  widget widget-text">
                            <div class=" widget-inner">
                                <h2 class="widget-title maincolor1 text-uppercase">{{Title}}</h2>
                                <div class="textwidget">{{ Description }}</div>
                            </div>
                        </div>

                        <div v-if="AccountType == 'admin'" class=" col-md-3  widget widget-recent-entries">
                            <div class=" widget-inner">
                                <h2 class="widget-title maincolor1">RECENT POSTS</h2>
                                <ul>
                                    <li v-for="blog in recent_posts.slice(0, 5)">
                                        <router-link :to="'/blog/'+blog.ID">
                                            {{blog.Title}}
                                        </router-link>
                                    </li>
                                </ul>
                            </div>
                        </div>

                        <div v-if="AccountType == 'admin'" class=" col-md-3  widget widget-recent-entries">
                            <div class=" widget-inner">
                                <h2 class="widget-title maincolor1">RECENT EVENTS</h2>
                                <ul>
                                    <li v-for="event in recent_events.slice(0, 5)">
                                        <router-link :to="'/event/'+event.ID">
                                            {{event.Title}}
                                        </router-link>
                                    </li>
                                </ul>
                            </div>
                        </div>

                        <div class=" col-md-3  widget widget-nav-menu">
                            <div class=" widget-inner">
                                <h2 class="widget-title maincolor1">LINKS</h2>
                                <div class="menu-others-container">
                                    <ul id="menu-others" class="menu">
                                        <li v-for="page in pages.slice(0, 5)" class="menu-item menu-item-type-custom menu-item-object-custom">
                                            <router-link :to="'/pages/'+page.ID">
                                                {{page.Title}}
                                            </router-link>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </div>

                        <address_view :address="Address"></address_view>
                    </div>
                </div>
            </div>
        </section>

        <div id="bottom-nav">
            <div class="container">
                <div class="link-center">
                    <div class="line-under"></div>
                    <a class="flat-button go-top-v1 style1" href="#top">TOP</a>
                </div>
                <div class="row footer-content">
                    <div class="copyright col-md-6">
                        © {{year}} {{AppName}}.
                    </div>
                    <nav class="col-md-6 footer-social">
                        <ul class="social-list">
                            <li v-if="Facebook">
                                <a :href="Facebook" target="_blank" class="btn btn-default social-icon">
                                    <i class="fab fa-facebook"></i>
                                </a>
                            </li>
                            <li v-if="Youtube">
                                <a :href="Youtube" target="_blank" class="btn btn-default social-icon">
                                    <i class="fab fa-youtube"></i>
                                </a>
                            </li>
                            <li v-if="Instagram">
                                <a :href="Instagram" target="_blank" class="btn btn-default social-icon">
                                    <i class="fab fa-instagram"></i>
                                </a>
                            </li>
                            <li v-if="Pinterest">
                                <a :href="Pinterest" target="_blank" class="btn btn-default social-icon">
                                    <i class="fab fa-pinterest"></i>
                                </a>
                            </li>
                            <li v-if="WhatsApp">
                                <a :href="WhatsApp" target="_blank" class="btn btn-default social-icon">
                                    <i class="fab fa-whatsapp"></i>
                                </a>
                            </li>
                        </ul>
                    </nav>
                </div><!--/row-->
            </div><!--/container-->
        </div>
    </footer>
</template>