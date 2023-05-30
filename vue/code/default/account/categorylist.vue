<script>
    export default {
        components: {
            'CategoryEditor': () => import("/vue/account/categoryeditor.js"),
            'SubcategoryList': () => import("/vue/account/subcategorylist.js"),
        },
        data() {
            return {
                title: 'Categories',
                breadcrumb: [
                    {
                        text: 'Category',
                    },
                    {
                        text: 'All',
                        active: true,
                    },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: "image" },
                    { key: "Icon" },
                    { key: "Name", sortable: true, sortkey: "name" },
                    { key: "Active", sortable: true, sortkey: "active", type: 'boolean' },
                    { key: "Subcategories", title: "Sub Categories", sortkey: "subcategories" },
                    { key: "Products", sortable: true, sortkey: "products" },
                    { key: "Revenue", sortable: true, sortkey: "revenue" },
                    { key: "Top", sortable: true, sortkey: "top" },
                    { key: "Featured", sortable: true, sortkey: "featured" },
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
                showsubcategoriesof: false
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            showpanel: {
                // getter
                get: function () {
                    if (this.showsubcategoriesof) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.showsubcategoriesof = false
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
                    default:
                        break;
                }
            },
            onActionDone(data) {
                this.loading = false
            },
            showsubcategories(record) {
                if (record && record.ID > 0) {
                    this.showsubcategoriesof = record
                }
            }
        },
        mounted() {
            if(this.account.AccountType != 'admin'){
                this.columns.forEach(column => {
                    if(column.key == "Icon"){
                        column.always_hide = true
                    }
                });
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="col-12">
        <div class="card">
            <List api="categories" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction" editor_size="lg" >
                <template v-slot:Icon="{ row, col }" class="icon_holder">
                    <div class="icon_holder" v-html="row.Icon"></div>
                </template>
                <template v-slot:Subcategories="{ row, col }">
                    <div class="form-check form-switch">
                        <button class="btn btn-primary" type="button" @click="showsubcategories(row)">
                            {{row.Subcategories}}
                        </button>
                    </div>
                </template>
                <template v-slot:editor="editing_item">
                    <CategoryEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                    </CategoryEditor>
                </template>
            </List>
        </div>
        <v-offcanvas v-model="showpanel" :title="'Subcategories of '+ showsubcategoriesof.Name">
            <SubcategoryList v-if="showsubcategoriesof" :category_id="showsubcategoriesof.ID" :basecategory_id="showsubcategoriesof.BaseCategoryID"></SubcategoryList>
        </v-offcanvas>
    </div>
</template>