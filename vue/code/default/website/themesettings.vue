<script>
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
                if (newValue) {
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
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="border rounded container py-2">
        <v-alert v-model="message" :error="error" />
        <label v-if="loading">Loading theme settings , please wait....</label>
        <template v-for="key in Object.keys(ThemeSettings)">
            <div v-if="ThemeSettings[key].type == 'array' || ThemeSettings[key].type == 'object'" class="mt-1 p-1">
                <h2 class="accordion-header d-flex">
                    <a class="accordion-button" data-bs-toggle="collapse" :href="'#'+key+'details'" role="button"
                        aria-expanded="false" aria-controls="collapseExample">
                        {{ThemeSettings[key].title}}
                    </a>
                    <button class="btn btn-outline-success" @click="addValueToArraySetting(ThemeSettings[key])">
                        <i class="fas fa-plus"></i>
                    </button>
                </h2>
                <div v-if="ThemeSettings[key].type == 'array'" class="collapse border rounded" :id="key+'details'">
                    <div class="m-1" v-for="(item, index) in getSettingArray(ThemeSettings[key])">
                        <h2 class="accordion-header d-flex">
                            <a class="accordion-button" data-bs-toggle="collapse" :href="'#'+key+'details'+index+'item'"
                                role="button" aria-expanded="false" aria-controls="collapseExample">
                                item {{index + 1}}
                            </a>
                            <button class="btn btn-outline-danger"
                                @click.prevent="removeValueFromArraySetting(ThemeSettings[key], index)">
                                <i class="fas fa-trash-alt"></i>
                            </button>
                        </h2>
                        <div class="collapse card card-body" :id="key+'details'+index+'item'">
                            <formitem v-for="subkey in Object.keys(item)" :customLayout="false"
                                :name="'input'+key+'_'+subkey" :label="item[subkey].title" :value="item[subkey].value"
                                @input="submit(item[subkey], $event)" :type="item[subkey].type" prefix="website" />
                        </div>
                    </div>
                </div>
                <div v-if="ThemeSettings[key].type == 'object'" class="collapse border rounded" :id="key+'details'">
                    <formitem v-for="subkey in Object.keys(ThemeSettings[key].value)" :customLayout="false"
                    :name="'input'+key+'_'+subkey" :label="ThemeSettings[key].value[subkey].title" :value="ThemeSettings[key].value[subkey].value" @input="submit(ThemeSettings[key].value[subkey], $event)" :type="ThemeSettings[key].value[subkey].type" prefix="website" />
                </div>
            </div>
            <formitem v-else :customLayout="false" :name="'input'+key" :label="ThemeSettings[key].title"
                :value="ThemeSettings[key].value" @input="submit(ThemeSettings[key], $event)"
                :type="ThemeSettings[key].type" prefix="website" />
        </template>
    </div>
</template>