
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                editor: false,
            }
        },
        mounted: function () {
            ace.require("ace/ext/language_tools");
            this.editor = ace.edit("editor");
            this.editor.session.setMode("ace/mode/html");
            this.editor.setTheme("ace/theme/twilight");
            this.editor.setOptions({
                enableBasicAutocompletion: true,
                enableSnippets: true,
                enableLiveAutocompletion: false,
                fontSize: "16px"
            });
        },
        template: `
    <div class="row">
        <div class="col-auto">
            <div class="card" style="width: 18rem;">
                <div class="card-header">
                    List of files
                </div>
                <divloading :loading="loading" class="card-body">

                </divloading>
            </div>
        </div>
        <div class="col">
            <div id="editor" style="height: calc(100vh - 160px ); width: 100%"></div>
        </div>
    </div>

`
    };
