<script>
    export default {
        components: {
            'ace_editor': () => import("/vue/website/ace_editor.js"),
        },
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
                filedata : "",
                js_data : "",
                html_data : "",
                vue_data : "",
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
                        this.convertData(this.value.Data)
                    }
                } else {
                    file_session = ace.createEditSession("");
                    this.editor.convertData(file_session);
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
                        }else{
                            this.value.Data = ""
                        }
                        this.convertData(this.value.Data)
                    }
                }).catch((error) => {
                    this.error = true
                    this.message = error.message
                }).finally(() => {
                    this.loading = false
                });
            },
            convertData(data){
                var parts = data.split("template: `")
                var js_part = ""
                var template = ""
                if(Array.isArray(parts) && parts.length == 2){
                    js_part = parts[0].trim()
                    var template_parts = parts[1].split("`")
                    template = template_parts[0]
                    js_part = js_part + "\n" + template_parts[1].trim()
                }
                this.html_data = template
                this.js_data = js_part
                this.vue_data = `<`+`script>
`+ this.js_data +`    
</`+`script>
<`+`template>
`+ this.html_data +`        
</`+`template>`
            },
            submit() {
                var vue_data = this.$refs.vue_editor.getValue()
                var data = vue_data.replace(`<`+`script>
`, "")
                data = data.replace(`<`+`template>
`, "")
                data = data.replace(`
</`+`template>`, "")
                data = data.split('</'+'script>')

                var js_data = data[0].trim()
                var html_data = data[1].trim()

                var filedata = js_data.replace(new RegExp("[" + "}" + "]+$"), "")
                filedata += `    template`+": `" + html_data + "`" + `
                `+"}"
                this.filedata = filedata
                this.$emit('input', {
                    FileName: this.filename,
                    Data: this.filedata,
                })
            },
            doIndent(){
                this.$refs.vue_editor.doIndent()
            },
            togggleSearch() {
                this.$refs.vue_editor.togggleSearch()
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
                    //var element = this.$refs.editor_holder.$el;
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
        template: `{{{template}}}`
    };
</script>
<template>
    <divloading :loading="loading" class="d-flex flex-column w-100 h-100" ref="editor_holder">
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
        </div>
        <ace_editor id="split-0" ref="vue_editor" v-model="vue_data" mode="ace/mode/html" class="flex-1"/>
    </divloading>
</template>