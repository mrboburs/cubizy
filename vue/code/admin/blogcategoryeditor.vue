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
                Name: "",
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
            NameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Name.trim()) {
                    return "Name can not  be empty"
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
                    if (this.value.Name) {
                        this.Name = this.value.Name
                        this.Status = this.value.Status
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError) { return }
                this.value.Name = this.Name
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
    <form @submit.prevent="submit">
        <formitem name="inputName" label="Name" :error="NameError" v-model="Name" />
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
    </form>
</template>