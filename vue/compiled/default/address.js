
    export default {
        props: {
            address: {
                type: Object,
                default: null
            },
        },
        data: () => {
            return { }
        },
        watch: { },
        methods: { },
        mounted: function () {},
        template: `
    <div v-if="address" class=" col-md-3 border widget widget-text">
        <div class=" widget-inner">
            <h2 class="widget-title maincolor1">CONTACT</h2>
            <div class="textwidget">
                <p>{{address.AddressLine1}}<br>
                   {{address.AddressLine2}}<br>
                   {{address.AddressLine3}}<br>
                   {{address.SubLocality}}, {{address.Locality}}, {{address.District}}<br>
                   {{address.Country}}, {{address.Code}}
                </p>
            </div>
        </div>
    </div>
`
    }
