
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
                StatusList: [
                    "Waiting for Pick Up",
                    "Undeliverable",
                    "Ready to Dispatch",
                    "Dispatched",
                    "Returned to Seller",
                    "Returning to seller",
                    "Rejected by Buyer",
                    "Picked Up",
                    "Out For Delivery",
                    "Lost in Transit",
                    "Label Cancelled",
                    "Delivered",
                    "Damaged in Transit",
                    "At Origin FC",
                    "At Destination FC",
                ],
                Status: "",
                TrackingID: "",
                TrackingDetail: "",
                ShippingDetails: {},
                cancelorder : false,
                userMessage : "",
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
            StatusError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Status.trim()) {
                    return "Status can not  be empty"
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
                    if (this.value.Status) {
                        this.Status = this.value.Status
                        this.TrackingID = this.value.TrackingID
                        try {
                            var ShippingDetails = JSON.parse(this.value.ShippingDetails)
                            this.ShippingDetails = {}
                            ShippingDetails.forEach(element => {
                                if (!this.ShippingDetails[this.getFullDate(element.CreatedAt)]) {
                                    this.ShippingDetails[this.getFullDate(element.CreatedAt)] = []
                                }
                                this.ShippingDetails[this.getFullDate(element.CreatedAt)].push(
                                    {
                                        time: this.getTime(element.CreatedAt),
                                        note: element.Note
                                    }
                                )
                            });
                        } catch (error) {
                            console.log(error)
                        }

                        setTimeout(() => {
                            var objDiv = this.$refs.order_reacking_info;
                            objDiv.scrollTo({
                                top: objDiv.scrollHeight,
                                left: 0,
                                behavior: 'smooth'
                            });
                        }, 50);

                        if(this.value.RequestedToCalcel> 0){
                            this.userMessage = "User requested seller to cancel order on " + this.getFullDateTime(this.value.RequestedToCalcel)
                        }

                    }
                }
            },
            submit() {
                this.submitted = true
                if (this.StatusError) { return }
                var value = {}
                if (this.value && this.value.ID) {
                    value.ID = this.value.ID
                }
                value.Status = this.Status
                value.TrackingID = this.TrackingID
                value.TrackingDetail = this.TrackingDetail
                if(this.cancelorder){
                    value.cancel_order = true
                    this.cancelorder = false
                }
                this.$emit('input', value)
            },
            cancel_order(){
                this.cancelorder = true
                this.submit()
            }
        },
        mounted: function () {
            this.SetData()
            this.$emit('onload', this)
        },
        template: `
    <divloading :fullpage="false" :loading="loading" class="m-1 row">
        <div  class="col-12">
            <v-alert v-model="message" :error="error" />
            <div class="alert alert-danger" role="alert" v-if="userMessage">
                {{userMessage}}
            </div>
        </div>
        <div class="col" style="overflow-y: scroll; max-height: 379px;" ref="order_reacking_info">
            <ul class="list-group border-bottom">
                <template v-for="key in Object.keys(ShippingDetails)">
                    <li class="list-group-item">
                        <strong>{{key}}</strong>
                    </li>
                    <li class="list-group-item border-bottom-0" v-for="item in ShippingDetails[key]">
                        <label class="ms-3">{{item.time}} :</label> {{item.note}}
                    </li>
                </template>
            </ul>
        </div>
        <form @submit.prevent="submit" class="col">
            <formitem name="inputStatus" label="Status" :error="StatusError" v-model="Status" type="select"
                :values="StatusList" />
            <formitem name="inputTrackingID" label="TrackingID" v-model="TrackingID" />
            <formitem name="inputTrackingDetail" label="TrackingDetail" v-model="TrackingDetail" type="textarea" />
            <div class="d-flex align-items-center end m-2">
                <button type="button" class="btn btn-danger m-1" @click.prevent="cancel_order"> Cancel the order </button>
                <button type="submit" class="btn btn-success m-1"> submit </button>
                <button class="btn btn-secondary m-1" @click="Reset">Cancel</button>
            </div>
        </form>
    </divloading>
`
    }
