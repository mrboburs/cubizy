
    export default {
        props: {
            chat_message: {
                default: function () {
                    return { 
                        ID : 0,
                        ChatID : 0,
                        Content : "",
                        Type : 0,
                        RecivedAt : 0,
                        SeenAt : 0,
                        ReplayTo : 0,
                        UpdatedBy : 0,
                        CreatedBy : 0,
                    }
                }
            },
            chat: {
                default: function () {
                    return { 
                        ID : 0,
                        Name : "",
                        Photo : "",
                        ChatID : 0,
                        CreatedFor : 0,
                        Messages : []
                    }
                }
            },
        },
        components: {
            'rtmessage': () => import("/vue/chat/rtmessage.js"),
        },
        data: () => {
            return {
                sender : {
                    ID : 0,
                    Name : "",
                    Photo : "",
                }
            }
        },
        computed: {
            ...Vuex.mapState(['user']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.chat_message = false
                    this.submitted = false
                }
            },
            chat_message : function(newValue, oldValue) {
                this.init()
            },
        },
        mounted: function () {
            this.init()
        },
        methods: {
            init(){
                if(!this.chat_message){
                    return
                }
                
                if (this.chat_message.CreatedBy == this.user.ID) {
                    this.sender = this.user
                }else if (this.chat.CreatedFor != 0 && this.chat_message.CreatedBy == this.chat.ID) {
                    this.sender = this.chat
                }
            },
            copy(){
                this.$emit('copy', this.chat_message.Content)
            },
            replay(){
                this.$emit('replay', this.chat_message.ID)
            },
            edit(){
                this.$emit('edit', this.chat_message)
            },
            delete_message(){
                this.$emit('delete_message', this.chat_message.ID)
            }
        },
        template: `
    <li class="clearfix" :class="{'odd': chat_message.CreatedBy == user.ID }">
        <div class="chat-avatar">
            <img :src="sender.Photo" class="rounded" alt="James Z">
            <i>{{getTime(chat_message.CreatedAt)}}</i>
        </div>
        <div class="conversation-text" v-if="!chat_message.DeletedAt">
            <rtmessage :replay_to="chat_message.ReplayTo" :chat="chat"/>
            <div class="ctext-wrap" :class="{'replay-message-wrap':chat_message.ReplayTo > 0}">
                <i>{{sender.Name}}</i>
                <p>{{chat_message.Content}}</p>
            </div>
        </div>
        <div class="conversation-text" v-if="chat_message.DeletedAt">
            <div class="ctext-wrap">
                <i>{{sender.Name}}</i>
                <p>Message deleted</p>
            </div>
        </div>
        <div class="conversation-actions"  v-if="!chat_message.DeletedAt">
            <a href="#" class="btn btn-outline-primary btn-sm" title="Reply Message" @click.prevent="replay()"><i class="fas fa-reply"></i></i></a>
            <a href="#" class="btn btn-outline-secondary btn-sm" title="Copy Message" v-if="chat_message.CreatedBy != user.ID" @click.prevent="copy()"><i class="far fa-copy"></i></a>
            <div class="dropdown" v-if="chat_message.CreatedBy == user.ID">
                <button data-bs-toggle="dropdown" aria-expanded="false"
                    class="btn btn-sm btn-outline-secondary">
                    <i class="fas fa-ellipsis-v"></i>
                </button>
                <div class="dropdown-menu">
                    <a href="#" class="dropdown-item" @click.prevent="copy()">Copy</a>
                    <a href="#" class="dropdown-item" @click.prevent="edit()">Edit</a>
                    <a href="#" class="dropdown-item" @click.prevent="delete_message()">Delete</a>
                </div>
            </div>
        </div>
    </li>
`
    }
    // v-for="message in selected_user.Messages"
