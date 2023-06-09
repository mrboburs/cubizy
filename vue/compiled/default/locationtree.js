
    export default {
        props: {
            value: {
                type: Object,
                default: function () {
                    return null;
                },
            },
        },
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                locationsearch: "",
                filtered_locations: [],
                countries: [
                    {
                        Country: 'Mauritius',
                        Districts: [
                            {
                                Country: 'Mauritius',
                                District: 'Riviere du Rempart',
                                Localities: [
                                    {
                                        Locality: 'Poudre DOr Hamlet',
                                        SubLocalities: [
                                            {
                                                Country: 'Mauritius',
                                                District: 'Riviere du Rempart',
                                                Locality: 'Poudre DOr Hamlet',
                                                SubLocality: 'Forbach Branch',
                                                Codes: [
                                                    {
                                                        Country: 'Mauritius',
                                                        District: 'Riviere du Rempart',
                                                        Locality: 'Poudre DOr Hamlet',
                                                        SubLocality: 'Forbach Branch',
                                                        Code: '31001',
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            }
                        ],
                    },
                ],
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
            filtered_locations: function (newValue, oldValue) {
                this.setLocations()
            },
            locations: function (newValue, oldValue) {
                this.filterLocations()
            },
            locationsearch: function (newValue, oldValue) {
                this.filterLocations()
            },
            $route(to, from) {
                this.init()
            },
        },
        computed: {
            ...Vuex.mapState(['user', 'account', 'locations']),
            country: {
                // getter
                get: function () {
                    return this.$store.state.country
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_country', newValue)
                }
            },
        },
        methods: {
            setLocations() {
                this.countries = []
                this.filtered_locations.forEach(location => {
                    var country = this.getCountry(location)
                    var District = this.getDistrict(country, location)
                    var Locality = this.getLocality(District, location)
                    var SubLocality = this.getSubLocality(Locality, location)
                })
                if (this.countries.length == 1) {
                    if (this.countries[0].Districts.length == 1) {
                        if (this.countries[0].Districts[0].Localities.length == 1) {
                            if (this.countries[0].Districts[0].Localities[0].SubLocalities.length == 1) {
                                if (!this.value || this.value.Country != this.countries[0].Country || this.value.District != this.countries[0].Districts[0].District || this.value.Locality != this.countries[0].Districts[0].Localities[0].Locality || this.value.SubLocality != this.countries[0].Districts[0].Localities[0].SubLocalities[0].SubLocality) {
                                    this.setLocation(this.countries[0].Districts[0].Localities[0].SubLocalities[0])
                                }
                            } else {
                                if (!this.value || this.value.Country != this.countries[0].Country || this.value.District != this.countries[0].Districts[0].District || this.value.Locality != this.countries[0].Districts[0].Localities[0].Locality) {
                                    this.setLocation(this.countries[0].Districts[0].Localities[0])
                                }
                            }
                        } else {
                            if (!this.value || this.value.Country != this.countries[0].Country || this.value.District != this.countries[0].Districts[0].District) {
                                this.setLocation(this.countries[0].Districts[0])
                            }
                        }
                    } else {
                        if (!this.value || this.value.Country != this.countries[0].Country) {
                            this.setLocation(this.countries[0])
                        }
                    }
                }
            },
            setLocation(location) {
                var value = {
                    Country: location.Country,
                    District: location.District,
                    Locality: location.Locality,
                    SubLocality: location.SubLocality,
                }
                this.$store.commit('set_location', value)
                this.$emit('input', value)
            },
            filterLocations() {
                var search = this.locationsearch
                if (search) {
                    this.filtered_locations = this.locations.filter((item) => {
                        return item.Country.includes(search) || item.District.includes(search) || item.Locality.includes(search) || item.SubLocality.includes(search) || item.Code.includes(search)
                    })
                } else {
                    this.filtered_locations = this.locations
                }
            },
            getCountry(location) {
                if (!location.Country) {
                    location.Country = "Other"
                }
                var country = this.countries.filter((item) => {
                    return item.Country == location.Country
                })
                if (!country.length) {
                    country = {
                        Country: location.Country,
                        District: '',
                        Locality: '',
                        SubLocality: '',
                        Code: '',
                        Districts: [],
                    }
                    this.countries.push(country)
                } else {
                    country = country[0]
                }
                return country
            },
            getDistrict(Country, location) {
                if (!location.District) {
                    location.District = "Other"
                }
                var District = Country.Districts.filter((item) => {
                    return item.District == location.District
                })
                if (!District.length) {
                    District = {
                        Country: location.Country,
                        District: location.District,
                        Locality: '',
                        SubLocality: '',
                        Code: '',
                        Localities: [],
                    }
                    Country.Districts.push(District)
                } else {
                    District = District[0]
                }
                return District
            },
            getLocality(District, location) {
                if (!location.Locality) {
                    location.Locality = "Other"
                }
                var Locality = District.Localities.filter((item) => {
                    return item.Locality == location.Locality
                })
                if (!Locality.length) {
                    Locality = {
                        Country: location.Country,
                        District: location.District,
                        Locality: location.Locality,
                        SubLocality: '',
                        Code: '',
                        SubLocalities: [],
                    }
                    District.Localities.push(Locality)
                } else {
                    Locality = Locality[0]
                }
                return Locality
            },
            getSubLocality(Locality, location) {
                if (!location.SubLocality) {
                    location.SubLocality = "_"
                }
                var SubLocality = Locality.SubLocalities.filter((item) => {
                    return item.SubLocality == location.SubLocality
                })
                if (!SubLocality.length) {
                    SubLocality = {
                        Country: location.Country,
                        District: location.District,
                        Locality: location.Locality,
                        SubLocality: location.SubLocality,
                        Code: '',
                        Codes: [],
                    }
                    Locality.SubLocalities.push(SubLocality)
                } else {
                    SubLocality = SubLocality[0]
                }
                return SubLocality
            },
            load() {
                if (this.loading) {
                    return
                }
                this.loading = true
                this.$store.dispatch('call', {
                    api: "locations",
                    data: {}
                }).then((data) => {
                    this.message = data.Message;
                    if (data.Status == 2) {
                        this.error = false
                        this.$store.commit('set_locations', data.data)
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
            init() {
                if (this.$route.query.country) {
                    var Location = {
                        Country: decodeURI(this.$route.query.country).replaceAll("_", " ")
                    }
                    if (this.$route.query.district) {
                        Location.District = decodeURI(this.$route.query.district).replaceAll("_", " ")
                        if (this.$route.query.locality) {
                            Location.Locality = decodeURI(this.$route.query.locality).replaceAll("_", " ")
                            if (this.$route.query.sublocality) {
                                Location.SubLocality = decodeURI(this.$route.query.sublocality).replaceAll("_", " ")
                            }
                        }
                    }
                    this.setLocation(Location)
                } else if (!this.value) {
                    if (this.$store.state.location) {
                        this.$emit('input', location)
                    }
                }
            }
        },
        mounted: function () {
            this.init()
            if (!this.locations.length) {
                this.load()
            } else {
                this.filterLocations()
            }
        },
        template: `
    <div class="card">
        <strong class="text-nowrap lead">
            Locations
        </strong>
        <input type="search" class="form-control" placeholder="Search Locations" v-model="locationsearch"
            style="margin-bottom: unset;" @click.prevent.stop="">
        <v-alert v-model="message" :error="error" />
        <divloading :fullpage="false" :loading="loading" class="card-body">
            <div class="list-group location_tree">
                <template v-for="(country, index) in countries">
                    <a :key="'country'+index" v-if="!value || value.Country != country.Country" href="#"
                        class="list-group-item" @click.prevent.stop="setLocation(country)">
                        {{country.Country}} <i v-if="country.Districts" class="fa fa-angle-right"></i>
                    </a>
                    <a :key="'country'+index" v-if="value && value.Country == country.Country" href="#"
                        class="list-group-item active" @click.prevent.stop="setLocation({})">
                        {{country.Country}} <i v-if="country.Districts" class="fa fa-angle-down"></i>
                    </a>
                    <template v-if="value && value.Country == country.Country"
                        v-for="(district, districtindex) in country.Districts">
                        <a :key="'district'+districtindex" v-if="value.District != district.District" href="#"
                            class="list-group-item district pl-1" @click.prevent.stop="setLocation(district)">
                            {{district.District}} <i v-if="district.Localities" class="fa fa-angle-right"></i>
                        </a>
                        <a :key="'district'+districtindex" v-if="value.District == district.District" href="#"
                            class="list-group-item district active pl-1" @click.prevent.stop="setLocation(country)">
                            {{district.District}} <i v-if="district.Localities" class="fa fa-angle-down"></i>
                        </a>
                        <template v-if="value.District == district.District"
                            v-for="(locality, localityindex) in district.Localities">
                            <a :key="'locality'+localityindex" v-if="value.Locality != locality.Locality" href="#"
                                class="list-group-item locality pl-1" @click.prevent.stop="setLocation(locality)">
                                {{locality.Locality}} <i v-if="locality.SubLocalities" class="fa fa-angle-right"></i>
                            </a>
                            <a :key="'locality'+localityindex" v-if="value.Locality == locality.Locality" href="#"
                                class="list-group-item locality active pl-1"
                                @click.prevent.stop="setLocation(district)">
                                {{locality.Locality}} <i v-if="locality.SubLocalities" class="fa fa-angle-down"></i>
                            </a>
                            <template v-if="value.Locality == locality.Locality"
                                v-for="(sublocality, sublocalityindex) in locality.SubLocalities">
                                <a :key="'sublocality'+sublocalityindex"
                                    v-if="value.SubLocality != sublocality.SubLocality" href="#"
                                    class="list-group-item sublocality pl-1"
                                    @click.prevent.stop="setLocation(sublocality)">
                                    {{sublocality.SubLocality}}
                                </a>
                                <a :key="'sublocality'+sublocalityindex"
                                    v-if="value.SubLocality == sublocality.SubLocality" href="#"
                                    class="list-group-item sublocality active pl-1"
                                    @click.prevent.stop="setLocation(locality)">
                                    {{sublocality.SubLocality}}
                                </a>
                            </template>
                        </template>
                    </template>
                </template>
            </div>
        </divloading>
    </div>
`
    }
