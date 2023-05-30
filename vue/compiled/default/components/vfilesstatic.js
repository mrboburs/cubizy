
    // https://mystudyfiles.s3.ap-southeast-2.amazonaws.com/empty.jpg
    export default {
        props: {
            value: {
                type: String,
                default: ""
            },
            prefix: {
                default: ""
            },
            multiple: {
                type: Boolean,
                default: false
            },
            only_image: {
                type: Boolean,
                default: true
            },
            priview: {
                type: Boolean,
                default: true
            },
        },
        data() {
            return {
                loading: 0,
                error: false,
                message: "",

                empty_image: "https://mystudyfiles.s3.ap-southeast-2.amazonaws.com/empty.jpg",
                accessURL: "/",
                files: [],
                upload_count: 0,
                uploded_files: [],
                filetypeimage: 'image/*',
                filetypeall: 'application/msword, application/vnd.ms-excel, application/vnd.ms-powerpoint, text/plain, application/pdf, image/*'
            };
        },
        mounted() {
            this.$emit('load', this)
            this.loadData();
        },
        methods: {
            isImageFile(file) {
                if (!file) {
                    return false
                }
                if(!this.priview){
                    return false
                }
                if (file.match(/.(jpg|jpeg|png|gif|svg)$/i)) {
                    return true
                } else {
                    return false
                }
            },
            getFileName(file) {
                var parts = file.split("/")
                return parts[parts.length - 1]
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
                        this.loadData()
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
                this.uploded_files.forEach(file => {
                    if (file.Key) {
                        this.upload_count++
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
                }).catch(error => {
                    console.log("Fail to upload file " + file.name + error)
                }).finally(() => {
                    this.loading -= 1;
                    if (!this.loading) {
                        this.loadData()
                    }
                });
            },
            loadData: function () {
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
            selecteImage: function (file) {
                if (file != "") {
                    file = this.accessURL + file.Key
                }
                if (this.multiple) {
                    this.$emit('input', [file])
                } else {
                    this.$emit('input', file)
                }
            },
        },
        template: `
        <divloading :loading="loading > 0">
            <v-alert v-model="message" :error="error" :hide_lable="true" :compact="true" />
            <div class="file-drop-control">
                <input multiple :accept="only_image?filetypeimage:filetypeall" :state="Boolean(uploded_files.length)"
                    type="file" placeholder="Choose a files or drop files here..." @change="uploadFiles">
                    <div class="d-flex flex-wrap mt-4">
                        <div v-for="file in files" :key="file.Key" class="m-1" style="z-index: 9;">
                            <div v-if="isImageFile(file.Key)">
                                <img class="img-thumbnail" fluid :src="accessURL + file.Key" :alt="file.Key"
                                    style="margin: auto; width: 100px;" @click.prevent="selecteImage(file)"
                                    :class="{ 'border border-primary' : value.includes(file.Key)}"></img>
                                <a :href="accessURL + file.Key" @click.prevent="deleteFile(file.Key)"><i
                                        class="fas fa-trash"></i></a>
                            </div>
                            <div v-else class="border rounded bg-faint p-1 file-div"
                                :class="{ 'border border-primary' : value.includes(file.Key)}">
                                <a :href="accessURL + file.Key" @click.prevent="selecteImage(file)">{{getFileName(file.Key)}}</a>
                                <a :href="accessURL + file.Key" target="_blank" style="width: fit-content; float: left; margin-right: 1em;">
                                    <i class="fas fa-eye"></i>
                                </a>
                                <a :href="accessURL + file.Key" @click.prevent="deleteFile(file.Key)" style="width: fit-content;">
                                    <i class="fas fa-trash"></i>
                                </a>
                            </div>
                        </div>
                    </div>
                </input>
            </div>
        </divloading>
`
    };
