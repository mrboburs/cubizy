
    export default {
        components: {
            'ThemeEditor': () => import("/vue/pthemeeditor.js"),
        },
        data() {
            return {
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
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: 'image', sortable: false },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "Status", title: "Status", sortable: true, sortkey: "status" },
                    { key: "CreatedAt", title: "SubmitedOn", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                    { key: "CreatedByName", title: "Updatedby", sortable: true, sortkey: "created_by_name" },
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
                Status_Types : ['', 'Fail to Upload', 'Submiting', 'Submitted', 'Rejected', 'Accepted', 'Published'],
                showsubthemesof: false,
                conditions : {}
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
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
                if(record && record.ID > 0){
                    this.showsubthemesof = record
                }
            }
        },
        template: `
        <div class="col-12">
            <div class="card">
                <List api="publishedthemes" :columns="columns" title_column="Code" :conditions="conditions" :can_select="true" :can_export="true" editor_type="offcanvas"
                    :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction" editor_size="end"> <!--fullscreen-->
                    <template v-slot:SubThemeCount="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showsubthemes(row)">
                                {{row.SubThemeCount}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <ThemeEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </ThemeEditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
`
    }
