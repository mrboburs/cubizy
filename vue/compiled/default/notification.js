
    export default {
        props: {
            notification: {
                type : Object
            },
        },
        data: () => {
            return {
                started : moment(),
                time_ago : "Just Now",
            }
        },
        mounted: function () {
            this.loop()
        },
        methods: {
            loop() {
                setInterval(() => {
                    this.time_ago = this.started.fromNow()
                }, 6000);
            }
        },
        template: `
    <router-link :to="notification.Url" class="dropdown-item notify-item active">
        <div class="notify-icon">
            <img :src="notification.Image" class="img-fluid rounded-circle"
                alt="notification.ID" />
        </div>
        <p class="notify-details">{{notification.Title}}</p>
        <p class="text-muted mb-0 user-msg">
            <small style="white-space: pre-line">{{notification.Content}}</small> <br/>
            <p class="text-muted text-end">{{time_ago}}</small>
        </p>
    </router-link>
`
    }
