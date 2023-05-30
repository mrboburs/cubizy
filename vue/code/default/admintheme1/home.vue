<script>
    export default {
        data: () => {
            return {
                loadingfeaturedaccounts: false,
                featuredaccountsmessage: "",
                featuredaccounts: [],
                loading_recent_sessions: false,
                loading_featured_accounts: false,
                types:[
                    'Academic Sellers',
                    'Institutions',
                    'Extra Curriculum',
                    'Educational Suppliers',
                ],
                selected_type : 'Academic Sellers',
                BackgroundImage : application.ThemeSettings.SearchBackgroundImage.value,
                RepeatBgi : application.ThemeSettings.RepeatSearchBackgroundImage.value
            }
        },
        components: {
            'search_row': () => import("/vue/search_row.js"),
            'featured': () => import("/vue/featured.js"),
            'featured_subjects': () => import("/vue/featured_subjects.js"),
            'EventCard': () => import("/vue/EventCard.js"),
            'CourseCard': () => import("/vue/CourseCard.js"),
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
            ...Vuex.mapState(['recent_posts', 'recent_events', 'recent_sessions', 'featured_accounts']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getDate', 'getTime']),
        },
        methods: {
            get_recent_sessions() {
                if (this.loading_recent_sessions) {
                    return
                }
                this.loading_recent_sessions = true
                this.$store.dispatch('call', {
                    api: "sessions",
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
                        this.$store.commit('set_recent_sessions', data.data)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading_sessions_posts = false
                })
            },
            get_featured_accounts() {
                if (this.loading_featured_accounts) {
                    return
                }
                this.loading_featured_accounts = true
                this.$store.dispatch('call', {
                    api: "sessions",
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
                        this.$store.commit('set_featured_accounts', data.data)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading_sessions_posts = false
                })
            },
        },
        mounted: function () {
            this.get_recent_sessions()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div>
        <featured class="flat-row no-padding"></featured>
        <div class="page-title parallax parallax1" :style="{ 'background-image' : 'url('+ BackgroundImage +')', 'background-repeat' : RepeatBgi?'repeat':'no-repeat', 'background-size':  RepeatBgi?'auto':'cover' }">
            <div class="container">
                <div class="flat-university">
                    <ul class="nav nav-pills large_searchbox" >
                        <li role="presentation" :class="{'active': _type == selected_type}" v-for="_type in types"><a href="#"  @click.prevent="selected_type = _type">{{_type}}</a></li>
                    </ul>
                    <search_row v-if="selected_type == 'Academic Sellers'" mode="horizontal" default_type="courses" type_locked="true" class="large_searchbox">
                    </search_row>
                    <ul v-else class="nav nav-pills large_searchbox" >
                        <li role="presentation">Comming soon..... </li>
                    </ul>
                </div>
                <!--/flat-university -->
            </div>
            <!--/container -->
        </div>

        <section class="flat-row">
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="flat-events">
                            <div class="grid-item color-full">
                                <div class="event-item">
                                    <div class="grid-item-content">
                                        <h1 class="title">Upcoming Events</h1>
                                        <router-link to="/events" class="flat-button">
                                            ALL EVENTS <i class="fa fa-angle-right"></i>
                                        </router-link>
                                    </div>
                                </div>
                            </div>
                            <EventCard v-for="event in recent_events.slice(0, 7)" :event="event"/>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <section class="flat-row">
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="flat-events">
                            <div class="grid-item color-full">
                                <div class="event-item">
                                    <div class="grid-item-content">
                                        <h1 class="title">Upcoming Courses</h1>
                                        <router-link to="/courses" class="flat-button">
                                            ALL COURSES <i class="fa fa-angle-right"></i>
                                        </router-link>
                                    </div>
                                </div>
                            </div>
                            <CourseCard v-for="session in recent_sessions.slice(0, 7)" :session="session" class="grid-item"/>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section class="flat-row flat-blog">
            <div class="container">
                <div class="row">
                    <div class="section-header col-md-12">
                        <div class="name-blog">
                            <h2 class="title">Blogs</h2>
                            <router-link to="/blogs" class="flat-button button-right">
                                VISIT BLOGS <i class="fa fa-angle-right"></i>
                            </router-link>
                        </div>
                    </div>

                    <div v-for="(blog, index) in recent_posts.slice(0, 4)" class="blog-item col-md-6">
                        <div class="post-item blog-post-item">
                            <div class="row">
                                <div class="col-md-6 col-sm-12">
                                    <div class="content-pad">
                                        <div class="blog-thumbnail">
                                            <div class="item-thumbnail-gallery">
                                                <div class="item-thumbnail">
                                                    <router-link :to="'/blog/'+blog.ID">
                                                        <img :src="blog.Image" alt="Title" style="max-height: 200px;">
                                                        <div class="thumbnail-hoverlay main-color-1-bg">
                                                        </div>
                                                        <div class="thumbnail-hoverlay-cross"></div>
                                                    </router-link>
                                                </div>
                                            </div>
                                        </div>
                                        <!--/blog-thumbnail-->
                                        <div class="thumbnail-overflow" style="left: 0;">
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
                                                <router-link :to="'/blog/'+blog.ID" class="main-color-1-hover">
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
            </div>
        </section>
    </div>
</template>