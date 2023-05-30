
    export default {
        props: { // 
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Variation: "",
                    }
                }
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.load()
                    this.$emit('onset', this.value)
                }
            },
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        methods: {
            getContent(data) {
                if (this.loading || !this.value || !this.value.ID) {
                    return
                }
                if (!data) {
                    data = {}
                }
                data.ProductID = this.value.ID
                this.loading = true
                return this.$store.dispatch('call', {
                    api: "description",
                    data: data
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        return data.Result.Content
                    } else {
                        this.error = true
                        return ""
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                })
            },
            prefix() {
                return "product/" + this.value.ID + "/images"
            },
            setEditor() {
                var editor = grapesjs.init({
                    container: this.$refs.gjs,
                    components: '<div class="txt-red">Loading.....!</div>',
                    style: '.txt-red{color: red}',
                    height: '100%',
                    width: '100%',
                    protectedCss: '',
                    fromElement: 1,
                    showOffsets: 1,
                    assetManager: {
                        embedAsBase64: 1,
                        //assets: images
                    },
                    selectorManager: { componentFirst: true },
                    styleManager: { clearProperties: 1 },
                    storageManager: {
                        id: 'gjs-',             // Prefix identifier that will be used on parameters
                        type: 'simple-storage',          // Type of the storage
                        autosave: false,         // Store data automatically
                        autoload: false,         // Autoload stored data on init
                        stepsBeforeSave: 1,     // If autosave enabled, indicates how many changes are necessary before store method is triggered
                    },
                    plugins: ['gjs-preset-webpage'],
                    pluginsOpts: {

                        'gjs-preset-webpage': {
                            modalImportTitle: 'Import Template',
                            modalImportLabel: '<div style="margin-bottom: 10px; font-size: 13px;">Paste here your HTML/CSS and click Import</div>',
                            modalImportContent: function (editor) {
                                return editor.getHtml() + '<style>' + editor.getCss() + '</style>'
                            },
                            filestackOpts: null, //{ key: 'AYmqZc2e8RLGLE7TGkX3Hz' },
                            aviaryOpts: false,
                            blocksBasicOpts: { flexGrid: 1 },
                            customStyleManager: [{
                                name: 'General',
                                buildProps: ['float', 'display', 'position', 'top', 'right', 'left', 'bottom'],
                                properties: [{
                                    name: 'Alignment',
                                    property: 'float',
                                    type: 'radio',
                                    defaults: 'none',
                                    list: [
                                        { value: 'none', className: 'fa fa-times' },
                                        { value: 'left', className: 'fa fa-align-left' },
                                        { value: 'right', className: 'fa fa-align-right' }
                                    ],
                                },
                                { property: 'position', type: 'select' }
                                ],
                            }, {
                                name: 'Dimension',
                                open: false,
                                buildProps: ['width', 'flex-width', 'height', 'max-width', 'min-height', 'margin', 'padding'],
                                properties: [{
                                    id: 'flex-width',
                                    type: 'integer',
                                    name: 'Width',
                                    units: ['px', '%'],
                                    property: 'flex-basis',
                                    toRequire: 1,
                                }, {
                                    property: 'margin',
                                    properties: [
                                        { name: 'Top', property: 'margin-top' },
                                        { name: 'Right', property: 'margin-right' },
                                        { name: 'Bottom', property: 'margin-bottom' },
                                        { name: 'Left', property: 'margin-left' }
                                    ],
                                }, {
                                    property: 'padding',
                                    properties: [
                                        { name: 'Top', property: 'padding-top' },
                                        { name: 'Right', property: 'padding-right' },
                                        { name: 'Bottom', property: 'padding-bottom' },
                                        { name: 'Left', property: 'padding-left' }
                                    ],
                                }],
                            }, {
                                name: 'Typography',
                                open: false,
                                buildProps: ['font-family', 'font-size', 'font-weight', 'letter-spacing', 'color', 'line-height', 'text-align', 'text-decoration', 'text-shadow'],
                                properties: [
                                    { name: 'Font', property: 'font-family' },
                                    { name: 'Weight', property: 'font-weight' },
                                    { name: 'Font color', property: 'color' },
                                    {
                                        property: 'text-align',
                                        type: 'radio',
                                        defaults: 'left',
                                        list: [
                                            { value: 'left', name: 'Left', className: 'fa fa-align-left' },
                                            { value: 'center', name: 'Center', className: 'fa fa-align-center' },
                                            { value: 'right', name: 'Right', className: 'fa fa-align-right' },
                                            { value: 'justify', name: 'Justify', className: 'fa fa-align-justify' }
                                        ],
                                    }, {
                                        property: 'text-decoration',
                                        type: 'radio',
                                        defaults: 'none',
                                        list: [
                                            { value: 'none', name: 'None', className: 'fa fa-times' },
                                            { value: 'underline', name: 'underline', className: 'fa fa-underline' },
                                            { value: 'line-through', name: 'Line-through', className: 'fa fa-strikethrough' }
                                        ],
                                    }, {
                                        property: 'text-shadow',
                                        properties: [
                                            { name: 'X position', property: 'text-shadow-h' },
                                            { name: 'Y position', property: 'text-shadow-v' },
                                            { name: 'Blur', property: 'text-shadow-blur' },
                                            { name: 'Color', property: 'text-shadow-color' }
                                        ],
                                    }],
                            }, {
                                name: 'Decorations',
                                open: false,
                                buildProps: ['opacity', 'background-color', 'border-radius', 'border', 'box-shadow', 'background'],
                                properties: [{
                                    type: 'slider',
                                    property: 'opacity',
                                    defaults: 1,
                                    step: 0.01,
                                    max: 1,
                                    min: 0,
                                }, {
                                    property: 'border-radius',
                                    properties: [
                                        { name: 'Top', property: 'border-top-left-radius' },
                                        { name: 'Right', property: 'border-top-right-radius' },
                                        { name: 'Bottom', property: 'border-bottom-left-radius' },
                                        { name: 'Left', property: 'border-bottom-right-radius' }
                                    ],
                                }, {
                                    property: 'box-shadow',
                                    properties: [
                                        { name: 'X position', property: 'box-shadow-h' },
                                        { name: 'Y position', property: 'box-shadow-v' },
                                        { name: 'Blur', property: 'box-shadow-blur' },
                                        { name: 'Spread', property: 'box-shadow-spread' },
                                        { name: 'Color', property: 'box-shadow-color' },
                                        { name: 'Shadow type', property: 'box-shadow-type' }
                                    ],
                                }, {
                                    property: 'background',
                                    properties: [
                                        { name: 'Image', property: 'background-image' },
                                        { name: 'Repeat', property: 'background-repeat' },
                                        { name: 'Position', property: 'background-position' },
                                        { name: 'Attachment', property: 'background-attachment' },
                                        { name: 'Size', property: 'background-size' }
                                    ],
                                },],
                            }, {
                                name: 'Extra',
                                open: false,
                                buildProps: ['transition', 'perspective', 'transform'],
                                properties: [{
                                    property: 'transition',
                                    properties: [
                                        { name: 'Property', property: 'transition-property' },
                                        { name: 'Duration', property: 'transition-duration' },
                                        { name: 'Easing', property: 'transition-timing-function' }
                                    ],
                                }, {
                                    property: 'transform',
                                    properties: [
                                        { name: 'Rotate X', property: 'transform-rotate-x' },
                                        { name: 'Rotate Y', property: 'transform-rotate-y' },
                                        { name: 'Rotate Z', property: 'transform-rotate-z' },
                                        { name: 'Scale X', property: 'transform-scale-x' },
                                        { name: 'Scale Y', property: 'transform-scale-y' },
                                        { name: 'Scale Z', property: 'transform-scale-z' }
                                    ],
                                }]
                            }, {
                                name: 'Flex',
                                open: false,
                                properties: [{
                                    name: 'Flex Container',
                                    property: 'display',
                                    type: 'select',
                                    defaults: 'block',
                                    list: [
                                        { value: 'block', name: 'Disable' },
                                        { value: 'flex', name: 'Enable' }
                                    ],
                                }, {
                                    name: 'Flex Parent',
                                    property: 'label-parent-flex',
                                    type: 'integer',
                                }, {
                                    name: 'Direction',
                                    property: 'flex-direction',
                                    type: 'radio',
                                    defaults: 'row',
                                    list: [{
                                        value: 'row',
                                        name: 'Row',
                                        className: 'fas fa-arrow-right',
                                        title: 'Row',
                                    }, {
                                        value: 'row-reverse',
                                        name: 'Row reverse',
                                        className: 'fas fa-arrow-left',
                                        title: 'Row reverse',
                                    }, {
                                        value: 'column',
                                        name: 'Column',
                                        title: 'Column',
                                        className: 'fas fa-arrow-down',
                                    }, {
                                        value: 'column-reverse',
                                        name: 'Column reverse',
                                        title: 'Column reverse',
                                        className: 'fas fa-arrow-up',
                                    }],
                                }, {
                                    name: 'Justify',
                                    property: 'justify-content',
                                    type: 'radio',
                                    defaults: 'flex-start',
                                    list: [{
                                        value: 'flex-start',
                                        className: 'mdi mdi-align-horizontal-left',
                                        title: 'Start',
                                    }, {
                                        value: 'flex-end',
                                        title: 'End',
                                        className: 'mdi mdi-align-horizontal-right',
                                    }, {
                                        value: 'space-between',
                                        title: 'Space between',
                                        className: 'mdi mdi-align-horizontal-distribute',
                                    }, {
                                        value: 'space-around',
                                        title: 'Space around',
                                        className: 'mdi mdi-align-horizontal-distribute',
                                    }, {
                                        value: 'center',
                                        title: 'Center',
                                        className: 'mdi mdi-align-horizontal-center',
                                    }],
                                }, {
                                    name: 'Align',
                                    property: 'align-items',
                                    type: 'radio',
                                    defaults: 'center',
                                    list: [{
                                        value: 'flex-start',
                                        title: 'Start',
                                        className: 'mdi mdi-align-vertical-top',
                                    }, {
                                        value: 'flex-end',
                                        title: 'End',
                                        className: 'mdi mdi-align-vertical-bottom',
                                    }, {
                                        value: 'stretch',
                                        title: 'Stretch',
                                        className: 'icons-flex icon-al-str',
                                    }, {
                                        value: 'center',
                                        title: 'Center',
                                        className: 'mdi mdi-align-vertical-center',
                                    }],
                                }, {
                                    name: 'Flex Children',
                                    property: 'label-parent-flex',
                                    type: 'integer',
                                }, {
                                    name: 'Order',
                                    property: 'order',
                                    type: 'integer',
                                    defaults: 0,
                                    min: 0
                                }, {
                                    name: 'Flex',
                                    property: 'flex',
                                    type: 'composite',
                                    properties: [{
                                        name: 'Grow',
                                        property: 'flex-grow',
                                        type: 'integer',
                                        defaults: 0,
                                        min: 0
                                    }, {
                                        name: 'Shrink',
                                        property: 'flex-shrink',
                                        type: 'integer',
                                        defaults: 0,
                                        min: 0
                                    }, {
                                        name: 'Basis',
                                        property: 'flex-basis',
                                        type: 'integer',
                                        units: ['px', '%', ''],
                                        unit: '',
                                        defaults: 'auto',
                                    }],
                                }, {
                                    name: 'Align',
                                    property: 'align-self',
                                    type: 'radio',
                                    defaults: 'auto',
                                    list: [{
                                        value: 'auto',
                                        name: 'Auto',
                                    }, {
                                        value: 'flex-start',
                                        title: 'Start',
                                        className: 'icons-flex icon-al-start',
                                    }, {
                                        value: 'flex-end',
                                        title: 'End',
                                        className: 'icons-flex icon-al-end',
                                    }, {
                                        value: 'stretch',
                                        title: 'Stretch',
                                        className: 'icons-flex icon-al-str',
                                    }, {
                                        value: 'center',
                                        title: 'Center',
                                        className: 'icons-flex icon-al-center',
                                    }],
                                }]
                            }
                            ],
                        },
                    },
                });
                editor.Panels.removeButton('options', 'export-template');
                editor.Panels.addButton('options', [{
                    id: 'save-db',
                    className: 'fa fa-floppy-o',
                    command: 'save-db',
                    attributes: { title: 'Save DB' }
                }]);
                editor.Panels.getButton('options', 'sw-visibility').set('active', 1).set('className', 'far fa-square' );
                // Add the command
                editor.Commands.add('save-db', {
                    run: function (editor, sender) {
                        sender && sender.set('active', 0); // turn off the button
                        editor.store();
                    }
                });
                var component = this
                editor.StorageManager.add('simple-storage', {
                    /**
                     * Load the data
                     * @param  {Array} keys Array containing values to load, eg, ['gjs-components', 'gjs-style', ...]
                     * @param  {Function} clb Callback function to call when the load is ended
                     * @param  {Function} clbErr Callback function to call in case of errors
                     */
                    load(keys, clb, clbErr) {

                        component.getContent().then((data) => {
                            var parts = data.split("<style>")
                            var html = parts[0]
                            var css = ""
                            if(parts.length > 1){
                                css = parts[1]
                                css = css.replace("</style>", "")
                            }
                            const result = {};
                            result["html"] = html
                            result["css"] = css
                            // Might be called inside some async method
                            clb(result);
                        })
                    },

                    /**
                     * Store the data
                     * @param  {Object} data Data object to store
                     * @param  {Function} clb Callback function to call when the load is ended
                     * @param  {Function} clbErr Callback function to call in case of errors
                     */
                    store(data, clb, clbErr) {
                        var htmldata = editor.getHtml();
                        var cssdata = editor.getCss();

                        var Description = htmldata  + '<style>' + cssdata + `</style>`;

                        var data = {
                            Content: Description
                        }
                        component.getContent(data)
                        console.log(htmldata);
                        console.log(cssdata);
                    }
                });

                editor.load()
                this.editor = editor
            }
        },
        mounted: function () {
            this.setEditor()
            this.$emit('onload')
        },
        template: `
    <divloading :fullpage="false" :loading="loading" class="mt-1 d-flex flex-column" style=" height: calc(100vh - 150px);" >
        <v-alert v-model="message" :error="error"></v-alert>
        <div ref="gjs" class="flex-fill"> </div>
    </divloading>
`
    }
