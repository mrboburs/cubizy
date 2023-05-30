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
                        Wallet: 0,
                        IsAdmin: false,
                        IsSuperAdmin: false,
                        IsSupportagent: false,
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
                Wallet: 0,
                IsAdmin: false,
                IsSuperAdmin: false,
                IsSupportagent: false,
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
            Readonly: function (){
                if(this.value.EmailVerified ){
                    return true
                }else{
                    return false
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
                    this.Wallet = this.value.Wallet
                    this.IsAdmin = this.value.IsAdmin
                    this.IsSuperAdmin = this.value.IsSuperAdmin
                    this.IsSupportagent = this.value.IsSupportagent
                }
            },
            submit() {
                this.submitted = true
                if (this.NameError || this.EmailError || this.MobileError) { return }
                this.value.Photo = this.Photo
                this.value.Name = this.Name
                this.value.Email = this.Email
                this.value.Mobile = this.Mobile
                this.value.Wallet = this.Wallet
                this.value.IsAdmin = this.IsAdmin
                this.value.IsSuperAdmin = this.IsSuperAdmin
                this.value.IsSupportagent = this.IsSupportagent
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
                <div class="form-check form-switch">
                    <input class="form-check-input" type="checkbox" id="inputIsAdmin" v-model="IsAdmin">
                    <label class="form-check-label" for="inputIsAdmin">Is Admin</label>
                </div>
                <div class="form-check form-switch">
                    <input class="form-check-input" type="checkbox" id="inputIsSuperAdmin" v-model="IsSuperAdmin">
                    <label class="form-check-label" for="inputIsSuperAdmin">Is Super Admin</label>
                </div>
                <div class="form-check form-switch">
                    <input class="form-check-input" type="checkbox" id="inputIsSupportagent" v-model="IsSupportagent">
                    <label class="form-check-label" for="inputIsSupportagent">Is Supportagent</label>
                </div>
            </div>
            <div class="col">
                <formitem :customLayout="true" name="inputName" label="Name" :error="NameError" v-model="Name"  :type="Readonly?'readonly':'text'" />
                <formitem :customLayout="true" name="inputEmail" label="Email" :error="EmailError" v-model="Email" :type="Readonly?'readonly':'text'"  />
                <formitem :customLayout="true" name="inputMobile" label="Mobile" :error="MobileError" v-model="Mobile" :type="Readonly?'readonly':'text'" />
                <formitem :customLayout="true" name="inputWallet" label="Wallet" v-model="Wallet" type="readonly"/>
                <formitem v-if="value && value.ID" :customLayout="true" name="inputEmailCode" label="Email Code" :value="value.EmailCode" type="readonly" />
                <formitem v-if="value && value.ID" :customLayout="true" name="inputMobileCode" label="Mobile Code" v-model="value.MobileCode" type="readonly" />
                <formitem v-if="IsSupportagent" :customLayout="true" name="inputSellerAccountID" label="Seller Account ID" v-model="value.SellerAccountID" />
            </div>
        </div>
        <formitem :customLayout="true" label="">
            <button type="submit" class="btn btn-success" :disabled="loading">  
                <b-spinner small v-if="loading"></b-spinner>
                Save
            </button>
            <button class="btn btn-danger ml-1" @click="Reset">Cancel</button>
        </formitem>
    </form>
</template>