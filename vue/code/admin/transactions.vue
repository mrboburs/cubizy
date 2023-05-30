<script>
    export default {
        components: {
            'TransactionEditor': () => import("/vue/transactioneditor.js"),
        },
        data() {
            return {
                title: 'Transactions',
                breadcrumb: [
                    {
                        text: 'Transaction',
                    },
                    {
                        text: 'All',
                        active: true,
                    },
                ],
                /*
                            <th>Date</th>
                            <th>Amount</th>
                            <th>Method</th>
                            <th>For</th>
                            <th>TransactionID</th>
                            <th>Status</th>
                            <th>Action</th>
                 */
                columns: [
                    { key: "ID", sortable: true, sortkey: "transactions`.`id" },
                    { key: "CreatedAt", type: 'date', sortable: true, sortkey: "transactions`.`created_at" },
                    { key: "Amount", type: 'money', sortable: true, sortkey: "transactions`.`amount" },
                    { key: "TransactionID", sortkey: "transactions`.`transaction_id" },
                    { key: "UserID", title: "UID", sortable: true, sortkey: "transactions`.`user_id" },
                    { key: "UserName", sortable: true, sortkey: "user_name" },
                    { key: "SellerID", title: "SellerID", sortable: true, sortkey: "transactions`.`account_id" },
                    { key: "SellerName", title: "Seller", sortable: true, sortkey: "seller_name" },
                    { key: "For", title: "For", sortable: true },
                    { key: "Method", sortable: true, sortkey: "transactions`.`method" },
                    { key: "Status", sortable: true, sortkey: "transactions`.`status" },
                    { key: "UpdatedAt", type: 'date', sortable: true, sortkey: "transactions`.`updated_at" },
                ],
                table: false,
                error: "",
                message: "",
                actions: [
                    // {
                    //     key: "import",
                    //     icon: "ri-upload-line",
                    //     text: "Import"
                    // }
                ],
                conditions: {
                    For: "wallet",
                    DateFrom: 0,
                    DateTill: 0,
                },
                new: {
                    ID: 0,
                    Name: "",
                    Status: true
                },
                TransactionTypes: [
                    "All",
                    "Order",
                    "Cancel Item",
                    "Delivered Item"
                ],
                SelectedTransactionType: "All",
                DateFrom: "",
                DateTill: "",
                TotalCount: 0,
                TotalAmount: 0,
            }
        },
        computed: {
            ...Vuex.mapGetters(['getFullDate', 'getMonth', 'getFullDateTime', 'getDate', 'getTime']),
        },
        watch: {
            $route: function (newValue, oldValue) {
                if (this.$route.name == "wallet_transactions") {
                    this.SelectedTransactionType = "wallet"
                } else {
                    this.SelectedTransactionType = "All"
                }
                this.init()
            },
            DateFrom: function (newValue, oldValue) {
                this.init()
            },
            DateTill: function (newValue, oldValue) {
                this.init()
            },
            SelectedTransactionType: function (newValue, oldValue) {
                this.init()
            },
        },
        mounted: function () {
            if (this.$route.name == "wallet_transactions") {
                this.SelectedTransactionType = "wallet"
            } else {
                this.SelectedTransactionType = "All"
            }
            var now = new Date()
            var day = ("0" + now.getDate()).slice(-2);
            var month = ("0" + (now.getMonth() + 1)).slice(-2);
            var today = now.getFullYear() + "-" + (month) + "-" + (day);
            this.DateFrom = today
            this.DateTill = today
            this.init()
        },
        methods: {
            init: _.debounce(function () {
                if (this.$refs.DateFrom.valueAsDate == null) {
                    this.$refs.DateFrom.valueAsDate = new Date()
                }

                if (this.$refs.DateTill.valueAsDate == null) {
                    this.$refs.DateTill.valueAsDate = new Date()
                }

                var start_date = this.$refs.DateFrom.valueAsDate;
                start_date.setHours(0, 0, 0, 0);
                var start = Math.floor(start_date / 1000)

                var end_date = this.$refs.DateTill.valueAsDate;
                end_date.setHours(23, 59, 59, 999);
                var end = Math.floor(end_date / 1000)

                if (this.$route.name == "wallet_transactions") {
                    this.conditions = {
                        For: "wallet",
                        From: start,
                        Till: end
                    }
                } else {
                    this.conditions = {
                        For: this.SelectedTransactionType,
                        From: start,
                        Till: end
                    }
                }
            }, 100),
            onAction(action, arg) {
                switch (action) {
                    case 'loading':
                        this.loading = true
                        this.table = arg
                        break;
                    case 'add_new':
                        this.table.editing_item = Object.assign({}, this.new)
                        break;
                    default:
                        break;
                }
            },
            onActionDone(data) {
                debugger
                this.loading = false
                this.TotalCount = data.recordsTotal
                if(data.Result.TotalAmount){
                    this.TotalAmount = data.Result.TotalAmount
                }else{
                    this.TotalAmount = 0
                }
                
            },
            getStatusClass(status) {
                status = status.trim().toLowerCase()
                switch (status) {
                    case "successful":
                        return "bg-success";
                    case "failed":
                        return "bg-danger";
                    case "pending":
                        return "bg-warning";
                    case "cancelled":
                        return "bg-info";
                    default:
                        return "bg-info";
                }
            },
        },
        template: `{{{template}}}`
    }
</script>

<template>
    <div class="col-12 row">
        <div class="col-12">
            <div class="card">
                <div class="card-body">
                    <div class="row">
                        <div class="col-auto">
                            <div class="row g-3">
                                <div class="col-auto">
                                    <label>From Date</label>
                                    <input ref="DateFrom" class="form-control" type="date" placeholder="From Date"
                                        aria-label="From Date" v-model="DateFrom">
                                </div>
                                <div class="col-auto">
                                    <label>Till Date</label>
                                    <input ref="DateTill" class="form-control" type="date" placeholder="Till Date"
                                        aria-label="Till Date" v-model="DateTill">
                                </div>
                                <div v-if="$route.name != 'wallet_transactions'" class="col-auto">
                                    <label>Transaction Types</label>
                                    <select ref="TransactionTypes" class="form-select" placeholder="Transaction Type"
                                        aria-label="Transaction Type" v-model="SelectedTransactionType">
                                        <option v-for="item in TransactionTypes" :value="item">{{item}}</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                        <div class="col">
                            <div class="row g-3 flex-row-reverse">
                                <div class="col-auto">
                                    <label>Total Transactions : </label> <br />
                                    <label>{{TotalCount}}</label>
                                </div>
                                <div v-if="SelectedTransactionType == 'wallet'" class="col-auto">
                                    <label>Wallet Total</label> <br />
                                    <label><span>$</span>{{TotalAmount}}</label>
                                </div>
                                <div v-if="SelectedTransactionType == 'Order'" class="col-auto">
                                    <label>Order Total</label> <br />
                                    <label><span>$</span>{{TotalAmount}}</label>
                                </div>
                                <div v-if="SelectedTransactionType == 'Cancel Item'" class="col-auto">
                                    <label>Cancel Items Total</label> <br />
                                    <label><span>$</span>{{TotalAmount}}</label>
                                </div>
                                <div v-if="SelectedTransactionType == 'Delivered Item'" class="col-auto">
                                    <label>Delivered Items Total</label> <br />
                                    <label><span>$</span>{{TotalAmount}}</label>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-12">
            <div class="card">
                <List api="transactions" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                    editor_type="offcanvas" :conditions="conditions" :can_import="true" :actions="actions"
                    @done="onActionDone" @onaction="onAction" editor_size="end"
                    default_sort_by="transactions\`.\`created_at" :default_desc="true">
                    <template v-slot:editor="editing_item">
                        <TransactionEditor v-if="editing_item.item" :value="editing_item.item"
                            @input="editing_item.submit">
                        </TransactionEditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
</template>