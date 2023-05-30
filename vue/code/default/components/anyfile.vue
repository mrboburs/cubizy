<script>
    export default {
        props: {
            title: {
                type: String,
                default: ""
            },
            value: {
                type: String,
                default: ""
            },
            prefix: {
                type: String,
                default: ""
            },
            button_image: {
                type: String,
                default: ""
            },
            multiple: {
                type: Boolean,
                default: false
            },
            maxHeight: {
                type: String,
                default: "200px"
            },
            maxWidth: {
                type: String,
                default: "200px"
            },
            removable: {
                type: Boolean,
                default: true
            }
        },
        data() {
            return {
                empty_image: "https://mystudyfiles.s3.ap-southeast-2.amazonaws.com/empty.jpg",
                showmodal: false,
                actualWidth : 0,
            };
        },
        mounted() {
            this.actualWidth = this.$refs.holder.clientHeight
        },
        methods: {
            removeImage: function(file){
                this.$emit('input', "")
            },
            selecteImage: function (file) {
                if (!file) {
                    this.showmodal = false
                }else if (this.multiple) {
                    this.$emit('input', [file])
                } else {
                    this.$emit('input', file)
                    this.showmodal = false
                }
            },
        },
        template: `{{{template}}}`
    };
</script>

<template>
    <div ref="holder">
        <div class="d-flex" >
            <label>{{value.split("/").pop()}}</label>
            <button type="button" class="btn btn-sm btn-primary m-1" @click="showmodal = !showmodal">
                <i class="mdi mdi-sync"></i>
                <span v-if="actualWidth > 199" class="list-btn-label">Change</span>
            </button>
            <button v-if="removable" type="button" class="btn btn-sm btn-primary m-1" @click="removeImage">
                <i class="mdi mdi-delete"></i>
                <span v-if="actualWidth > 199" class="list-btn-label">Remove</span>
            </button>
        </div>
        <v-files v-if="prefix" :title="title" :showmodal="showmodal" :value="value" :multiple="multiple" :only_image="false" :prefix="prefix" @input="selecteImage($event)" @close="showmodal = false" />
    </div>
</template>