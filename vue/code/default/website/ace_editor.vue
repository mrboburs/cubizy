<script>
    export default {
        props: {
            value: {
                type: String,
                default: ""
            },
            mode:{
                type: String,
                default: "ace/mode/html"
            },
        },
        data: () => {
            return {
                editor: null,
                editor_options : {
                    enableBasicAutocompletion: true,
                    enableSnippets: true,
                    enableLiveAutocompletion: true,
                    fontSize: 16,
                    wrap: "off",
                    wrapBehavioursEnabled: true
                },
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                this.setSession()
            },
        },
        mounted: function () {
            ace.require("ace/ext/language_tools");
            this.editor = ace.edit(this.$refs.editor);
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
            this.setSession()
        },
        methods: {
            setSession() {
                var file_session
                file_session = ace.createEditSession(this.value);
                file_session.setMode(this.mode)
                this.editor.setSession(file_session);
            },
            getValue() {
                var data = this.editor.getValue()
                return data
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
        },
        template: `{{{template}}}`
    };
</script>
<template>
    <div class="d-flex flex-column">
        <div ref="editor" class="flex-1" spellcheck="true"></div>
        <div ref="statusBar" class="bg-light text-end">
        </div>
    </div>
</template>