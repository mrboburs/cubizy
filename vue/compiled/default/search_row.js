
    export default {
        props: {
            default_type: {
                default: "courses",
            },
            mode: {
                default: "vertical",  //horizontal
            },
            type_locked: {
                default: false,
            }
        },
        components: {
            'locationtree': () => import("/vue/locationtree.js"),
        },
        data: () => {
            return {
                types: [
                    "courses",
                    "events",
                    "blogs",
                    "sellers",
                ],
                searchtype: "courses",
                AccountType: window.application.Account.AccountType,
                Level: null,
                SubLevel: null,
                Subject: null,
            }
        },
        watch: {

            Level: function (newValue, oldValue) {
                this.SubLevel = null
            },
            default_type: function (newValue, oldValue) {
                this.searchtype = this.default_type
            },
            $route(to, from) {
                this.init()
            },
        },
        methods: {
            dosearch: _.debounce(function () {
                this.do_search();
            }, 1000),
            do_search() {
                var uri = ""
                if (this.searchtype != "blogs") {
                    if (this.searchtype != "events") {
                        if (this.Level) {
                            uri = this.getUrlName(this.Level)
                            if (this.SubLevel) {
                                uri += "/" + this.getUrlName(this.SubLevel)
                                if (this.Subject) {
                                    uri += "/" + this.getUrlName(this.Subject)
                                }
                            }
                        }
                    }
                    if (this.Location && this.Location.Country) {
                        uri = this.addPostFix(uri) + "country=" + encodeURI(this.Location.Country)
                    }
                    if (this.Location && this.Location.District) {
                        uri = this.addPostFix(uri) + "district=" + encodeURI(this.Location.District)
                    }
                    if (this.Location && this.Location.Locality) {
                        uri = this.addPostFix(uri) + "locality=" + encodeURI(this.Location.Locality)
                    }
                    if (this.Location && this.Location.SubLocality) {
                        uri = this.addPostFix(uri) + "sublocality=" + encodeURI(this.Location.SubLocality)
                    }
                }
                if (this.search.trim()) {
                    uri = this.addPostFix(uri) + "search=" + encodeURI(this.search.trim())
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
            getUrlName(level) {
                var name = level.Name
                return encodeURI(name.toLowerCase().replaceAll(' ', '_'))
            },
            init() {
                var flag = false
                if (this.$route.params.level) {
                    var level = decodeURI(this.$route.params.level).replaceAll("_", " ")
                    Levels = this.levels.filter((item) => {
                        return item.Name == level
                    })
                    if (Levels.length == 1) {
                        this.Level = Levels[0]
                    }
                    if (this.$route.params.sublevel) {
                        var sublevel = decodeURI(this.$route.params.sublevel).replaceAll("_", " ")
                        SubLevels = this.Level.SubLevels.filter((item) => {
                            return item.Name == sublevel
                        })
                        if (SubLevels.length == 1) {
                            this.SubLevel = SubLevels[0]
                        }
                        if (this.$route.params.subject) {
                            var subject = decodeURI(this.$route.params.subject).replaceAll("_", " ")
                            Subjects = this.SubLevel.Subjects.filter((item) => {
                                return item.Name == subject
                            })
                            if (Subjects.length == 1) {
                                this.Subject = Subjects[0]
                            }
                        }
                    }
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
            ...Vuex.mapState(['levels']),
            Location: {
                // getter
                get: function () {
                    return this.$store.state.location
                },
                // setter
                set: function (newValue) {
                    this.$store.commit('set_location', newValue)
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
            //this.init()
        },
        template: `
    <div style="background-color: #eaeaea;">
        <div :class="{'container':(mode == 'horizontal')}">
            <form class="search_row" :class="{'search_column' :(mode == 'vertical')}" @submit.prevent="dosearch">
                <strong class="text-nowrap lead" v-if="!type_locked">
                    I am Looking for .....
                </strong>
                <div v-if="AccountType == 'admin' && !type_locked" class="dropdown">
                    <button class="btn btn-default dropdown-toggle text-capitalize" type="button" id="dropdownMenu1"
                        data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span class="selected_text">{{searchtype}}</span>
                        <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <li v-for="_type in types">
                            <a class="text-capitalize" href="#" @click.prevent="searchtype = _type">{{_type}}</a>
                        </li>
                    </ul>
                </div>
                <div class="dropdown" v-if="searchtype == 'courses' || searchtype == 'sellers' ">
                    <button class="btn btn-default dropdown-toggle capitalize" type="button" id="dropdownMenu1"
                        data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span class="selected_text">
                            <template v-if="!SubLevel">All</template>
                            <template v-else>{{SubLevel.Name}}</template>
                            <template v-if="Level">{{Level.Name}}</template>
                            <template v-else>Levels</template>
                        </span>
                        <span class="caret"></span>
                    </button>
                    <div class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <div class="list-group location_tree">
                            <template v-for="(level, index) in levels">
                                <a :key="'level'+index" v-if="!Level || Level.Name != level.Name" href="#"
                                    class="list-group-item" @click.prevent.stop="Level = level">
                                    {{level.Name}} <i v-if="level.SubLevels.length" class="fa fa-angle-right"></i>
                                </a>
                                <a :key="'level'+index" v-if="Level && Level.Name == level.Name" href="#"
                                    class="list-group-item active" @click.prevent.stop="Level = null">
                                    {{level.Name}} <i v-if="level.SubLevels" class="fa fa-angle-down"></i>
                                </a>
                                <template v-if="Level && Level.Name == level.Name && Level.SubLevels"
                                    v-for="(sublevel, sublevelindex) in level.SubLevels">
                                    <a :key="'sublevel'+sublevelindex"
                                        v-if="!SubLevel || SubLevel.Name != sublevel.Name" href="#"
                                        class="list-group-item district pl-1" @click.prevent.stop="SubLevel = sublevel">
                                        {{sublevel.Name}}
                                    </a>
                                    <a :key="'sublevel'+sublevelindex" v-if="SubLevel && SubLevel.Name == sublevel.Name"
                                        href="#" class="list-group-item district active pl-1"
                                        @click.prevent.stop="SubLevel = null">
                                        {{sublevel.Name}}
                                    </a>
                                </template>
                            </template>
                        </div>
                    </div>
                </div>
                <div class="dropdown" v-if="searchtype == 'courses' || searchtype == 'sellers' ">
                    <button class="btn btn-default dropdown-toggle capitalize" type="button" id="dropdownMenu1"
                        data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span v-if="!Subject" class="selected_text">All Subjects</span>
                        <span v-else class="selected_text">{{Subject.Name}}</span>
                        <span class="caret"></span>
                    </button>
                    <div v-if="SubLevel" class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <a href="#" class="list-group-item district pl-1" @click.prevent.stop="Subject = null"> All
                            Subjects </a>
                        <template v-if="SubLevel.Subjects" v-for="(subject, subjectindex) in SubLevel.Subjects">
                            <a :key="'subject'+subjectindex" v-if="!Subject || Subject.Name != subject.Name" href="#"
                                class="list-group-item district pl-1" @click.prevent.stop="Subject = subject">
                                {{subject.Name}}
                            </a>
                            <a :key="'subject'+subjectindex" v-if="Subject && Subject.Name == subject.Name" href="#"
                                class="list-group-item district active pl-1" @click.prevent.stop="Subject = null">
                                {{subject.Name}}
                            </a>
                        </template>
                    </div>
                    <div v-else class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <a href="#" class="list-group-item district pl-1" @click.prevent.stop="Subject = null"> All
                            Subjects </a>
                    </div>
                </div>
                <div class="dropdown"
                    v-if="searchtype == 'events' || searchtype == 'courses' || searchtype == 'sellers' ">
                    <button class="btn btn-default dropdown-toggle text-nowrap capitalize" type="button"
                        id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        <span v-if="!Location || !Location.Country" class="selected_text">All
                            Locations</span>
                        <span v-else class="selected_text">
                            at
                            <template v-if="Location.SubLocality">
                                {{Location.SubLocality }}in
                            </template>
                            <template v-if="Location.Locality && Location.SubLocality != Location.Locality">
                                {{Location.Locality }} in
                            </template>
                            <template v-if="Location.District && Location.District != Location.Locality">
                                {{Location.District }} in
                            </template>
                            <template v-if="Location.Country">
                                {{Location.Country }}
                            </template>
                        </span>
                        <span class="caret"></span>
                    </button>
                    <div class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <locationtree class="my-1" v-model="Location"></locationtree>
                    </div>
                </div>
                <input class="form-control" id="inputEmail3" placeholder="Search phrase" v-model="search">
                <button type="submit" class="btn btn-primary">Search</button>
            </form>
        </div>
    </div>
`
    }
