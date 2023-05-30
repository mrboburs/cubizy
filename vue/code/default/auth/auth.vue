<script>
    export default {
        data: () => {
            return {
                loading: false,
                submitted: false,
                error: false,
                message: "",
                logo : window.application.Account.WideLogo,
                theme_color : window.application.Account.ThemeColor,
                slides : [],
                glide : false,
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
        methods: {
            load() {
                var component = this
                component.loading = true
                this.$store.dispatch('call', {
                    api: "loginpagesliders",
                    data: {}
                }).then((data) => {
                    component.message = data.Message;
                    if (data.Status == 2) {
                        this.slides = data.data
                        if(this.slides && this.slides.length){
                            setTimeout(() => {
                                this.glide = new Glide('.glide', {
                                    autoplay: 7000,
                                    hoverpause: true,
                                    perView: 1,
                                    type: 'carousel' ,
                                }).mount()
                            }, 300)
                        }
                    } else {
                        component.error = true
                    }
                }).catch((error) => {
                    console.error('Error:', error);
                    component.error = true
                    component.message = error
                }).finally(() => {
                    component.loading = false
                })
            },
        },
        mounted: function () {
            this.load()
        },
        template: `{{{template}}}`
    }
</script>
<template>
    <div class="auth-fluid">
        <!-- Auth fluid right content -->
        <div class="auth-fluid-right">

            <div class="glide" style="height: 100%;">
                <div class="glide__track" data-glide-el="track">
                    <ul class="glide__slides">
                        <li v-for="(slide,index) in slides" class="glide__slide" :style="{ 'background-image' :'url('+slide.Image+')'}">
                                <div class="auth-user-testimonial" >
                                    <h3 class="mb-3 text-white">{{slide.Title}}</h3>
                                    <p v-if="slide.Content" class="lead fw-normal">
                                        <i class="mdi mdi-format-quote-open"></i>
                                        {{slide.Content}}
                                        <i class="mdi mdi-format-quote-close"></i>
                                    </p>
                                    <h5 v-if="slide.Footer" class="text-white">
                                        - {{slide.Footer}}
                                    </h5>
                                </div> <!-- end auth-user-testimonial-->
                        </li>
                    </ul>
                </div>
                <div class="glide__bullets" data-glide-el="controls[nav]">
                    <button v-for="(slide,index) in slides" class="glide__bullet" :data-glide-dir="'='+index"></button>
                </div>
            </div>
        </div>
        <!-- end Auth fluid right content -->

        <!--Auth fluid left content -->
        <div class="auth-fluid-form-box">
            <div class="align-items-center d-flex h-100">
                <div class="card-body">

                    <!-- Logo -->
                    <div class="auth-brand text-center text-lg-start"  :style="{'background-color': theme_color?theme_color:'#fff'}">
                        <div class="auth-logo">
                            <a href="/" class="logo text-center">
                                <span class="logo-lg">
                                    <img :src="logo" alt="" height="100">
                                </span>
                            </a>
                        </div>
                    </div>
                    <transition name="custom-classes-transition" mode="out-in"
                        enter-active-class="animate__animated animate__slideInRight animate__faster">
                        <router-view></router-view>
                    </transition>
                </div> <!-- end .card-body -->
            </div> <!-- end .align-items-center.d-flex.h-100-->
        </div>
        <!-- end auth-fluid-form-box-->
    </div>
</template>