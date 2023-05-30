<script>
    export default {
        data() {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                title: 'Themes',
                breadcrumb: [{
                    text: 'Theme',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "Image", type: 'image', sortable: false, sortkey: "image" },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "CreatedByName", title: "Updatedby", sortable: true, sortkey: "created_by_name" },
                ],
                table: false,
                item_actions: [
                    {
                        key: "get",
                        icon: "ri-download",
                        text: "Get"
                    }
                ],
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                conditions : {
                    status : "Published",
                },
                showsubthemesof: false,
                updating: [],
                Themes: [],
                first_loading: true,
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        mounted: function () {
            this.LoadAccountThemes()
        },
        methods: {
            onAction(action, arg) {
                switch (action) {
                    case 'loading':
                        this.loading = true
                        this.table = arg
                        break;
                    case 'add_new':
                        this.table.editing_item = Object.assign({}, this.new)
                        break;
                    default:
                        break;
                }
            },
            onActionDone(data) {
                this.loading = false
            },
            showsubthemes(record) {
                if (record && record.ID > 0) {
                    this.showsubthemesof = record
                }
            },
            owned(theme) {
                var themes = this.Themes.filter(item => (item.PublishedThemeID == theme.ID));
                return themes.length
            },
            haveUpdates(theme) {
                var themes = this.Themes.filter(item => (item.PublishedThemeID == theme.ID));
                if(themes.length){
                    if(themes[0].UpdatedAt < theme.UpdatedAt){
                        return true
                    }
                }
                return false
            },
            getTheme(theme) {
                this.updating.push(theme.ID)
                this.LoadAccountThemes({
                    GetPublishedThemeID: theme.ID
                }).finally(() => {
                    this.updating.splice(this.updating.indexOf(theme.ID), 1)
                });
            },
            LoadAccountThemes(data) {
                if (!data) {
                    data = {}
                    this.loading = true
                }
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
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                    if (!this.Themes.length) {
                        this.message = "Please get one theme first to enable website for this account"
                        this.error = true
                    }
                    if (!this.first_loading && this.Themes.length == 1) {
                        this.$router.push('/website/settings')
                    }
                    this.first_loading = false
                })
            },
            showTestingThems(){
                this.conditions = {
                    status : "Testing"
                }
            },
            showPublishedThems(){
                this.conditions = {
                    status : "Published"
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="col-12">
        <divloading :fullpage="false" :loading="loading" class="card">
            <v-alert v-model="message" :error="error" />
            <div v-if="account.TestAccount" class="d-flex justify-content-center p-1">
                <button v-if="conditions.status == 'Published'" class="btn btn-outline-primary" @click.prevent="showTestingThems"> Show themes under testing </button>
                <button v-if="conditions.status == 'Testing'" class="btn btn-outline-primary" @click.prevent="showPublishedThems"> Show published themes </button>
            </div>
            <v-list api="publishedthemes" :columns="columns" :can_select="false" @done="onActionDone" @onaction="onAction"
                editor_size="end" :conditions="conditions">
                <!--fullscreen-->
                <template v-slot:recordcard="{record}">
                    <divloading class="d-flex flex-column" :fullpage="false" :loading="updating.includes(record.ID)">
                        <h6 class="text-center text-capitalize">{{record.Title}}</h6>
                        <div :style="{'background-image': 'url('+ encodeURI(record.Logo)+')'}"
                            class="img-thumbnail flex-1"
                            style="width: 200px;min-height: 250px;background-position: left top;background-size: cover;">
                        </div>
                        <label v-if="haveUpdates(record)"> New version available</label>
                        <label v-else-if="owned(record)">You own it</label>
                        <button type="button" class="btn btn-sm btn-outline-primary border-0 me-1 flex-1"
                            :disabled="updating.includes(record.ID)" @click.prevent="getTheme(record)">
                            <span v-if="haveUpdates(record)">Update</span>
                            <span v-else-if="owned(record)">Reset Theme</span>
                            <span v-else>Get Theme</span>
                        </button>
                    </divloading>
                </template>
            </v-list>
        </divloading>
    </div>
</template>