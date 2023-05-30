
    // https://mystudyfiles.s3.ap-southeast-2.amazonaws.com/empty.jpg
    export default {
        props: {
            title: {
                type: String,
                default: ""
            },
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
            showmodal: {
                type: Boolean,
                default: false
            },
        },
        data() {
            return {
                loading: 0,
                error: false,
                message: "",

                empty_image: "https://mystudyfiles.s3.ap-southeast-2.amazonaws.com/empty.jpg",
                selectedtab: 2,
                accessURL: "/",
                external_link: "",
                files: [],
                upload_count: 0,
                uploded_files: [],
            };
        },
        mounted() {
            this.$emit('load', this)
        },
        methods: {
            applyExternalLink: function (event) {
                if (this.multiple) {
                    this.$emit('input', [this.external_link])
                } else {
                    this.$emit('input', this.external_link)
                }
                this.external_link = ""
            },
        },
        template: `
    <div>
        <v-modal :value="showmodal" @input="$emit('close')" :title="title?'Select Image for'+ title: 'Select image'" header-close-variant="light" title-class="font-18"
            hide-footer>
            <ul class="nav nav-tabs">
                <li class="nav-item">
                    <a class="nav-link" :class="{active: selectedtab == 2}" @click.prevent="selectedtab = 2"
                        href="#">Selecte file</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" :class="{active: selectedtab == 3}" @click.prevent="selectedtab = 3"
                        href="#">Set from external link</a>
                </li>
            </ul>
            <v-files-static v-if="selectedtab == 2" @input="$emit('input', $event)" :multiple="multiple" :value="value" :only_image="only_image" :prefix="prefix" />
            <div v-if="selectedtab == 3" title="Set from external link">
                <div class="form-group">
                    <label for="parent_category">External Link</label>
                    <input id="external_link" v-model="external_link" type="text" class="form-control"
                        placeholder="Image url" />
                </div>
                <div class="text-right">
                    <button type="button" class="btn btn-success" @click="applyExternalLink">
                        Apply
                    </button>
                </div>
            </div>
        </v-modal>
    </div>
`
    };
