<script>
    export default {
        components: {
            'ProductEditor': () => import("/vue/products/edit.js"),
        },
        data() {
            return {
                title: 'Products',
                breadcrumb: [{
                    text: 'Product',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: 'image', sortable: false },
                    { key: "Name", sortable: true, sortkey: "title" },
                    { key: "Condition" },
                    { key: "Price" },
                    { key: "MaxPrice" },
                    { key: "Discount" },
                    { key: "Cost" },
                    { key: "MaxCost" },
                    { key: "Status", title: "Active", type: 'boolean', sortable: true, sortkey: "status" },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                    { key: "UpdatedByName", title: "Updatedby", sortable: true, sortkey: "updated_by_name" },
                ],
                table: false,
                error: "",
                message: "",
                item_actions: [
                    {
                        key: "edit",
                        icon: "fas fa-edit",
                        text: "Edit"
                    }
                ],
                conditions:{
                    Service : false
                },
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                product : null,
            }
        },
        
        computed: {
            ...Vuex.mapState(['user', 'account']),
            editPanel: {
                // getter
                get: function () {
                    if (this.product) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.product = false
                    }
                }
            },
            editPanelTitle: function () {
                if (this.product) {
                    return this.product.Name
                } else {
                    return ''
                }
            }
        },
        watch: {
            product: function (newValue) {
                if (!newValue && this.table) {
                    this.table.reload()
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
                    case 'edit':
                        this.product = arg
                        break
                    case 'add_new':
                        this.$router.push('/products/add')
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
            showsubproducts(record) {
                if(record && record.ID > 0){
                    this.showsubproductsof = record
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
        <div class="col-12">
            <div class="card">
                <List api="products" :columns="columns" title_column="Code" :can_select="true" :can_export="true" :can_edit="false"
                    :can_import="true" :item_actions="item_actions" :conditions="conditions" @done="onActionDone" @onaction="onAction" editor_size="end"> <!--fullscreen-->
                    <template v-slot:Status="{ row, col }">
                        <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" :id="'inputStatus_'+row.ID"
                                v-model="row.Status" @change="updateStatus(row)">
                            <!-- <label class="form-check-label" :for="'inputStatus_'+row.ID">Enabled</label> -->
                        </div>
                    </template>
                </List>
            </div>
            <v-offcanvas v-model="editPanel" :title="editPanelTitle" >
                <ProductEditor v-if="product" v-model="product" class="vw90"/>
            </v-offcanvas>
        </div>
</template>