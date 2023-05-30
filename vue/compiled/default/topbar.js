
    export default {
        data: () => {
            return {
                Youtube: window.application.Account.Youtube,
                Facebook: window.application.Account.Facebook,
                Instagram: window.application.Account.Instagram,
                Pinterest: window.application.Account.Pinterest,
                WhatsApp: window.application.Account.WhatsApp,
                BaseDomin: window.application.BaseDomin,
                Protocol: window.location.protocol,
            }
        },
        watch: { },
        methods: { },
        mounted: function () { },
        template: `
    <div class="header-inner-pages">
        <div class="top flex">
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="topnav-sidebar">
                            <ul class="textwidget">
                                <li v-if="Facebook">
                                    <a :href="Facebook" target="_blank">
                                        <i class="fab fa-facebook"></i>
                                    </a>
                                </li>
                                <li v-if="Youtube">
                                    <a :href="Youtube" target="_blank">
                                        <i class="fab fa-youtube"></i>
                                    </a>
                                </li>
                                <li v-if="Instagram">
                                    <a :href="Instagram" target="_blank">
                                        <i class="fab fa-instagram"></i>
                                    </a>
                                </li>
                                <li v-if="Pinterest">
                                    <a :href="Pinterest" target="_blank">
                                        <i class="fab fa-pinterest"></i>
                                    </a>
                                </li>
                                <li v-if="WhatsApp">
                                    <a :href="WhatsApp" target="_blank">
                                        <i class="fab fa-whatsapp"></i>
                                    </a>
                                </li>
                                <!-- <li>
                                    <a href="#"><i class="fab fa-twitter"></i></a>
                                </li> -->
                            </ul>
                        </div>
                        <nav class="navbar-right navbar menu-top">
                            <ul class="menu clearfix">
                                <li>
                                    <a :href="Protocol + '//student.'+ BaseDomin" target="_blank">
                                        <i class="fa fa-user"></i>
                                        Student Dashboard
                                    </a>
                                </li>
                                <li>
                                    <a :href="Protocol + '//seller.'+ BaseDomin" target="_blank">
                                        <i class="fa fa-user"></i>
                                        Seller Dashboard
                                    </a>
                                </li>
                            </ul><!-- /.menu -->
                        </nav><!-- /.mainnav -->

                    </div><!-- col-md-12 -->
                </div><!-- row -->
            </div><!-- container -->
        </div><!-- Top -->
    </div><!-- header-inner-pages -->
`
    }
