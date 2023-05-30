
    export default {
        data: () => {
            return {
                min : 0,
                max : 100,
                min_value : 0,
                max_value : 24,
            }
        },
        watch: {
            min_value: function (newValue, oldValue) {
                new_max_value = (newValue - oldValue) + this.max_value
                if(new_max_value > 100){
                    new_max_value = 100
                }
                this.max_value = new_max_value
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        methods: { },
        mounted: function () {
            
         },
        template: `
    
`
    }
