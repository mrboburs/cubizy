<script>
    export default {
        data: () => {
            return {
                loading_recent_sessions: false,
            }
        },
        components: {
            'featured_subjects': () => import("/vue/featured_subjects.js"),
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
            ...Vuex.mapState(['recent_sessions']),
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
        },
        mounted: function () {
            this.get_recent_sessions()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div>
        <featured_subjects class="flat-row no-padding"></featured_subjects>
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
    </div>
</template>