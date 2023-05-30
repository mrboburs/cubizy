<script>
    export default {
        props: {
            category_id: {
                type: Number,
                default: 0
            },
            basecategory_id: {
                type: Number,
                default: 0
            },
        },
        components: {
            'SubcategoryEditor': () => import("/vue/account/subcategoryeditor.js"),
            'ChildcategoryList': () => import("/vue/account/childcategorylist.js"),
            'AttributeList': () => import("/vue/attributelist.js"),
        },
        data() {
            return {
                title: 'Subcategories',
                breadcrumb: [{
                    text: 'Subcategory',
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
                    { key: "CategoryID", title: "CategoryID", sortable: true, sortkey: "category_id", always_hide: false  },
                    { key: "Childcategories", title: "Child Categories", sortkey: "childcategories" },
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
                showChildcategoriesof : false,
                showAttributesof : false
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            showpanel: {
                // getter
                get: function () {
                    if (this.showChildcategoriesof) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.showChildcategoriesof = false
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
            showChildcategories(record) {
                if(record && record.ID > 0){
                    this.showChildcategoriesof = record
                }
            }
        },
        mounted() {
            if(this.category_id){
                this.new.CategoryID = this.category_id
                this.conditions = {
                    category_id : this.category_id
                }
                this.columns.forEach(column => {
                    if(column.key == "CategoryID"){
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
        template: `{{{template}}}`
    }
</script>

<template>
    <!-- Start Content-->
    <div class="row">
        <div class="col-12">
            <div class="card">
                <List api="subcategories" :columns="columns" title_column="Code" :can_select="true" :can_export="true" :can_add="category_id > 0"
                    :can_import="true" :conditions="conditions" @done="onActionDone" @onaction="onAction"  editor_size="lg" :item_actions="item_actions" >
                    <template v-slot:Icon="{ row, col }" class="icon_holder">
                        <div class="icon_holder" v-html="row.Icon"></div>
                    </template>
                    <template v-slot:Childcategories="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showChildcategories(row)">
                                {{row.Childcategories}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <SubcategoryEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit" :category_id="category_id"  :basecategory_id="basecategory_id">
                        </SubcategoryEditor>
                    </template>
                </List>
            </div>
        </div>
        <!-- Modal  -->
        <v-offcanvas v-model="showpanel" :title="'Childcategories of '+ showChildcategoriesof.Name">
            <ChildcategoryList v-if="showChildcategoriesof" :subcategory_id="showChildcategoriesof.ID" :basesubcategory_id="showChildcategoriesof.BaseSubcategoryID"></ChildcategoryList>
        </v-offcanvas>
        <v-offcanvas v-model="showAttributesPanel" :title="'Attributes of '+ showAttributesof.Name">
            <AttributeList v-if="showAttributesof" :subcategory_id="showAttributesof.ID" ></AttributeList>
        </v-offcanvas>
    </div>
</template>