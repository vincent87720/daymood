<script>
export default {
    name: 'Alert',
    data() {
        return {
            alertSuccess: false,
            alertFail: false,
        };
    },
    props: {
        prop_alert: {
            type: Boolean,
            required: true
        },
        prop_alertType: {
            type: String,
            required: true
        },
        prop_alertText: {
            type: String,
            required: true
        },
    },
    computed: {
        alert: {
            get() {
                return this.prop_alert
            },
            set(val) {
                this.$emit('update:prop_alert', val)

            }
        },
        alertType: {
            get() {
                return this.prop_alertType
            },
            set(val) {
                this.$emit('update:prop_alertType', val)

            }
        },
        alertText: {
            get() {
                return this.prop_alertText
            },
            set(val) {
                this.$emit('update:prop_alertText', val)

            }
        },
    },
    watch: {
        alert: function () {
            let vthis = this;
            if (vthis.alertType == "Success") {
                vthis.alertSuccess = true;
                var timer = setInterval(function () {
                    vthis.alert = false;
                    vthis.alertSuccess = false;
                    vthis.alertType = "";
                    vthis.alertText = "";
                    clearInterval(timer);
                }, 3000);
            }
            else if (vthis.alertType == "Fail") {
                vthis.alertFail = true;
                var timer = setInterval(function () {
                    vthis.alert = false;
                    vthis.alertFail = false;
                    vthis.alertType = "";
                    vthis.alertText = "";
                    clearInterval(timer);
                }, 3000);
            }
        },
        //   alertSuccess: function(){

        //     let vthis = this;

        //     if(vthis.alertSuccess == true){
        //       var timer = setInterval(function() {
        //         vthis.alertSuccess = false;
        //         clearInterval(timer);
        //       }, 3000);
        //     }
        //   },
        //   alertFail: function(){

        //     let vthis = this;

        //     if(vthis.alertFail == true){
        //       var timer = setInterval(function() {
        //         vthis.alertFail = false;
        //         clearInterval(timer);
        //       }, 3000);
        //     }
        //   },
    }

}

</script>

<style scoped src="./style.css">

</style>
<template src="./template.html"></template>