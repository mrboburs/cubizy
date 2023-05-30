
    export default {
        props: {
            value: {
                default: "",
            },
            prefix: {
                default : false,
            },
        },
        data: () => {
            return {
                quill: false,
                showmodal: false,
            };
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue != oldValue) {
                    if (this.quill && this.quill.root.innerHTML != newValue) {
                        this.quill.root.innerHTML = newValue
                    }
                }
            },
        },
        methods: {
            load() {
                this.quill = new Quill(this.$refs["editor"], {
                    modules: {
                        formula: true,
                        syntax: true,
                        toolbar: this.$refs["toolbar-container"]
                    },
                    placeholder: 'Compose an epic...',
                    theme: 'snow'
                });
                var toolbar = this.quill.getModule('toolbar');
                toolbar.addHandler('image', this.imageHandler);
                if (this.value) {
                    this.quill.root.innerHTML = this.value
                }
                var component = this
                this.quill.on('text-change', function (delta, oldDelta, source) {
                    if(component.quill.root.innerHTML != component.value){
                        component.$emit("input", component.quill.root.innerHTML);
                    }
                });
            },
            imageHandler() {
                if (!this.prefix) {
                    alert("Please save new record once to add images")
                } else {
                    this.showmodal = true
                }
            },
            setImage(value) {
                if (value) {
                    var range = this.quill.getSelection();
                    if(!range){
                        range = {
                            index : this.quill.getLength()
                        }
                    }
                    this.quill.insertEmbed(range.index, 'image', value, Quill.sources.USER);
                }
                this.showmodal = false
            },
        },
        mounted: function () {
            this.load();
        },
        template: `
    <div>
        <div ref="toolbar-container">
            <span class="ql-formats">
                <select class="ql-font"></select>
                <select class="ql-size"></select>
            </span>
            <span class="ql-formats">
                <button class="ql-bold"></button>
                <button class="ql-italic"></button>
                <button class="ql-underline"></button>
                <button class="ql-strike"></button>
            </span>
            <span class="ql-formats">
                <select class="ql-color"></select>
                <select class="ql-background"></select>
            </span>
            <span class="ql-formats">
                <button class="ql-script" value="sub"></button>
                <button class="ql-script" value="super"></button>
            </span>
            <span class="ql-formats">
                <button class="ql-header" value="1"></button>
                <button class="ql-header" value="2"></button>
                <button class="ql-blockquote"></button>
                <button class="ql-code-block"></button>
            </span>
            <span class="ql-formats">
                <button class="ql-list" value="ordered"></button>
                <button class="ql-list" value="bullet"></button>
                <button class="ql-indent" value="-1"></button>
                <button class="ql-indent" value="+1"></button>
            </span>
            <span class="ql-formats">
                <button class="ql-direction" value="rtl"></button>
                <select class="ql-align"></select>
            </span>
            <span class="ql-formats">
                <button class="ql-link"></button>
                <button class="ql-image"></button>
                <button class="ql-video"></button>
                <button class="ql-formula"></button>
            </span>
            <span class="ql-formats">
                <button class="ql-clean"></button>
            </span>
        </div>
        <div ref="editor" style="min-height: 390px;">
        </div>
        <v-files v-if="prefix" :showmodal="showmodal" :prefix="prefix" @input="setImage($event)" @close="showmodal = false" />
    </div>
`,
    };
