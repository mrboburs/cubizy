
    export default {
        props: {
            session: {
                type: Object
            },
        },
        components: {
            'AccountReviewsEditor': () => import("/vue/accountreviewseditor.js"),
        },
        data() {
            return {
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Title", sortable: false },
                    { key: "Details", sortable: false },
                    { key: "Replay", sortable: false },
                    { key: "Rating", title:"Avg.Rating"},
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                ],
                table: false,
                error: "",
                message: "",
                item_actions: [
                    // {
                    //     key: "view",
                    //     icon: "fas fa-eye",
                    //     text: "View"
                    // }
                ],
                conditions: {},
                new: {
                    ID: 0,
                    Content: "",
                },
            }
        },
        watch: {
            columns: function (newValue, oldValue) {
                if (newValue) {
                    this.$emit("onset", this.value);
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
                        if (this.session && this.session.ID && this.session.ID > 0) {
                            this.table.editing_item = Object.assign({
                                SessionID: this.session.ID,
                            }, this.new);
                        }else{
                            this.table.editing_item = Object.assign({}, this.new);
                        }
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
                        })
                    }
                }, 100);
            },
            init(){
                if (this.session && this.session.ID && this.session.ID > 0) {
                    this.conditions = {
                        session_id: this.session.ID
                    }
                    if(this.columns[2].key == "Title"){
                        this.columns.splice(2, 1)
                    }
                }else{
                    if(this.columns[2].key != "Title"){
                        this.columns.splice(2, 0, { key: "Title", sortable: true })
                    }
                }
            }
        },
        mounted: function () {
            this.init()
        },
        template: `
    <div class="card">
        <div class="card-body">
            <List api="accountreviews" :columns="columns" title_column="Title" :can_select="true" :can_delete="false" :can_add="false" :can_export="true"
                :can_import="false" :item_actions="item_actions" @done="onActionDone" @onaction="onAction"
                :conditions="conditions">
                <template v-slot:editor="editing_item">
                    <AccountReviewsEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                    </AccountReviewsEditor>
                </template>
            </List>
        </div>
    </div>
`
    }
