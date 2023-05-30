
    export default {
        props: {
            default_type: {
                default: "courses",
            },
            mode: {
                default: "vertical",  //horizontal
            },
        },
        data: () => {
            return {
                types: [
                    "courses",
                    "events",
                    "blogs",
                    "sellers",
                ],
                TuitionTypes : [],

                searchtype: "courses",
                subjectsearch: "",
                sublevels: [],
                filtered_subjects: [],
                filtered_locations: [],
                AccountType: window.application.Account.AccountType,
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
                loadinglocations: false,
                locationsearch: "",
            }
        },
        components: {
            'tuitiontypes': () => import("/vue/tuitiontypes.js"),
        },
        watch: {
            subjects: function (newValue, oldValue) {
                this.filterSubjects()
            },
            subjectsearch: function (newValue, oldValue) {
                this.filterSubjects()
            },
            filtered_subjects: function (newValue, oldValue) {
                this.setSublevels()
            },
            locationsearch: function (newValue, oldValue) {
                this.filterLocations()
            },
            locations: function (newValue, oldValue) {
                this.filterLocations()
            },
            filtered_locations: function (newValue, oldValue) {
                this.setLocations()
            },
            default_type: function (newValue, oldValue) {
                this.searchtype = this.default_type
            },
            $route(to, from) {
                this.init()
            },
        },
        methods: {
            setLocation(location) {
                if (location && location.Country != undefined) {
                    this.country = location.Country
                    this.district = location.District
                    this.locality = location.Locality
                    this.sublocality = location.SubLocality
                    this.code = location.Code
                } else {
                    this.country = ""
                    this.district = ""
                    this.locality = ""
                    this.sublocality = ""
                    this.code = ""
                }
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
            setLocations() {
                this.countries = []
                this.filtered_locations.forEach(location => {
                    var country = this.getCountry(location)
                    var District = this.getDistrict(country, location)
                    var Locality = this.getLocality(District, location)
                    var SubLocality = this.getSubLocality(Locality, location)
                    var Code = this.getCode(SubLocality, location)
                })
            },
            getCountry(location) {
                if (!location.Country) {
                    location.Country = "_"
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
                    location.District = "_"
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
                    location.Locality = "_"
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
            getCode(SubLocality, location) {
                if (!location.Code) {
                    location.Code = "_"
                }
                var Codes = SubLocality.Codes.filter((item) => {
                    return item.Code == location.Code
                })
                if (!Codes.length) {
                    SubLocality.Codes.push(location)
                    Codes = [location]
                }

                return Codes[0]
            },
            filterSubjects() {
                var search = this.subjectsearch
                if (search) {
                    this.filtered_subjects = this.subjects.filter((item) => {
                        return item.Name.includes(search) || item.SubLevelName.includes(search) || item.LevelName.includes(search)
                    })
                } else {
                    this.filtered_subjects = this.subjects
                }
            },
            setSublevels(_sublevel) {
                var sublevelmap = {}
                this.filtered_subjects.forEach(subject => {
                    if (sublevelmap[subject.SubLevelName]) {
                        sublevelmap[subject.SubLevelName].SubjectCount++
                        sublevelmap[subject.SubLevelName].Subjects.push(subject)
                    } else {
                        sublevelmap[subject.SubLevelName] = {
                            ID: subject.SubLevelID,
                            Name: subject.SubLevelName,
                            LevelName: subject.LevelName,
                            LevelID: subject.LevelID,
                            SessionCount: subject.SubLevelSessionCount,
                            SubjectCount: 1,
                            Subjects: [subject]
                        }
                    }
                });
                sublevelmap = Object.values(sublevelmap)
                sublevelmap.sort(function (a, b) { return b.SessionCount - a.SessionCount });
                sublevelmap.forEach(sublevel => {
                    sublevel.Subjects.sort(function (a, b) { return b.SessionCount - a.SessionCount });
                });
                this.sublevels = sublevelmap
            },
            setSubject(_subject) {
                if (_subject) {
                    this.subject = _subject.Name
                    this.sublevel = _subject.SubLevelName
                    this.level = _subject.LevelName
                } else {
                    this.subject = ""
                    this.sublevel = ""
                    this.level = ""
                }
            },
            getLocations() {
                this.loadinglocations = true
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
                    this.loadinglocations = false
                })
            },
            dosearch() {

                var uri = ""

                if (this.search.trim()) {
                    uri += this.addPostFix(uri) + "search=" + encodeURI(this.search.trim())
                }
                if (this.searchtype != "blogs") {
                    if (this.searchtype != "events") {
                        if (this.level) {
                            uri = this.addPostFix(uri) + "level=" + this.level
                        }
                        if (this.sublevel) {
                            uri = this.addPostFix(uri) + "sub_level=" + this.sublevel
                        }
                        if (this.subject) {
                            uri = this.addPostFix(uri) + "subject=" + this.subject
                        }
                    }
                    if (this.country) {
                        uri = this.addPostFix(uri) + "country=" + this.country
                    }
                    if (this.district) {
                        uri = this.addPostFix(uri) + "district=" + this.district
                    }
                    if (this.locality) {
                        uri = this.addPostFix(uri) + "locality=" + this.locality
                    }
                    if (this.sublocality) {
                        uri = this.addPostFix(uri) + "sublocality=" + this.sublocality
                    }
                    if (this.code) {
                        uri = this.addPostFix(uri) + "code=" + this.code
                    }
                }
                uri = "/" + this.searchtype + "/" + uri
                if (uri) {
                    if (this.$route.path != "/" && uri.includes(this.$route.path)) {
                        this.$emit('search')
                    } else {
                        this.$router.push(uri)
                    }
                }
            },
            addPostFix(uri) {
                if (uri.includes("?")) {
                    uri += "&"
                } else {
                    uri += "?"
                }
                return uri
            },
            init() {
                var flag = false
                if (this.$route.query.level && this.$route.query.level != this.level) {
                    this.level = this.$route.query.level
                    flag = true
                }
                if (this.$route.query.sublevel && this.$route.query.sublevel != this.sublevel) {
                    this.sublevel = this.$route.query.sublevel
                    flag = true
                }
                if (this.$route.query.subject && this.$route.query.subject != this.subject) {
                    this.subject = this.$route.query.subject
                    flag = true
                }

                if (this.$route.query.country && this.$route.query.country != this.country) {
                    this.country = this.$route.query.country
                    flag = true
                }
                if (this.$route.query.district && this.$route.query.district != this.district) {
                    this.district = this.$route.query.district
                    flag = true
                }
                if (this.$route.query.locality && this.$route.query.locality != this.locality) {
                    this.locality = this.$route.query.locality
                    flag = true
                }
                if (this.$route.query.sublocality && this.$route.query.sublocality != this.sublocality) {
                    this.sublocality = this.$route.query.sublocality
                    flag = true
                }
                if (this.$route.query.code && this.$route.query.code != this.code) {
                    this.code = this.$route.query.code
                    flag = true
                }
                if (this.$route.query.search && this.$route.query.search != this.search) {
                    this.search = this.$route.query.search
                    flag = true
                }
                if (flag) {
                    this.dosearch()
                }
            }
        },
        computed: {
            ...Vuex.mapState(['subjects', 'locations']),
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
            district: {
                // getter
                get: function () {
                    return this.$store.state.district
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_district', newValue)
                }
            },
            locality: {
                // getter
                get: function () {
                    return this.$store.state.locality
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_locality', newValue)
                }
            },
            sublocality: {
                // getter
                get: function () {
                    return this.$store.state.sublocality
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_sublocality', newValue)
                }
            },
            code: {
                // getter
                get: function () {
                    return this.$store.state.code
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_code', newValue)
                }
            },
            level: {
                // getter
                get: function () {
                    return this.$store.state.level
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_level', newValue)
                }
            },
            sublevel: {
                // getter
                get: function () {
                    return this.$store.state.sublevel
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_sublevel', newValue)
                }
            },
            subject: {
                // getter
                get: function () {
                    return this.$store.state.subject
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_subject', newValue)
                }
            },
            search: {
                // getter
                get: function () {
                    return this.$store.state.search
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_search', newValue)
                }
            },
        },
        mounted: function () {
            this.searchtype = this.default_type
            if (!this.locations.length) {
                this.getLocations()
            } else {
                this.filterLocations()
            }
            this.filterSubjects()
            this.init()
        },
        template: `
    <div class="sticky p-1">
        <div :class="{'container':(mode == 'horizontal')}">
            <form class="search_row" :class="{'search_column' :(mode == 'vertical')}" @submit.prevent="dosearch">
                <tuitiontypes class="my-1" v-model="TuitionTypes"></tuitiontypes>
                
                <div class="dropdown" v-if="searchtype == 'courses' || searchtype == 'sellers' ">
                    <button class="btn btn-default dropdown-toggle capitalize" type="button" id="dropdownMenu1"
                        data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span v-if="!subject && !sublevel && !level " class="selected_text">All Subjects</span>
                        <span v-else class="selected_text">{{subject}}-{{sublevel}}-{{level}}</span>
                        <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <li @click.stop.prevent>
                            <input class="form-control" id="inputSearchSubject" placeholder="Search Subject"
                                v-model="subjectsearch" style="margin: 0.5rem; width: auto;">
                        </li>
                        <li role="separator" class="divider"></li>
                        <li>
                            <a class="text-capitalize" href="#" @click.prevent="setSubject()">all
                                subjects</a>
                        </li>
                        <template v-for="_sublevel in sublevels">
                            <li>
                                <a class="text-capitalize" href="#"
                                    @click.prevent="setSubject(_sublevel)"><strong>{{_sublevel.Name}}</strong><span
                                        class="caret"></span></a>
                            </li>
                            <li v-for="_subject in _sublevel.Subjects">
                                <a class="text-capitalize" href="#"
                                    @click.prevent="setSubject(_subject)">{{_subject.Name}}</a>
                            </li>
                        </template>
                    </ul>
                </div>
                <div class="dropdown"
                    v-if="searchtype == 'events' || searchtype == 'courses' || searchtype == 'sellers' ">
                    <button class="btn btn-default dropdown-toggle text-nowrap capitalize" type="button"
                        id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span>@</span>
                        <span v-if="!code && !sublocality && !locality && ! country" class="selected_text">All
                            Locations</span>
                        <span v-else class="selected_text"> {{code}}-{{sublocality}}-{{locality}}-{{country}}</span>
                        <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <li @click.stop.prevent>
                            <input class="form-control" id="inputSearchLocation" placeholder="Search Location"
                                v-model="locationsearch" style="margin: 0.5rem; width: auto;">
                        </li>
                        <li role="separator" class="divider"></li>
                        <li>
                            <a class="text-capitalize" href="#" @click.prevent="setLocation()">all
                                locations</a>
                        </li>
                        <li v-if="loadinglocations">
                            <span aria-hidden="true" class="spinner-border">
                                <!---->
                            </span>
                        </li>
                        <template v-for="_country in countries">
                            <li class="text-nowrap">
                                <a class="text-uppercase" href="#" @click.prevent="setLocation(_country)">
                                    <strong>{{_country.Country}}</strong>
                                    <span class="caret"></span>
                                </a>
                            </li>
                            <template v-for="_district in _country.Districts">
                                <li class="text-nowrap">
                                    <a class="text-capitalize" href="#" @click.prevent="setLocation(_district)"
                                        style="margin-left: 0.5rem;">
                                        <strong>{{_district.District}}</strong>
                                        <span class="caret"></span>
                                    </a>
                                </li>
                                <template v-for="_locality in _district.Localities">
                                    <li class="text-nowrap">
                                        <a class="text-capitalize" href="#" @click.prevent="setLocation(_locality)"
                                            style="margin-left: 1rem;">
                                            {{_locality.Locality}}
                                            <span class="caret"></span>
                                        </a>
                                    </li>
                                    <template v-for="_sublocality in _locality.SubLocalities">
                                        <li class="text-nowrap">
                                            <a class="text-capitalize" href="#" style="margin-left: 1.5rem;"
                                                @click.prevent="setLocation(_sublocality)">
                                                {{_sublocality.SubLocality}}
                                                <span class="caret"></span>
                                            </a>
                                        </li>
                                        <template v-for="_code in _sublocality.Codes">
                                            <li class="text-nowrap">
                                                <a class="text-capitalize" href="#" style="margin-left: 2rem;"
                                                    @click.prevent="setLocation(_code)">
                                                    {{_code.Code}}
                                                </a>
                                            </li>
                                        </template>
                                    </template>
                                </template>
                            </template>
                        </template>
                    </ul>
                </div>
                <input class="form-control" id="inputEmail3" placeholder="Search phrase" v-model="search">
                <button type="submit" class="btn btn-primary">Search</button>
            </form>
        </div>
    </div>
`
    }
