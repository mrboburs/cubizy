

    import routes from "/vue/admintheme1/routes.js"
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
        },
        components: {
            'topbar': () => import("/vue/topbar.js"),
            'vheader': () => import("/vue/admintheme1/header.js"),
            'vfooter': () => import("/vue/footer.js"),
        },
        watch: {
            $route(to, from) {
                setTimeout(() => {
                    scrollTo(0,0)
                }, 300);
            },
        },
    })

