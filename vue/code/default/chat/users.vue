<script>
    export default {
        props: {
            selected_user_id: {
                default: 0,
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                popupitem: false,
                search: "",
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
            ...Vuex.mapState(['user', 'account', 'usersinchat', 'conn']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
            users: function () {
                var users = []
                if (this.usersinchat) {
                    if (this.search) {
                        users = this.usersinchat.filter(user => user.Name.toLowerCase().includes(this.search.toLowerCase()));
                    } else {
                        users = this.usersinchat
                    }
                }
                users = users.sort((a, b) => { return a.CreatedAt - b.CreatedAt; });
                return users
            },
        },
        methods: {
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "chatlist",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_chats', data.data)
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
            unseen_count(userinchat) {
                var count = 0
                if (!userinchat.Seen) {
                    count = 1
                }
                if (userinchat.Messages && userinchat.Messages.length) {
                    var unseen_messages = userinchat.Messages.filter(message => message.Seen);
                }
            },
        },
        mounted: function () {
            this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="card d-md-block p-0 mb-0" :class="{'d-lg-block d-md-none d-none': selected_user_id}">
        <divloading :fullpage="false" :loading="loading" class="card-body">
            <v-alert v-model="message" :error="error" />
            <div v-if="conn" class="alert alert-primary" role="alert">
                Online
            </div>
            <div v-else class="alert alert-secondary" role="alert">
                Offline
            </div>
            <!-- start search box -->
            <div class="search-bar mb-3">
                <div class="position-relative">
                    <input type="text" class="form-control form-control-light" placeholder="Supportagents"
                        v-model="search">
                    <span class="mdi mdi-magnify"></span>
                </div>
            </div>
            <!-- end search box -->

            <h6 v-if="user.IsSupportagent" class="font-13 text-muted text-uppercase mb-2">Connected</h6>

            <h6 v-if="!user.IsSupportagent" class="font-13 text-muted text-uppercase mb-2">Supportagents</h6>

            <!-- users -->
            <div class="row supportagents_list">
                <div class="col" style="max-height: 498px">
                    <router-link v-for="userinchat in users" :key="'user_'+userinchat.ID"
                        :to="'/support/'+ userinchat.ID" class="text-body">
                        <div class="d-flex align-items-start p-2">
                            <div class="position-relative">
                                <span v-if="userinchat.Online" class="user-status online"></span>
                                <img :src="userinchat.Photo" class="me-2 rounded-circle user_icon" alt="user">
                            </div>
                            <div class="flex-1">
                                <h5 class="mt-0 mb-0 font-14">
                                    <span
                                        class="float-end text-muted fw-normal font-12">{{getTime(userinchat.UpdatedAt)}}</span>
                                    {{userinchat.Name}}
                                </h5>
                                <p class="mt-1 mb-0 text-muted font-14">
                                    <span class="w-25 float-end text-end"><span
                                            class="badge badge-soft-danger">{{unseen_count(userinchat)}}</span></span>
                                    <span class="w-75">{{userinchat.Message}}</span>
                                </p>
                            </div>
                        </div>
                    </router-link>
                </div> <!-- End col -->
            </div>
            <!-- end users -->
        </divloading> <!-- end card-body-->
        <!-- end card-->
    </div>
</template>