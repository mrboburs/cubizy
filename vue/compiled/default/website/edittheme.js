
    export default {
        props: ["theme_id"],
        components: {
            'fileeditor': () => import("/vue/website/fileeditor.js"),
            'vuefileeditor': () => import("/vue/website/vuefileeditor.js"),
        },
        data() {
            return {
                loading: false,
                error: false,
                message: "",
                accessURL: "/",
                files: [],
                selected_file: null,
                newfilename: "",
                upload_count: 0,
                uploded_files: [],
                zipfile: false,
            };
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            prefix: function () {
                var id = 0
                if (this.theme_id) {
                    id = this.theme_id
                } else if (this.account) {
                    id = this.account.ThemeID
                }
                if (id) {
                    return 'themes/theme_' + id
                } else {
                    return ''
                }
            },
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
            user: function (newValue, oldValue) {
                if (!oldValue && (!this.files || !this.files.length)) {
                    this.loadData();
                }
            },
        },
        mounted() {
            if (!this.theme_id && (!this.account || !this.account.ThemeID)) {
                this.$router.push('/website/settings')
            }
            if (window.innerWidth >= 993) {
                setTimeout(() => {
                    document.body.setAttribute('data-sidebar-size', 'condensed');
                    Split([this.$refs.files, this.$refs.file], { sizes: [20, 80] })
                }, 30);
            }
            this.loadData();
        },
        methods: {
            isImageFile(file) {
                if (!file) {
                    return false
                }
                if (file.match(/.(jpg|jpeg|png|gif|svg)$/i)) {
                    return true
                } else {
                    return false
                }
            },
            isVueFile(file) {
                if (file.FileName.includes("vue.js")) {
                    return true
                }
                return false
            },
            getFileName(file) {
                if (!file.FileName) {
                    file.FileName = file.Key.split(this.prefix+ "/").pop()
                }
                return file.FileName
            },
            deleteFile: function (key) {
                this.loading += 1;
                this.message = "";
                this.error = false;
                this.$store.dispatch('call', {
                    api: "deletefile",
                    data: {
                        key: key,
                    },
                }).then((data) => {
                    this.message = data.Message
                    if (data.Status == 2) {
                        var index_to_delete = -1
                        for (let index = 0; index < this.files.length; index++) {
                            const element = this.files[index];
                            if (element.Key == key) {
                                index_to_delete = index
                                break
                            }
                        }
                        if (index_to_delete > -1) {
                            this.files.splice(index_to_delete, 1);
                        }
                        if (this.selected_file.Key == key) {
                            this.selected_file = null
                        }
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    this.error = true
                    this.message = error.Message
                    if (error) console.log(error)
                }).finally(() => {
                    this.loading -= 1;
                });
            },
            uploadFiles: function (event) {
                this.upload_count = 0
                this.uploded_files = Array.from(event.target.files)
                this.uploded_files.forEach(file => this.uploadFile(file));
            },
            uploadFile(file) {
                if (!file.name.trim()) {
                    return
                }
                if (file.Key) {
                    this.upload_count++
                    return
                }
                if(!this.prefix){
                    return
                }
                this.loading += 1;
                this.message = "";
                this.error = false;
                this.$store.dispatch('call', {
                    api: "getpresignedputurl",
                    data: {
                        prefix: this.prefix,
                        key: file.name,
                        type: file.type,
                    },
                }).then((data) => {
                    this.message = data.Message
                    if (data.Status == 2) {
                        file.Key = data.Result.key
                        this.putFile(file, data.Result.presignedUrl)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    this.error = true
                    this.message = error.Message
                    if (error) console.log(error)
                }).finally(() => {
                    this.loading -= 1;
                });
            },
            putFile(file, presignedUrl) {
                this.loading += 1
                fetch(presignedUrl, { // Your PUT endpoint
                    method: 'PUT',
                    headers: {
                        "x-amz-acl": "public-read"
                    },
                    body: file // This is your file object
                }).then(response => response.text()).then((data) => {
                    if (data == "") {
                        if (this.uploded_files.length == 1 && !this.multiple) {
                            //this.$emit('input', this.accessURL + file.Key)
                        } else {
                            this.upload_count++
                            if (this.upload_count == this.uploded_files.length) {
                                if (this.multiple) {
                                    var output_files = []
                                    this.uploded_files.forEach(uploded_file => {
                                        output_files.push(this.accessURL + uploded_file.Key)
                                    })
                                    if (output_files.length > 0) {
                                        //this.$emit('input', output_files)
                                    }
                                }
                            }
                        }
                    }
                    if (file.name == "index.html") {
                        this.updateTemplate()
                    }
                }).catch(error => {
                    console.log("Fail to upload file " + file.name + error)
                }).finally(() => {
                    this.loading -= 1;
                    if(!this.loading){
                        if(!this.files || !this.files.length){
                            this.loadData()
                        }
                    }
                });
            },
            updateTemplate: function () {
                this.message = "";
                this.error = false
                this.$store.dispatch('call', {
                    api: "updatetemplate",
                    data: {},
                }).then((data) => {
                    if (data.Status == 2) {
                        this.message = "";
                    } else {
                        this.message = data.Result.Error
                    }
                }).catch((error) => {
                    this.error = true;
                    this.message = error.Message
                    console.log(error)
                });
            },
            loadData: function () {
                if(!this.prefix){
                    return
                }
                this.loading += 1
                this.message = "";
                this.error = false
                this.$store.dispatch('call', {
                    api: "files",
                    data: { prefix: this.prefix },
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.files = data.Result.files
                        if (!this.files) {
                            this.files = []
                        }
                        this.accessURL = data.Result.accessURL
                    } else {
                        this.error = true;
                    }
                }).catch((error) => {
                    this.error = true;
                    this.message = error.Message
                    console.log(error)
                }).finally(() => {
                    this.loading -= 1;
                    this.$emit('oncount', this.files.length)
                });
            },
            addFile() {
                var file = {
                    Key: "",
                    FileName: "NewFile",
                    Data: "",
                }
                this.files.push(file)
                this.selected_file = file
            },
            updateFile(file) {
                if (!file.FileName) {
                    alert("No file Name, it is required to save file")
                }
                this.selected_file.FileName = file.FileName
                this.selected_file.Data = file.Data
                file = this.selected_file
                var type_of_file = "text/plain"
                if (file.FileName.includes(".css")) {
                    type_of_file = "text/css"
                }
                if (file.FileName.includes(".html")) {
                    type_of_file = "text/html"
                }
                if (file.FileName.includes(".js")) {
                    type_of_file = "text/javascript"
                }
                var _file = new File([file.Data], file.FileName, { type: type_of_file, })
                this.uploadFile(_file)
                if (file.Key && !file.Key.includes(file.FileName)) {
                    this.deleteOldFile(file)
                } else if (!file.Key) {
                    file.Key = this.prefix + "/" + file.FileName
                }
            },
            deleteOldFile(file) {
                this.deleteFile(file.Key)
                setTimeout(() => {
                    file.Key = file.Key.replace(file.Key.split("/").pop(), file.FileName);
                }, 200);
            },
            export_theme() {
                debugger
                if(!this.prefix){
                    return
                }
                this.loading += 1;
                this.$store.dispatch('call', {
                    api: "export_theme",
                    data: { prefix: this.prefix },
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.zipfile = data.Result.zipfile
                        setTimeout(() => {
                            this.$refs.file_link.click()
                            setTimeout(() => {
                               this.zipfile = false 
                            }, 100);
                        }, 100);
                    } else {
                        this.error = true;
                    }
                }).catch((error) => {
                    this.error = true;
                    this.message = error.Message
                    console.log(error)
                }).finally(() => {
                    this.loading = false
                    this.$emit('oncount', this.files.length)
                });
            }
        },
        template: `
    <div class="card" id="fullscreen_window">
        <v-alert v-model="message" :error="error" :hide_lable="true" />
        <divloading :loading="loading > 0" class="card-body h-100 w-100">
            <div class="split">
                <div ref="files" class="file_list">
                    <button class="btn btn-primary text-nowrap w-100" type="button" @click.prevent="addFile()">Add New
                        File</button>
                            <ul class="list-group">
                                <li class="list-group-item" v-for="file in files" :key="file.Key">
                                    <div class="d-flex justify-content-between">
                                        <a :href="accessURL + file.Key" class="flex-1"
                                            @click.prevent="selected_file = file">{{getFileName(file)}}</a>
                                        <a href="#" v-if="!file.Key.includes(file.FileName)"
                                            @click.prevent="deleteOldFile(file)" style="width: fit-content; color: red">
                                            <i class="fas fa-exclamation-triangle"></i>
                                        </a>
                                        <a href="#" @click.prevent="deleteFile(file.Key)" style="width: fit-content;">
                                            <i class="fas fa-trash"></i>
                                        </a>
                                    </div>
                                </li>
                            </ul>
                    <div class="file-drop-control">
                        <input multiple :state="Boolean(uploded_files.length)" type="file" placeholder="Choose a files or drop files here..." @change="uploadFiles"></input>
                    </div>
                    <button v-if="!zipfile"  class="btn btn-primary text-nowrap w-100" type="button" @click.prevent="export_theme()">
                        Export Theme
                    </button>
                    <a ref="file_link" v-if="zipfile" :href="zipfile" target="_blank" rel="noopener">Download theme</a>
                </div>
                <div ref="file" class="file_holder">
                    <div v-if="selected_file" class="h-100 w-100 d-flex flex-column">
                        <div v-if="isImageFile(selected_file.Key)">
                            <h4>{{getFileName(selected_file)}}</h4>
                            <img :src="accessURL + selected_file.Key"
                                style="width: auto; max-width: 100%; height: auto; max-height: 100%;" />
                        </div>
                        <vuefileeditor v-else-if="isVueFile(selected_file)" :value="selected_file"
                            :accessurl="accessURL" @input="updateFile" class="flex-1" />
                        <fileeditor v-else :value="selected_file" :accessurl="accessURL" @input="updateFile"
                            class="flex-1" />
                    </div>
                </div>
            </div>
        </divloading>
    </div>
`
    };
