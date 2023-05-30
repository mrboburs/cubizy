export default [
    {
        label: "MENU",
        items: [
            {
                icon: "mdi mdi-view-grid-outline",
                label: "Dashbord",
                permission: "",
                to: "/student"
            },
            {
                icon: "mdi mdi-information-outline",
                label: "Info"
            },
            {
                icon: "mdi mdi-cog",
                label: "Administrator",
                items: [
                    {
                        icon: "mdi mdi-account",
                        label: "Users",
                        to: "/student/users"
                    },
                    {
                        icon: "mdi mdi-cellphone-link",
                        label: "Devices",
                        to: "/student/admins"
                    },
                    {
                        icon: "mdi mdi-account",
                        label: "My Account",
                    },
                    {
                        icon: "mdi mdi-link",
                        label: "Expo",
                        link: "/expo",
                    },
                ]
            },
            {
                icon: "mdi mdi-cog",
                label: "Admin",
                items: [
                    {
                        icon: "mdi mdi-account",
                        label: "Users",
                        to: "/"
                    },
                    {
                        icon: "mdi mdi-cellphone-link",
                        label: "Devices",
                        to: "/admins"
                    },
                    {
                        icon: "mdi mdi-account",
                        label: "My Account",
                    },
                    {
                        icon: "mdi mdi-link",
                        label: "Expo",
                        link: "/expo",
                    },
                ]
            }
        ],
    },
    {
        label: "User",
        items: [
            {
                icon: "mdi mdi-logout",
                label: "Logout",
                action: "logout"
            },
        ]
    }
]