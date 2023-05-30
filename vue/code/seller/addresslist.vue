<script>
export default {
  components: {
    AddressEditor: () => import("/vue/addresseditor.js"),
  },
  data() {
    return {
      title: "Addresses",
      breadcrumb: [
        {
          text: "Address",
        },
        {
          text: "All",
          active: true,
        },
      ],
      columns: [
        { key: "ID", sortable: true, sortkey: "id" },
        { key: "Title", sortable: true, sortkey: "title" },
        { key: "Mobile", sortable: true, sortkey: "mobile" },
        { key: "AddressLine1", title: "Address" },
        { key: "Code", sortable: true, sortkey: "code" },
        { key: "SubLocality", sortable: true, sortkey: "sub_locality" },
        { key: "Locality", sortable: true, sortkey: "locality" },
        { key: "District", sortable: true, sortkey: "district" },
        { key: "Country", sortable: true, sortkey: "country" },
        {
          key: "CreatedAt",
          type: "date",
          sortable: true,
          sortkey: "created_at",
        },
        {
          key: "UpdatedAt",
          type: "date",
          sortable: true,
          sortkey: "updated_at",
        },
        {
          key: "UpdatedByName",
          title: "Updatedby",
          sortable: true,
          sortkey: "updated_by_name",
        },
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
      conditions: {},
      new: {
        ID: 0,
        Title: "New Address",
        Mobile: "",
        AddressLine1: "",
        AddressLine2: "",
        AddressLine3: "",
        Longitude: "",
        Latitude: "",
        Code: "",
        SubLocality: false,
        Locality: false,
        District: false,
        Country: false,
      },
    };
  },
  computed: {
    ...Vuex.mapState(["user"]),
  },
  methods: {
    onAction(action, arg) {
      switch (action) {
        case "loading":
          this.loading = true;
          this.table = arg;
          break;
        case "add_new":
          this.new.Mobile = this.user.Mobile;
          this.new.Title = this.user.Name;
          this.table.editing_item = Object.assign({}, this.new);
          break;
        default:
          break;
      }
    },
    onActionDone(data) {
      this.loading = false;
    },
  },
  mounted() {},
  template: `{{{template}}}`,
};
</script>

<template>
  <!-- Start Content-->
  <div class="row">
    <div class="col-12">
      <div class="card">
        <List
          api="addresses"
          :columns="columns"
          editor_size="lg"
          title_column="Code"
          :can_select="true"
          :can_export="true"
          :can_import="true"
          :actions="actions"
          :conditions="conditions"
          @done="onActionDone"
          @onaction="onAction"
        >
          <template v-slot:editor="editing_item">
            <AddressEditor
              v-if="editing_item.item"
              :value="editing_item.item"
              @input="editing_item.submit"
            >
            </AddressEditor>
          </template>
        </List>
      </div>
    </div>
    <!-- Modal  -->
  </div>
</template>