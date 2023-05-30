<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                records: [],
                recordsTotal: 0,
                recordsFiltered: 0,

                page: 1,
                limit: 10,
                sortdesc: true,
                sort_by: "updated_at",

                searched: "",
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
            page: function(newValue, oldValue) {
                if(!this.records[newValue] || !this.records[newValue].length ){
                    this.load()
                }
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account', 'search']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getDate']),
            showpopup: {
                // getter
                get: function () {
                    if (this.popupitem) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.popupitem = false
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'danger'
                } else {
                    return 'success'
                }
            }
        },
        methods: {
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "blogs",
                    data: {
                        sort: this.sort_by,
                        sortdesc: this.sortdesc,
                        search: this.search,
                        limit: this.limit,
                        page: this.page-1,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.records[this.page] = data.data
                        this.recordsTotal = data.recordsTotal
                        this.recordsFiltered = data.recordsFiltered
                        this.searched = data.Request.search
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
            dosearch() {
                this.records = []
                this.page = 1
                this.recordsFiltered = 0
                this.load()
            },
            clear_search() {
                this.$store.commit('clear_search')
                this.dosearch()
            }
        },
        mounted: function () {
            this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>

    <divloading :fullpage="false" :loading="loading" style="height: 100%;">
        <div class="page-title full-color">
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="page-title-heading">
                            <h2 class="title">Blogs</h2>
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
                        <div class="project-listing">
                            <div class="filter-cat">
                                <a v-if="searched" href="#" class="flat-button" @click.prevent="">Searched for :
                                    {{searched}}</a>
                                <a v-if="searched" href="#" class="flat-button" @click.prevent="clear_search">Clear
                                    Search</a>
                            </div>
                        </div>
                        <div class="blog-listing">
                            <div v-if="!recordsFiltered" class="no_records">
                                No blogs found!
                            </div>
                            <div v-else v-for="(blog, index) in records[page]" class="blog-item">
                                <div class="blog-item">
                                    <div class="post-item blog-post-item">
                                        <div class="row">
                                            <div class="col-md-6 col-sm-12">
                                                <div class="content-pad">
                                                    <div class="blog-thumbnail">
                                                        <div class="item-thumbnail-gallery">
                                                            <div class="item-thumbnail">
                                                                <router-link :to="'/blog/'+blog.ID">
                                                                    <img :src="blog.Image" alt="Title">
                                                                    <div class="thumbnail-hoverlay main-color-1-bg">
                                                                    </div>
                                                                    <div class="thumbnail-hoverlay-cross"></div>
                                                                </router-link>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <!--/blog-thumbnail-->
                                                    <div class="thumbnail-overflow">
                                                        <div class="date-block main-color-2-bg dark-div">
                                                            <div class="month">{{getMonth(blog.UpdatedAt)}}</div>
                                                            <div class="day">{{getDate(blog.UpdatedAt)}}</div>
                                                        </div>
                                                    </div>
                                                </div>
                                                <!--/blog-thumbnail-->
                                            </div>

                                            <div class="col-md-6 col-sm-12">
                                                <div class="content-pad">
                                                    <div class="item-content">
                                                        <h3 class="title">
                                                            <router-link :to="'/blog/'+blog.ID"
                                                                class="main-color-1-hover">
                                                                {{blog.Title}}
                                                            </router-link>
                                                        </h3>
                                                        <div class="item-meta blog-item-meta">
                                                            <span>By<span class="sep">|</span> </span>
                                                            <span>{{blog.UpdatedByName}}</span>
                                                            <br />
                                                            <span>On<span class="sep">|</span> </span>
                                                            <span>{{getFullDate(blog.UpdatedAt)}}</span>
                                                        </div>
                                                        <router-link :to="'/blog/'+blog.ID" class="button">
                                                            DETAILS
                                                            <i class="fa fa-angle-right"></i>
                                                        </router-link>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <!--/post-item-->
                                    </div>
                                </div><!-- /blog-item -->
                            </div>
                            <v-pagination v-model="page" :totalrows="recordsFiltered"
                                :perpage="parseInt(limit)">
                            </v-pagination>
                        </div><!-- /blog-listing -->
                    </div>
                    <div class="col-md-3">
                        <div class="sidebar">
                            <searchwidget default_type="blogs" @search="dosearch"></searchwidget>
                            <letestevents></letestevents>
                            <letestposts></letestposts>
                        </div><!-- /col-md-9 -->
                    </div>
                </div>
            </div>
        </section>
    </divloading>
</template>