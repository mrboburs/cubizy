<script>
    import AsyncComponent from "/vue/async_component.js"
    import _store from "/vue/store.js"
    export default [
        {
            name: "courses",
            path: '/courses/:level?/:sublevel?/:subject?',
            component: () => AsyncComponent("/vue/courses.js"),
            props: true
        },
        {
            name: "course",
            path: '/course/:course_id',
            component: () => AsyncComponent("/vue/course.js"),
            props: true
        },
        {
            name: "sellers",
            path: '/sellers',
            component: () => AsyncComponent("/vue/sellers.js"),
            props: true
        },
        {
            name: "home",
            path: '/',
            component: () => AsyncComponent("/vue/theme1/home.js"),
            props: true
        },
        {
            name: "pages",
            path: '/pages/:page_id',
            component: () => AsyncComponent("/vue/pages.js"),
            props: true
        },
        {
            name: "notfound",
            path: '/*',
            component: () => AsyncComponent("/vue/notfound.js")
        },
    ]
</script>