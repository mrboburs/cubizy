<script>
    export default {
        props: {
            value: {
                default : "",
            },
            placeholder: {
                default : "",
            },
        },
        data: () => {
            return {
                tags: [],
                output: "",
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.tags = newValue.split(",");
                }
            },
        },
        methods: {
            remove(value) {
                this.tags = this.tags.filter(item => item !== value)
                this.setValue()
            },
            add(event) {
                if (event.code == "Comma" || event.code == "Enter") {
                    var value = event.target.value.trim()
                    if (value && !this.tags.includes(value)) {
                        this.tags.push(event.target.value)
                    }
                    event.target.value = ""
                    event.preventDefault();
                    this.setValue()
                }
            },
            setValue() {
                this.output = this.tags.join()
                this.$emit('input', this.output)
            },
            loadValue() {
                if (this.output != this.value) {
                    if (this.value) {
                        this.tags = this.value.split(",");
                    } else {
                        this.tags = [];
                    }
                }
            }
        },
        mounted: function () {
            this.loadValue()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="input-group flex-wrap border rounded-2">
        <!-- <span class="input-group-text"> -->
        <span class="border border-primary rounded-pill d-flex centered m-1 px-2" v-for="(tag, index) in tags">
            {{tag}}
            <button type="button" class="btn-close btn-sm" aria-label="Close" @click.prevent="remove(tag)"></button>
        </span>
        <!-- </span> -->
        <input type="text" class="form-control" :placeholder="placeholder" @keydown="add($event)">
    </div>
</template>