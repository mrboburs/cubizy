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

                searched: null,
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
            page: function (newValue, oldValue) {
                if (!this.records[newValue] || !this.records[newValue].length) {
                    this.load()
                }
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account', 'search', 'country', 'district', 'locality', 'sublocality', 'code', 'level', 'sublevel', 'subject']),
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
                    api: "accounts",
                    data: {
                        account_type: 'seller',
                        sort: this.sort_by,
                        sortdesc: this.sortdesc,
                        search: this.search,
                        country: this.country,
                        district: this.district,
                        locality: this.locality,
                        sublocality: this.sublocality,
                        code: this.code,
                        level: this.level,
                        sublevel: this.sublevel,
                        subject: this.subject,
                        limit: this.limit,
                        page: this.page - 1,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.records[this.page] = data.data
                        this.recordsTotal = data.recordsTotal
                        this.recordsFiltered = data.recordsFiltered
                        if (data.Request.country || data.Request.district || data.Request.locality || data.Request.sublocality || data.Request.search || data.Request.level || data.Request.sublevel || data.Request.subject) {
                            this.searched = {}
                            this.searched.country = data.Request.country
                            this.searched.district = data.Request.district
                            this.searched.locality = data.Request.locality
                            this.searched.sublocality = data.Request.sublocality
                            this.searched.level = data.Request.level
                            this.searched.sublevel = data.Request.sublevel
                            this.searched.subject = data.Request.subject
                            this.searched.search = data.Request.search
                        } else {
                            this.searched = null
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
                            <h2 class="title">sellers</h2>
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
                                    <router-link to="/sellers">
                                        sellers
                                    </router-link>
                                </li>
                            </ul>
                        </div>
                    </div><!-- /.col-md-12 -->
                </div><!-- /.row -->
            </div><!-- /.container -->
        </div>
        <section class="flat-row padding-small-v1">
            <div class="container">
                <div class="row">
                    <div class="col-md-9">
                        <div class="project-listing">
                            <div v-if="searched" class="filter-cat">
                                <label v-if="!recordsFiltered" class="no_records">
                                    No sellers found
                                </label>
                                <label v-if="recordsFiltered == 1" class="no_records">
                                    {{recordsFiltered}} sellers found
                                </label>
                                <label v-if="recordsFiltered > 1" class="no_records">
                                    {{recordsFiltered}} sellers found
                                </label>
                                for
                                <a v-if="searched.level" href="#" class="flat-button" @click.prevent="">
                                    {{searched.level}}
                                </a>
                                <a v-if="searched.sublevel" href="#" class="flat-button" @click.prevent="">
                                    {{searched.sublevel}}
                                </a>
                                <a v-if="searched.subject" href="#" class="flat-button" @click.prevent="">
                                    {{searched.subject}}
                                </a>
                                <a v-if="searched.sublocality" href="#" class="flat-button" @click.prevent="">
                                    {{searched.sublocality}}
                                </a>
                                <a v-else-if="searched.locality" href="#" class="flat-button" @click.prevent="">
                                    {{searched.locality}}
                                </a>
                                <a v-else-if="searched.district" href="#" class="flat-button" @click.prevent="">
                                    {{searched.district}}
                                </a>
                                <a v-else-if="searched.country" href="#" class="flat-button" @click.prevent="">
                                    {{searched.country}}
                                </a>
                                <a v-if="searched.search" href="#" class="flat-button" @click.prevent="">
                                    {{searched.search}}
                                </a>
                                <a v-if="searched.level || searched.sublevel || searched.subject || searched.sublocality || searched.locality || searched.district || searched.country  || searched.search"
                                    href="#" class="flat-button" @click.prevent="clear_search">
                                    Clear Search
                                </a>
                            </div>
                        </div>
                        <div v-if="recordsFiltered" class="item-row">
                            <div v-for="(account, index) in records[page]" class="item">
                                <div class="thumb-item">
                                    <div class="item-thumbnail">
                                        <router-link :to="'/account/'+account.ID">
                                            <img :src="account.Logo" alt="image">
                                        </router-link>
                                    </div><!-- /item-thumbnail -->

                                    <div class="item-content">
                                        <h3 class="item-title h3">
                                            <router-link :to="'/account/'+account.ID">
                                                {{account.Title}}
                                            </router-link>
                                        </h3>
                                        <h4 class="small-text h4">{{account.Keywords}}</h4>
                                        <p>{{account.Description}}</p>
                                        <star-rating :rating="account.Rating" :showRating="false" :starSize="20" :readOnly="true"></star-rating>
                                        <ul class="list-inline social-light">
                                            <li v-if="account.Facebook"><a class="btn btn-default social-icon"
                                                    target="_blank" :href="account.Facebook">
                                                    <i class="fab fa-facebook"></i></a>
                                            </li>
                                            <li v-if="account.Twitter"><a class="btn btn-default social-icon"
                                                    target="_blank" :href="account.Twitter">
                                                    <i class="fab fa-twitter"></i></a>
                                            </li>
                                            <li v-if="account.Linkedin"><a class="btn btn-default social-icon"
                                                    :href="account.Linkedin"><i class="fab fa-linkedin"></i></a>
                                            </li>
                                        </ul>
                                    </div><!-- /item-content -->
                                </div><!-- /thumb-item -->
                            </div>
                        </div><!-- /item -->
                        <v-pagination v-model="page" :totalrows="recordsFiltered" :perpage="parseInt(limit)">
                        </v-pagination>
                    </div><!-- /col-md-9 -->

                    <div class="col-md-3">
                        <div class="sidebar">
                            <searchwidget default_type="sellers" @search="dosearch"></searchwidget>
                            <letestevents></letestevents>
                            <letestposts></letestposts>
                        </div><!-- /col-md-9 -->
                    </div><!-- /col-md-3 -->
                </div>
            </div>
        </section>
    </divloading>
</template>