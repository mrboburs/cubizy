
    export default {
        props: {
            value: {
                type: Object,
                default: null
            },
            accessurl: {
                type: String,
                default: "",
            }
        },
        data: () => {
            return {
                loading: false,
                error: false,
                message: "",
                filename: "",
                editor: null,
                editor_options : {
                    enableBasicAutocompletion: true,
                    enableSnippets: true,
                    enableLiveAutocompletion: true,
                    fontSize: 16,
                },
                selectedtheme: "chrome",
                thems: ["twilight", "ambiance", "chrome"],
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                this.SetData()
            },
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        mounted: function () {
            ace.require("ace/ext/language_tools");
            this.editor = ace.edit("editor");
            var editor_options = localStorage.getItem('editor_options');
            try {
                if(editor_options) editor_options = JSON.parse(editor_options)
            } catch (error) {
                console.log(error)
                editor_options = false
            }
            if(!editor_options){
                editor_options = this.editor_options
            }
            this.editor.setOptions(editor_options)
            this.beautify = ace.require("ace/ext/beautify");
            var StatusBar = ace.require("ace/ext/statusbar").StatusBar;
            // create a simple selection status indicator
            var statusBar = new StatusBar(this.editor, this.$refs.statusBar);
            setInterval(() => {
                localStorage.setItem('editor_options',JSON.stringify(this.editor.getOptions()))
            }, 30000);
            this.SetData()
            window.addEventListener("keydown", (e) => {
                if((e.metaKey || e.ctrlKey)  && e.keyCode == 83){
                    e.preventDefault()
                    this.submit()
                }
            })
        },
        methods: {
            SetData() {
                if (this.value && this.accessurl) {
                    this.filename = this.value.FileName
                    if (this.value.Data == undefined) {
                        this.loadData()
                    } else {
                        this.setSession(this.value.Data)
                    }
                } else {
                    file_session = ace.createEditSession("");
                    this.editor.setSession(file_session);
                }
            },
            loadData() {
                this.loading = true
                var key = this.value.Key
                fetch(this.accessurl + this.value.Key + "?v="+Date.now()).then(function (response) {
                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    return response.text();
                }).then(response => {
                    if (key == this.value.Key) {
                        if (response) {
                            this.value.Data = response
                            this.setSession(response)
                        }
                    }
                }).catch((error) => {
                    this.error = true
                    this.message = error.message
                }).finally(() => {
                    this.loading = false
                });
            },
            setSession(data) {
                var file_session
                file_session = ace.createEditSession(data);
                var modelist = ace.require("ace/ext/modelist")
                var mode = modelist.getModeForPath(this.filename).mode
                file_session.setMode(mode)
                this.editor.setSession(file_session);
            },
            submit() {
                var data = this.editor.getValue()
                this.$emit('input', {
                    FileName: this.filename,
                    Data: data,
                })
            },
            doIndent(){
                if(this.beautify && this.editor)
                this.beautify.beautify(this.editor.session);
            },
            togggleSettings() {
                this.editor.execCommand('showSettingsMenu');
            },
            togggleSearch() {
                this.editor.execCommand('find');
            },
            togggleFullScreen() {
                // if already full screen; exit
                // else go fullscreen
                if (
                    document.fullscreenElement ||
                    document.webkitFullscreenElement ||
                    document.mozFullScreenElement ||
                    document.msFullscreenElement
                ) {
                    if (document.exitFullscreen) {
                        document.exitFullscreen();
                    } else if (document.mozCancelFullScreen) {
                        document.mozCancelFullScreen();
                    } else if (document.webkitExitFullscreen) {
                        document.webkitExitFullscreen();
                    } else if (document.msExitFullscreen) {
                        document.msExitFullscreen();
                    }
                } else {
                    //var element = this.$refs.editor_holder;
                    var element = document.getElementById("fullscreen_window")
                    if (element.requestFullscreen) {
                        element.requestFullscreen();
                    } else if (element.mozRequestFullScreen) {
                        element.mozRequestFullScreen();
                    } else if (element.webkitRequestFullscreen) {
                        element.webkitRequestFullscreen(Element.ALLOW_KEYBOARD_INPUT);
                    } else if (element.msRequestFullscreen) {
                        element.msRequestFullscreen();
                    }
                }
            }
        },
        template: `
    <div class="d-flex flex-column" ref="editor_holder">
        <v-alert v-model="message" :error="error" :hide_lable="true"/>
        <div class="input-group">
            <span class="input-group-text" id="basic-addon1">File Name : </span>
            <input type="text" class="form-control" placeholder="New file name" aria-label="New file name"
                v-model="filename" style="min-width: 182px;">
            <button class="btn btn-primary" @click.prevent="submit">Save</button>
            <button class="btn btn-primary" type="button" @click.prevent="togggleSearch">
                <i class="fe-search"></i>
            </button>
            <button class="btn btn-primary" type="button" @click.prevent="doIndent">
                <i class="fas fa-indent"></i>
            </button>
            <button class="btn btn-primary" type="button" @click.prevent="togggleFullScreen">
                <i class="fe-maximize noti-icon"></i>
            </button>
            <button class="btn btn-primary" type="button" @click.prevent="togggleSettings">
                <i class="fe-settings noti-icon"></i>
            </button>
        </div>
        <divloading :loading="loading" class="flex-1">
            <div id="editor" ref="editor" style="height: 100%; width: 100%" spellcheck="true"></div>
        </divloading>
        <div ref="statusBar" class="bg-light text-end">
        </div>
    </div>
`
    };
