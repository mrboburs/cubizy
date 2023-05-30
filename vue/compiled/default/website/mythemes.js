
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
                columns: [
                    { key: "Image", type: 'image', sortable: false, sortkey: "image" },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "CreatedByName", title: "Updatedby", sortable: true, sortkey: "created_by_name" },
                ],
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                publish_theme: false,
            }
        },
        watch: {},
        computed: {
            ...Vuex.mapState(['user','account']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
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
            can_add: function () {
                var id = 0
                if(this.user.IsAdmin){
                    return true
                }else{
                    return false
                }
            },
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
            apply(theme) {
                if (!theme || !theme.ID) {
                    return
                }
                var value = {}
                if (this.account) {
                    value.ID = this.account.ID
                }
                value.Active = this.Active
                value.ThemeID = theme.ID
                this.loading = true
                this.$store.dispatch('call', {
                    api: "account",
                    data: {
                        account: value
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        if (this.account.ThemeID > 0) {
                            this.$emit('done', data)
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
            onActionDone(data) {
                this.loading = false
                if (!this.$refs.vlist.recordsTotal && this.account.AccountType != 'admin') {
                    this.$router.push('/website/themes')
                }
            },
            update_publishing_theme(theme) {
                if (theme) {
                    this.publish_theme.PublishedID = theme.PublishedID
                    this.publish_theme.Status = theme.Status
                    this.publish_theme.SubmitedOn = theme.SubmitedOn
                }
                this.publish_theme = false
            },            
            creat_a_copy(theme) {
                this.table.loadData({
                    CopyFromThemeID: theme.ID
                }, this.table)
            },
        },
        mounted: function () { },
        template: `
    <div>
        <v-alert v-model="message" :error="error" />
        <v-list ref="vlist" api="themes" :columns="columns" :can_add="can_add" :can_edit="true" :can_delete="true"
            :can_bulk_delete="true" @done="onActionDone" @onaction="onAction" editor_size="end" :can_select="false">
            <!--fullscreen-->
            <template v-slot:recordcard="{record}">
                <divloading :fullpage="false" :loading="loading" class="d-flex justify-content-between flex-column">
                    <h5 class="text-center" style="max-width: 100%;">
                        ({{record.ID}})
                        {{record.Title}}
                        <span v-if="record.ID == account.ThemeID">(Current Theme)</span>
                    </h5>
                    <div :style="{'background-image': 'url('+ encodeURI(record.Logo)+')'}" class="img-thumbnail flex-1"
                        style="width: 100%;min-height: 250px;min-width: 250px;background-position: left top;background-size: cover;">
                    </div>
                    <div class="d-flex pt-1">
                        <button v-if="record.ID != account.ThemeID" class="btn btn-sm btn-outline-primary mb-1 me-1 flex-1" @click.prevent="apply(record)">
                            <i class="fas fa-brush"></i>
                            <span v-else>Apply</span>
                        </button>
                        <button class="btn btn-sm btn-outline-primary flex-1 mb-1"
                                        @click.prevent="creat_a_copy(record)">
                                        <i class="fas fa-copy"></i>
                                        <span>Copy</span>
                                    </button>
                        <div class="btn-group dropup mb-1 " v-if="record.PublishedThemeID == 0 || record.SubmitedOn > 0">
                            <button type="button" class="btn btn-sm btn-outline-primary dropdown-toggle"
                                data-bs-toggle="dropdown" aria-expanded="false">
                                <i class="fas fa-ellipsis-v"></i>
                            </button>
                            <div class="dropdown-menu">
                                <div class="d-flex flex-column m-1">
                                    <button class="btn btn-sm btn-outline-primary flex-1 mb-1"
                                        @click.prevent="publish_theme = record">
                                        <i class="fas fa-upload"></i>
                                        <span v-if="record.PublishedThemeID">Publish Updates</span>
                                        <span v-else>Publish</span>
                                    </button>
                                    <label class="badge bg-primary p-1 m-1" v-if="record.Status"> {{record.Status}}</label>
                                    <label class="badge bg-warning p-1 m-1" v-if="record.SubmitedOn"> 
                                        Submited : {{getFullDateTime(record.SubmitedOn)}}
                                    </label>
                                    <label class="badge bg-success p-1 m-1" v-if="record.PublishedOn"> 
                                        Published : {{getFullDateTime(record.PublishedOn)}}
                                    </label>
                                </div>
                            </div>
                        </div>
                    </div>
                </divloading>
            </template>
            <template v-slot:editor="editing_item">
                <themeeditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                </themeeditor>
            </template>
        </v-list>
        <v-modal v-model="showpublishpanel" :title="'Publish '+publish_theme.Title">
            <publishtheme v-if="publish_theme" :value="publish_theme" @input="update_publishing_theme" />
        </v-modal>
    </div>
`
    }
