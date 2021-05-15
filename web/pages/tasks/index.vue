<template>
    <section class="container">

    </section>
</template>
<script>
export default {
    data : () => ({
        
    }),
    methods : {

    },
    mounted : function() {
        this.$axios.setToken(localStorage.getItem("token"), 'Bearer')
        this.$axios.$get("http://la4z.xyz:8080/api/auth/accessible").catch((err) => {
            this.$axios.setToken(localStorage.getItem("refresh_token"), 'Bearer')
            this.$axios.$get("http://la4z.xyz:8080/api/auth/refresh_token").then((response) => {
                localStorage.setItem("token", response.token)
                localStorage.setItem("refresh_token", response.refresh_token)
            }).catch((err) => {
                console.log(err)
                location.href = "/auth/login"  
            })
        })
    }
}
</script>