export default [
    {
        label: "MENU",
        items: [
            {
                icon: "fas fa-hands-helping", // mdi mdi-view-handshake
                label: "Support",
                to: "/support"
            },
            {
                icon: "mdi mdi-view-grid-outline",
                label: "Dashbord",
                to: "/"
            },
            {
                icon: "mdi mdi-book",
                label: "Courses",
                to: "/sessions"
            },
            {
                icon: "far fa-bell",
                label: "Notices",
                to: "/notices"
            },
            {
                icon: "mdi mdi-application",
                label: "Notes",
                to: "/notes"
            },
        ],
    },
    {
        label: "Account",
        items: [
            {
                icon: "mdi mdi-account",
                label: "Profile",
                to: "/me"
            },
        ]
    }
]