
    export default {
        props: {
            subcategory_id: {
                type: Number,
                default: 0
            },
            basesubcategory_id: {
                type: Number,
                default: 0
            },
        },
        components: {
            'ChildcategoryEditor': () => import("/vue/account/childcategoryeditor.js"),
            'ProductList': () => import("/vue/account/productlist.js"),
            'AttributeList': () => import("/vue/attributelist.js"),
        },
        data() {
            return {
                title: 'Childcategories',
                breadcrumb: [{
                    text: 'Childcategory',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: "image" },
                    { key: "Name", sortable: true, sortkey: "name" },
                    { key: "Active", sortable: true, sortkey: "active", type: 'boolean' },
                    { key: "SubcategoryID", title: "SubcategoryID", sortable: true, sortkey: "subcategory_id", always_hide: false  },
                    { key: "Products", sortable: true, sortkey: "products" },
                    { key: "Revenue", sortable: true, sortkey: "revenue" },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                    { key: "UpdatedByName", title: "Updatedby", sortable: true, sortkey: "updated_by_name" },
                ],
                table: false,
                error: "",
                message: "",
                item_actions:[],
                conditions:{},
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                showProductsOf : false,
                showAttributesof : false
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            showpanel: {
                // getter
                get: function () {
                    if (this.showProductsOf) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.showProductsOf = false
                    }
                }
            },
            showAttributesPanel: {
                // getter
                get: function () {
                    if (this.showAttributesof) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.showAttributesof = false
                    }
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
                        this.table.editing_item = Object.assign({}, this.new)
                        break;
                    case 'showAttributes':
                        this.showAttributesof = arg
                        break;
                    default:
                        break;
                }
            },
            onActionDone(data) {
                this.loading = false
            },
            showProducts(record) {
                if(record && record.ID > 0){
                    this.showProductsOf = record
                }
            }
        },
        mounted() {
            if(this.subcategory_id){
                this.new.SubcategoryID = this.subcategory_id
                this.conditions = {
                    subcategory_id : this.subcategory_id
                }
                this.columns.forEach(column => {
                    if(column.key == "SubcategoryID"){
                        column.always_hide = true
                    }
                });
            }
            if(this.account.AccountType == 'admin'){
                this.item_actions=[
                    {
                        key : "showAttributes",
                        text: "Attributes"
                    }
                ]
            }
        },
        template: `
    <!-- Start Content-->
    <div class="row">
        <div class="col-12">
            <div class="card">
                <List api="childcategories" :columns="columns" title_column="Code" :can_select="true" :can_export="true" :can_add="subcategory_id > 0"
                    :can_import="true" :conditions="conditions" @done="onActionDone" @onaction="onAction"  editor_size="lg" :item_actions="item_actions" >
                    <template v-slot:Icon="{ row, col }" class="icon_holder">
                        <div class="icon_holder" v-html="row.Icon"></div>
                    </template>
                    <template v-slot:Products="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showProducts(row)">
                                {{row.Products}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <ChildcategoryEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit" :subcategory_id="subcategory_id" :basesubcategory_id="basesubcategory_id">
                        </ChildcategoryEditor>
                    </template>
                </List>
            </div>
        </div>
        <!-- Modal  -->
        <v-offcanvas v-model="showpanel" :title="'Products of '+ showProductsOf.Name">
            <ProductList v-if="showProductsOf" :childcategory_id="showProductsOf.ID"></ProductList>
        </v-offcanvas>
        <v-offcanvas v-model="showAttributesPanel" :title="'Attributes of '+ showAttributesof.Name">
            <AttributeList v-if="showAttributesof" :subcategory_id="showAttributesof.SubcategoryID"  :childcategory_id="showAttributesof.ID" ></AttributeList>
        </v-offcanvas>
    </div>
`
    }
