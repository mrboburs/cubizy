
    export default {
        components: {
            'ordereditor': () => import("/vue/products/ordereditor.js"),
        },
        data() {
            return {
                title: 'Orders',
                breadcrumb: [{
                    text: 'Orders',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Logo", type: 'image', sortable: false },
                    { key: "Name", sortable: true, sortkey: "name" },
                    { key: "SKU" },
                    { key: "Cost" , type: 'money' },
                    { key: "Quantity" },
                    { key: "CostTotal" , type: 'money' },
                    { key: "ShippingMethod" },
                    { key: "ShippingPrice" , type: 'money' },
                    { key: "ShippingCost" , type: 'money' },
                    { key: "TrackingID"},
                    { key: "EDTMin" },
                    { key: "EDTMax" },
                    { key: "Status" },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "created_at" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "updated_at" },
                ],
                error: "",
                message: "",
                conditions:{},
            }
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        watch: {
        },
        methods: {
        },
        template: `
        <div class="col-12">
            <div class="card">
                <List api="orders" :columns="columns" title_column="Code" :can_select="true" :can_export="false" :can_edit="true" :can_import="false" :can_delete="false" :conditions="conditions" editor_size="lg"> <!--fullscreen-->
                    <template v-slot:editor="editing_item">
                        <ordereditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </ordereditor>
                    </template>
                </List>
            </div>
        </div>
`
    }
