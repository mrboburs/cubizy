<script>
    export default {
        props: {
            subcategory_id: {
                type: Number,
                default: 0
            },
            childcategory_id: {
                type: Number,
                default: 0
            },
        },
        components: {
            'AttributeEditor': () => import("/vue/attributeeditor.js"),
        },
        data() {
            return {
                title: 'Attributes',
                breadcrumb: [{
                    text: 'Attribute',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Name", sortable: false },
                    { key: "FieldType", sortable: true, sortkey: "field_type" },
                    { key: "ProductColumn", sortable: false },
                    { key: "Options", title: "Default/Options", sortable: false  },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                    { key: "UpdatedByName", title: "Updatedby", sortable: true, sortkey: "updated_by_name" },
                ],
                table: false,
                error: "",
                message: "",
                conditions:{
                    subcategory_id : this.subcategory_id,
                    childcategory_id : this.childcategory_id,
                },
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
                showsubattributesof: false
            }
        },
        watch: {
            subcategory_id: function (newValue, oldValue) {
                this.init()
            },
            childcategory_id: function (newValue, oldValue) {
                this.init()
            },
        },
        mounted() {
            this.init()
        },
        methods: {
            init: _.debounce(function (data) {
                this.new.SubcategoryID = this.subcategory_id
                this.new.ChildcategoryID = this.childcategory_id
                this.conditions = {
                    subcategory_id : this.subcategory_id,
                    childcategory_id : this.childcategory_id
                }
            }, 200),
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
            showsubattributes(record) {
                if(record && record.ID > 0){
                    this.showsubattributesof = record
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
        <div class="col-12">
            <div class="card">
                <List api="attributes" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                    :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction" :conditions="conditions" > <!--fullscreen-->
                    <template v-slot:SubAttributeCount="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showsubattributes(row)">
                                {{row.SubAttributeCount}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <AttributeEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </AttributeEditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
</template>