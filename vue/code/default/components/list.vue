<script>
    export default {
        props: {
            title: {
                type: String,
                default: ""
            },
            title_column: {
                type: String,
                default: ""
            },
            columns: {
                type: Array,
                default: function () { return [{ key: "ID" }] }
            },
            conditions: {
                type: Object,
                default: function () { return {} }
            },
            api: {
                type: String,
                default: "accounts"
            },
            default_sort_by: {
                type: String,
                default: "id"
            },
            can_add: {
                type: Boolean,
                default: true
            },
            default_desc: {
                type: Boolean,
                default: false
            },
            can_edit: {
                type: Boolean,
                default: true
            },
            can_delete: {
                type: Boolean,
                default: true
            },
            can_bulk_delete: {
                type: Boolean,
                default: true
            },
            can_select: {
                type: Boolean,
                default: true
            },
            can_export: {
                type: Boolean,
                default: false
            },
            can_import: {
                type: Boolean,
                default: false
            },
            actions: {
                type: Array,
                default: function () { return [] }
            },
            item_actions: {
                type: Array,
                default: function () { return [] }
            },
            bulk_actions: {
                type: Array,
                default: function () { return [] }
            },
            editor_size:{
                default: "",
            },
            editor_type:{
                type: String,
                default: "modal",
            },
            default_per_page:{
                type: Number,
                default : 10
            }
        },
        data() {
            return {
                loading: false,
                error: false,
                message: "",
                updating: false,
                selected: {},
                records: [],
                recordsTotal: 0,
                recordsFiltered: 0,
                currentPage: 1,
                perPage: 10,
                hidden: [],
                pageOptions: [2, 10, 25, 50, 100, 1000],
                search: "",
                searched: "",
                desc: false,
                sort_by: "id",
                editing_item: false,
                show_export: false,
                show_import: false,
                exportPages: 0,
                exportPageSize: 1000,
                exportType: "json",
                exporting: false,
                items_to_delete: [],
                loadData: _.debounce(function (data, component) {
                        if (!data) {
                            data = {}
                        }
                        component.loading = true;
                        component.error = false
                        component.message = ""
                        component.$emit('onaction', "loading", component)
                        data.sort = component.sort_by
                        data.sortdesc = component.desc
                        data.search = component.search.trim().replace(": ", '":"')
                        data.limit = component.perPage
                        data.page = component.currentPage - 1
                        data.fix_condition = component.conditions
                        return component.$store.dispatch('call', {
                            api: component.api,
                            data: data,
                        }).then((data) => {
                            component.message = data.Message
                            if (data.Status == 2) {
                                component.records[component.currentPage] = data.data
                                component.recordsTotal = data.recordsTotal
                                component.recordsFiltered = data.recordsFiltered
                                component.setExportPagesCount()
                                component.searched = component.search.trim()
                            } else {
                                component.error = true
                            }
                            component.$emit('done', data)
                            return data
                        }).catch((error) => {
                            console.log(error)
                            component.$emit('done', false)
                        }).finally(() => {
                            component.loading = false;
                        });
                }, 500),
            };
        },
        mounted() {
            this.sort_by = this.default_sort_by
            this.desc = this.default_desc
            this.perPage = this.default_per_page
            this.loadData(null, this);
            var hidden_string = localStorage.getItem("columns_" + this.api)
            if (hidden_string != null || hidden_string != undefined) {
                this.hidden = hidden_string.split(",")
            }
        },
        watch: {
            show_export(newValue, oldValue) {
                if (!newValue && this.exporting) {
                    this.show_export = true
                }
            },
            show_import(newValue, oldValue) {
                if (!newValue) {
                    this.loadData(null, this)
                }
            },
            exportPageSize(newValue, oldValue) {
                this.setExportPagesCount()
            },
            currentPage(newValue, oldValue) {
                if (this.records[this.currentPage] == undefined) {
                    this.records[this.currentPage] = []
                }
                if (this.records[this.currentPage].length == 0) {
                    this.loadData(null, this)
                }
            },
            default_per_page(newValue, oldValue) {
                this.perPage = newValue
            },
            perPage(newValue, oldValue) {
                if (newValue != oldValue) {
                    this.currentPage = 1
                    this.records = {}
                }
                this.loadData(null, this)
            },
            default_desc(newValue, oldValue) {
                this.desc = newValue
            },
            default_sort_by(newValue, oldValue) {
                this.sort_by = newValue
            },
            conditions(newValue, oldValue) {
                this.loadData(null, this)
            },
        },
        computed: {
            ...Vuex.mapState(["user"]),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
            maxPage: function () {
                return Math.ceil(this.recordsFiltered / this.perPage)
            },
            hidablecolumns: function () {
                if (this.columns){
                    return this.columns.filter(column => !column.always_hide);
                }else{
                    return []
                }
            },
            visiblecolumns: function () {
                var component = this
                if (this.columns){
                    return this.columns.filter(column => (!component.hidden.includes(column.key) && !column.always_hide));
                }else{
                    return []
                }
            },
            ask_to_delete: {
                // getter
                get: function () {
                    return this.items_to_delete.length > 0
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.items_to_delete = []
                    }
                }
            },
            editing: {
                // getter
                get: function () {
                    return this.editing_item != false
                },
                // setter
                set: function (newValue) {
                    this.editing_item = false
                }
            }
        },
        methods: {
            onAction(action) {
                this.message = ""
                this.error = false
                this.$emit('onaction', action)
            },
            onItemAction(action, record) {
                this.message = ""
                this.error = false
                this.$emit('onaction', action, record)
            },
            onBulkAction(action) {
                this.message = ""
                this.error = false
                this.$emit('onaction', action, this.selected)
            },
            onEdit(record) {
                this.message = ""
                this.error = false
                if(!!this.$scopedSlots.editor){
                    this.editing_item = record
                }else{
                    this.$emit('onaction', "edit", record)
                }
            },
            delete_item(record) {
                this.message = ""
                this.error = false
                this.items_to_delete.push(record)
            },
            deleteSelected() {
                this.message = ""
                this.error = false
                for (const key in this.selected) {
                    if (Object.hasOwnProperty.call(this.selected, key)) {
                        this.items_to_delete.push(this.selected[key]);
                    }
                }
            },
            onDelete() {
                var todelete = []
                this.items_to_delete.forEach(element => {
                    if (element.ID) {
                        todelete.push(element.ID)
                    }
                });
                var data = {
                    todelete: todelete
                }
                this.loadData(data, this).then((data) => {
                    if (data.Status == 2) {
                        if (this.items_to_delete.length < 1) {
                            this.selected = {}
                        }
                        this.items_to_delete = []

                    }
                });
            },
            select(record) {
                if (!this.selected[record.ID]) {
                    this.selected[record.ID] = record
                } else {
                    delete this.selected[record.ID]
                }
                this.$forceUpdate();
            },
            selectall() {
                var flag = 0
                this.records[this.currentPage].forEach(record => {
                    if (!flag) {
                        if (!this.selected[record.ID]) {
                            this.selected[record.ID] = record
                            flag = 1
                        } else {
                            delete this.selected[record.ID]
                            flag = 2
                        }
                    } else if (flag == 1) {
                        this.selected[record.ID] = record
                    } else if (flag == 2) {
                        if (this.selected[record.ID]) {
                            delete this.selected[record.ID]
                        }
                    }
                });
                this.$forceUpdate();
            },
            onSort(column) {
                if (column.sortable) {
                    if (!column.sortkey) {
                        column.sortkey = column.key
                    }
                    this.sort_by = column.sortkey
                    this.desc = !this.desc
                    this.currentPage = 1
                    this.records = {}
                    this.loadData(null, this)
                }
            },
            clearSearch() {
                this.search = ""
                this.onSearch()
            },
            onSearch() {
                if (this.search.trim() != this.searched.trim()) {
                    this.currentPage = 1
                    this.records = {}
                    this.selected = {}
                }
                this.loadData(null, this)
            },
            getColumnName(column) {
                var columnTitle = ""
                if (column.title) {
                    columnTitle = column.title
                } else {
                    columnTitle = column.key
                }
                return this.getTitle(columnTitle)
            },
            getTitle(columnTitle) {
                if(columnTitle){
                    columnTitle = columnTitle.replace(/([a-z])([A-Z])/g, '$1 $2');
                    columnTitle = columnTitle.replace("_", ' ').trim()
                    return this.camalize(columnTitle)
                }else{
                    debugger
                }
            },
            camalize(str) {
                return str
                    .replace(/\s(.)/g, function ($1) { return $1.toUpperCase(); })
                    .replace(/^(.)/, function ($1) { return $1.toLowerCase(); });

            },
            getLastIndex(index) {
                if (index < this.recordsFiltered) {
                    return index
                } else {
                    return this.recordsFiltered
                }
            },
            toggleVisibility(column) {
                if (this.hidden.includes(column.key)) {
                    const index = this.hidden.indexOf(column.key);
                    if (index > -1) {
                        this.hidden.splice(index, 1);
                    }
                } else {
                    this.hidden.push(column.key)
                }
                localStorage.setItem("columns_" + this.api, this.hidden.toString())
                this.$forceUpdate();
            },
            setExportPagesCount() {
                if (this.recordsFiltered == 0) {
                    this.exportPages = 0
                } else {
                    this.exportPages = Math.ceil(this.recordsFiltered / this.exportPageSize)
                }
            },
            start_exporting(index) {
                this.exporting = true;
                var data = {}
                data.sort = this.sort_by
                data.sortdesc = this.desc
                data.search = this.search.trim().replace(": ", '":"')
                data.limit = this.exportPageSize
                data.page = index - 1
                data.fix_condition = this.conditions
                return this.$store.dispatch('call', {
                    api: this.api,
                    data: data,
                }).then((data) => {
                    var msg_title = "Message"
                    var msg_class = "primary"
                    if (data.Status == 2) {
                        msg_title = "Successful"
                        msg_class = "success"
                        this.export_records(index, data.data)
                    } else {
                        msg_title = "Failed"
                        msg_class = "danger"
                    }
                    if (data.Message && !this.editing) {
                        this.$bvModal.msgBoxOk(data.Message, {
                            title: msg_title,
                            size: 'sm',
                            buttonSize: 'sm',
                            okVariant: msg_class,
                            centered: true
                        })
                    }
                    return data
                }).catch((error) => {
                    console.log(error)
                }).finally(() => {
                    this.exporting = false;
                });
            },
            export_records(index, content) {
                var contentType = 'text/plain'
                if (this.exportType == "csv") {
                    contentType = 'text/csv'
                    var json = content
                    var fields = Object.keys(json[0])
                    var replacer = function (key, value) { return value === null ? '' : value }
                    var csv = json.map(function (row) {
                        return fields.map(function (fieldName) {
                            return JSON.stringify(row[fieldName], replacer)
                        }).join(',')
                    })
                    csv.unshift(fields.join(',')) // add header column
                    csv = csv.join('\r\n');
                    content = csv
                } else {
                    contentType = 'text/json'
                    content = JSON.stringify(content)
                }
                this.download(content, this.api + "_" + index + "." + this.exportType, contentType)
            },
            download(content, filename, contentType) {
                if (!contentType) contentType = 'application/octet-stream';
                var a = document.createElement('a');
                var blob = new Blob([content], { 'type': contentType });
                a.href = window.URL.createObjectURL(blob);
                a.download = filename;
                a.click();

            },
            reload(){
                this.loadData(null, this)
            },
            submit(record) {
                if (!record) {
                    this.editing_item = false
                    return
                }
                this.loadData({
                    items: [record]
                }, this).then((data) => {
                    if (data.Status == 2) {
                        this.editing_item = false
                    }
                })
            },
        },
        template: `{{{template}}}`
    };
</script>

<template>
    <div>
        <!-- start table -->
        <divloading :fullpage="false" :loading="loading" class="container-fluid">
            <div class="row justify-content-between">
                <!-- Search -->
                <label class="col d-flex align-items-center justify-content-md-start">
                    <input type="search" class="form-control form-control-sm m-1 flex-grow-1 w-auto" id="searchInput1"
                        placeholder="Search..." v-model="search">
                    <button type="button" class="btn btn-sm btn-success m-1" @click="onSearch">
                        <i class="mdi mdi-magnify"></i>
                    </button>
                    <button type="button" class="btn btn-sm btn-success m-1" @click="reload"><i
                            class="mdi mdi-reload"></i></button>
                    <span v-if="searched" class="m-1 text-nowrap flex-grow-1"> Searching "{{searched}}"</span>
                    <button v-if="searched" type="button" class="btn btn-sm btn-outline-danger" :disabled="updating"
                        @click.prevent="clearSearch()">
                        <i class="fas fa-times"></i>
                    </button>
                </label>
                <!-- End search -->
                <div v-if="message" class="col">
                    <div class="alert d-flex align-items-center alert-dismissible fade show m-1"
                        :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
                        <strong v-if="!error">Success : </strong>
                        <strong v-if="error">Error : </strong>
                        <span class="ms-1"> {{message}} </span>
                        <button type="button" class="btn-close" @click.prevent="message = false"
                            aria-label="Close"></button>
                    </div>
                </div>
                <div
                    class="col-12 col-sm-auto d-flex flex-wrap align-items-center justify-content-evenly justify-content-md-end">

                    <div class="btn-group" id="dropdown-1">
                        <button type="button" class="btn btn-sm btn-primary dropdown-toggle " data-bs-toggle="dropdown"
                            aria-expanded="false">
                            <i class="mdi mdi-table-column"></i>
                            <span class="list-btn-label">Columns</span>
                        </button>
                        <ul class="dropdown-menu">
                            <li v-for="(column, index) in hidablecolumns" :key="column.key+'_'+index">
                                <div class="form-check form-switch text-capitalize">
                                    <input class="form-check-input" type="checkbox"
                                        :id="'flexSwitchCheckChecked'+column.key" :checked="!hidden.includes(column.key)"
                                        @input="toggleVisibility(column)">
                                    <label class="form-check-label"
                                        :for="'flexSwitchCheckChecked'+column.key">{{getColumnName(column)}}</label>
                                </div>
                            </li>
                        </ul>
                    </div>
                    <button v-if="can_add" type="button" class="btn text-white btn-sm btn-success m-1"
                        href="javascript: void(0);" @click="onAction('add_new')">
                        <i class="mdi mdi-plus-circle mr-1"></i>
                        <span class="list-btn-label">Add New</span>
                    </button>
                    <button v-if="can_export" type="button" class="btn btn-sm  btn-outline-success m-1"
                        :disabled="updating" @click.prevent="show_export = true">
                        <i class="ri-download-line"></i>
                        <span class="list-btn-label">Export</span>
                    </button>
                    <button v-if="can_import" type="button" class="btn btn-sm  btn-outline-success m-1"
                        :disabled="updating" @click.prevent="show_import = true">
                        <i class="ri-upload-line"></i>
                        <span class="list-btn-label">Import</span>
                    </button>
                    <button v-for="action in actions" :key="action.key" type="button"
                        class="btn btn-sm  btn-outline-success m-1" :disabled="updating"
                        @click.prevent="onAction( action.key )">
                        <i v-if="action.icon" :class="action.icon"></i>
                        <span v-if="action.text" class="text-capitalize">{{getTitle( action.text )}}</span>
                    </button>
                    <button
                        v-if="Object.keys(selected).length > 0 && recordsFiltered > 0 && can_delete && can_bulk_delete"
                        type="button" class="btn text-white btn-sm btn-danger m-1" href="javascript: void(0);"
                        @click="deleteSelected">
                        <i class="fas fa-trash-alt mr-1"></i>
                        <span class="list-btn-label">Delete</span>
                    </button>
                    <template v-if="Object.keys(selected).length > 0" >
                        <button v-for="action in bulk_actions" :key="action.key"
                            type="button" class="btn btn-sm  btn-outline-success m-1" :disabled="updating"
                            @click.prevent="onBulkAction(action.key )">
                            <i v-if="action.icon" :class="action.icon"></i>
                            <span class="list-btn-label" v-if="action.text">{{ action.text}}</span>
                        </button>
                    </template>
                </div>
            </div>
            <div class="table-responsive">
                <table class="table table-sm table-hover">
                    <thead>
                        <tr class="text-capitalize">
                            <th v-if="can_select" style="max-width: 30px;" @click="selectall">
                                <i class="mdi mdi-select-all"></i>
                            </th>
                            <th v-for="(column, index) in visiblecolumns" :key="column.key+ '_'+ index" @click="onSort(column)" >
                                <div class="d-flex " style="cursor: pointer;">
                                    <span class="me-1">{{getColumnName(column)}}</span>
                                    <div v-if="column.sortable">
                                        <i v-if="sort_by != column.sortkey"
                                            class="fa fa-arrows-alt-v text-black-50"></i>
                                        <i v-if="sort_by == column.sortkey && desc"
                                            class="fa fa-arrow-down text-primary"></i>
                                        <i v-if="sort_by == column.sortkey && !desc"
                                            class="fa fa-arrow-up text-primary"></i>
                                    </div>
                                </div>
                            </th>
                            <th v-if="can_edit || can_delete || item_actions.length ">
                                Options
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="record in records[currentPage]" :key="record.ID"
                            :class="selected[record.ID] ? 'table-primary' : ''">
                            <td v-if="can_select" style="max-width: 30px;" @click="select(record)">
                                <i v-if="selected[record.ID]" class="far fa-check-square"></i>
                                <i v-if="!selected[record.ID]" class="far fa-square"></i>
                            </td>
                            <td v-for="(column, index) in visiblecolumns" :key="column.key + '_' + index + '_' + record.ID " >

                                <slot :name="column.key" v-bind:row="record" v-bind:col="column">
                                    <!-- Fallback content -->
                                    <template v-if="column.type == 'image'">
                                        <img :src="record[column.key]" :alt="record[title_column]"
                                            class="avatar-md rounded" />
                                    </template>
                                    <template v-else-if="column.type == 'date'">
                                        {{getFullDate(record[column.key])}}
                                    </template>
                                    <template v-else-if="column.type == 'date_time'">
                                        {{getFullDateTime(record[column.key])}}
                                    </template>
                                    <template v-else-if="column.type == 'money'">
                                        <span class="text-nowrap"><span>$</span>{{record[column.key]}}<span>.00</span></span>
                                    </template>
                                    <template v-else-if="column.type == 'discount'">
                                        {{record[column.key]}} %
                                    </template>
                                    <template v-else-if="column.type == 'boolean'">
                                        <span v-if='record[column.key]' class="badge bg-success">Yes</span>
                                        <span v-if='!record[column.key]' class="badge bg-danger">No</span>
                                    </template>
                                    <template v-else>
                                        <span class="text-capitalize"> {{record[column.key]}}</span>
                                    </template>

                                </slot>
                            </td>
                            <td v-if="can_edit || can_delete || item_actions.length ">
                                <div class="d-flex">
                                    <button v-for="action in item_actions" :key="action.key" type="button"
                                        class="btn btn-sm btn-outline-primary border-0 me-1" :disabled="updating"
                                        @click.prevent="onItemAction(action.key, record)">
                                        <i v-if="action.icon" :class="action.icon"></i>
                                        <span v-if="action.text">{{ action.text}}</span>
                                    </button>
                                    <button v-if="can_edit" type="button"
                                        class="btn btn-sm btn-outline-success border-0 me-1" :disabled="updating"
                                        @click.prevent="onEdit(record)">
                                        <i class="fas fa-edit"></i>
                                    </button>
                                    <button v-if="can_delete" type="button"
                                        class="btn btn-sm btn-outline-danger border-0 me-1" :disabled="updating"
                                        @click.prevent="delete_item(record)">
                                        <i class="fas fa-trash-alt"></i>
                                    </button>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="row m-1 text-nowrap">
                <div class="col-sm-12 col-md-4">
                    <label class="d-flex align-items-center"> Display
                        <select class="form-select form-select-sm" aria-label="Default select example" v-model="perPage"
                            style="width: auto;">
                            <option v-for="pageOption in pageOptions" :key="pageOption"
                                :selected="pageOption == perPage">{{pageOption}}</option>
                        </select>
                        Items
                    </label>
                </div>
                <div class="col-sm-12 col-md-4 text-center">
                    <label>
                        Showing from {{ (perPage * (currentPage - 1)) + 1 }}
                        to {{ getLastIndex(perPage * currentPage) }}
                        out of {{ recordsFiltered }}
                        <span v-if="recordsFiltered != recordsTotal">
                            ( Total : {{recordsTotal}} )
                        </span>
                        <span v-if="Object.keys(selected).length">
                            (Selected : {{Object.keys(selected).length}})
                        </span>
                    </label>
                </div>
                <div class="col-sm-12 col-md-4">
                    <v-pagination v-model="currentPage" :totalrows="recordsFiltered" :perpage="parseInt(perPage)"
                        class="">
                    </v-pagination>
                </div>
            </div>
        </divloading>
        <!-- end table -->
        <!-- Modal  -->
        <v-modal id="modal-2" v-model="show_export" title="Export Records in CSV/TSV/JSON files"
            header-close-variant="light" title-class="font-18" hide-footer>
            <divloading :fullpage="false" :loading="exporting" class="modal-body">
                <div class="row">
                    <div class="col">
                        <label>Page Size :</label>
                    </div>
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="radio" name="rdoexportPageSize" id="rdoexportPageSize"
                                :value="1000" v-model="exportPageSize">
                            <label class="form-check-label" for="rdoexportPageSize">
                                1000
                            </label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="radio" name="rdoexportPageSize" id="rdoexportPageSize"
                                :value="5000" v-model="exportPageSize">
                            <label class="form-check-label" for="rdoexportPageSize">
                                5000
                            </label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="radio" name="rdoexportPageSize" id="rdoexportPageSize"
                                :value="10000" v-model="exportPageSize">
                            <label class="form-check-label" for="rdoexportPageSize">
                                10000
                            </label>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <label>Export Type :</label>
                    </div>
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="radio" name="rdoexportType" id="rdoexportType"
                                value="json" v-model="exportType">
                            <label class="form-check-label" for="rdoexportType">
                                JSON
                            </label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="radio" name="rdoexportType" id="rdoexportType"
                                value="csv" v-model="exportType">
                            <label class="form-check-label" for="rdoexportType">
                                CSV
                            </label>
                        </div>
                    </div>
                    <div class="col"></div>
                </div>
                <div class="row" v-if="exportPages > 0">
                    <div class="col" v-for="index in exportPages " :key="index">
                        <button type="button" class="btn btn-sm  btn-outline-success m-1"
                            @click="start_exporting(index)">
                            Export File {{index}} Of {{ exportPageSize }} records</button>
                    </div>
                </div>
            </divloading>
        </v-modal>

        <v-modal id="modal-2" v-model="show_import" title="Import Records from CSV, TSV, JSON"
            header-close-variant="light" title-class="font-18" hide-footer>
            <v-import :columns="columns" :api="api" />
        </v-modal>

        <v-modal id="modal-2" v-model="ask_to_delete" title="Delete Records" header-close-variant="light"
            title-class="font-18" :footer="true" :header="false">
            <divloading :fullpage="false" :loading="loading" class="modal-body">
                <label v-if="items_to_delete.length > 0" class="fs-3 text-secondary">
                    Do you realy want to delete
                    <span v-if="items_to_delete.length == 1 && title_column">
                        {{items_to_delete[0][title_column]}}
                    </span>
                    <span v-else>{{Object.keys(selected).length}} selected items</span>
                    ?
                </label>
                <h3 v-else>Please select items to delete</h3>

                <div v-if="message" class="alert d-flex align-items-center alert-dismissible fade show"
                    :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
                    <strong v-if="!error">Success : </strong>
                    <strong v-if="error">Error : </strong>
                    <span class="ms-1"> {{message}} </span>
                    <button type="button" class="btn-close" @click.prevent="message = false"
                        aria-label="Close"></button>
                </div>
            </divloading>
            <template v-slot:footer>
                <button class="btn btn-danger" @click="onDelete" :disabled="loading">Yes</button>
                <button class="btn btn-warning ml-1" @click="ask_to_delete = false" :disabled="loading">No</button>
            </template>
        </v-modal>

        <v-modal v-if="can_edit && editor_type == 'modal'" v-model="editing" :size="editor_size"
            :title="editing_item.ID? 'Edit Record': 'Add Record'" header-close-variant="light" title-class="font-18">
            <div v-if="message" class="alert d-flex align-items-center alert-dismissible fade show"
                :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
                <strong v-if="!error">Success : </strong>
                <strong v-if="error">Error : </strong>
                <span class="ms-1"> {{message}} </span>
                <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
            </div>
            <slot name="editor" v-bind:item="editing_item" v-bind:submit="submit">
                Please contact support if no ui for editing this record
            </slot>
        </v-modal>

        <v-offcanvas v-if="can_edit && editor_type == 'offcanvas'" v-model="editing" :placement="editor_size" :title="editing_item.ID? 'Edit Record': 'Add Record'" header-close-variant="light" title-class="font-18">
            <div v-if="message" class="alert d-flex align-items-center alert-dismissible fade show"
                :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
                <strong v-if="!error">Success : </strong>
                <strong v-if="error">Error : </strong>
                <span class="ms-1"> {{message}} </span>
                <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
            </div>
            <slot name="editor" v-bind:item="editing_item" v-bind:submit="submit">
                Please contact support if no ui for editing this record
            </slot>
        </v-offcanvas>
    </div>
</template>