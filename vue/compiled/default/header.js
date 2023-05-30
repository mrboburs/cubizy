
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                Title: window.application.Account.Title,
                LogoHeight: window.application.ThemeSettings.LogoHeight.value,
                WideLogo: window.application.Account.WideLogo,
                Logo: window.application.Account.Logo,
                AccountType: window.application.Account.AccountType,
                Email: window.application.Account.Email,
                Mobile: window.application.Account.Mobile,
                Logo: window.application.Account.Logo,
                showMenu: false,
            }
        },
        computed: {
            ...Vuex.mapState(['pages','levels'])
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
            showMenu: function (newValue) {
                $('#mainnav-mobi').slideToggle(300);
                $(this).toggleClass('active');
            },
            $route(to, from) {
                $('#mainnav-mobi').slideToggle(300);
                $(this).toggleClass('active');
            },
        },
        methods: {
            load() {
                this.loading = true
                this.$store.dispatch('call', {
                    api: "pages",
                    data: {
                        sort: "weightage",
                        sortdesc: true,
                        limit: 20,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if (Array.isArray(data.data)) {
                            this.$store.commit('set_pages', data.data)
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
            pages_from(from) {
                if (this.pages.length > from) {
                    return this.pages.slice(from, from + 5);
                } else {
                    return []
                }
            },
            getUrlName(level){
                var name = level.Name
                return encodeURI(name.toLowerCase().replaceAll(' ', '_'))
            }
        },
        mounted: function () {
            this.load()
        },
        template: `
    <header id="header" class="header sticky">
        <div class="menu-hover">
            <div class="btn-menu" :class="{'active': showMenu}" @click.prevent="showMenu = !showMenu">
                <span></span>
            </div><!-- //mobile menu button -->
        </div>
        <div class="header-wrap">
            <div class="container">
                <div class="header-wrap clearfix">
                    <div id="logo" class="logo">
                        <router-link to="/" style=" display: grid; grid-auto-flow: column; align-items: center; gap: 1em; ">
                            <img :src="WideLogo" alt="image" :style="{height : LogoHeight+'px' }">
                            <h2 v-if="Logo == WideLogo" class="h2">{{Title}}</h2>
                        </router-link>
                    </div><!-- /.logo -->


                    <div class="nav-wrap">

                        <nav id="mainnav" class="mainnav">
                            <ul class="menu">
                                <li class="home">
                                    <router-link to="/">
                                        Home
                                    </router-link>
                                </li>
                                <li v-if="AccountType == 'admin'">
                                    <router-link :to="'/events'">
                                        Events
                                    </router-link>
                                </li>
                                <li v-if="AccountType == 'admin'">
                                    <router-link :to="'/sellers'">
                                        Academic
                                    </router-link>
                                </li>
                                <li>
                                    <router-link :to="'/courses'">
                                        Courses
                                    </router-link>
                                    <ul class="submenu">
                                        <li v-for="level in levels">
                                            <router-link :to="'/courses/'+getUrlName(level)">
                                                {{level.Name}}
                                            </router-link>
                                            <ul class="submenu">
                                                <li v-for="sublevel in level.SubLevels">
                                                    <router-link :to="'/courses/' + getUrlName(level)+'/' + getUrlName(sublevel)">
                                                        {{sublevel.Name}}
                                                    </router-link>
                                                    <ul class="submenu">
                                                        <li v-for="subject in sublevel.Subjects">
                                                            <router-link :to="'/courses/' + getUrlName(level)+'/' + getUrlName(sublevel) +'/' + getUrlName(subject)">
                                                                {{subject.Name}}
                                                            </router-link>
                                                        </li>
                                                    </ul>
                                                </li>
                                            </ul>
                                        </li>
                                    </ul>
                                </li>
                                <li v-if="AccountType == 'admin'">
                                    <router-link :to="'/blogs'">
                                        Blogs
                                    </router-link>
                                </li>
                                <li v-if="pages.length">
                                    <router-link :to="'/pages/'+pages[0].ID">
                                        {{pages[0].Title}}
                                    </router-link>
                                </li>
                                <li v-if="pages.length > 1">
                                    <router-link :to="'/pages/'+pages[1].ID">
                                        {{pages[1].Title}}
                                    </router-link>
                                </li>
                                <li v-if="pages.length == 3">
                                    <router-link :to="'/pages/'+pages[2].ID">
                                        {{pages[2].Title}}
                                    </router-link>
                                </li>
                                <li class="has-mega-menu" v-if="pages.length > 3">
                                    <a href="#" @click.prevent="">
                                        More...
                                    </a>
                                    <ul class="submenu submenu-right mega-menu clearfix">
                                        <li class="menu-column">
                                            <ul>
                                                <li v-for="page in pages_from(2)">
                                                    <router-link :to="'/pages/'+page.ID"
                                                        class="menu-link  sub-menu-link">
                                                        {{page.Title}}
                                                    </router-link>
                                                </li>
                                            </ul>
                                        </li>
                                        <li class="menu-column" v-if="pages.length > 7">
                                            <ul>
                                                <li v-for="page in pages_from(7)">
                                                    <router-link :to="'/pages/'+page.ID"
                                                        class="menu-link  sub-menu-link">
                                                        {{page.Title}}
                                                    </router-link>
                                                </li>
                                            </ul>
                                        </li>
                                        <li class="menu-column" v-if="pages.length > 12">
                                            <ul>
                                                <li v-for="page in pages_from(12)">
                                                    <router-link :to="'/pages/'+page.ID"
                                                        class="menu-link  sub-menu-link">
                                                        {{page.Title}}
                                                    </router-link>
                                                </li>
                                            </ul>
                                        </li>
                                        <li class="menu-column" v-if="pages.length > 17">
                                            <ul>
                                                <li v-for="page in pages_from(17)">
                                                    <router-link :to="'/pages/'+page.ID"
                                                        class="menu-link  sub-menu-link">
                                                        {{page.Title}}
                                                    </router-link>
                                                </li>
                                            </ul>
                                        </li>
                                    </ul><!-- /.submenu -->
                                </li>
                            </ul><!-- /.menu -->
                        </nav><!-- /.mainnav -->
                    </div><!-- /.nav-wrap -->
                </div><!-- /.header-wrap -->
            </div><!-- /.container-->
        </div><!-- /.header-wrap-->
    </header><!-- /.header -->
`
    }
