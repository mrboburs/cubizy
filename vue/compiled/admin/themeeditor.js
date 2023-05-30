
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
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

                Title: "",
                Image: "",
                Tags: "",
                Description: "",
                Published : false,
                Admin: false,
                Seller: false,
                mode: "editor",
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.SetData()
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
            TitleError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Title.trim()) {
                    return "Title can not  be empty"
                }
            },
            prefix: function () {
                if (this.value && this.value.ID) {
                    return 'themes/theme_' + this.value.ID
                } else {
                    return 'files'
                }
            },
            themeLink: function () {
                var link = ""
                if (this.value && this.value.ID > 0) {
                    link = '/theme/' + this.value.ID
                    link = window.location.protocol + "//" + window.application.BaseDomin + link
                }
                return link
            }
        },
        methods: {
            highlighter(code) {
                // js highlight example
                return Prism.highlight(code, Prism.languages.js, "json");
            },
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    if (this.value.Title) {
                        this.Title = this.value.Title
                        this.Image = this.value.Image
                        this.Tags = this.value.Tags
                        this.Description = this.value.Description
                        this.Published = this.value.Published
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TitleError) { return }
                this.value.Title = this.Title
                this.value.Image = this.Image
                this.value.Tags = this.Tags
                this.value.Description = this.Description
                this.value.Published = !this.Published
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `
    <form @submit.prevent="submit" class="vw90">
        <div class="row">
            <div class="col-auto" v-if="mode != 'priview'">
                <label class="form-label"> Banner Image : </label>
                <ImageFile v-model="Image" :prefix="prefix" maxHeight="250px" maxWidth="400px">
                </ImageFile>
            </div>
            <div class="col">
                <formitem name="inputName" label="Title" :error="TitleError" v-model="Title" />
                <formitem name="inputTags" label="Tags" v-model="Tags" :prefix="prefix" type="tags"  />
                <formitem name="inputDescription" label="Description" v-model="Description" :prefix="prefix" type="textarea"  />
                <div class="d-flex align-items-center end m-2">
                    <button type="submit" class="btn btn-success m-1" :disabled="loading" :class="{'btn-success': !Published, 'btn-danger': Published}">
                        <span v-if="Published">Unpublish</span> 
                        <span v-else>Publish</span> 
                    </button>
                    <button class="btn btn-secondary m-1" @click="Reset">Cancel</button>
                </div>
            </div>
        </div>
    </form>
`
    }
