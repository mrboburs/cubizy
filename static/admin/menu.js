export default [
    {
        label: "MENU",
        items: [
            {
                icon: "mdi mdi-view-grid-outline",
                label: "Dashbord",
                permission: "",
                to: "/"
            },
            {
                icon: "fas fa-store",
                label: "Accounts",
                items: [
                    {
                        label: "Active",
                        permission: "",
                        to: "/accounts",
                    },
                    {
                        label: "Pending",
                        permission: "",
                        to: "/accounts/pending"
                    },
                    {
                        label: "Expired",
                        permission: "",
                        to: "/accounts/expired"
                    },
                    {
                        label: "All",
                        permission: "",
                        to: "/accounts/all"
                    },
                ]
            },
            {
                icon: "ri-dashboard-line",
                label: "Categories",
                to: "/categories",
            },
            {
                icon: "fas fa-gifts",
                label: "Products",
                to: "/products"
            },
            {
                icon: "fas fa-file-invoice-dollar",
                label: "Orders",
                permission: "",
                to: "/orders"
            },
            {
                icon: "fas fa-star",
                label: "Reviews",
                permission: "",
                to: "/reviews"
            },
            {
                icon: "mdi mdi-blogger",
                label: "Blogs",
                permission: "",
                items: [
                    {
                        icon: "mdi mdi-blogger",
                        label: "All",
                        to: "/blogs/all"
                    },
                    {
                        icon: "mdi mdi-shape-outline",
                        label: "Blog Categories",
                        to: "/blogs/blogcategories"
                    },
                ]
            },
            {
                icon: "fas fa-hands-helping", // mdi mdi-view-handshake
                label: "Support",
                permission: "support",
                to: "/support"
            },
            {
                icon: "mdi mdi-web",
                label: "Website",
                items: [
                    {
                        icon: "mdi mdi-cog",
                        label: "Settings",
                        permission: "",
                        to: "/Website/settings"
                    },
                    {
                        icon: "mdi mdi-application",
                        label: "Pages",
                        permission: "",
                        to: "/Website/pages"
                    },
                    {
                        icon: "mdi mdi-palette",
                        label: "Website Editor",
                        permission: "",
                        to: "/Website/website_editor"
                    },
                    {
                        icon: "mdi mdi-palette-swatch",
                        label: "Themes market",
                        permission: "",
                        to: "/Website/themes"
                    },
                ]
            },
            {
                icon: "mdi mdi-account-group",
                label: "Users",
                items: [
                    {
                        icon: "mdi mdi-account-group",
                        label: "All Users",
                        permission: "",
                        to: "/users"
                    },
                    {
                        icon: "mdi mdi-account-group",
                        label: "Admins",
                        permission: "",
                        to: "/users/admins"
                    },
                    {
                        icon: "mdi mdi-account-group",
                        label: "Sellers",
                        permission: "",
                        to: "/users/sellers"
                    },
                ]
            },
            {
                icon: "mdi mdi-cog",
                label: "Application",
                items: [
                    {
                        icon: "mdi mdi-cog",
                        label: "Settings",
                        permission: "",
                        to: "/application/settings"
                    },
                    {
                        icon: "mdi mdi-account",
                        label: "Details",
                        to: "/application"
                    },
                    {
                        icon: "mdi mdi-image",
                        label: "Login Page Sliders",
                        permission: "",
                        to: "/application/loginpagesliders"
                    },
                    {
                        icon: "mdi mdi-help",
                        label: "Questions",
                        permission: "",
                        to: "/application/questions"
                    },
                    {
                        icon: "mdi mdi-map-marker",
                        label: "Locations",
                        permission: "",
                        to: "/application/locations"
                    },
                    {
                        icon: "mdi mdi-table",
                        label: "Session Types",
                        permission: "",
                        to: "/application/sessiontypes"
                    },
                ]
            },
            {
                icon: "mdi mdi-shape-outline",
                label: "Published Themes",
                to: "/themes",
            },
            {
                icon: "mdi mdi-repeat",
                label: "Transactions",
                items: [
                    {
                        icon: "mdi mdi-wallet",
                        label: "wallet",
                        permission: "",
                        to: "/transactions/wallet"
                    },
                    {
                        icon: "mdi mdi-receipt",
                        label: "orders",
                        to: "/transactions/orders"
                    },
                ]
            },
        ],
    },
]