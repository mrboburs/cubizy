<script>
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return {
                        ID: 0,
                    }
                }
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",

                Amount: 0,
                Method: "",
                TransactionID: "",
                UserID: 0,
                AccountID: 0,
                For: "",
                Status: "",
                Note: "",
                Accepted: false,

                Methods : [ "Coupon", "Checks", "Debitcard", "Creditcard", "Electronic bank transfers" ],
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                if (newValue) {
                    this.SetData()
                    this.$emit('onset', this.value)
                }
            },
            loading: function (newValue, oldValue) {
                if (newValue) {
                    this.error = false
                    this.message = false
                    this.submitted = false
                }
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account']),
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
            TransactionIDError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.TransactionID.trim()) {
                    return "TransactionID can not  be empty"
                }
            },
            AmountError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Amount || this.Amount < 1) {
                    return "Amount can not  be empty"
                }
            },
            MethodError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Method) {
                    return "Method can not  be empty"
                }
            },
        },
        methods: {
            Reset() {
                this.SetData()
                this.$emit('input')
            },
            SetData() {
                if (this.value) {
                    this.submitted = false
                    if (this.value.Amount) {
                        this.Amount = this.value.Amount
                        this.Method = this.value.Method
                        this.TransactionID = this.value.TransactionID
                        this.UserID = this.value.UserID
                        this.Status = this.value.Status
                        this.Note = this.value.Note
                        this.Accepted = this.value.Accepted
                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.TransactionIDError || this.MethodError || this.AmountError) { return }
                var value = {}
                debugger
                if(this.value){
                    value.ID = this.value.ID
                }
                value.Amount = this.Amount
                value.Method = this.Method
                value.TransactionID = this.TransactionID
                value.UserID = this.UserID
                value.Status = this.Status
                value.Note = this.Note
                value.Accepted = this.Accepted
                this.$emit('input', value)
                this.Accepted = false
            },
            accept(){
                this.Accepted = true
                this.submit()
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
    <form @submit.prevent="submit" class="vw90">
        <div class="row">
            <div class="col">
                <formitem name="inputAmount" label="Amount" :error="AmountError" v-model="Amount" suffix="$" :inputgroup="true" :type="value.UserID?'readonly':'number'"/>
                <formitem name="inputMethod" label="Method" :error="MethodError" v-model="Method" :values="Methods"  :type="value.UserID?'readonly':'select'"/>
                <formitem name="inputTransactionID" label="TransactionID" :error="TransactionIDError" v-model="TransactionID"  :type="value.UserID?'readonly':'text'"/>
                <formitem name="inputUserID" label="UserID" v-model="UserID" :type="value.UserID?'readonly':'number'"/>
                <formitem name="inputNote" label="Note" v-model="Note" type="textarea"/>
                <div class="d-flex align-items-center end m-2">
                    <button type="submit" class="btn btn-success m-1" :disabled="loading" @click.prevent="accept()">
                        Approve
                    </button>
                    <button type="submit" class="btn btn-success m-1" :disabled="loading">
                        Save
                    </button>
                    <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
                </div>
            </div>
        </div>
    </form>
</template>