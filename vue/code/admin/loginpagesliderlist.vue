<script>
    export default {
        components: {
            'LoginpagesliderEditor': () => import("/vue/loginpageslidereditor.js"),
        },
        data() {
            return {
                title: 'Loginpagesliders',
                breadcrumb: [{
                    text: 'Loginpageslider',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Image", type: 'image', sortable: false, sortkey: "image" },
                    { key: "Title", sortable: true, sortkey: "title" },
                    { key: "Content", always_hide: true  },
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
                showsubloginpageslidersof: false
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
            showsubloginpagesliders(record) {
                if(record && record.ID > 0){
                    this.showsubloginpageslidersof = record
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
        <div class="col-12">
            <div class="card">
                <List api="loginpagesliders" :columns="columns" title_column="Code" :can_select="true" :can_export="true" editor_type="offcanvas"
                    :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction" editor_size="end"> <!--fullscreen-->
                    <template v-slot:Status="{ row, col }">
                        <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" :id="'inputStatus_'+row.ID"
                                v-model="row.Status" @change="updateStatus(row)">
                            <!-- <label class="form-check-label" :for="'inputStatus_'+row.ID">Enabled</label> -->
                        </div>
                    </template>
                    <template v-slot:SubLoginpagesliderCount="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showsubloginpagesliders(row)">
                                {{row.SubLoginpagesliderCount}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <LoginpagesliderEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </LoginpagesliderEditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
</template>