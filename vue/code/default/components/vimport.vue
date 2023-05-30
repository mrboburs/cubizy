<script>
    /**
     * Page-header component
     */
    export default {
        components: {},
        props: {
            api: {
                type: String,
                default: ""
            },
            title: {
                type: String,
                default: ""
            },
            items: {
                type: Array,
                default: () => {
                    return [];
                }
            },
            columns: {
                type: Array,
                default: function () { return [{ key: "ID" }] }
            },
        },
        data() {
            return {
                file: false,
                filetext: "",
                json: false,
                conectedColumns: {},
                fixedValueColumns: {},
                importing: false,
                message: "",
                error: false,
                records_to_import: [],
                pageno: 0,
                pagesize: 1000,
            }
        },
        computed: {
            json_columns() {
                if (this.json.length) {
                    return Object.keys(this.json[0])
                }
                return []
            },
            selected_json_columns() {
                return Object.values(this.conectedColumns)
            },
            progress() {
                var number_of_pages = Math.ceil(this.records_to_import.length / this.pagesize)
                var progress = ((this.pageno + 1) / number_of_pages) * 100
                return progress
            }
        },
        methods: {
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
                columnTitle = columnTitle.replace(/([a-z])([A-Z])/g, '$1 $2');
                columnTitle = columnTitle.replace("_", ' ').trim()
                return this.camalize(columnTitle)
            },
            camalize(str) {
                return str
                    .replace(/\s(.)/g, function ($1) { return $1.toUpperCase(); })
                    .replace(/^(.)/, function ($1) { return $1.toLowerCase(); });

            },
            onfileselect(event) {
                if (event.target.files.length) {
                    this.file = event.target.files[0]
                    var fileExt = this.file.name.split('.').pop();
                    var component = this
                    component.conectedColumns = {}
                    var fr = new FileReader();
                    fr.onload = function () {
                        var filetext = fr.result;
                        switch (fileExt) {
                            case "csv":
                                component.json = component.csvJSON(filetext)
                                break;
                            case "tsv":
                                component.json = component.tsvJSON(filetext)
                                break;
                            case "json":
                                try {
                                    component.json = JSON.parse(filetext)
                                } catch (error) {
                                    console.log(error)
                                }
                                break;
                            default:
                                break;
                        }
                    }
                    fr.readAsText(this.file);
                }
            },
            //var csv is the CSV file with headers
            csvJSON(csv) {

                var lines = csv.split("\n");

                var result = [];

                var headers = lines[0].split(",");

                for (var i = 1; i < lines.length; i++) {

                    var obj = {};
                    var currentline = lines[i].split(",");

                    for (var j = 0; j < headers.length; j++) {
                        obj[headers[j]] = currentline[j];
                    }

                    result.push(obj);
                }
                //return result; //JavaScript object
                return result; //JSON
            },
            //var tsv is the TSV file with headers
            tsvJSON(tsv) {

                var lines = tsv.split("\n");

                var result = [];

                var headers = lines[0].split("\t");

                for (var i = 1; i < lines.length; i++) {

                    var obj = {};
                    var currentline = lines[i].split("\t");

                    for (var j = 0; j < headers.length; j++) {
                        obj[headers[j]] = currentline[j];
                    }
                    result.push(obj);
                }
                //return result; //JavaScript object
                return result; //JSON
            },
            trim(str, ch) {
                var start = 0,
                    end = str.length;

                while (start < end && str[start] === ch)
                    ++start;

                while (end > start && str[end - 1] === ch)
                    --end;

                return (start > 0 || end < str.length) ? str.substring(start, end) : str;
            },
            start_importing() {
                this.message = ""
                this.error = false

                var result = []
                var component = this
                component.json.forEach(element => {
                    var obj = {}
                    component.columns.forEach(column => {
                        if (component.conectedColumns[column.key] == "fixed") {
                            if (component.fixedValueColumns[column.key]) {
                                obj[column.key] = component.fixedValueColumns[column.key]
                            }
                        } else if (component.conectedColumns[column.key] != "none") {
                            if (element[component.conectedColumns[column.key]]) {
                                obj[column.key] = component.trim(element[component.conectedColumns[column.key]], '"')
                            }
                        }
                    });
                    result.push(obj)
                });
                this.records_to_import = result

                if (this.records_to_import.length > 0) {
                    if (this.records_to_import.length > this.pagesize) {
                        var batch = this.records_to_import.slice(this.pageno * this.pagesize, this.pagesize);
                        this.loadData({
                            items: batch
                        })
                    } else {
                        this.loadData({
                            items: this.records_to_import
                        })
                    }
                }
            },

            loadData(data) {
                if (!this.api) {
                    return
                }
                if (!data) {
                    return
                }
                var component = this
                this.importing = true;
                data.limit = -1
                this.$store.dispatch('call', {
                    api: this.api,
                    data: data,
                }).then((data) => {
                    component.message += data.Message
                    if (data.Status == 2) {
                        component.error = false
                    } else {
                        component.error = true
                    }
                }).catch((error) => {
                    component.error = true
                    component.message += error
                }).finally(function () {
                    if (component.progress < 100) {
                        component.pageno++
                        var batch = component.records_to_import.slice(component.pageno * component.pagesize, (component.pageno * component.pagesize) + component.pagesize);
                        if (batch.length > 0) {
                            component.loadData({
                                items: batch
                            })
                        }
                    } else {
                        component.importing = false;
                    }
                });
            },
        },
        mounted() {
        },
        template: `{{{template}}}`
    };
</script>

<template>
    <!-- start page title -->
    <div class="row">
        <div class="col-12" v-if="!json">
            <label for="file">File to import </label>
            <input id="inputfile" @change="onfileselect" type="file" class="form-control"
                placeholder="Select valid CSV or TSV or JSON File" :class="{ 'is-invalid': !file }" />
            <div v-if="!file" class="invalid-feedback">
                CSV or TSV or JSON File is required to start import.
            </div>
        </div>

        <div class="col-12" v-if="json">
            <div class="row">
                <div class="col">
                    Found {{ json.length }} number of records in file , metch column and start importing.
                </div>
            </div>
            <div class="row p-2" v-for="column in columns">
                <div class="col">
                    <label class="text-right">{{getColumnName(column)}}</label> :
                </div>
                <div class="col">
                    <select class="form-control" v-model="conectedColumns[column.key]">
                        <option value="">None</option>
                        <option value="fixed">Fixed</option>
                        <option v-for="key in json_columns" :value="key"
                            :disabled="selected_json_columns.includes(key)">{{key}} ({{json[0][key]}})</option>
                    </select>
                </div>
                <div class="col">
                    <input v-if="conectedColumns[column.key] == 'fixed'" class="form-control"
                        v-model="fixedValueColumns[column.key]" type="text" />
                </div>
            </div>
            <div v-if="importing" class="row">
                <div class="col">
                    <div class="progress">
                        <div class="progress-bar progress-bar-striped progress-bar-animated" role="progressbar"
                            :style="{ width: progress + '%' }"></div>
                    </div>
                </div>
            </div>
            <div class="row" v-if="!importing">
                <div class="col">
                    <button @click="start_importing" type="button" class="btn btn-success"> Start Importing</button>
                </div>
            </div>
        </div>
        <div class="col-12" v-if="message">
            <div class="alert d-flex align-items-center alert-dismissible fade show m-1"
                :class="{'alert-success': !error, 'alert-danger': error }" role="alert">
                <strong v-if="!error">Success : </strong>
                <strong v-if="error">Error : </strong>
                <span class="ms-1"> {{message}} </span>
                <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
            </div>
        </div>
    </div>
    <!-- end page title -->
</template>