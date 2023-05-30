
    export default {
        props: {
            value: {
                type: Object,
            },
        },
        components: {
            'AccountEditor': () => import("/vue/account/accounteditor.js"),
        },
        data: () => {
            return {
                submitted: false,
                CanActive: true,
                IDProof: 0,
                IDStatus: false,
                AddressProof: 0,
                AddressStatus: false,
                RegistretionProof: 0,
                RegistretionStatus: false,
                OtherDocument: 0,
                Status: false,
                StatusComment: false,
                location: {
                    lat: 0,
                    lng: 0,
                },
                statuslist : [
                    {Value : 1, Title : "Under Review" },
                    {Value : 2, Title : "Rejected" },
                    {Value : 3, Title : "OnHold" },
                    {Value : 4, Title : "Expired" },
                    {Value : 10, Title : "Active" },
                ],
            }
        },
        computed: {
            StatusCommentError: function () {
                if (!this.submitted) {
                    return false
                }
                if (this.Status == 10) {
                    if (this.StatusComment.trim()) {
                        return "Status comment must be empty"
                    }
                } else {
                    if (!this.StatusComment.trim()) {
                        return "Status comment can not be empty"
                    }
                }
                return false
            },
            StatusError: function () {
                if (!this.submitted) {
                    return false
                }
                if (!this.Status) {
                    return "Status can not be empty"
                }
                if (this.Status == 10) {
                    if (this.StatusComment.trim()) {
                        return "Status comment must be empty"
                    }
                } else {
                    if (!this.StatusComment.trim()) {
                        return "Status comment can not be empty"
                    }
                }
                return false
            },
        },
        methods: {
            SetData() {
                if (this.value && this.value.Title) {
                    this.submitted = false
                    this.IDProof = this.value.IDProof
                    this.IDStatus = this.value.IDStatus
                    this.AddressProof = this.value.AddressProof
                    this.AddressStatus = this.value.AddressStatus
                    this.RegistretionProof = this.value.RegistretionProof
                    this.RegistretionStatus = this.value.RegistretionStatus
                    this.OtherDocument = this.value.OtherDocument
                    this.Status = this.value.Status
                    this.StatusComment = this.value.StatusComment
                    this.CanActive = this.value.CanActive
                    this.location.lat = this.value.Latitude;
                    this.location.lng = this.value.Longitude;
                }
            },
            submit() {
                this.submitted = true
                if (this.StatusCommentError || this.StatusError) {
                    return
                }
                var account = {
                    CanActive: this.CanActive,
                    ID: this.value.ID,
                    IDStatus: this.IDStatus,
                    AddressStatus: this.AddressStatus,
                    RegistretionStatus: this.RegistretionStatus,
                    Status: this.Status,
                    StatusComment: this.StatusComment,
                }
                
                if(this.value.Status < 2 && this.Status > 9){
                    account.CanActive = true
                }
                this.$emit('input', account)
            },
            cancel() {
                this.$emit('input')
            }
        },
        mounted: function () {
            this.SetData()

        },
        template: `
    <section class="section is-medium card ">
        <div class="card-body container">
            <div class="row">
                <div class="col">
                    <dl class="summary_grid">
                        <dt class="">AccountType : </dt>
                        <dd class=" text-capitalize ">{{value.AccountType}}</dd>
                        <dt class="">Title : </dt>
                        <dd class=" text-capitalize ">{{value.Title}}</dd>
                        <dt class="">Description : </dt>
                        <dd class=" text-capitalize ">{{value.Description}}</dd>
                        <dt class="">Keywords : </dt>
                        <dd class=" text-capitalize ">{{value.Keywords}}</dd>
                        <dt class="">Email : </dt>
                        <dd class=" text-capitalize ">{{value.Email}}</dd>
                        <dt class="">Mobile : </dt>
                        <dd class=" text-capitalize ">{{value.Mobile}}</dd>
                        <dt class="">Subdomain : </dt>
                        <dd class=" text-capitalize ">{{value.Subdomain}}</dd>
                        <dt class="">Domain : </dt>
                        <dd class=" text-capitalize ">{{value.Domain}}</dd>
                        <dt class="">Website Status : </dt>
                        <dd class=" text-capitalize d-flex">
                            <span v-if="value.Active && value.CanActive">Active</span>
                            <span v-else>Not Active</span>
                            <formitem name="inputCanActive" class="ms-2" label="Can Active" v-model="CanActive"
                                type="checkbox" />
                        </dd>
                        <dt class="">Question1 : </dt>
                        <dd class=" text-capitalize ">{{value.Question1}}</dd>

                        <dt class="">Question2 : </dt>
                        <dd class=" text-capitalize ">{{value.Question2}}</dd>

                        <dt class="">Question3 : </dt>
                        <dd class=" text-capitalize ">{{value.Question3}}</dd>
                    </dl>
                </div>
                <div class="col">
                    <dl class="summary_grid">
                        <dt class="">Owner Name : </dt>
                        <dd class=" text-capitalize ">
                            {{value.Name}}
                            <span v-if="parseInt(value.Joined)">Joined</span>
                        </dd>
                        <dt class="">Owner Email : </dt>
                        <dd class=" text-capitalize ">{{value.OwnerEmail}}
                            <span v-if="parseInt(value.EmailVerified)">Verified</span>
                        </dd>
                        <dt class="">Owner Mobile : </dt>
                        <dd class=" text-capitalize ">{{value.OwnerEmail}}
                            <span v-if="parseInt(value.MobileVerified)">Verified</span>
                        </dd>
                        <dt class="">Wallet : </dt>
                        <dd class=" text-capitalize ">{{value.Wallet}}</dd>
                        <dt class="">Address : </dt>
                        <dd class=" text-capitalize ">
                            {{value.AddressTitle}} <span v-if="value.AddressMobile">({{value.AddressMobile}})</span>
                            <br />
                            {{value.AddressLine1}} {{value.AddressLine2}} {{value.AddressLine3}}
                        </dd>
                        <dt class="">SubLocality : </dt>
                        <dd class=" text-capitalize ">{{value.SubLocality}}</dd>
                        <dt class="">Locality : </dt>
                        <dd class=" text-capitalize ">{{value.Locality}}</dd>
                        <dt class="">District : </dt>
                        <dd class=" text-capitalize ">{{value.District}}</dd>
                        <dt class="">Code : </dt>
                        <dd class=" text-capitalize ">{{value.Code}}</dd>
                        <dt class="">Country : </dt>
                        <dd class=" text-capitalize ">{{value.Country}}</dd>

                        <dt class="">Answer1 : </dt>
                        <dd class=" text-capitalize ">{{value.Answer1}}</dd>

                        <dt class="">Answer2 : </dt>
                        <dd class=" text-capitalize ">{{value.Answer2}}</dd>

                        <dt class="">Answer3 : </dt>
                        <dd class=" text-capitalize ">{{value.Answer3}}</dd>
                    </dl>
                </div>
                <div class="col-auto">
                    <v-map v-model="location" style="width: 500px; height: 300px; max-width: 100%;"></v-map>
                </div>
                <div class="col-12">
                    <h5>Documents</h5>
                    <div class="row">
                        <div class="col-6 col-md-4 col-lg-3">
                            <formitem name="inputIDProof" type="file" label="IDProof *">
                                <v-files-static @oncount="IDProof = $event" :only_image="false" :priview="false"
                                    :prefix="'account_'+ value.ID +'/account/idproof'" />
                            </formitem>
                            <formitem name="inputIDStatus" label="IDStatus" v-model="IDStatus" type="checkbox" />
                        </div>
                        <div class="col-6 col-md-4 col-lg-3">
                            <formitem name="inputAddressProof" type="file" label="AddressProof *">
                                <v-files-static @oncount="AddressProof = $event" :only_image="false" :priview="false"
                                    :prefix="'account_'+ value.ID +'/account/addressproof'" />
                            </formitem>
                            <formitem name="inputAddressStatus" label="Address Status" v-model="AddressStatus"
                                type="checkbox" />
                        </div>
                        <div class="col-6 col-md-4 col-lg-3">
                            <formitem name="inputRegistretionProof" type="file" label="RegistretionProof *">
                                <v-files-static @oncount="RegistretionProof = $event" :only_image="false"
                                    :priview="false" :prefix="'account_'+ value.ID +'/account/registretionproof'" />
                            </formitem>
                            <formitem name="inputRegistretionStatus" label="Registretion Status"
                                v-model="RegistretionStatus" type="checkbox" />
                        </div>
                        <div class="col-6 col-md-4 col-lg-3">
                            <formitem name="inputOtherDocument" type="file" label="OtherDocument *">
                                <v-files-static @oncount="OtherDocument = $event" :only_image="false" :priview="false"
                                    :prefix="'account_'+ value.ID +'/account/otherdocuments'" />
                            </formitem>
                        </div>
                    </div>
                </div>
                <div class="col-12">
                    <div class="row">
                        <div class="col-sm-12 col-md">
                            <formitem name="inputStatusComment" label="Status Comment" :error="StatusCommentError"
                                v-model="StatusComment" type="textarea" />
                        </div>
                        <div class="col">
                            <formitem name="inputStatus" label="Status" :error="StatusError" v-model="Status">
                                <select class="form-select" v-model="Status">
                                    <option v-for="statusitem in statuslist" :value="statusitem.Value">{{statusitem.Title}}</option>
                                </select>
                            </formitem>
                            <div class="d-flex align-items-center end m-2">
                                <button type="submit" class="btn btn-success m-1" @click.prevent="submit">
                                    Save </button>
                                <button class="btn btn-danger m-1" @click="cancel"> Cancel </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
`
    }
