<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
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
        components: {
            'UserEditor': () => import("/vue/usereditor.js"),
        },
        computed: {
            ...Vuex.mapState(['user']),
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
            }
        },
        methods: {
            submit(value) {
                if(!value){
                    return
                }
                var component = this
                component.loading = true
                setTimeout(() => {
                    this.$store.dispatch('call', {
                        api: "me",
                        data: {
                            user : value
                        }
                    }).then(function (data) {
                        component.message = data.Message;
                        if (data.Status == 2) {
                            component.error = false
                        } else {
                            component.error = true
                        }
                    }).catch((error) => {
                        console.error('Error:', error);
                        component.error = true
                        component.message = error
                    }).finally(() => {
                        component.loading = false
                    })
                }, 1000);
            },
        },
        mounted: function () {
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="col-12">
        <div class="card">
            <divloading class="card-body" :fullpage="false" :loading="loading">
                <v-alert v-model="message" :error="error"/>
                <UserEditor v-if="user" :value="user" @input="submit"></UserEditor>
            </divloading>
        </div>
    </div>
</template>