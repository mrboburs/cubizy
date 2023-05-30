<script>
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
                Logo : "",
                Images: "",
                Tags: "",
                Description: "",
                Status : "",
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
                debugger
                if (this.value) {
                    this.submitted = false
                    if (this.value.Title) {
                        this.Title = this.value.Title
                        this.Logo = this.value.Logo
                        this.Images = this.value.Images
                        this.Tags = this.value.Tags
                        this.Description = this.value.Description
                        this.Status = this.value.Status
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TitleError) { return }
                this.value.Title = this.Title
                this.value.Logo = this.Logo
                this.value.Images = this.Images
                this.value.Tags = this.Tags
                this.value.Description = this.Description
                this.value.Status = this.Status
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <form @submit.prevent="submit" class="vw90">
        <divloading :loading="loading" class="row">
            <div class="col-auto" v-if="mode != 'priview'">
                <label class="form-label"> Banner Image : </label>
                <ImageFile v-model="Logo" :prefix="prefix" maxHeight="250px" maxWidth="400px">
                </ImageFile>
            </div>
            <div class="col">
                <formitem name="inputName" label="Title" :error="TitleError" v-model="Title" />
                <formitem name="inputTags" label="Tags" v-model="Tags" type="tags"  />
                <formitem name="inputDescription" label="Description" v-model="Description" type="textarea"  />
                <formitem name="inputStatus" label="Status" v-model="Status" type="select" :values="['Submitted','Testing','Rejected', 'Accepted', 'Published']"/>
                <div class="d-flex align-items-center end m-2">
                    <button type="submit" class="btn btn-success m-1">Save</button>
                    <button class="btn btn-secondary m-1" @click="Reset">Cancel</button>
                </div>
            </div>
        </divloading>
    </form>
</template>