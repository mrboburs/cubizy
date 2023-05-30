<script>
    export default {
        props: {
            selected_user_id: {
                default: 0,
            },
        },
        components: {
            'users': () => import("/vue/chat/users.js"),
            'chathistory': () => import("/vue/chat/chathistory.js"),
        },
        data: () => {
            return {}
        },
        mounted: function () {
            if (window.innerWidth >= 993) {
                setTimeout(() => {
                    document.body.setAttribute('data-sidebar-size', 'condensed');
                }, 30);
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="row support_page">
        <!-- start chat users-->
        <users :class="{'col-xl-3 col-lg-4': selected_user_id}" :selected_user_id="selected_user_id"/>
        <!-- end chat users-->
        <!-- chat area -->
        <chathistory class="col-xl-9 col-lg-8" v-if="selected_user_id" :selected_chat_id="selected_user_id"/>
        <!-- end chat area-->
    </div>
</template>