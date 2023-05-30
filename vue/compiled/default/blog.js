
    export default {
        props: {
            blog_id: {
                default: 0,
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                blog: {
                    Title: "",
                    Content: "",
                    Image: "",
                },
                previous: null,
                next: null,
                author: null,
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
            blog_id: function (newValue, oldValue) {
                if (newValue != oldValue) {
                    this.load()
                }
            },
            blog: function(newValue){
                scroll(0,0)
            },
        },
        computed: {
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getDate', 'getFBLink', 'getTwitterLink', 'getGooglePlusLink', 'getPinterestLink']),
        },
        methods: {
            load() {
                if (this.loading) { return }
                this.loading = true
                this.oldid = this.blog_id
                this.$store.dispatch('call', {
                    api: "blog",
                    data: {
                        blog_id: this.blog_id
                    }
                }).then((data) => {
                    this.content = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.blog = data.Result.blog
                        if (data.Result.author) this.author = data.Result.author
                        else this.author = null
                        if (data.Result.previous) this.previous = data.Result.previous
                        else this.previous = null
                        if (data.Result.next) this.next = data.Result.next
                        else this.next = null
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                    if (this.oldid != this.blog_id) {
                        this.load()
                    }
                })
            },
        },
        mounted: function () {
            this.load()
        },
        template: `
    <divloading :fullpage="false" :loading="loading" style="height: 100%;">
        <div class="page-title full-color">
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="page-title-heading">
                            <h2 class="title">Blog: {{blog.Title}}</h2>
                        </div>
                        <div class="breadcrumbs">
                            <ul>
                                <li class="home">
                                    <router-link to="/">
                                        Home
                                    </router-link>
                                </li>
                                <li class="home">
                                    <span class="mh-1">/</span>
                                    <router-link to="/blogs">
                                        Blogs
                                    </router-link>
                                </li>
                            </ul>
                        </div>
                    </div><!-- /.col-md-12 -->
                </div><!-- /.row -->
            </div><!-- /.container -->
        </div>
        <section class="flat-row padding-v1">
            <div class="container">
                <div class="row">
                    <div id="content" class="col-md-9">
                        <article class="post">
                            <div class="entry-wrapper">
                                <div class="entry-box">
                                    <router-link :to="'/blog/'+blog.ID">
                                        <img :src="blog.Image" alt="images">
                                    </router-link>
                                </div>
                                <div class="post-content">
                                    <div v-html="blog.Content"></div>
                                </div>
                                <div class="content-pad">
                                    <div class="item-content">
                                        <div class="item-meta blog-item-meta">
                                            <span>{{getFullDate(blog.UpdatedAt)}}<span class="sep">|</span> </span>
                                        </div>
                                    </div>
                                </div>
            
                                <div class="list-inline item-content">
                                    <ul class="list-inline social-light">
                                        <li v-if="getFBLink('/blog/'+blog.ID)">
                                            <a :href="getFBLink(blog.ID)" target="_blank"
                                                class="btn btn-default social-icon">
                                                <i class="fab fa-facebook"></i>
                                            </a>
                                        </li>
                                        <li v-if="getTwitterLink('/blog/'+blog.ID)">
                                            <a :href="getTwitterLink(blog.ID)" target="_blank"
                                                class="btn btn-default social-icon">
                                                <i class="fab fa-twitter"></i>
                                            </a>
                                        </li>
                                        <li v-if="getGooglePlusLink('/blog/'+blog.ID)">
                                            <a :href="getGooglePlusLink(blog.ID)" target="_blank"
                                                class="btn btn-default social-icon">
                                                <i class="fab fa-google-plus"></i>
                                            </a>
                                        </li>
                                        <li v-if="getPinterestLink('/blog/'+blog.ID)">
                                            <a :href="getPinterestLink(blog.ID)" target="_blank"
                                                class="btn btn-default social-icon">
                                                <i class="fab fa-pinterest"></i>
                                            </a>
                                        </li>
                                    </ul>
                                </div>
            
                                <div class="about-author" v-if="author">
                                    <div class="author-avatar">
                                        <img :src="author.Photo" alt="image">
                                    </div>
                                    <div class="author-info">
                                        <h4>{{author.Name}}</h4>
                                    </div>
                                    <div class="clearfix"></div>
                                </div>
            
                                <div class="simple-navigation">
                                    <div class="row">
                                        <div v-if="previous" class="simple-navigation-item col-md-6 col-sm-6 col-xs-6 main-color-1-bg-hover ">
                                            <router-link  :to="'/blog/'+previous.ID"
                                                class="maincolor2hover">
                                                <i class="fa fa-angle-left pull-left"></i>
                                                <div class="simple-navigation-item-content">
                                                    <span>Previous</span>
                                                    <h4>{{previous.Title}}</h4>
                                                </div>
                                            </router-link>
                                        </div>
                                        <div v-if="next" class="simple-navigation-item col-md-6 col-sm-6 col-xs-6 main-color-1-bg-hover pull-right">
                                            <router-link :to="'/blog/'+next.ID"
                                                class="maincolor2hover pull-right">
                                                <i class="fa fa-angle-right pull-right"></i>
                                                <div class="simple-navigation-item-content">
                                                    <span>Next</span>
                                                    <h4>{{next.Title}}</h4>
                                                </div>
                                            </router-link>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </article>
                    </div>
                    <div class="col-md-3">
                        <div class="sidebar">
                            <searchwidget default_type="blogs"></searchwidget>
                            <letestevents></letestevents>
                            <letestposts></letestposts>
                        </div><!-- /col-md-9 -->
                    </div>
                </div>
            </div>
        </section>
    </divloading>
`
    }
