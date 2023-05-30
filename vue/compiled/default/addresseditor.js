
  export default {
    props: {
      value: {
        type: Object,
        default: function () {
          return {
            ID: 0,
          };
        },
      },
    },
    data: () => {
      return {
        loading: false,
        submitted: false,
        error: false,
        message: "",

        Title: "New Address",
        Mobile: "",
        AddressLine1: "",
        AddressLine2: "",
        AddressLine3: "",
        LocationID: "",
        Code: "",
        SubLocality: false,
        Locality: false,
        District: false,
        Country: false,
        location: {
          lat: 0,
          lng: 0,
        },
      };
    },
    watch: {
      value: function (newValue, oldValue) {
        if (newValue) {
          this.SetData();
          this.$emit("onset", this.value);
        }
      },
      loading: function (newValue, oldValue) {
        if (newValue) {
          this.error = false;
          this.message = false;
          this.submitted = false;
        }
      },
      Country: function (newValue, oldValue) {
        if (newValue != oldValue) {
          if (newValue != this.value.Country) {
            this.District = false;
          }
        }
      },
      District: function (newValue, oldValue) {
        if (newValue != oldValue) {
          if (newValue != this.value.District) {
            this.Locality = false;
          }
        }
      },
      Locality: function (newValue, oldValue) {
        if (newValue != oldValue) {
          if (newValue != this.value.Locality) {
            this.SubLocality = false;
          }
        }
      },
      SubLocality: function (newValue, oldValue) {
        if (newValue != oldValue) {
          if (newValue != this.value.SubLocality) {
            this.LocationID = false;
          }
        }
      },
    },
    computed: {
      ...Vuex.mapState(["user"]),
      TitleError: function () {
        if (!this.submitted) {
          return false;
        }
        if (!this.Title) {
          return "Please enter title";
        }
      },
      MobileError: function () {
        if (!this.submitted) {
          return false;
        }
        if (!this.Mobile) {
          return "Please enter mobile";
        }
      },
      AddressLine1Error: function () {
        if (!this.submitted) {
          return false;
        }
        if (!this.AddressLine1) {
          return "Please enter AddressLine1";
        }
      },
      LatitudeError: function () {
        if (!this.submitted) {
          return false;
        }
        if (!this.location.lat) {
          return "Please enter Latitude, or Select it on map";
        }
      },
      LongitudeError: function () {
        if (!this.submitted) {
          return false;
        }
        if (!this.location.lng) {
          return "Please enter Longitude, or Select it on map";
        }
      },
    },
    methods: {
      Reset() {
        this.SetData();
        this.$emit("input");
      },
      SetData() {
        if (this.value) {
          this.submitted = false;
          this.Title = this.value.Title;
          this.Mobile = this.value.Mobile;
          this.AddressLine1 = this.value.AddressLine1;
          this.AddressLine2 = this.value.AddressLine2;
          this.AddressLine3 = this.value.AddressLine3;
          this.LocationID = this.value.LocationID;
          this.Country = this.value.Country;
          this.District = this.value.District;
          this.Locality = this.value.Locality;
          this.SubLocality = this.value.SubLocality;
          this.Code = this.value.Code;
          this.location.lat = this.value.Latitude;
          this.location.lng = this.value.Longitude;
        }
      },
      submit() {
        this.submitted = true;
        if (this.TitleError || this.MobileError || this.AddressLine1Error || this.LatitudeError || this.LongitudeError) {
          return;
        }
        this.value.Title = this.Title;
        this.value.Mobile = this.Mobile;
        this.value.AddressLine1 = this.AddressLine1;
        this.value.AddressLine2 = this.AddressLine2;
        this.value.AddressLine3 = this.AddressLine3;
        this.value.LocationID = this.LocationID;
        this.value.Latitude = this.location.lat.toString()
        this.value.Longitude = this.location.lng.toString()
        this.value.Country = this.Country;
        this.value.District = this.District;
        this.value.Locality = this.Locality;
        this.value.SubLocality = this.SubLocality;
        this.value.Code = this.Code;
        this.$emit("input", this.value);
      },
    },
    mounted: function () {
      this.SetData();
      this.$emit("onload");
    },
    template: `
  <form @submit.prevent="submit">
    <div class="row">
      <div class="col">
        <formitem v-if="Title != 'Account Address'" :customLayout="true" name="inputTitle" label="Title"
          :error="TitleError" v-model="Title" />
        <formitem :customLayout="true" name="inputCountry" label="Country" v-model="Country" type="select"
          service="countries" displayby="Country" selectby="Country" />
        <formitem :customLayout="true" name="inputDistrict" label="District" v-model="District" type="select"
          service="districts" :filter="{ country: Country }" displayby="District" selectby="District" />
        <formitem :customLayout="true" name="inputLocality" label="Locality" v-model="Locality" type="select"
          service="localities" :filter="{ district: District }" displayby="Locality" selectby="Locality" />
        <formitem :customLayout="true" name="inputSubLocality" label="SubLocality" v-model="SubLocality" type="select"
          service="sublocalities" :filter="{ locality: Locality }" displayby="SubLocality" selectby="SubLocality" />
        <formitem :customLayout="true" name="inputCode" label="Code" v-model="Code" type="select" service="codes"
          :filter="{ sub_locality: SubLocality }" displayby="Code" selectby="Code" />
        <formitem :customLayout="true" name="inputAddressLine1" label="Address Line 1" :error="AddressLine1Error"
          v-model="AddressLine1" />
        <formitem :customLayout="true" name="inputAddressLine2" label="Address Line 2" v-model="AddressLine2" />
        <formitem :customLayout="true" name="inputAddressLine3" label="Address Line 3" v-model="AddressLine3" />
        <formitem :customLayout="true" name="inputMobile" label="Mobile" :error="MobileError" v-model="Mobile" />
        <formitem :customLayout="true" name="inputLatitude" label="Latitude" v-model="location.lat"
          :error="LatitudeError" />
        <formitem :customLayout="true" name="inputLongitude" label="Longitude" v-model="location.lng"
          :error="LongitudeError" />
      </div>
      <div class="col">
        <p>You can pan, zoom map & drag marker to set location</p>
        <v-map v-model="location" style="height: 450px;"></v-map>
      </div>
    </div>
    <div v-if="$route.name == 'setlocation'" class="d-flex justify-content-between">
      <router-link to="/setup/account" class="btn btn-primary">
        Back
      </router-link>
      <button type="submit" class="btn btn-success" :disabled="loading">
        Next
      </button>
    </div>
    <div v-else class="d-flex centered">
      <button type="submit" class="btn btn-success m-1" :disabled="loading">
        <b-spinner small v-if="loading"></b-spinner>
        Save
      </button>
      <button class="btn btn-danger m-1" @click="Reset">Cancel</button>
    </div>
  </form>
`,
  };
