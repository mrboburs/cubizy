<script>
  export default {
    props: {
      value: {
        default: "",
      },
      service: "",
      filter: {},
      selectby: {
        type: String,
        default: "ID",
      },
      displayby: {
        type: String,
        default: "Title",
      },
      values: {
        type: Array,
        default: null,
      },
      nothing_selected_text: {
        type: String,
        default: "Nothing selected",
      },
      select_out_of_single: {
        type: Boolean,
        default: true,
      },
    },
    data: () => {
      return {
        loading: false,
        reload: false,
        message: "",
        error: false,
        placeholder: "placeholder",
        options: [],
        optionsTotal: 0,
        optionsFiltered: 0,
        selectedOption: false,
        search: "",
        limit: 10,
        load: _.debounce(function () {
          if (this.loading || !this.service || !this.canload()) {
            return;
          }
          this.loading = true;
          this.error = false;
          this.message = "";
          
          var data = {};
          data.search = this.search.trim().replace(": ", '":"');
          data.sort = this.displayby;
          data.sortdesc = false;
          data.limit = this.limit;
          data.page = Math.ceil(this.options.length / this.limit);
          if (this.filter) {
            data.fix_condition = this.filter;
          }
          this.$store
            .dispatch("call", {
              api: this.service,
              data: data,
            })
            .then((data) => {
              this.message = data.Message;
              if (data.Status == 2) {
                this.setValues(data.data, false)
                this.optionsFiltered = data.recordsFiltered;
                this.optionsTotal = data.recordsTotal;
              } else {
                this.error = true;
              }
            }).catch((error) => {
              console.error("Error:", error);
              this.error = true;
              this.message = error;
            }).finally(() => {
              this.loading = false;
              if (this.reload) {
                this.options = []
                this.load()
                this.reload = false
              }
            });
        }, 300),
      };
    },
    watch: {
      value: function (newValue, oldValue) {
        var component = this;
        var availableOptions = this.options.filter(
          (option) => {
            if (option instanceof Object) {
              return option[component.selectby] == newValue;
            } else {
              return option == newValue;
            }
          }
        );
        if (newValue === oldValue) {
          alert("got same value on value change");
        } else if (availableOptions.length == 1) {
          this.selectedOption = availableOptions[0];
        } else if ((this.selectedOption instanceof Object && this.selectedOption[this.selectby] != newValue) && this.selectedOption != newValue) {
          this.selectedOption = false;
        }
      },
      values: function (newValue, oldValue) {
        this.setValues(newValue, true)
      },
      selectedOption: function (newValue, oldValue) {
        this.$emit('onselect', this.selectedOption)
      },
      service: function (newValue, oldValue) {
        if (newValue && !this.values) {
          this.load();
        }
      },
      search: function (newValue, oldValue) {
        this.options = [];
        if (!this.values) {
          this.load();
        } else {
          this.options = this.values.filter((value) => value.includes(newValue))
        }
      },
      filter: function (newValue, oldValue) {
        if (!_.isEqual(newValue, oldValue)) {
          if (!this.loading) {
            this.options = [];
            this.load();
          } else {
            this.reload = true
          }
        }
      },
    },
    methods: {
      select(value) {
        this.selectedOption = value;
        if (this.selectby && value instanceof Object) {
          this.$emit("input", value[this.selectby]);
        } else {
          this.$emit("input", value);
        }
      },
      canload() {
        if (this.filter && Object.keys(this.filter).length) {
          if (!this.filter[Object.keys(this.filter)[0]]) {
            //return false;
          }
        }
        return true;
      },
      setValues(values, reset) {
        if (reset) {
          this.options = []
          this.selectedOption = false
        }
        if (Array.isArray(values)) {
          this.options = this.options.concat(values);
          this.optionsFiltered = this.options.length;
          this.optionsTotal = this.options.length;
          var availableOptions = this.options.filter(
            (option) => {
              if (option instanceof Object) {
                return option[this.selectby] == this.value;
              } else {
                return option == this.value;
              }
            }
          );
          if (availableOptions.length) {
            this.selectedOption = availableOptions[0];
          } else if (this.options.length == 1 && this.select_out_of_single) {
            this.select(this.options[0]);
          }
        }
      },
      loadValue() {
        if (this.output != this.value) {
          if (this.value) {
            this.tags = this.value.split(",");
          } else {
            this.tags = [];
          }
        }
      },
    },
    mounted: function () {
      if (this.values) {
        this.setValues(this.values, true)
      } else {
        this.load();
      }
    },
    template: `{{{template}}}`,
  };
</script>
<template>
  <div :fullpage="false" :loading="loading" class="dropdown">
    <button type="button" class="d-flex centered form-control dropdown-toggle p-0" data-bs-toggle="dropdown"
      aria-expanded="false" style="overflow-x: hidden;">
      <!-- <span class="input-group-text border-0 bg-transparent flex-1"> {{ value }} </span> -->
      <span class="input-group-text border-0 bg-transparent" v-if="!selectedOption">
        {{nothing_selected_text}}
      </span>
      <span class="input-group-text border-0 bg-transparent flex-1"
        v-else-if="typeof(selectedOption) == 'object' && selectedOption[displayby]">
        {{selectedOption[displayby]}}
      </span>
      <span class="input-group-text border-0 bg-transparent"
        v-else-if="typeof(selectedOption) == 'object' && !selectedOption[displayby]">
        Selected empty value
      </span>
      <span class="input-group-text border-0 bg-transparent" v-else>
        {{selectedOption}}
      </span>
      <span class="input-group-text border-0 bg-transparent" v-if="loading">
        <span class="spinner-border spinner-border-sm" role="status">
          <span class="visually-hidden">Loading...</span>
        </span>
      </span>
      <span class="visually-hidden">Toggle Dropdown</span>
      <i class="mdi mdi-chevron-down"></i>
    </button>
    <ul class="dropdown-menu">
      <li v-if="optionsFiltered > limit || search" class="p-1">
        <input type="search" class="form-control" v-model="search" />
      </li>
      <li>
        <a class="dropdown-item" href="#" @click.prevent="select(null)">{{nothing_selected_text}}</a>
      </li>
      <li v-for="(option, index) in options" :key="index">
        <a class="dropdown-item" href="#" v-if="typeof(option) == 'object' && option[displayby]"
          @click.prevent="select(option)">
          {{ option[displayby] }}
        </a>
        <a class="dropdown-item" href="#" v-else-if="typeof(option) == 'object' && !option[displayby]"
          @click.prevent="select(option)">
          Empty Value
        </a>
        <a class="dropdown-item" href="#" v-else @click.prevent="select(option)">
          {{option}}
        </a>
      </li>
      <li v-if="loading" class="d-flex centered">
        Loading
        <div class="spinner-border spinner-border-sm" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
      </li>
      <li>
        <hr class="dropdown-divider" />
      </li>
      <li class="d-flex flex-column px-1">
        <button v-if="optionsFiltered > options.length" type="button" class="btn btn-primary" @click.prevent="load">
          Load More
        </button>
        <span v-if="options.length && optionsFiltered > options.length" class="text-muted">shown {{ options.length }}
          out of {{ optionsFiltered }}
          <span v-if="optionsTotal > optionsFiltered">({{ optionsTotal }})</span></span>
        <span v-if="message" :class="{ 'text-danger': error }">{{ message }}</span>
      </li>
    </ul>
  </div>
</template>