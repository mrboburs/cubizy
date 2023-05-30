
    export default {
        props: {
            replay_to: {
                default: 0,
            },
            chat: {
                default: function () {
                    return {
                        ID: 0,
                        Name: "",
                        Photo: "",
                        ChatID: 0,
                        CreatedFor: 0,
                        Messages: []
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
                ChatMessage: {
                    ID: 0,
                    ChatID: 0,
                    Content: "",
                    Type: 0,
                    RecivedAt: 0,
                    SeenAt: 0,
                    ReplayTo: 0,
                    UpdatedBy: 0,
                    CreatedBy: 0,
                },
                sender: {
                    ID: 0,
                    Name: "",
                    Photo: "",
                }
            }
        },
        computed: {
            ...Vuex.mapState(['user']),
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.chat_message = false
                    this.submitted = false
                }
            },
            ChatMessage: function (newValue, oldValue) {
                if (this.ChatMessage.CreatedBy == this.user.ID) {
                    this.sender = this.user
                } else if (this.chat.CreatedFor != 0 && this.ChatMessage.CreatedBy == this.chat.ID) {
                    this.sender = this.chat
                }
            },
        },
        mounted: function () {
            this.init()
        },
        methods: {
            init() {
                if (!this.replay_to || !this.chat || !this.chat.Messages) {
                    return
                }
                var messages = this.chat.Messages.filter(message => message.ID == this.replay_to);
                if (messages.length) {
                    this.ChatMessage = messages[0]
                } else {
                    this.load({
                        MessageID: this.replay_to
                    })
                }
            },
            load: _.debounce(function (data) {
                if (this.loading || !this.replay_to) {
                    return
                }
                if (!data) {
                    data = {}
                }
                if (this.chat.ChatID > 0) {
                    data.ChatID = this.chat.ChatID
                } else {
                    return
                }

                this.loading = true
                this.$store.dispatch('call', {
                    api: "chathistory",
                    data: data
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if (data.Result.chathistory && Array.isArray(data.Result.chathistory)) {
                            this.ChatMessage = data.Result.chathistory[0]
                        }
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
            }, 300),
        },
        template: `
    <div v-if="ChatMessage.ID > 0" class="replay-to-message-wrap">
        <i>Replayed To : {{sender.Name}}'s message : </i>
        <p class="text-end">{{ChatMessage.Content}}</p>
        <v-alert v-model="message" :error="error" :compact="true" />
    </div>
`
    }
    // v-for="message in selected_user.Messages"
