<script>
    export default {
        computed: {
            ...Vuex.mapState(['recent_posts']),
            ...Vuex.mapGetters(['getFullDate']),
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="widget widget-posts">
        <h2 class="widget-title">postS LIST</h2>
        <ul class="recent-posts clearfix">
            <li v-for="post in recent_posts.slice(0, 5)">
                <div class="thumb item-thumbnail">
                    <router-link :to="'/blog/'+post.ID">
                        <img :src="post.Image" alt="image">
                        <div class="thumbnail-hoverlay main-color-1-bg"></div>
                        <div class="thumbnail-hoverlay-cross"></div>
                    </router-link>
                </div>
                <div class="text">
                    <router-link :to="'/blog/'+post.ID">
                        {{post.Title}}
                    </router-link>
                    <p>{{getFullDate(post.UpdatedAt)}}</p>
                </div>
            </li>
        </ul><!-- /popular-news clearfix -->
    </div>
</template>