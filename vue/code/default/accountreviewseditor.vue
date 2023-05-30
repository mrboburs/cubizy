<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return null
                }
            },
        },
        data: () => {
            return {
                Replay: "",
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.SetData()
                    this.$emit('onset', this.value)
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    if (this.value.Replay) {
                        this.Replay = this.value.Replay
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.ReplayError) { return }
                this.value.Replay = this.Replay
                this.$emit('input', this.value)
            },
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <form @submit.prevent="submit">
        <dl class="row">
            <dt class="col-sm-3">Review</dt>
            <dd class="col-sm-9">{{value.Review}}</dd>
            <dt class="col-sm-3">Pros</dt>
            <dd class="col-sm-9">{{value.Pros}}</dd>
            <dt class="col-sm-3">Cons</dt>
            <dd class="col-sm-9">{{value.Cons}}</dd>
            <dt class="col-sm-6">Rating </dt>
            <dd class="col-sm-6"><star-rating :rating="value.Rating" :showRating="false" :starSize="20" :readOnly="true"></star-rating></dd>
        </dl>
        <formitem name="inputReplay" type="textarea" label="Replay" v-model="Replay" />
        <div class="d-flex align-items-center end m-2">
            <button type="submit" class="btn btn-success m-1">
                Save
            </button>
            <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
        </div>
    </form>
</template>