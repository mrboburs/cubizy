
    export default {
        props: {
            value: {
                required: true
            }
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                ThemeSettings: {
                    key: {
                        title: "property",
                        value: "value",
                        type: "text",
                    }
                }
            }
        },
        watch: {
            value: function (newValue, oldValue) {
                this.init()
            },
            loading: function (newValue, oldValue) {
                if(newValue){
                    this.error = false
                    this.message = ""
                }
            }
        },
        methods: {
            init() {
                this.ThemeSettings = []
                var need_to_load_file = true
                try {
                    if (this.value) {
                        this.ThemeSettings = JSON.parse(this.value)
                        need_to_load_file = false
                    }
                } catch (error) {
                    console.log("error while loading theme setting json")
                    console.log(error)
                }
                if (need_to_load_file) {
                    this.loadData()
                }
            },
            reset() {
                this.loadData(true)
            },
            loadData: function (reset) {
                this.loading = true
                this.message = "";
                this.error = false
                var data = {}
                if (reset) {
                    data["ResetThemeSettings"] = true
                }
                this.$store.dispatch('call', {
                    api: "themesettings",
                    data: data,
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        if (data.Result.ThemeSettings) {
                            try {
                                this.ThemeSettings = JSON.parse(data.Result.ThemeSettings)
                            } catch (error) {
                                this.error = true;
                                this.message = error
                            }
                        }
                    } else {
                        this.error = true;
                    }
                }).catch((error) => {
                    this.error = true;
                    this.message = error
                    console.log(error)
                }).finally(() => {
                    this.loading = false
                });
            },
            submit(setting, value) {
                setting.value = value
            },
            getSettingArray(setting) {
                if (Array.isArray(setting.value) && setting.value.length) {
                    return setting.value
                } else {
                    return []
                }
            },
            addValueToArraySetting(setting) {
                if (!Array.isArray(setting.value)) {
                    setting.value = []
                }
                setting.value.push(JSON.parse(JSON.stringify(setting.template)))
            },
            removeValueFromArraySetting(setting, index) {
                setting.value.splice(index, 1);
            },
            validate() {
                try {
                    var ThemeSettings = JSON.stringify(this.ThemeSettings)
                    this.$emit('input', ThemeSettings)
                } catch (error) {
                    console.log("error while submiting theme setting json")
                    console.log(error)
                    return false
                }
                return true
            }
        },
        mounted: function () {
            this.init()
        },
        template: `
    <div class="border rounded container py-2">
        <v-alert v-model="message" :error="error" />
        <label v-if="loading">Loading theme settings , please wait....</label>
        <template v-for="key in Object.keys(ThemeSettings)">
            <div v-if="ThemeSettings[key].type == 'array'" class="row">
                <label class="form-label col-md-4 col-lg-3 text-end">{{ThemeSettings[key].title}}</label>
                <div class="col-md-8 col-lg-9">
                    <div class="m-1 p-1 border rounded position-relative"
                        v-for="(item, index) in getSettingArray(ThemeSettings[key])">
                        <label class="position-absolute"> {{index + 1}} </label>
                        <formitem v-for="subkey in Object.keys(item)" :customLayout="true"
                            :name="'input'+key+'_'+subkey" :label="item[subkey].title" :value="item[subkey].value"
                            @input="submit(item[subkey], $event)" :type="item[subkey].type" prefix="website" />
                        <button class="btn btn-outline-danger w-100"
                            @click="removeValueFromArraySetting(ThemeSettings[key], index)">
                            <i class="fas fa-trash-alt"></i>
                        </button>
                    </div>
                    <button class="btn btn-outline-success w-100" @click="addValueToArraySetting(ThemeSettings[key])">
                        Add Item
                    </button>
                </div>
            </div>
            <formitem v-else :customLayout="true" :name="'input'+key" :label="ThemeSettings[key].title"
                :value="ThemeSettings[key].value" @input="submit(ThemeSettings[key], $event)"
                :type="ThemeSettings[key].type" prefix="website" />
        </template>
    </div>
`
    }
