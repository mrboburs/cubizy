
    export default {
        props: {
            selected_chat_id: {
                default: 0,
            },
        },
        components: {
            'message': () => import("/vue/chat/message.js"),
            'rtmessage': () => import("/vue/chat/rtmessage.js"),
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                showchat : true,
                ChatMessages : [],
                ChatMessage : {
                    ID : 0,
	                ChatID : 0,
                    Content : "",
                    Type : 0,
                    RecivedAt : 0,
                    SeenAt : 0,
                    ReplayTo : 0,
                    UpdatedBy : 0,
                    CreatedBy : 0,
                },
                newChatMessage : {
                    ID : 0,
	                ChatID : 0,
                    Content : "",
                    Type : 0,
                    RecivedAt : 0,
                    SeenAt : 0,
                    ReplayTo : 0,
                    UpdatedBy : 0,
                    CreatedBy : 0,
                },
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
            selected_chat : function(newValue, oldValue) {
                if(oldValue && oldValue.ID && newValue && newValue.ID && oldValue.ID == newValue.ID ){
                    return
                }
                this.init()
            },
            "selected_chat.Messages" : function(newValue, oldValue){
                this.go_to_end()
            },
            letest_message: function(newValue, oldValue) {
                this.selected_chat.Messages = this.selected_chat.Messages.sort((a, b) => { return a.ID - b.ID; });
                this.ChatMessages = this.selected_chat.Messages
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account', 'usersinchat']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
            selected_chat: function () {
                var users = []
                if (this.selected_chat_id > 0 && this.usersinchat){
                        users = this.usersinchat.filter(user => user.ID == (this.selected_chat_id));
                }
                if (users.length > 0) {
                    var user = users[0]
                    if (!user.Messages) {
                        user.Messages = []
                    }
                    return user
                }
            },
        },
        methods: {
            go_to_end (){
                setTimeout(() => {
                    if(this.$refs.conversation){
                        this.$refs.conversation.scrollTo(0, this.$refs.conversation.scrollHeight, 'smooth');
                    }
                }, 100);
            },
            load: _.debounce(function (data) {
                if (this.loading || ! this.selected_chat_id || ! this.selected_chat) {
                    return
                }
                if(!data){
                    data = {}
                }
                if(this.selected_chat.ChatID > 0){
                    data.ChatID = this.selected_chat.ChatID
                    if(this.selected_chat.Messages.length > 0 && !data.BeforID){
                        data.AfterID =  this.selected_chat.MessageID
                    }
                }else if (this.selected_chat.CreatedFor > 0 && this.selected_chat.ID > 0){
                    data.CreatedFor = this.selected_chat.ID
                }

                this.loading = true
                this.$store.dispatch('call', {
                    api: "chathistory",
                    data: data
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        if(data.Result.chat){
                            if(this.selected_chat.ID == data.Result.chat.CreatedFor){
                                this.selected_chat.ChatID = data.Result.chat.ID
                                this.selected_chat.CreatedFor = data.Result.chat.CreatedFor
                                this.selected_chat.CreatedAt = data.Result.chat.CreatedAt
                                this.selected_chat.UpdatedAt = data.Result.chat.UpdatedAt
                                this.selected_chat.DeletedAt = data.Result.chat.DeletedAt
                            }
                        }
                        if(data.Result.chathistory && Array.isArray(data.Result.chathistory)){
                                this.selected_chat.Messages = this.selected_chat.Messages.concat(data.Result.chathistory)
                                this.selected_chat.Messages = this.selected_chat.Messages.sort((a, b) => { return a.ID - b.ID; });
                                this.ChatMessages = this.selected_chat.Messages
                                this.go_to_end()
                        }
                        this.ChatMessage = JSON.parse(JSON.stringify(this.newChatMessage))
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
            }, 200),
            init: _.debounce(function (data) {
                console.log("time")
                console.log(Date.now());
                if(!this.selected_chat){
                    return
                }
                if (!this.selected_chat.Messages) {
                    this.selected_chat.Messages = []
                }
                this.ChatMessages = this.selected_chat.Messages
                this.newChatMessage.ChatID = this.selected_chat.ChatID
                this.ChatMessage = JSON.parse(JSON.stringify(this.newChatMessage))
                if(!this.selected_chat.Messages.length){
                    this.load()
                }
            }, 200),
            textareasubmit(e){
                if(e.which === 13 && !e.shiftKey) {
                    e.preventDefault();
                    this.sendMessage()
                }else{
                    e.target.style.height = (e.target.rows * 38) + "px"
                }
            },
            sendMessage(){
                if(this.loading || !this.ChatMessage.Content){
                    return
                }
                if(this.loading){
                    return
                }
                this.load(this.ChatMessage)
            },
            delete_message(message_id){
                this.load({
                    delete_message : message_id
                })
            },
            edit(message){
                this.ChatMessage = message 
            },
            copy(Content){
                this.ChatMessage = JSON.parse(JSON.stringify(this.newChatMessage))
                this.ChatMessage.Content = Content 
            },
            replay(ReplayTo){
                if(this.ChatMessage.id > 0){
                    this.ChatMessage = JSON.parse(JSON.stringify(this.newChatMessage))
                }
                this.ChatMessage.ReplayTo = ReplayTo 
            }
        },
        mounted: function () {
            this.init()
        },
        template: `
    <div class=" p-0 mb-0 card" v-if="selected_chat">
        <div class="card-header py-2 px-3 border-bottom border-light bg-white">
            <div class="d-flex py-1">
                <button class="d-lg-none btn btn-outline-primary rounded-circle me-1"
                    @click.prevent="selected_chat = false"> <i class="fas fa-arrow-left"></i></button>
                <img :src="selected_chat.Photo" class="me-2 rounded-circle user_icon" alt="Brandon Smith">
                <div class="flex-1">
                    <h5 class="mt-0 mb-0 font-15">
                        <a target="blank" href="#" class="text-reset">{{selected_chat.Name}}</a>
                    </h5>
                    <p class="mt-1 mb-0 text-muted font-12" v-if="selected_chat.Online">
                        <small class="mdi mdi-circle text-success"></small> Online
                    </p>
                    <v-alert v-model="message" :error="error" />
                </div>
                <div id="tooltip-container">
                    <a href="javascript: void(0);" class="text-reset font-19 py-1 px-2 d-inline-block">
                        <i class="fe-phone-call" data-bs-container="#tooltip-container" data-bs-toggle="tooltip"
                            data-bs-placement="top" title="" data-bs-original-title="Voice Call"
                            aria-label="Voice Call"></i>
                    </a>
                    <a href="javascript: void(0);" class="text-reset font-19 py-1 px-2 d-inline-block">
                        <i class="fe-video" data-bs-container="#tooltip-container" data-bs-toggle="tooltip"
                            data-bs-placement="top" title="" data-bs-original-title="Video Call"
                            aria-label="Video Call"></i>
                    </a>
                    <a href="javascript: void(0);" class="text-reset font-19 py-1 px-2 d-inline-block">
                        <i class="fe-user-plus" data-bs-container="#tooltip-container" data-bs-toggle="tooltip"
                            data-bs-placement="top" title="" data-bs-original-title="Add Users"
                            aria-label="Add Users"></i>
                    </a>
                    <a href="javascript: void(0);" class="text-reset font-19 py-1 px-2 d-inline-block">
                        <i class="fe-trash-2" data-bs-container="#tooltip-container" data-bs-toggle="tooltip"
                            data-bs-placement="top" title="" data-bs-original-title="Delete Chat"
                            aria-label="Delete Chat"></i>
                    </a>
                </div>
            </div>
        </div>
        <div ref="conversation" class="card-body conversation-list-box" v-if="selected_chat.ChatID > 0 && showchat">
            <ul class="conversation-list chat-app-conversation w-100 overflow-auto">
                    <message v-for="message in ChatMessages" :chat_message="message" :key="'message'+ message.ID"  :chat="selected_chat" 
                    @copy="copy($event)" @replay="replay($event)" @edit="edit($event)" @delete_message="delete_message($event)"/>
            </ul>
        </div>
        <divloading v-if="selected_chat.ChatID > 0" :fullpage="false" :loading="loading" class="card-footer bg-light p-3 rounded">
            <v-alert v-model="message" :error="error" />
            <form class="needs-validation" novalidate="" name="chat-form" id="chat-form"
                @submit.prevent="sendMessage">
                <div class="row">
                    <div class="col mb-2 mb-sm-0">
                        <div v-if="ChatMessage.ID > 0 " class="clearfix d-flex align-items-start">
                            <span   class="text-muted flex-fill"> Editing message</span>
                            <button type="button" class="btn-close" aria-label="Close" title="Cancel Edit"  @click.prevent="copy('')"></button>
                        </div>
                        
                        <div v-if="ChatMessage.ReplayTo"  class="d-flex">
                            <rtmessage :replay_to="ChatMessage.ReplayTo" :chat="selected_chat"/>
                            <button v-if="ChatMessage.ID == 0" type="button" class="btn-close" aria-label="Close" title="Cancel Replay"  @click.prevent="copy('')"></button>
                        </div>
                        <textarea type="text" class="form-control border-0 position-relative" @keypress="textareasubmit"
                            placeholder="Enter your text" required="" v-model="ChatMessage.Content"></textarea>
                        <div class="invalid-feedback mt-2">
                            Please enter your messsage
                        </div>
                    </div>
                    <div class="col-sm-auto">
                        <div class="btn-group">
                            <a href="#" class="btn btn-light"><i class="fe-paperclip"></i></a>
                            <div class="d-grid">
                                <button type="submit" class="btn btn-success chat-send"><i
                                        class="fe-send"></i></button>
                            </div>
                        </div>
                    </div> <!-- end col -->
                </div> <!-- end row-->
            </form>
        </divloading>
    </div>
`
    }
