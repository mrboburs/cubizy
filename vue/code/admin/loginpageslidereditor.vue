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
                Image: "",
                Content: "",
                Footer : "",
                Status: true,
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
                    return 'loginpagesliders/loginpageslider_' + this.value.ID
                } else {
                    return 'loginpagesliders'
                }
            },
            loginpagesliderLink: function () {
                var link = ""
                if (this.value && this.value.ID > 0) {
                    link = '/loginpageslider/' + this.value.ID
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
                        this.Image = this.value.Image
                        this.Footer = this.value.Footer
                        this.Content = this.value.Content
                        this.Status = this.value.Status
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TitleError) { return }
                this.value.Title = this.Title
                this.value.Image = this.Image
                this.value.Footer = this.Footer
                this.value.Content = this.Content
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
        <div class="row">
            <div class="col">
                <label class="form-label"> Banner Image : </label>
                <ImageFile v-model="Image" :prefix="prefix" maxHeight="100%" maxWidth="100%">
                </ImageFile>
            </div>
            <div class="col">
                <formitem name="inputTitle" label="Title" :error="TitleError" v-model="Title" />

                <formitem name="loginpageslider_editor" label="Content" type="textarea" v-model="Content"/>

                <formitem name="inputFooter" label="Footer" v-model="Footer" />

                <div class="d-flex align-items-center end m-2">
                    <div class="form-check form-switch ml-2">
                        <input class="form-check-input" type="checkbox" id="inputIsSuperAdmin" v-model="Status">
                        <label class="form-check-label" for="inputIsSuperAdmin">Enabled</label>
                    </div>
                    <button type="submit" class="btn btn-success m-1" :disabled="loading">
                        Save
                    </button>
                    <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
                </div>
            </div>
        </div>
    </form>
</template>