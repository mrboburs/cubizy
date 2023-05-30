
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
                    { key: "Logo", type: 'image', sortable: false },
                    { key: "Name", sortable: true, sortkey: "title" },
                    { key: "Review", sortable: false },
                    { key: "Replay", sortable: false },
                    { key: "Verefied", type: 'boolean', sortable: false },
                    { key: "Rating", title:"Rating"},
                    { key: "CreatedByName"},
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
            init(){
            }
        },
        mounted: function () {
            this.init()
        },
        template: `
    <div class="card">
        <div class="card-body">
            <List api="reviews" :columns="columns" title_column="ID" :can_select="true" :can_delete="true" :can_add="false" :can_export="true"
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
