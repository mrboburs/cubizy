<script>
    export default {
        components: {
            'PageEditor': () => import("/vue/website/pageeditor.js"),
        },
        data() {
            return {
                title: 'Pages',
                breadcrumb: [{
                    text: 'Page',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "Weightage", sortable: true, sortkey: "weightage" },
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
                showsubpagesof: false
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
            showsubpages(record) {
                if (record && record.ID > 0) {
                    this.showsubpagesof = record
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="card">
        <List api="pages" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
            editor_type="offcanvas" :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction"
            editor_size="end" default_sort_by="weightage" :default_desc="true">
            <!--fullscreen-->
            <template v-slot:Status="{ row, col }">
                <div class="form-check form-switch">
                    <input class="form-check-input" type="checkbox" :id="'inputStatus_'+row.ID" v-model="row.Status"
                        @change="updateStatus(row)">
                    <!-- <label class="form-check-label" :for="'inputStatus_'+row.ID">Enabled</label> -->
                </div>
            </template>
            <template v-slot:SubPageCount="{ row, col }">
                <div class="form-check form-switch">
                    <button class="btn btn-primary" type="button" @click="showsubpages(row)">
                        {{row.SubPageCount}}
                    </button>
                </div>
            </template>
            <template v-slot:editor="editing_item">
                <PageEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                </PageEditor>
            </template>
        </List>
    </div>
</template>