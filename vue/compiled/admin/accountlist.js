
    export default {
        components: {
            'AccountEditor': () => import("/vue/accounteditor.js"),
        },
        data() {
            return {
                title: 'Accounts',
                breadcrumb: [{
                    text: 'Account',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: 'image', sortable: false },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "Photo", type: 'image', sortable: false },
                    { key: "Name", title:"Owner" , sortable: true, sortkey: "name" },
                    // { key: "AccountType", sortable: true, sortkey: "account_type" },
                    { key: "TestAccount", title: "Is Test Account", sortable: true, sortkey: "test_account" },
                    { key: "Status", title: "Status", sortable: true, sortkey: "status" },
                    { key: "CanActive", type: 'boolean', title: "Website", sortable: true, sortkey: "can_active" },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                    { key: "UpdatedByName", title: "Updatedby", sortable: true, sortkey: "updated_by_name" },
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
                    Name: "",
                    Status: true
                },
                status : -1,
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
            updateTestAccount(record) {
                setTimeout(() => {
                    if (this.table) {
                        this.table.loadData({
                            items: [record]
                        }, this.table)
                    }
                }, 100);
            },
            init() {
                switch (this.$route.name) {
                    case "active":
                        this.status = 10
                        break;
                    case "pending":
                        this.status = 1
                        break;
                    case "expired":
                        this.status = 4
                        break;
                    default:
                        this.status = -1
                        break;
                }
                if(this.status > 0){
                    this.conditions = {
                        status: this.status
                    }
                }else{
                    this.conditions = {}
                }
            },
        },
        mounted() {
            this.init()
        },
        template: `
    <!-- Start Content-->
    <div class="row">
        <div class="col-12">
            <div class="card">
                <List api="accounts" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                    :can_add="false" :can_import="true" :actions="actions" :conditions="conditions" @done="onActionDone"
                    @onaction="onAction" editor_type="offcanvas" editor_size="end">
                    <template v-slot:TestAccount="{ row, col }">
                        <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" :id="'inputTestAccount_'+row.ID"
                                v-model="row.TestAccount" @change="updateTestAccount(row)" >
                            <!-- <label class="form-check-label" :for="'inputStatus_'+row.ID">Enabled</label> -->
                        </div>
                    </template>
                    <template v-slot:Status="{ row }">
                        <span class="text-capitalize text-warning" v-if="row.Status == 1">Under Review</span>
                        <span class="text-capitalize text-danger" v-if="row.Status == 2">Rejected</span>
                        <span class="text-capitalize text-warning" v-if="row.Status == 3">OnHold</span>
                        <span class="text-capitalize text-danger" v-if="row.Status == 4">Expired</span>
                        <span class="text-capitalize text-success" v-if="row.Status == 10">Active</span>
                    </template>
                    <template v-slot:editor="editing_item">
                        <AccountEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </AccountEditor>
                    </template>
                </List>
            </div>
        </div>
        <!-- Modal  -->
    </div>
`
    }
