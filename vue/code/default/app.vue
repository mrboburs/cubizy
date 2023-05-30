<script>

    import routes from "/vue/routes.js"
    import _store from "/vue/store.js"
    Vue.directive('tooltip', function (el, binding) {
        new bootstrap.Tooltip(el,{
            title: binding.value,
            placement: binding.arg,
            trigger: 'hover'
        })
    })
    const store = new Vuex.Store(_store)
    window.MyVueStore = store
    const router = new VueRouter({
        mode: 'history',
        routes // short for `routes: routes`
    })
    new Vue({
        el: '#app',
        router,
        store,
        data: {
            message: 'Hello To admin panel',
        },
    })

</script>
<template>

</template>