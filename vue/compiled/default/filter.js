
    export default {
        components: {
            'tuitiontypes': () => import("/vue/tuitiontypes.js"),
            'locationtree': () => import("/vue/locationtree.js"),
            'timeslot': () => import("/vue/timeslot.js"),
        },
        data: () => {
            return {
                TuitionTypes: [],
                Location: {},
                TimesRange: {
                    min: 7,
                    max: 21
                }
            }
        },
        watch: {
            TuitionTypes: function (newValue, oldValue) {
                this.dosearch()
            },
            Location: function (newValue, oldValue) {
                this.dosearch()
            },
            TimesRange: function (newValue, oldValue) {
                this.dosearch()
            },
        },
        methods: {
            dosearch: _.debounce(function () {
                this.search();
            }, 1000),
            search() {
                var search = {}
                if (this.TuitionTypes.length) {
                    search.TuitionTypes = JSON.parse(JSON.stringify(this.TuitionTypes))
                }
                if (this.Location.Country) {
                    search.Location = JSON.parse(JSON.stringify(this.Location))
                }
                if (this.TimesRange.min || this.TimesRange.max) {
                    search.TimesRange = {
                        min : this.paddZero(this.TimesRange.min) + ":00",
                        max : this.paddZero(this.TimesRange.max) + ":00",
                    }
                }
                this.$emit('search', search)
            },
            paddZero(number){
                if (number <10){
                    return "0"+number + ""
                }else{
                    return number + ""
                }
            }
        },
        mounted: function () {
        },
        template: `
    <div class="p-1">
        <tuitiontypes class="my-1" v-model="TuitionTypes"></tuitiontypes>
        <div class="card">
            <strong class="text-nowrap lead">
                Time Slot
            </strong>
            <div class="card-body">
                <v-range_slider v-model="TimesRange" :min="0" :max="24"></v-range_slider>
            </div>
        </div>
        <locationtree class="my-1" v-model="Location"></locationtree>
    </div>
`
    }
