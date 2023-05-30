
    import AsyncComponent from "/vue/async_component.js"
    import _store from "/vue/store.js"
    export default [
        {
            name: "blogs",
            path: '/blogs',
            component: () => AsyncComponent("/vue/blogs.js"),
            props: true
        },
        {
            name: "blog",
            path: '/blog/:blog_id',
            component: () => AsyncComponent("/vue/blog.js"),
            props: true
        },
        {
            name: "events",
            path: '/events',
            component: () => AsyncComponent("/vue/events.js"),
            props: true
        },
        {
            name: "event",
            path: '/event/:event_id',
            component: () => AsyncComponent("/vue/event.js"),
            props: true
        },
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
            name: "account",
            path: '/account/:account_id',
            component: () => AsyncComponent("/vue/account.js"),
            props: true
        },
        {
            name: "pages",
            path: '/pages/:page_id',
            component: () => AsyncComponent("/vue/pages.js"),
            props: true
        },
        {
            name: "home",
            path: '/',
            component: () => AsyncComponent("/vue/admintheme1/home.js")
        },
        {
            name: "notfound",
            path: '/*',
            component: () => AsyncComponent("/vue/notfound.js")
        },
    ]
