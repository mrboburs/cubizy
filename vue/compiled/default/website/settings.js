
    export default {
        components: {
            'themeeditor': () => import("/vue/website/themeeditor.js"),
            'publishtheme': () => import("/vue/website/publishtheme.js"),
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Protocol: window.location.protocol,
                BaseDomin: window.application.BaseDomin,
                Subdomain: "",
                Domain: "",
                Themes: [],
                ThemePath: "",
                Theme: "",
                ThemeSettings: false,
                ThemeID: 0,
                Active: true,
                theme_to_edit: false,
                publish_theme: false,
                deletingtheme: [],
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                }
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
                if(oldValue && newValue && oldValue.ID == newValue.ID){
                    return
                }
                this.init()
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            showpanel: {
                // getter
                get: function () {
                    if (this.theme_to_edit) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.theme_to_edit = false
                    }
                }
            },
            showpublishpanel: {
                // getter
                get: function () {
                    if (this.publish_theme) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.publish_theme = false
                    }
                }
            },
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
            DomainError: function () {
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
                if (!this.ThemeID) {
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
                if (!this.ThemeID) {
                    return "Please selecte theme for website "
                }
            },
        },
        methods: {
            can_load() {
                if (this.loading) {
                    return false
                }
                return true
            },
            load(data) {
                if (!data) {
                    data = {}
                }
                data.fix_condition = {
                    created_by: this.user.ID
                }
            },
            submit(record) {
                if (!record) {
                    return
                }
                this.submitted = true
                if (this.ActiveError || this.ThemeError) {
                    return
                }
                var value = {}
                if (this.account) {
                    value.ID = this.account.ID
                }
                value.Subdomain = this.Subdomain
                value.Domain = this.Domain
                value.Active = this.Active
                value.ThemeID = this.ThemeID

                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        account: value
                    }
                }).then(function (data) {
                    component.message = data.Message;
                    if (data.Status == 2) {
                    } else {
                        component.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    component.error = true
                    component.message = error
                }).finally(() => {
                    component.loading = false
                })
            },
            LoadThemes(data) {
                if (!data) {
                    data = {}
                }
                this.loading = true
                return this.$store.dispatch('call', {
                    api: "themes",
                    data: data
                }).then((data) => {
                    this.message = data.Message;
                    if (Array.isArray(data.data)) {
                        this.Themes = data.data
                    }
                    if (data.Status == 2) {
                        this.theme_to_edit = false
                    }else{
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                    this.SetData()
                })
            },
            SetData() {
                if (this.account) {
                    if(!this.Themes.length && this.account.AccountType != 'admin'){
                        this.$router.push('/website/themes')
                    }
                    this.submitted = false
                    if (this.account && this.account.ID > 0) {
                        this.submitted = false
                        this.Subdomain = this.account.Subdomain
                        this.Domain = this.account.Domain
                        this.ThemeID = this.account.ThemeID
                        this.Active = this.account.Active
                    }
                    if (!this.Subdomain.trim()) {
                        this.Subdomain = this.account.Title.replace(/\s/g, "")
                    }
                    if(this.Themes.length == 1 && !this.ThemeID){
                        this.ThemeID = this.Themes[0].ID
                    }
                }
            },
            CreateTheme() {
                this.theme_to_edit = Object.assign({}, this.new)
            },
            submitTheme(theme) {
                if(theme){
                    this.theme_to_edit = theme
                    this.LoadThemes({
                        items: [theme]
                    });
                }else{
                    this.theme_to_edit = false
                }
            },
            delete_theme(theme){
                if(theme){
                    this.deletingtheme.push(theme.ID)
                    this.LoadThemes({
                        "todelete": [
                            theme.ID
                        ]
                    }).finally(() => {
                        this.deletingtheme.splice(this.deletingtheme.indexOf(theme.ID), 1)
                    });
                }
            },
            update_publishing_theme(theme){
                if(theme){
                    this.publish_theme.PublishedID = theme.PublishedID
                    this.publish_theme.Published = theme.Published
                }
                this.publish_theme = false
            },
            init(){
                if (this.account && this.account.CanActive) {
                    this.LoadThemes()
                }
            },
        },
        mounted: function () {
            this.init()
        },
        template: `
    <section class="section is-medium">
        <v-alert v-model="message" :error="error" />
        <div class="card" v-if="account">
            <div class="card-body">
                <p v-if="!account.CanActive" class="text-info">
                    Website service is not active for your account, please contact support to activate it
                </p>
                <divloading v-else :fullpage="false" :loading="loading" class="container">
                    <form @submit.prevent="submit" class="flex" v-if="account">
                        <formitem :customLayout="true" name="inputSubdomain" label="Subdomain" v-model="Subdomain"
                            inputclass="w-auto" :error="SubdomainError" :suffix="'.'+BaseDomin" placeholder="Subdomin"
                            :inputgroup="true" v-if="account.AccountType != 'admin'" />
                        <formitem :customLayout="true" name="inputDomain" label="Domain" v-model="Domain"
                            inputclass="w-auto" :error="DomainError" placeholder="Domin"
                            :type="(account.AccountType == 'admin')?'readonly':'text'" />
                        <formitem :customLayout="true" name="inputThemeID" label="Theme" v-model="ThemeID"
                            :error="ThemeError">
                            <div class="theme_list d-flex flex-wrap p-1 gap-1">
                                <divloading v-for="_theme in Themes" :key="'theme'+_theme.ID" :fullpage="false" :loading="deletingtheme.includes(_theme.ID)"  class="d-flex justify-content-between flex-column border border-4 rounded" :class="{ 'border-primary' : _theme.ID == ThemeID }">
                                    <label class="text-center" style="max-width: 250px;">{{_theme.Title}}</label>
                                    <div :style="{'background-image': 'url('+ encodeURI(_theme.Image)+')'}" @click.prevent="ThemeID = _theme.ID" class="img-thumbnail flex-1" style="width: 250px;min-height: 300px;background-position: left top;background-size: cover;"></div>
                                    <span class="text-muted" v-if="_theme.ID != ThemeID">Click on image to select</span>
                                    <div class="d-flex justify-content-between align-items-center p-1">
                                        <button type="button" class="btn btn-sm btn-primary"
                                            @click="publish_theme = _theme">
                                            <span v-if="_theme.PublishedID">Republish</span> 
                                            <span v-else>Publish</span>
                                        </button>
                                        <button type="button" class="btn btn-sm btn-danger" @click="delete_theme(_theme)" v-if="_theme.ID != ThemeID">
                                            <i class="fas fa-trash"></i>
                                        </button>
                                        <button type="button" class="btn btn-sm btn-info" @click="theme_to_edit = _theme">
                                            <i class="fas fa-edit"></i>
                                        </button>
                                    </div>
                                </divloading>
                            </div>
                        </formitem>
                        <formitem :customLayout="true" name="inputActive" label="Active" v-model="Active"
                            :error="ActiveError" type="switch" />
                        <div class="d-flex justify-content-end">
                            <a v-if="account.Subdomain && account.AccountType != 'admin'"
                                class="btn btn-outline-secondary m-2" type="button" target="_blank"
                                :href="Protocol+'//'+account.Subdomain + '.' +  BaseDomin">
                                Visit Subdomain
                            </a>
                            <a v-if="account.Domain" class="btn btn-outline-secondary m-2" type="button" target="_blank"
                                :href="Protocol+'//'+account.Domain">
                                Visit Domain
                            </a>
                            <button type="button" class="btn btn-success m-2" :disabled="loading"
                                @click.prevent="CreateTheme()"> Create New theme </button>
                            <button type="submit" class="btn btn-success m-2" :disabled="loading"> Save </button>
                            <button class="btn btn-danger m-2" @click.prevent="SetData">Cancel</button>
                        </div>
                    </form>
                </divloading>
            </div>
        </div>
        <v-modal v-model="showpanel" :title="'Edit ' + theme_to_edit.Title">
            <divloading :fullpage="false" :loading="loading">
                <themeeditor v-if="theme_to_edit" :value="theme_to_edit" @input="submitTheme" :themes="Themes"></themeeditor>
            </divloading>
        </v-modal>
        <v-modal v-model="showpublishpanel" :title="'Publish '+publish_theme.Title">
            <publishtheme v-if="publish_theme" :value="publish_theme" @input="update_publishing_theme()"/>
        </v-modal>
    </section>
`
    }
