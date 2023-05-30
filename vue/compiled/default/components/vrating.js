
    export default {
        props: {
            value: {
                type: Number,
                default: 0
            },
            maxRating: {
                type: Number,
                default: 5
            },
            icon_class: {
                type : String,
                default : "fa fa-star"
            },
            active_class: {
                type : String,
                default : "text-warning"
            },
            inactive_class: {
                type : String,
                default : "text-muted"
            },
            editable:{
                type : Boolean,
                default : false,
            },
            display:{
                type : Boolean,
                default : false,
            }
        },
        data: () => {
            return {
                rating : 0,
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if(this.rating != newValue){
                    this.rating = newValue
                }
            },
        }, 
        methods: {
            mouseLeave(){
                this.rating = this.value
            },
            mouseOver(arg){
                if(this.editable){
                    this.rating = arg
                }
            },
            click(){
                if(this.editable){
                    this.$emit('input', this.rating)
                }
            },
            rating_class(val){
                if(val > this.rating){
                    return this.inactive_class
                }else{
                    return this.active_class
                }
            }
        },
        template: `
    <div class="d-flex flex-row">
        <div class="d-flex flex-row">
            <span v-for="n in maxRating" :key="'rating'+n" class="icon"  @mouseover="mouseOver(n)" @mouseleave="mouseLeave" @click="click" :class="rating_class(n)"><i :class="icon_class" aria-hidden="true" style="font-size: large;"></i></span>
        </div>
        <span v-if="display">({{rating}})</span>
    </div>
`
    }
