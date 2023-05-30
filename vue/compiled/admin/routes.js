
    import AsyncComponent from "/vue/async_component.js"
    import _store from "/vue/store.js"
    export default [
        {
            name: "test",
            path: '/test/',
            component: () => AsyncComponent("/vue/test.js")
        },
        {
            path: '/auth',
            component: () => AsyncComponent("/vue/auth/auth.js"),
            beforeEnter: (to, from, next) => {
                if (_store.state.user) next({ name: 'dashboard' })
                else next()
            },
            children: [
                {
                    name: "login",
                    path: 'login',
                    component: () => AsyncComponent("/vue/auth/login.js"),
                },
                {
                    name: "recoverpw",
                    path: 'recoverpw',
                    component: () => AsyncComponent("/vue/auth/recoverpw.js"),
                },
            ]
        },
        {
            path: '/setup',
            component: () => AsyncComponent("/vue/setup/setup.js"),
            beforeEnter: (to, from, next) => {
                if (!_store.state.user) next({ name: 'login' })
                else if (_store.state.account) {
                    next({ name: 'dashboard' })
                } else {
                    next()
                }
            },
            children: [
                {
                    name: "creataccount",
                    path: 'account',
                    component: () => AsyncComponent("/vue/account/account.js"),
                },
            ]
        },
        {
            path: '/',
            meta: { title: "Admin Panel" },
            component: () => AsyncComponent("/vue/mintonindex.js"),
            beforeEnter: (to, from, next) => {
                if (!_store.state.user) next({ name: 'login' })
                if (!_store.state.account) next({ name: 'creataccount' })
                else next()
            },
            children: [
                {
                    name: "support",
                    path: '/support/:selected_user_id?',
                    meta: { title: "Support" },
                    component: () => AsyncComponent("/vue/support.js"),
                    props: true
                },
                {
                    path: '/',
                    meta: { title: "Admin Dashboard" },
                    component: () => AsyncComponent("/vue/dashboard.js"),
                },
                {
                    name: "me",
                    path: 'me',
                    component: () => AsyncComponent("/vue/me.js"),
                },
                {
                    name: "categories",
                    path: 'categories',
                    component: () => AsyncComponent("/vue/account/categorylist.js"),
                },
                {
                    path: 'products',
                    meta: { title: "Products" },
                    component: () => AsyncComponent("/vue/products/list.js"),
                },
                {
                    name: "orders",
                    path: 'orders',
                    meta: { title: "Orders" },
                    component: () => AsyncComponent("/vue/products/orders.js"),
                },
                {
                    name: "reviews",
                    path: 'reviews',
                    meta: { title: "Reviews" },
                    component: () => AsyncComponent("/vue/reviews.js"),
                },
                {
                    path: 'blogs',
                    meta: { title: "blogs" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "allblogs",
                            path: 'all',
                            component: () => AsyncComponent("/vue/bloglist.js"),
                        },
                        {
                            name: "blogcategories",
                            path: 'blogcategories',
                            meta: { title: "Blog Categories" },
                            component: () => AsyncComponent("/vue/blogcategorylist.js"),
                        },
                    ]
                },
                {
                    name: "blogs",
                    path: 'blogs',
                    component: () => AsyncComponent("/vue/bloglist.js"),
                },
                {
                    name: "events",
                    path: 'events',
                    component: () => AsyncComponent("/vue/eventlist.js"),
                },
                {
                    name: "eventsx24",
                    path: 'sql',
                    component: () => AsyncComponent("/vue/eventsx24.js"),
                },
                {
                    path: 'transactions',
                    meta: { title: "Transactions" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            path: '/',
                            component: () => AsyncComponent("/vue/transactions.js"),
                        },
                        {
                            name: "wallet_transactions",
                            path: 'wallet',
                            meta: { title: "Wallet Transactions" },
                            component: () => AsyncComponent("/vue/transactions.js"),
                        },
                        {
                            name: "order_transactions",
                            path: 'orders',
                            meta: { title: "Order Transactions" },
                            component: () => AsyncComponent("/vue/transactions.js"),
                        },
                    ]
                },
                {
                    path: 'accounts',
                    meta: { title: "Accounts" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "active",
                            path: '/',
                            component: () => AsyncComponent("/vue/accountlist.js"),
                        },
                        {
                            name: "pending",
                            path: 'pending',
                            component: () => AsyncComponent("/vue/accountlist.js"),
                        },
                        {
                            name: "expired",
                            path: 'expired',
                            component: () => AsyncComponent("/vue/accountlist.js"),
                        },
                        {
                            name: "all",
                            path: 'all',
                            component: () => AsyncComponent("/vue/accountlist.js"),
                        },
                    ],
                },
                {
                    path: 'users',
                    meta: { title: "Users" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "allusers",
                            path: '/',
                            component: () => AsyncComponent("/vue/userlist.js"),
                        },
                        {
                            name: "admins",
                            path: 'admins',
                            component: () => AsyncComponent("/vue/userlist.js"),
                        },
                        {
                            name: "sellers",
                            path: 'sellers',
                            component: () => AsyncComponent("/vue/userlist.js"),
                        },
                    ],
                },
                {
                    name: "website",
                    path: 'website',
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "websettings",
                            path: 'settings',
                            meta: { title: "website settings" },
                            component: () => AsyncComponent("/vue/website/website.js"),
                        },
                        {
                            name: "pages",
                            path: 'pages',
                            meta: { title: "website pages" },
                            component: () => AsyncComponent("/vue/website/pagelist.js"),
                        },
                        {
                            name: "website_editor",
                            path: 'website_editor',
                            meta: { title: "Website Editor" },
                            component: () => AsyncComponent("/vue/website/edittheme.js"),
                        },
                        {
                            name: "account_themes",
                            path: 'themes',
                            meta: { title: "Themes market" },
                            component: () => AsyncComponent("/vue/website/themes.js"),
                        },
                    ]
                },
                {
                    path: 'application',
                    meta: { title: "application" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "application",
                            path: '/',
                            meta: { title: "details" },
                            component: () => AsyncComponent("/vue/account/account.js"),
                        },
                        {
                            name: "settings",
                            path: 'settings',
                            component: () => AsyncComponent("/vue/settings.js"),
                        },
                        {
                            name: "loginpagesliders",
                            path: 'loginpagesliders',
                            meta: { title: "Login Page Sliders" },
                            component: () => AsyncComponent("/vue/loginpagesliderlist.js"),
                        },
                        {
                            name: "questions",
                            path: 'questions',
                            component: () => AsyncComponent("/vue/questionlist.js"),
                        },
                        {
                            name: "locations",
                            path: 'locations',
                            component: () => AsyncComponent("/vue/locationlist.js"),
                        },
                        {
                            name: "sessiontypes",
                            path: 'sessiontypes',
                            component: () => AsyncComponent("/vue/sessiontypelist.js"),
                        },
                    ]
                },
                {
                    path: 'themes',
                    meta: { title: "themes" },
                    component: () => AsyncComponent("/vue/themes.js"),
                },
                {
                    path: '/*',
                    component: () => AsyncComponent("/vue/notfound.js")
                },
            ]
        },
        {
            name: "notfound",
            path: '/*',
            component: () => AsyncComponent("/vue/notfound.js")
        },
    ]
