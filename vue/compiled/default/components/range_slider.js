
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return null;
                },
            },
            max: {
                type: Number,
                default: 100
            },
            min: {
                type: Number,
                default: 0
            }
        },
        data: () => {
            return {
                value1: 7,
                value2: 16,
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                this.setValue()
            },
            value1: function (newValue, oldValue) {
                this.value1 = Math.min(this.value1, this.value2 - 1);
                this.updateValue()
            },
            value2: function (newValue, oldValue) {
                this.value2 = Math.max(this.value2, this.value1 + 1);
                this.updateValue()
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            left_width: function () {
                return (100 / (parseInt(this.max) - parseInt(this.min))) * parseInt(this.value1) - (100 / (parseInt(this.max) - parseInt(this.min))) * parseInt(this.min);
            },
            right_width: function () {
                return (100 / (parseInt(this.max) - parseInt(this.min))) * parseInt(this.value2) - (100 / (parseInt(this.max) - parseInt(this.min))) * parseInt(this.min);
            }
        },
        methods: {
            setValue() {
                if (this.value && this.min != undefined && this.max != undefined) {
                    this.value1 = this.value.min
                    this.value2 = this.value.max
                } else {
                    this.value1 = this.min
                    this.value2 = this.max
                }
            },
            updateValue() {
                var value = {
                    min: this.value1,
                    max: this.value2,
                }
                if (this.value.min != value.min || this / value.max != value.max) {
                    this.$emit('input', value)
                }
            },
        },
        mounted: function () {},
        template: `
    <div slider>
        <div>
            <div ref="child1" inverse-left :style="{width:left_width+'%'}"></div>
            <div ref="child2" inverse-right :style="{width:(100 - right_width)+'%'}"></div>
            <div ref="child3" range :style="{left:left_width+'%', right: (100 - right_width)+'%'}"></div>
            <span ref="child4" thumb :style="{left:left_width+'%'}"></span>
            <span ref="child5" thumb :style="{left:right_width+'%'}"></span>
            <div sign :style="{left:left_width+'%'}">
                <span id="value">{{value1}}</span>
            </div>
            <div sign :style="{left:right_width+'%'}">
                <span id="value">{{value2}}</span>
            </div>
        </div>
        <input ref="input1" type="range" tabindex="0" v-model="value1" :max="max" :min="min" step="1" />

        <input ref="input2" type="range" tabindex="0" v-model="value2" :max="max" :min="min" step="1" />
    </div>
`
    }
