
    export default {
        components: {
            'LocationEditor': () => import("/vue/locationeditor.js"),
        },
        data() {
            return {
                title: 'Locations',
                breadcrumb: [{
                    text: 'Location',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Code", sortable: true, sortkey: "code" },
                    { key: "SubLocality", sortable: true, sortkey: "sub_locality" },
                    { key: "Locality", sortable: true, sortkey: "locality" },
                    { key: "District", sortable: true, sortkey: "district" },
                    { key: "Country", sortable: true, sortkey: "country" },
                    { key: "Longitude", sortable: true, sortkey: "longitude" },
                    { key: "Latitude", sortable: true, sortkey: "latitude" },
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
                    Code: "",
                    SubLocality: "",
                    Locality: "",
                    District: "",
                    Country: "Mauritius",
                    Latitude: 0,
                    Longitude: 0,
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
        },
        template: `
    <div class="col-12">
        <div class="card">
            <List api="locations" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction">
                <template v-slot:editor="editing_item">
                    <LocationEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                    </LocationEditor>
                </template>
            </List>
        </div>
    </div>
`
    }
