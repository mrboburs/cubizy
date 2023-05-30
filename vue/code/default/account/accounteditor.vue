<script>
    export default {
        props: {
            value: {},
            title: {
                type: String,
                required: false
            },
        },
        data: () => {
            return {
                submitted: false,
                Title: "You Business Name",
                Description: "",
                Keywords: "",
                Email: "",
                Mobile: "",
                Youtube: "",
                Facebook: "",
                Instagram: "",
                Pinterest: "",
                WhatsApp: "",

                Logo: "",
                WideLogo: "",
                Banner: "",

                regEmailExp: /\S+@\S+\.\S+/,
                regMobileExp: /^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\s\./0-9]*$/,
            }
        },
        computed: {
            ...Vuex.mapState(['user']),
            TitleError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Title.trim()) {
                    return "Title can not  be empty"
                }
            },
            LogoError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Logo) {
                    return "Logo can not  be empty"
                }
            },
            DescriptionError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Description.trim()) {
                    return "Description can not  be empty"
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
                    if (this.value && this.value.ID > 0) {
                        this.submitted = false
                        this.Title = this.value.Title
                        this.Description = this.value.Description
                        this.Keywords = this.value.Keywords
                        this.Email = this.value.Email
                        this.Mobile = this.value.Mobile
                        this.Logo = this.value.Logo
                        this.WideLogo = this.value.WideLogo
                        this.Banner = this.value.Banner
                        this.Youtube = this.value.Youtube
                        this.Facebook = this.value.Facebook
                        this.Instagram = this.value.Instagram
                        this.Pinterest = this.value.Pinterest
                        this.WhatsApp = this.value.WhatsApp
                    } else {
                        this.Title = this.user.Name
                        this.Email = this.user.Email
                        this.Mobile = this.user.Mobile
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TitleError || this.LogoError || this.DescriptionError) {
                    return
                }
                var value = {}
                if (this.account) {
                    value.ID = this.account.ID
                }
                value.Title = this.Title
                value.Description = this.Description
                value.Keywords = this.Keywords
                value.Email = this.Email
                value.Mobile = this.Mobile
                value.Logo = this.Logo
                value.WideLogo = this.WideLogo
                value.Banner = this.Banner
                value.Youtube = this.Youtube
                value.Facebook = this.Facebook
                value.Instagram = this.Instagram
                value.Pinterest = this.Pinterest
                value.WhatsApp = this.WhatsApp
                this.$emit('input', value)
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
    <form @submit.prevent="submit" class="row">
        <div class="col col-md-6">
            <formitem name="inputTitle" label="Title *" v-model="Title" :error="TitleError" />
            <formitem type="textarea" name="inputDescription" label="Description" v-model="Description"
                :error="DescriptionError" />
            <formitem name="inputKeywords" type="tags" label="Keywords" v-model="Keywords" />
            <formitem name="inputEmail" label="Email *" v-model="Email" />
            <formitem name="inputMobile" label="Mobile *" v-model="Mobile" />
            <formitem name="inputFacebook" label="Facebook Link" v-model="Facebook" />
            <formitem name="inputInstagram" label="Instagram Link" v-model="Instagram" />
            <formitem name="inputWhatsApp" label="WhatsApp Link" v-model="WhatsApp" />
        </div>
        <div class="col col-md-6">
            <div class="row">
                <div class="col-auto">
                    <div class="m-1">
                        <formitem name="inputLogo" label="Logo *" :error="LogoError">
                            <ImageFile v-model="Logo" :prefix="'account'"
                                maxWidth="100px" maxHeight="100px">
                            </ImageFile>
                        </formitem>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="m-1">
                        <label class="form-label"> WideLogo : </label>
                        <ImageFile v-model="WideLogo" :prefix="'account'"
                            maxWidth="250px" maxHeight="100px">
                        </ImageFile>
                    </div>
                </div>
            </div>
            <div class="m-1 mb-3">
                <label class="form-label"> Banner : </label>
                <ImageFile v-if="user.ID" v-model="Banner" :prefix="'account'" class="p-1"
                    maxWidth="100%" maxHeight="300px"> </ImageFile>
            </div>
            <formitem name="inputYoutube" label="Youtube Link" v-model="Youtube" />
            <formitem name="inputPinterest" label="Pinterest Link" v-model="Pinterest" />
        </div>

        <div v-if="$route.name == 'creataccount'" class="d-flex justify-content-end">
            <button type="submit" class="btn btn-success">
                Next
            </button>
        </div>
        <div v-else class="d-flex justify-content-end">
            <button type="button" class="btn btn-danger m-1" @click.prevent="Reset">
                Cancel
            </button>
            <button type="submit" class="btn btn-success m-1">
                Save
            </button>
        </div>
    </form>
</template>