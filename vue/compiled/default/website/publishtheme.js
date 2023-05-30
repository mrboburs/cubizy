
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
                Image: "",
                Content: "",
                FromTheme : "",
                Admin: false,
                Tutor: false,
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
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                }
            },
            submit() {
                this.submitted = true
                if(!this.value || !this.value.ID){
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "publish",
                    data: {
                        from_theme_id: this.value.ID
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.$emit('input', this.value)
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                })
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `
    <form @submit.prevent="submit">
        <divloading :fullpage="false" :loading="loading" class="text-center d-flex flex-column">
            <v-alert v-model="message" :error="error" />
            <p>This action will publish your theme to use by all other sellers and service providers. Once approved by cubizy team anyone with cubizy account can use and modify your theme only for there own use. You can not send or shear any data genereted by cubizy with any other third party server other theme cubizy server. You can not collect any personal information of any user applying theme theme or using it. Voileting that rule will lead to remove your theme from public listing. </p>
            <p>By publishing this theme you aggree to terms and conditions and protecting user pricacy.  </p>
            <div class="d-flex justify-content-center m-2">
                <button class="btn m-1 btn-primary" type="submit">
                    Publish
                </button>
                <button class="btn btn-danger m-1" @click="Reset" type="reset">Cancel</button>
            </div>
        </divloading>
    </form>
`
    }
