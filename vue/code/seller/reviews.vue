<script>
    export default {
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
        methods: {
            onAction(action, arg) {
                switch (action) {
                    case 'loading':
                        this.loading = true
                        this.table = arg
                        break;
                    default:
                        break;
                }
            },
            onActionDone(data) {
                this.loading = false
            },
        },
        mounted: function () {},
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="card">
        <div class="card-body">
            <List api="reviews" :columns="columns" title_column="Title" :can_select="true" :can_delete="false" :can_add="false" :can_export="true" :can_import="false" :item_actions="item_actions" @done="onActionDone" @onaction="onAction" :conditions="conditions">
                <template v-slot:editor="editing_item">
                    <AccountReviewsEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                    </AccountReviewsEditor>
                </template>
            </List>
        </div>
    </div>
</template>