
    import AsyncComponent from "/vue/async_component.js"
    import _store from "/vue/store.js"
    function validateTuterAccount(to, from, next) {
        if (_store.state.account.Status < 9) {
            if (window.MyVueStore) {
                window.MyVueStore.commit("set_last_message", {
                    Message: "Account must be active to access " + to.name,
                    Status: 1
                })
            }
            next({ name: 'dashboard' })
        } else next()
    }
    export default [
        {
            name: "test",
            path: '/test',
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
                    name: "register",
                    path: 'register',
                    component: () => AsyncComponent("/vue/auth/register.js"),
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
                else if (_store.state.user &&
                    _store.state.user.EmailVerified &&
                    _store.state.user.MobileVerified &&
                    _store.state.user.Question1 &&
                    _store.state.user.Answer1 &&
                    _store.state.user.Question2 &&
                    _store.state.user.Answer2 &&
                    _store.state.user.Question3 &&
                    _store.state.user.Answer3 &&
                    _store.state.account &&
                    _store.state.account.AddressID &&
                    _store.state.account.IDProof &&
                    _store.state.account.AddressProof &&
                    _store.state.account.RegistretionProof &&
                    _store.state.account.Status) {
                    next({ name: 'dashboard' })
                } else {
                    next()
                }
            },
            children: [
                {
                    name: "verifyemail",
                    path: 'verifyemail',
                    component: () => AsyncComponent("/vue/setup/verifyemail.js"),
                },
                {
                    name: "verifymobile",
                    path: 'verifymobile',
                    component: () => AsyncComponent("/vue/setup/verifymobile.js"),
                },
                {
                    name: "setquestionanswers",
                    path: 'setquestionanswers',
                    component: () => AsyncComponent("/vue/setup/setquestionanswers.js"),
                },
                {
                    name: "creataccount",
                    path: 'account',
                    component: () => AsyncComponent("/vue/account/account.js"),
                },
                {
                    name: "setlocation",
                    path: 'location',
                    component: () => AsyncComponent("/vue/account/location.js"),
                },
                {
                    name: "setdocuments",
                    path: 'documents',
                    component: () => AsyncComponent("/vue/account/documents.js"),
                },
                {
                    name: "welcome",
                    path: 'welcome',
                    component: () => AsyncComponent("/vue/setup/welcome.js"),
                },
            ]
        },
        {
            path: '/',
            component: () => AsyncComponent("/vue/mintonindex.js"),
            beforeEnter: (to, from, next) => {
                if (!_store.state.user) next({ name: 'login' })
                else if (!_store.state.user.EmailVerified) next({ name: 'verifyemail' })
                else if (!_store.state.user.MobileVerified) next({ name: 'verifymobile' })
                else if (!_store.state.user.Question1 || !_store.state.user.Answer1 || !_store.state.user.Question2 || !_store.state.user.Answer2 || !_store.state.user.Question3 || !_store.state.user.Answer3) next({ name: 'setquestionanswers' })
                else if (_store.state.account && _store.state.account.Status == 0 && _store.state.user.ID != _store.state.account.CreatedBy) next({ name: 'welcome' })
                else if (!_store.state.account || _store.state.account.Status == 0) next({ name: 'creataccount' })
                else if (!_store.state.account || _store.state.account.AddressID == 0) next({ name: 'setlocation' })
                else if (!_store.state.account || _store.state.account.IDProof == "" || _store.state.account.AddressProof == "" || _store.state.account.RegistretionProof == "") next({ name: 'setdocuments' })
                else if (!_store.state.user.Joined) next({ name: 'welcome' })
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
                    name: "dashboard",
                    path: '/',
                    meta: { title: "Seller dashboard" },
                    component: () => AsyncComponent("/vue/dashboard.js"),
                },
                {
                    name: "categories",
                    path: 'categories',
                    component: () => AsyncComponent("/vue/account/categorylist.js"),
                },
                {
                    name: "addresses",
                    path: 'addresses',
                    beforeEnter: validateTuterAccount,
                    component: () => AsyncComponent("/vue/addresslist.js"),
                },
                {
                    name: "me",
                    path: 'me',
                    component: () => AsyncComponent("/vue/me.js"),
                },
                {
                    path: 'account',
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "account",
                            path: '/',
                            component: () => AsyncComponent("/vue/account/account.js"),
                        },
                        {
                            name: "location",
                            path: 'location',
                            component: () => AsyncComponent("/vue/account/location.js"),
                        },
                        {
                            name: "documents",
                            path: 'documents',
                            component: () => AsyncComponent("/vue/account/documents.js"),
                        },
                    ],
                },
                {
                    path: 'products',
                    meta: { title: "Products" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "allproducts",
                            path: '/',
                            meta: { title: "All Products" },
                            component: () => AsyncComponent("/vue/products/list.js"),
                        },
                        {
                            name: "addproducts",
                            path: 'add',
                            meta: { title: "Add Product" },
                            component: () => AsyncComponent("/vue/products/add.js"),
                        },
                    ],
                }, // 
                {
                    name: "orders",
                    path: 'orders',
                    meta: { title: "Orders" },
                    component: () => AsyncComponent("/vue/products/orders.js"),
                },
                {
                    name: "reviews",
                    path: 'reviews',
                    meta: { title: "reviews" },
                    component: () => AsyncComponent("/vue/reviews.js"),
                },
                {
                    path: 'services',
                    meta: { title: "Services" },
                    component: () => AsyncComponent("/vue/page.js"),
                    children: [
                        {
                            name: "allservices",
                            path: '/',
                            meta: { title: "All Services" },
                            component: () => AsyncComponent("/vue/services/list.js"),
                        },
                        {
                            name: "addservices",
                            path: 'add',
                            meta: { title: "Add Service" },
                            component: () => AsyncComponent("/vue/services/add.js"),
                        },
                    ],
                },
                {
                    name: "website",
                    path: 'website',
                    beforeEnter: validateTuterAccount,
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
                    name: "notfounduserpage",
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
