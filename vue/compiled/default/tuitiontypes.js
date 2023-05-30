
    export default {
        props: {
            values: {
                type: Array,
                default: function () {
                    return [];
                },
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                popupitem: false
            }
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account', 'sessiontypes']),
            showpopup: {
                // getter
                get: function () {
                    if (this.popupitem) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.popupitem = false
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'danger'
                } else {
                    return 'success'
                }
            }
        },
        methods: {
            select(Title){
                if(this.values.includes(Title)){
                    this.values.splice(this.values.indexOf(Title), 1);
                }else{
                    this.values.push(Title); 
                }
                this.$emit('input', this.values)
            },
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "sessiontypes",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_sessiontypes', data.data)
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
            this.load()
        },
        template: `
    <div class="card">
        <strong class="text-nowrap lead">
            Session Types
        </strong>
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="card-body">
            <div class="list-group tuitiontypes">
                <a href="#" class="list-group-item" v-for="sessiontype in sessiontypes"
                    :class="{'active': values.includes(sessiontype.Title)}" @click.prevent="select(sessiontype.Title)">
                    {{sessiontype.Title}}
                </a>
            </div>
        </divloading>
    </div>
`
    }
