
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Question1: "",
                Answer1: "",

                Question2: "",
                Answer2: "",

                Question3: "",
                Answer3: "",
            }
        },
        watch: {
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
            user: function (newValue, oldValue) {
                if (newValue) {
                    this.init()
                }
            },
        },
        computed: {
            ...Vuex.mapState(["user"]),
            Question1Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Question1) {
                    return "Please set Question1"
                }
            },
            Answer1Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Answer1) {
                    return "Please set Answer1"
                }
            },
            Question2Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Question2) {
                    return "Please set Question2"
                }
            },
            Answer2Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Answer2) {
                    return "Please set Answer2"
                }
            },
            Question3Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Question3) {
                    return "Please set Question3"
                }
            },
            Answer3Error: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Answer3) {
                    return "Please set Answer3"
                }
            },
        },
        methods: {
            init() {
                if (this.user) {
                    this.Question1 = this.user.Question1
                    this.Answer1 = this.user.Answer1
                    this.Question2 = this.user.Question2
                    this.Answer2 = this.user.Answer2
                    this.Question3 = this.user.Question3
                    this.Answer3 = this.user.Answer3
                }
            },
            submit() {
                this.submitted = true
                if (this.Question1Error || this.Answer1Error || this.Question2Error || this.Answer2Error || this.Question3Error || this.Answer3Error) {
                    return
                }
                this.$store.dispatch('call', {
                    api: "setquestionanswers",
                    data: {
                        Question1: this.Question1,
                        Answer1: this.Answer1,
                        Question2: this.Question2,
                        Answer2: this.Answer2,
                        Question3: this.Question3,
                        Answer3: this.Answer3,
                    }
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$router.push('/')
                    } else {
                        this.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    this.error = true
                    this.message = error
                }).finally(() => {
                    this.loading = false
                })
            },
        },
        mounted: function () {
            this.init()
        },
        template: `
    <div class="m-auto w-50">
        <!-- title-->
        <h4 class="mt-0"> Secret Questions </h4>
        <p class="text-muted mb-4">Secret question is part of our 2 factor authentication which will help to make your account more secure </p>
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="container">
            <formitem name="inputQuestion1" label="Question 1" :error="Question1Error" v-model="Question1" type="select" service="questions" displayby="Question" selectby="Question" />
            <formitem name="inputAnswer1" label="Answer 1" :error="Answer1Error" v-model="Answer1" />
            <formitem name="inputQuestion2" label="Question 2" :error="Question2Error" v-model="Question2" type="select" service="questions" displayby="Question" selectby="Question" />
            <formitem name="inputAnswer2" label="Answer 2" :error="Answer2Error" v-model="Answer2" />
            <formitem name="inputQuestion3" label="Question 3" :error="Question3Error" v-model="Question3" type="select" service="questions" displayby="Question" selectby="Question" />
            <formitem name="inputAnswer3" label="Answer 3" :error="Answer3Error" v-model="Answer3" />
            <div class="d-grid text-center mb-3">
                <button type="button" class="btn btn-primary" @click="submit">
                    Submit
                </button>
            </div>
        </divloading>
    </div>
`
    }
