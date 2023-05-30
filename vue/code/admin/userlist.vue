<script>
    export default {
        components: {
            'UserEditor': () => import("/vue/usereditor.js"),
        },
        data() {
            return {
                title: 'Users',
                breadcrumb: [{
                    text: 'User',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "Photo", type: 'image', sortable: false },
                    { key: "Name", sortable: true, sortkey: "name" },
                    { key: "LastLoginOn", type: 'date_time', sortable: true, sortkey: "last_login_on" },
                    { key: "LoginOn", type: 'date_time', sortable: true, sortkey: "login_on" },
                    { key: "LastActiveOn", type: 'date_time', sortable: true, sortkey: "last_active_on" },
                    { key: "Wallet", type: 'money', sortable: true, sortkey: "wallet" },
                    { key: "IsSuperAdmin", type: 'boolean', sortable: true, sortkey: "is_super_admin" },
                    { key: "IsAdmin", type: 'boolean', sortable: true, sortkey: "is_admin" },
                    { key: "SellerAccountID", title: "IsSeller", type: 'boolean', sortable: true, sortkey: "seller_account_id" },
                    { key: "IsSupportagent", type: 'boolean', sortable: true, sortkey: "is_supportagent" },
                ],
                table: false,
                error: "",
                message: "",
                actions: [
                    // {
                    //     key: "import",
                    //     icon: "ri-upload-line",
                    //     text: "Import"
                    // }
                ],
                conditions: {},
                new: {
                    ID: 0,
                    Code: "",
                    SubLocality: "",
                    Locality: "",
                    District: "",
                    Country: "Mauritius",
                    Latitude: 0,
                    Longitude: 0,
                },
            }
        },
        watch: {
            $route(to, from) {
                this.init()
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
            onActionDone(data) {
                this.loading = false
            },
            init(key) {
                switch (this.$route.name) {
                    case 'admins':
                        this.conditions = {
                            user_type: "admins"
                        }
                        this.columns.forEach(element => {
                            if(element.key == "IsAdmin"){
                                element.always_hide = true
                            }
                            if(element.key == "SellerAccountID"){
                                element.always_hide = false
                            }
                        });
                        this.columns = JSON.parse(JSON.stringify(this.columns))
                        break;
                    case 'sellers':
                        this.conditions = {
                            user_type: "sellers"
                        }
                        this.columns.forEach(element => {
                            if(element.key == "IsAdmin"){
                                element.always_hide = false
                            }
                            if(element.key == "SellerAccountID"){
                                element.always_hide = true
                            }
                        });
                        this.columns = JSON.parse(JSON.stringify(this.columns))
                        break;
                    default:
                        this.conditions = {}
                        this.columns.forEach(element => {
                            if(element.key == "IsAdmin"){
                                element.always_hide = false
                            }
                            if(element.key == "SellerAccountID"){
                                element.always_hide = false
                            }
                        });
                        this.columns = JSON.parse(JSON.stringify(this.columns))
                        break;
                }
            }
        },
        mounted() {
            this.init()
        },
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="col-12">
        <div class="card">
            <List api="users" :columns="columns" title_column="Code" :can_select="true" :can_export="true" :conditions="conditions" :can_import="true" :actions="actions" editor_size="lg" @done="onActionDone" @onaction="onAction">
                <template v-slot:Photo="{ row, col }">
                    <div class="position-relative me-2">
                        <span v-if="row.Online" class="user-status online"></span>
                        <img :src="row.Photo" class="rounded-circle avatar-sm"
                            alt="user-pic">
                    </div>
                </template>
                <template v-slot:Name="{ row, col }">
                    <span>
                        {{row.Name}}<br/>
                        <a :href="'mailto: '+row.Email">{{row.Email}}</a><br/>
                        <a :href="'tel: '+row.Mobile">{{row.Mobile}}</a>
                    </span>
                </template>
                <template v-slot:editor="editing_item">
                    <UserEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                    </UserEditor>
                </template>
            </List>
        </div>
    </div>
</template>