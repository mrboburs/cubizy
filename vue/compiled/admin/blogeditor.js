
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
                Category : 0,
                Image: "",
                Content: "",
                Status: true,
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
            CategoryError: function () {
                if (!this.submitted) {
                    return false
                }
                var Category = parseInt(this.Category)
                if (!this.Category) {
                    return "Category can not  be empty"
                }
            },
            ContentError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Content.trim()) {
                    return "Content can not  be empty"
                }
            },
            prefix: function () {
                if (this.value && this.value.ID) {
                    return 'blogs/blog_' + this.value.ID
                } else {
                    return 'blogs'
                }
            },
            blogLink: function () {
                var link = ""
                if (this.value && this.value.ID > 0) {
                    link = '/blog/' + this.value.ID
                    link = window.location.protocol + "//" + window.application.BaseDomin + link
                }
                return link
            }
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    if (this.value.Title) {
                        this.Title = this.value.Title
                        this.Tags = this.value.Tags
                        this.Category = this.value.Category
                        this.Image = this.value.Image
                        this.Content = this.value.Content
                        this.Status = this.value.Status
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TitleError || this.CategoryError || this.ContentError) { return }
                this.value.Title = this.Title
                this.value.Tags = this.Tags
                this.value.Category = this.Category
                this.value.Image = this.Image
                this.value.Content = this.Content
                this.value.Status = this.Status
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
                <formitem name="inputName" label="Tags" v-model="Tags" type="tags" />
                <formitem name="inputCategory" label="Category" v-model="Category" type="select"
                    service="blogcategories" displayby="Name" selectby="ID":error="CategoryError"  />
                <div class="d-flex align-items-center end m-2">
                    <div class="form-check form-switch ml-2">
                        <input class="form-check-input" type="checkbox" id="inputIsSuperAdmin" v-model="Status">
                        <label class="form-check-label" for="inputIsSuperAdmin">Enabled</label>
                    </div>
                    <button type="submit" class="btn btn-success m-1" :disabled="loading">
                        Save
                    </button>
                    <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
                    <button v-if="value && value.ID > 0 && mode != 'priview'" type="button" class="btn btn-primary m-1"
                        @click.prevent="mode = 'priview'">
                        Show Priview
                    </button>
                    <button v-if="mode != 'raw'" type="button" class="btn btn-primary m-1"
                        @click.prevent="mode = 'raw'">
                        Show Raw
                    </button>
                    <button v-if="mode != 'editor'" type="button" class="btn btn-primary m-1"
                        @click.prevent="mode = 'editor'">
                        Show Editor
                    </button>
                </div>
            </div>
        </div>
        <formitem name="blog_editor" label="Content" :error="ContentError" v-model="Content">
            <iframe id="blog_editor" v-if="mode == 'priview'" :src="blogLink" class="blog_editor"></iframe>
            <v-quill id="blog_editor" v-if="mode == 'editor'" v-model="Content" :prefix="prefix" class="blog_editor">
            </v-quill>
            <textarea id="blog_editor" v-if="mode == 'raw'" v-model="Content"
                class="form-control blog_editor" style="min-height: 400px;"></textarea>
        </formitem>
    </form>
`
    }
