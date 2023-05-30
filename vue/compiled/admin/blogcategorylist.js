
    export default {
        components: {
            'blogcategoryeditor': () => import("/vue/blogcategoryeditor.js"),
        },
        data() {
            return {
                title: 'Blogs',
                breadcrumb: [{
                    text: 'Blog',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true },
                    { key: "Name", sortable: true },
                    { key: "Blogs", always_hide: true  },
                    { key: "Status", title: "Active", type: 'boolean', sortable: true, sortkey: "status" },
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
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                showsubblogsof: false
            }
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
            updateStatus(record) {
                setTimeout(() => {
                    if (this.table) {
                        this.table.loadData({
                            items: [record]
                        },this.table)
                    }
                }, 100);
            },
            showsubblogs(record) {
                if(record && record.ID > 0){
                    this.showsubblogsof = record
                }
            }
        },
        template: `
        <div class="col-12">
            <div class="card">
                <List api="blogcategories" :columns="columns" title_column="Code" :can_select="true" :can_export="true" 
                    :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction" editor_size="sm"> <!--fullscreen-->
                    <template v-slot:Status="{ row, col }">
                        <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" :id="'inputStatus_'+row.ID"
                                v-model="row.Status" @change="updateStatus(row)">
                            <!-- <label class="form-check-label" :for="'inputStatus_'+row.ID">Enabled</label> -->
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <blogcategoryeditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </blogcategoryeditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
`
    }
