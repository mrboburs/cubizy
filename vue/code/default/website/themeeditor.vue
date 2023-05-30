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
            themes: {
                type: Array,
                default: function () {
                    return [];
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
                Logo: "",
                Images: "",
                Tags : "",
                Description : "",
                FromTheme : "",
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
                        this.Logo = this.value.Logo
                        this.Images = this.value.Images
                        this.Tags = this.value.Tags
                        this.Description = this.value.Description
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
    <form @submit.prevent="submit">
        <div class="d-flex flex-column">
            <formitem :customLayout="true" name="inputName" label="Title" :error="TitleError" v-model="Title" />
            <formitem :customLayout="true" name="inputLogo" label="Logo" v-model="Logo" :prefix="prefix" type="image"  />
            <formitem :customLayout="true" name="inputTags" label="Tags" v-model="Tags" :prefix="prefix" type="tags"  />
            <formitem :customLayout="true" name="inputDescription" label="Description" v-model="Description" :prefix="prefix" type="textarea"/>
            <div class="d-flex justify-content-center m-2">
                <button type="submit" class="btn btn-success m-1" :disabled="loading">
                    Save
                </button>
                <button class="btn btn-danger m-1" @click="Reset" type="reset">Cancel</button>
            </div>
        </div>
    </form>
</template>