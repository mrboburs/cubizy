<script>
    export default {
        props: {
            value: {
                type: Number,
                default: 0
            },
            totalrows: {
                type: Number,
                required: true
            },
            perpage: {
                type: Number,
                required: true
            },
        },
        data: () => {
            return {
                priviousPage : 0,
                nextPage : 2,
                maxPage : 0,
            }
        },
        watch: {
            value(newValue, oldValue) {
                this.priviousPage = newValue - 1
                this.nextPage = newValue + 1
            },
            totalrows(newValue, oldValue) {
                this.setMaxPage()
            },
            perpage(newValue, oldValue) {
                this.setMaxPage()
            },
        },
        methods: {
            setMaxPage() {
                this.maxPage = Math.ceil(this.totalrows / this.perpage)
            },
            setPage(page){
                if(page < 1){
                    page = 1
                }
                if(page > this.maxPage){
                    page = this.maxPage
                }
                if(page != this.value){
                    this.$emit('input', page)
                }
            }
        },
        mounted: function () {
            this.priviousPage = this.value - 1
            this.nextPage = this.value + 1
            this.setMaxPage() 
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <nav v-if="maxPage > 1" aria-label="...">
        <ul class="pagination pagination-rounded my-0">
            <li class="page-item" :class="{'disabled': (value == 1)}">
                <a class="page-link" href="#" @click.prevent="setPage(1)">First</a>
            </li>
            <li class="page-item" :class="{'disabled': (priviousPage == 0)}">
                <a class="page-link" href="#" @click.prevent="setPage(priviousPage)">Previous</a>
            </li>
            <li v-if="value > 3">...</li>
            <li class="page-item" v-if="value > 1">
                <a class="page-link" href="#" @click.prevent="setPage(priviousPage)">{{priviousPage}}</a>
            </li>
            <li class="page-item active" aria-current="page">
                <a class="page-link" href="#" @click.prevent="setPage(value)">{{value}}</a>                
            </li>
            <li class="page-item" v-if="maxPage >= nextPage">
                <a class="page-link" href="#" @click.prevent="setPage(nextPage)">{{nextPage}}</a>
            </li>
            <li v-if="maxPage  > (value + 3) ">...</li>
            <li class="page-item" :class="{'disabled': (value == maxPage)}">
                <a class="page-link" href="#" @click.prevent="setPage(nextPage)">Next</a>
            </li>
            <li class="page-item" v-if="value != maxPage">
                <a class="page-link" href="#" @click.prevent="setPage(maxPage)">Last</a>
            </li>
        </ul>
    </nav>
</template>