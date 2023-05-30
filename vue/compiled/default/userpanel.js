
    import menus from "/menu.js"
    export default {
        data: () => {
            return {
                FullHD: 1408,
                Widescreen: 1216,
                Desktop: 1024,
                Tablet: 769,
                deviceWidth: 1024,
                DeviceType: "",
                open: true,
                position: "static",
                overlay: true,
                reduce: false,
                fullheight: true,
                fullwidth: false,
                right: false,
                menus
            }
        },
        computed: {
            ...Vuex.mapState(['user']),
        },
        watch: {
            user: function (newValue, oldValue) {
                if (!newValue) {
                    this.$router.push('login')
                }
            },
            $route(to, from) {
                this.show = false;
            }
        },
        mounted: function () {
            this.ping()
            this.setDeviceType()
        },
        created() {
            window.addEventListener("resize", this.onScreenSizeChange);
        },
        destroyed() {
            window.removeEventListener("resize", this.onScreenSizeChange);
        },
        methods: {
            onScreenSizeChange(e) {
                this.deviceWidth = window.innerWidth
                this.setDeviceType()
            },
            setDeviceType() {
                if (window.innerWidth > this.Desktop) {
                    this.DeviceType = "Desktop"
                } else {
                    this.DeviceType = "Mobile"
                }

                if (window.innerWidth > this.Desktop) {
                    this.reduce = false
                    this.open = true
                    this.overlay = false
                    this.position = "static"
                } else if (window.innerWidth > this.Tablet) {
                    this.reduce = true
                    this.open = true
                    this.overlay = false
                    this.position = "static"
                } else if (window.innerWidth < this.Tablet) {
                    this.reduce = false
                    this.open = false
                    this.overlay = true
                    this.position = "fixed"
                }
            },
            ping() {
                if (!this.user) { return }
                this.$store.dispatch('call', { api: "test", data: {} })
            },
            logout() {
                if (!this.user) { return }
                this.$store.dispatch('call', { api: "logout", data: {} })
            },
            toggle_menu() {
                if (window.innerWidth > this.Desktop) {
                    this.reduce = !this.reduce
                } else if (window.innerWidth > this.Tablet) {
                    this.reduce = !this.reduce
                } else if (window.innerWidth < this.Tablet) {
                    this.open = !this.open
                }
            },
            isActive(path) {
                if (!path) {
                    return false
                }
                if (path == this.$route.path) {
                    return true
                }
                return false
            }
        },
        template: `
    <div class="mypanel columns m-0">
        <div class="column is-narrow panel-column nav">
            <b-sidebar v-if="reduce && (this.Tablet > deviceWidth)" type="is-light" :overlay="overlay" v-model="open" :position="position" :fullheight="true"
                style="height:100%" :reduce="reduce">
                <div class="p-1">
                    <!-- reduce != animate__animated animate__bounce -->
                    <b-menu :activable="true">
                        <b-menu-list v-for="menu in menus" :key="menu.label" :label="menu.label">
                            <b-menu-item v-for="item in menu.items" :icon="item.icon" :label="reduce? '':item.label"
                                :to="item.to" :tag="item.to?'router-link':'a'" :key="item.label"
                                :active.sync="item.to == $route.path">
                                <b-menu-item v-if="item.items" v-for="subitem in item.items" :icon="subitem.icon"
                                    :label="subitem.label" :to="subitem.to" :active.sync="subitem.to == $route.path"
                                    :tag="subitem.to?'router-link':'a'" :key="subitem.label">
                                    <!-- reduce? '': -->
                                </b-menu-item>
                            </b-menu-item>
                        </b-menu-list>
                    </b-menu>
                </div>
            </b-sidebar>
            <div v-else class="box p-2 h-100" :class="{'reduced':reduce}">
            <b-menu  :activable="true" >
                <b-menu-list v-for="menu in menus" :label="menu.label" :key="menu.label">
                    <b-menu-item v-for="item in menu.items" :icon="item.icon" :label="reduce? '':item.label"
                        :to="item.to" :tag="item.to?'router-link':'a'" :key="item.label"
                        :active.sync="item.to == $route.path">
                        <b-menu-item v-if="item.items" v-for="subitem in item.items" :icon="subitem.icon"
                            :label="subitem.label" :to="subitem.to" :active.sync="subitem.to == $route.path"
                            :tag="subitem.to?'router-link':'a'" :key="subitem.label">
                            <!-- reduce? '': -->
                        </b-menu-item>
                    </b-menu-item>
                </b-menu-list>
            </b-menu>
        </div>
        </div>
        <div class="column panel-column animate__animated animate__fadeIn">

            <b-navbar class="animate__animated animate__slideInDown">
                <template #brand>
                    <button class="button m-1" @click.prevent="toggle_menu">
                        <b-icon icon="menu">
                        </b-icon>
                    </button>
                    <b-navbar-item tag="router-link" :to="{ path: '/' }">
                        <img src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
                            alt="Lightweight UI components for Vue.js based on Bulma">
                    </b-navbar-item>
                </template>
                <template #start>
                    <b-navbar-item href="#">
                        Home
                    </b-navbar-item>
                    <b-navbar-item href="#">
                        Documentation
                    </b-navbar-item>
                    <b-navbar-dropdown label="Info">
                        <b-navbar-item href="#">
                            About
                        </b-navbar-item>
                        <b-navbar-item href="#">
                            Contact
                        </b-navbar-item>
                    </b-navbar-dropdown>
                </template>
                <template #end>
                    <b-navbar-dropdown :right="true">
                        <template #label>
                            <span class="columns is-vcentered">
                                <b-icon icon="account"></b-icon>
                                <span>{{user.Name}}</span>
                            </span>
                        </template>
                        <b-navbar-item href="#">
                            About
                        </b-navbar-item>
                        <b-navbar-item @click.prevent="logout">
                            Logout
                        </b-navbar-item>

                    </b-navbar-dropdown>
                </template>
            </b-navbar>
            <transition name="custom-classes-transition" mode="out-in"
                enter-active-class="animate__animated animate__slideInRight animate__faster">
                <router-view></router-view>
            </transition>
        </div>
    </div>
`
    }
