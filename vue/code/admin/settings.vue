<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                settings: [],
                settingObjects: {},
                saving: 0,
                showtype : false,
                FieldTypes: [ 'text', 'image' ],
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
        },
        computed: {
            showmessage: {
                // getter
                get: function () {
                    if (this.message) {
                        return true;
                    } else {
                        return false
                    }
                },
                // setter
                set: function (newValue) {
                    if (!newValue) {
                        this.message = ""
                    }
                }
            },
            messagetype: function () {
                if (this.error) {
                    return 'alert-danger'
                } else {
                    return 'alert-success'
                }
            }
        },
        methods: {
            save(setting) {
                this.saving = setting.ID;
                this.$store.dispatch('call', {
                    api: "settings",
                    data: {
                        setting: setting
                    },
                }).then((data) => {
                    if (data.Status == 2) {
                        this.settingObjects[setting.ID].Value = setting.Value
                    }
                }).catch((error) => {
                    if (error) console.log(error)
                }).finally(() => {
                    this.saving = false;
                })
            },
            reset(setting) {
                setting.Value = this.settingObjects[setting.ID].Value
            },
            fillData(settings) {
                settings.forEach(setting => {
                    setting.loading = false
                    this.settingObjects[setting.ID] = Object.assign({}, setting)
                    this.settingObjects[setting.ID].loading = false
                });
                this.settings = settings
            },
            load() {
                this.loading = true;
                this.$store.dispatch('call', {
                    api: "settings",
                    data: {},
                }).then((data) => {
                    if (data.Status == 2) {
                        this.fillData(data.Result.settings)
                    } else {
                        throw data.Message;
                    }
                }).catch((error) => {
                    if (error) console.log(error)
                }).finally(() => {
                    this.loading = false;
                })
            },
            toggleloading() {
                this.loading = !this.loading
            },
            error_in(setting) {
                if(!this.submitted){
                    return false
                }
                if (!this.setting.Value) {
                    return "Please enter the value for " + setting.Name
                }
            },
        },
        mounted: function () {
            this.load()
            if(location.search == "?showtype"){
                this.showtype = true
            }
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="col-12">
        <divloading :fullpage="false" :loading="loading" class="row">
            <div class="col-12">
                <div class="card">
                    <div class="card-body">
                        <!-- <divloading  :show="saving == setting.ID"
                            rounded="sm"> -->
                        <divloading class="d-flex" v-for="(setting, index) in settings" :key="'setting'+setting.ID"
                            :loading="saving == setting.ID">
                            <formitem :customLayout="true" name="inputFieldValue" :label="setting.Name" 
                                    :type="setting.Type?setting.Type:'text'" :error="error_in(setting)"   
                                    v-model="setting.Value" :description="setting.Details" class="mx-2 col" :class="{ 'col-lg-4' : showtype}"/>
                            <formitem v-if="showtype" :customLayout="true"  :name="'input_setting_'+index" label="FieldType" v-model="setting.Type" 
                                    type="select" displayby="Type" :values="FieldTypes" class="mx-2 col-lg-2" />
                            <div class="col-lg-4">
                                <div class="text-lg-right">
                                    <button type="button" class="btn btn-success mr-1" @click="save(setting)">
                                        <i class="mdi mdi-content-save"></i>
                                    </button>
                                    <button type="button" class="btn btn-danger mr-1" @click="reset(setting)">
                                        <i class="mdi mdi-cancel"></i>
                                    </button>
                                </div>
                            </div>
                            <!-- end col-->
                        </divloading>
                        <!-- end row -->
                        <!-- </divloading> -->
                    </div>
                    <!-- end card-box -->
                </div>
            </div>
            <!-- end col-->
        </divloading>
    </div> <!-- content -->
</template>