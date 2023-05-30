
    import { Loader } from '/google/js-api-loader.js';
    /**
     * Loading component
     */
    const loader = new Loader({
        apiKey: "AIzaSyAP-MIF7d5orEyYQqAs7s6Qgi4ZsENeDpE",
        version: "weekly",
        libraries: ["places"]
    });
    export default {
        props: ["value"],
        data: () => ({
            loading: false,
            error: false,
            message: "",
            mapOptions: {
                center: {
                    lat: -20.272862824329096,
                    lng: 57.585073702734384
                },
                zoom: 10
            },
            loader,
            geocoder: null,
            map: null,
        }),
        watch: {
            "value.lat"(newValue, oldValue) {
                if (this.map) {
                    this.reset()
                } else {
                    if (!Number.isNaN(Number.parseFloat(this.value.lat))) {
                        this.mapOptions.center.lat = Number.parseFloat(this.value.lat)
                    }
                }
            },
            "value.lng"(newValue, oldValue) {
                if (this.map) {
                    this.reset()
                } else {
                    if (!Number.isNaN(Number.parseFloat(this.value.lng))) {
                        this.mapOptions.center.lng = Number.parseFloat(this.value.lng)
                    }
                }
            },
        },
        mounted() {
            if (this.value.lat && this.value.lng) {
                this.mapOptions.center.lat = parseFloat(this.value.lat)
                this.mapOptions.center.lng = parseFloat(this.value.lng)
            }
            this.init()
        },
        methods: {
            getFloat(num) {
                if (Number.isNaN(Number.parseFloat(num))) {
                    return 0;
                }
                return parseFloat(num)
            },
            init() {
                this.loading = true
                this.error = false
                this.message = "Loading Map"
                this.loader.load().then(() => {
                    this.message = ""
                    this.map = new google.maps.Map(this.$refs.map, this.mapOptions);
                    this.geocoder = new google.maps.Geocoder();
                    this.reset()
                }).catch(e => {
                    this.error = true
                    this.message = e.message ? e.message : "Unable to conect google map"
                }).finally(() => {
                    this.loading = false
                });
            },
            reset() {
                if (this.map && this.value.lat && this.value.lng) {
                    if (this.Marker) {
                        if (this.Marker.position.lat() == this.value.lat && this.Marker.position.lng() == this.value.lng) {
                            return
                        } else {
                            this.Marker.setMap(null);
                        }
                    }
                    var position = {
                        lat: parseFloat(this.value.lat),
                        lng: parseFloat(this.value.lng),
                    }
                    this.Marker = new google.maps.Marker({
                        position: position,
                        map: this.map,
                        title: "Your Business Location",
                        label: "C",
                        draggable: true,
                        animation: google.maps.Animation.Iq,
                    });
                    this.infowindow = new google.maps.InfoWindow({
                        content: "Your Business Location",
                    });
                    this.Marker.addListener("click", () => {
                        this.infowindow.open(this.map, this.Marker);
                    });
                    this.Marker.addListener("dragend", (event) => {
                        const pos = {
                            lat: event.latLng.lat(),
                            lng: event.latLng.lng(),
                        };
                        this.emit(pos, this)
                    });

                    var bound = this.map.getBounds()
                    var flag = false
                    if (bound) {
                        if (bound.La.g > this.value.lng) {
                            flag = true
                            //var old_g = bound.La.g
                            //bound.La.g = this.value.lng - ((bound.La.i - bound.La.g) / 2)
                            //bound.La.i = bound.La.i - (old_g - bound.La.g)
                        }
                        if (bound.La.i < this.value.lng) {
                            flag = true
                            //var old_i = bound.La.i
                            //bound.La.i = this.value.lng + ((bound.La.i - bound.La.g) / 2)
                            //bound.La.g = bound.La.g + (bound.La.i - old_i)
                        }

                        if (bound.Ua.g > this.value.lat) {
                            flag = true
                        }
                        if (bound.Ua.i < this.value.lat) {
                            flag = true
                        }
                        if (flag) {
                            this.map.setCenter(position)
                        }
                    }
                } else if (this.map && this.Marker) {
                    this.Marker.setMap(null);
                    this.tryGeolocation()
                } else if (this.map && !this.Marker && (!this.value.lat || !this.value.lng)) {
                    this.tryGeolocation()
                }
            },
            tryGeolocation() {
                if (navigator.geolocation) {
                    navigator.geolocation.getCurrentPosition(
                        (position) => {
                            const pos = {
                                lat: position.coords.latitude,
                                lng: position.coords.longitude,
                            };
                            this.emit(pos, this)
                            this.map.setCenter(pos);
                            this.message = "Location Detected by Geolocation service"
                        },
                        () => {
                            this.setDefaultMarker(true)
                        }
                    );
                } else {
                    this.setDefaultMarker(false)
                }
            },
            setDefaultMarker(browserHasGeolocation) {
                this.message = browserHasGeolocation
                    ? "Error: The Geolocation service failed."
                    : "Error: Your browser doesn't support geolocation."
                this.error = true
                this.emit(this.mapOptions.center, this)
            },
            emit:_.debounce(function (pos, component) {
                component.$emit('input', pos)
                component.geocoder.geocode({ location: pos }).then((response) => {
                    if (response.results[0]) {
                        component.$emit('address', response.results[0])
                    } else {
                        console.log("No results found");
                    }
                }).catch((e) => console.log("Geocoder failed due to: " + e));
            }, 300),
        },
        template: `
    <!-- Loader -->
    <divloading :show="loading" rounded="sm" style="height: fit-content;">
        <div ref="map" style="position: relative; overflow: hidden;width: 100%;height: 100%;min-width: 300px;min-height: 500px;display: block;"></div>
        <p v-if="message" :class="{'alert-success': !error, 'alert-danger': error }">
            {{message}}
            <button type="button" class="btn-close" @click.prevent="message = false" aria-label="Close"></button>
        </p>
    </divloading>
`
    };
