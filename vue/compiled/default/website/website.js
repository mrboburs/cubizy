
    export default {
        components: {
            'mythemes': () => import("/vue/website/mythemes.js"),
            'themesettings': () => import("/vue/website/themesettings.js"),
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                SelectedTab: 1,
                ThemeSettings: "",
                Protocol: window.location.protocol,
                BaseDomin: window.application.BaseDomin,
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
            account: function (newValue, oldValue) {
                if (oldValue && newValue && oldValue.ID == newValue.ID) {
                    return
                }
                this.SetData()
            },
        },
        computed: {
            ...Vuex.mapState(['account']),
            SubdomainError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Active) {
                    return false
                }
                if (!this.Subdomain.trim() && !this.Domain.trim()) {
                    return "Add atlist subdomin or domin to activate website"
                }
            },
            ThemeError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Active) {
                    return false
                }
                if (!this.account.ThemeID) {
                    return "Please selecte theme for website "
                }
            },
            ActiveError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Active) {
                    return false
                }
                if (!this.Subdomain.trim() && !this.Domain.trim()) {
                    return "Add atlist subdomin or domin to activate website"
                }
                if (!this.account.ThemeID) {
                    return "Please selecte theme for website "
                }
            },
        },
        methods: {
            SetData() {
                if (this.account) {
                    this.submitted = false
                    if (this.account && this.account.ID > 0) {
                        this.submitted = false
                        this.Subdomain = this.account.Subdomain
                        this.Domain = this.account.Domain
                        this.Active = this.account.Active
                    }
                    if (!this.Subdomain.trim()) {
                        this.Subdomain = this.account.Title.replace(/\s/g, "")
                    }
                }
            },
            submit(record) {
                if (!record) {
                    return
                }
                this.submitted = true
                if (this.SubdomainError || this.ActiveError || this.ThemeError) {
                    return
                }
                var value = {}
                if (this.account) {
                    value.ID = this.account.ID
                }
                value.Subdomain = this.Subdomain
                value.Domain = this.Domain
                value.Active = this.Active

                if (this.$refs.theme_setting_editor.validate()) {
                    value.ThemeSettings = this.ThemeSettings
                } else {
                    return
                }


                this.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        account: value
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.$refs.ifrem_priview.src = this.$refs.ifrem_priview.src;
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
            ResetThemeSettings() {
                this.$refs.theme_setting_editor.reset()
            },

            togggleFullScreen() {
                // if already full screen; exit
                // else go fullscreen
                if (
                    document.fullscreenElement ||
                    document.webkitFullscreenElement ||
                    document.mozFullScreenElement ||
                    document.msFullscreenElement
                ) {
                    if (document.exitFullscreen) {
                        document.exitFullscreen();
                    } else if (document.mozCancelFullScreen) {
                        document.mozCancelFullScreen();
                    } else if (document.webkitExitFullscreen) {
                        document.webkitExitFullscreen();
                    } else if (document.msExitFullscreen) {
                        document.msExitFullscreen();
                    }
                } else {
                    //var element = this.$refs.editor_holder.$el;
                    var element = document.getElementById("fullscreen_priview")
                    if (element.requestFullscreen) {
                        element.requestFullscreen();
                    } else if (element.mozRequestFullScreen) {
                        element.mozRequestFullScreen();
                    } else if (element.webkitRequestFullscreen) {
                        element.webkitRequestFullscreen(Element.ALLOW_KEYBOARD_INPUT);
                    } else if (element.msRequestFullscreen) {
                        element.msRequestFullscreen();
                    }
                }
            }
        },
        mounted: function () {
            this.SetData()
            if (this.account.ThemeID > 0) {
                this.SelectedTab = 2
            } else {
                this.SelectedTab = 1
            }
            if (window.innerWidth >= 993) {
                setTimeout(() => {
                    document.body.setAttribute('data-sidebar-size', 'condensed');
                    Split([this.$refs.theme_setting_editor_holder, this.$refs.ifrem_priview], { sizes: [20, 80] })
                }, 30);
            }
        },
        template: `
    <section class="section is-medium">
        <v-alert v-model="message" :error="error" />
        <div class="card" v-if="account">
            <p v-if="!account.CanActive" class="text-info">
                Website service is not active for your account, please contact support to activate it
            </p>
            <divloading v-else :fullpage="false" :loading="loading">
                <ul class="nav nav-tabs">
                    <li class="nav-item">
                        <a class="nav-link" href="#" :class="{'active': SelectedTab == 1}"
                            @click.prevent="SelectedTab = 1">My Themes</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#" :class="{'active': SelectedTab == 2}"
                            @click.prevent="SelectedTab = 2">Settings</a>
                    </li>
                </ul>
                <div class="p-1" v-if="SelectedTab == 1">
                    <mythemes @done="SelectedTab = 2"></mythemes>
                </div>
                <div class="p-1" v-if="SelectedTab == 2">
                    <form @submit.prevent="submit" class="flex" v-if="account">
                        <div class="d-flex flex-wrap">
                            <formitem :customLayout="true" name="inputSubdomain" label="Subdomain" v-model="Subdomain"
                                inputclass="w-auto" :error="SubdomainError" :suffix="'.'+BaseDomin"
                                placeholder="Subdomin" :inputgroup="true" class="flex-fill"/>
                            <a v-if="account.Subdomain" class="btn btn-outline-secondary m-2" type="button"
                                target="_blank" :href="Protocol+'//'+account.Subdomain + '.' +  BaseDomin">
                                Visit Subdomain
                            </a>
                        </div>
                        <div class="d-flex flex-wrap">
                            <formitem :customLayout="true" name="inputDomain" label="Domain" v-model="Domain"
                                inputclass="w-auto" :error="SubdomainError" placeholder="Domin"
                                :type="(account.AccountType == 'admin')?'readonly':'text'" class="flex-fill"/>
                            <a v-if="account.Domain" class="btn btn-outline-secondary m-2" type="button" target="_blank"
                                :href="Protocol+'//'+account.Domain">
                                Visit Domain
                            </a>
                        </div>
                        <formitem :customLayout="true" name="inputActive" label="Active" v-model="Active"
                            :error="ActiveError" type="switch" />

                        <div class="d-flex flex-column bg-light" id="fullscreen_priview">
                            <div class="d-flex gap-1">
                                <button type="submit" class="btn btn-success flex-fill" :disabled="loading"> Apply and
                                    Update
                                </button>
                                <button type="button" class="btn btn-success" :disabled="loading"
                                    @click.prevent="ResetThemeSettings()"> Reset Theme Settings </button>
                                <button class="btn btn-primary" type="button" @click.prevent="togggleFullScreen">
                                    <i class="fe-maximize noti-icon"></i>
                                </button>
                            </div>
                            <div class="split">
                                <div ref="theme_setting_editor_holder" class="bg-body">
                                    <themesettings v-model="ThemeSettings" :theme_id="account.ThemeID"
                                        ref="theme_setting_editor"></themesettings>
                                </div>
                                <iframe ref="ifrem_priview" :src="Protocol+'//'+account.Subdomain + '.' +  BaseDomin"
                                    style="width: 100%;"></iframe>
                            </div>
                        </div>
                    </form>
                </div>
            </divloading>
        </div>
    </section>
`
    }
