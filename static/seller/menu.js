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
                icon: "ri-dashboard-line",
                label: "Categories",
                to: "/categories",
            },
            {
                icon: "fas fa-gifts",
                label: "Products Hub",
                items: [
                    {
                        icon: "fas fa-gift",
                        label: "Add Product",
                        permission: "",
                        to: "/products/add"
                    },
                    {
                        icon: "fas fa-gifts",
                        label: "Products",
                        permission: "",
                        to: "/products"
                    },
                    {
                        icon: "fas fa-file-invoice-dollar",
                        label: "Orders",
                        permission: "",
                        to: "/orders"
                    },
                ]
            },
            {
                icon: "fas fa-toolbox",
                label: "Services Hub",
                items: [
                    {
                        icon: "fas fa-tool",
                        label: "Add Service",
                        permission: "",
                        to: "/services/add"
                    },
                    {
                        icon: "fas fa-tools",
                        label: "Services",
                        permission: "",
                        to: "/services"
                    },
                    {
                        icon: "fas fa-file-invoice-dollar",
                        label: "Billing",
                        permission: "",
                        to: "/billing"
                    },
                ]
            },
            {
                icon: "fas fa-star",
                label: "Reviews",
                permission: "",
                to: "/reviews"
            },
            {
                icon: "fas fa-hands-helping", // mdi mdi-view-handshake
                label: "Support",
                permission: "support",
                to: "/support"
            },
            {
                icon: "mdi mdi-map-marker",
                label: "Locations",
                permission: "",
                to: "/addresses"
            },
            {
                icon: "mdi mdi-blogger",
                label: "Blogs",
                permission: "",
                to: "/blogs"
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
        ],
    },
    {
        label: "Account",
        items: [
            {
                icon: "mdi mdi-account",
                label: "Account",
                to: "/account"
            },
            {
                icon: "mdi mdi-cog",
                label: "Location",
                to: "/account/location"
            },
            {
                icon: "mdi mdi-cog",
                label: "Documents",
                to: "/account/documents"
            },
        ]
    }
]