

    import routes from "/vue/theme2/routes.js"
    import _store from "/vue/store.js"
    const router = new VueRouter({
        mode: 'history',
        routes // short for `routes: routes`
    })
    const store = new Vuex.Store(_store)
    new Vue({
        el: '#app',
        router,
        store,
        data: {
            message: 'Hello To admin panel',
            BackgroundImage : application.ThemeSettings.BackgroundImage.value,
            RepeatBgi : application.ThemeSettings.RepeatBackgroundImage.value
        },
        components: {
            'topbar': () => import("/vue/topbar.js"),
            'vheader': () => import("/vue/header.js"),
            'vfooter': () => import("/vue/footer.js"),
        },
        watch: {
            $route(to, from) {
                setTimeout(() => {
                    scrollTo(0, 143, 'smooth')
                }, 600);
            },
        },
        mounted: function () {
            setTimeout(() => {
                scrollTo(0, 143, 'smooth')
            }, 600);
        },
    })

