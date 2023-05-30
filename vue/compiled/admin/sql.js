
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                sql: "",
                output : false
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
            ...Vuex.mapState(['user', 'account']),
        },
        methods: {
            load(data) {
                if(this.loading){
                    return
                }
                if(!data){
                    data = {}
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "eventsx24",
                    data: data
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        debugger
                        this.output = data.data
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
            select(event){
                this.load({
                    select_sql : this.sql
                })
            },
            execute(event){
                this.load({
                    execute_sql : this.sql
                })
            }
        },
        mounted: function () {},
        template: `
    <section class="section is-medium">
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="container">
            <formitem name="inpuSql" label="SQL" v-model="sql" type="textarea" />
            <div class="d-flex align-items-center end m-2">
                <button type="button" class="btn btn-success m-1" @click.prevent="select">
                    Select
                </button>
                <button type="button" class="btn btn-success m-1" @click.prevent="execute">
                    Execute
                </button>
            </div>
            <div>
                {{output}}
            </div>
        </divloading>
    </section>
`
    }
