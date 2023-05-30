<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                        Photo: "",
                        Name: "",
                        Email: "",
                        Mobile: "",
                        IsAdmin: false,
                        IsSuperAdmin: false,
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

                Photo: "",
                Name: "",
                Email: "",
                Mobile: "",
                IsAdmin: false,
                IsSuperAdmin: false,
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
            showmessage: {
                // getter
                get: function () {
                    if (this.message) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.message = ""
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'alert-danger'
                } else {
                    return 'alert-success'
                }
            },
            NameError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Name.trim()) {
                    return "Please provide a valid Name"
                }
            },
            EmailError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Email.trim()) {
                    return "Please provide a valid Email"
                }
            },
            MobileError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Mobile.trim()) {
                    return "Please provide a valid Mobile"
                }
            },
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    this.Photo = this.value.Photo
                    this.Name = this.value.Name
                    this.Email = this.value.Email
                    this.Mobile = this.value.Mobile
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.EmailError || this.MobileError) { return }
                this.value.Photo = this.Photo
                this.value.Name = this.Name
                this.value.Email = this.Email
                this.value.Mobile = this.Mobile
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload')
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <form @submit.prevent="submit" class="container-fluid">
        <div class="row">
            <div class="col">
                <label class="form-label"> User Photo : </label>
                <ImageFile v-if="value.ID" v-model="Photo" prefix="user_photo" style="width: fit-content; max-width: 100%;"> </ImageFile>
            </div>
            <div class="col">
                <formitem :customLayout="true" name="inputName" label="Name" :error="NameError" v-model="Name" />
                <formitem :customLayout="true" name="inputEmail" label="Email" :error="EmailError" v-model="Email" />
                <formitem :customLayout="true" name="inputMobile" label="Mobile" :error="MobileError" v-model="Mobile" />
            </div>
        </div>
        <formitem :customLayout="true" label="">
            <button type="submit" class="btn btn-success" :disabled="loading">  
                <b-spinner small v-if="loading"></b-spinner>
                Save
            </button>
            <button class="btn btn-danger ml-1" @click.prevent="Reset">Cancel</button>
        </formitem>
    </form>
</template>