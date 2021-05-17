<template>
    <section class="container">
        <h1 v-if="user">
            Почта : {{ user.email }} <br>
            Возраст : {{ user.age }} <br>
            Статус : {{ user.user_type.toLowerCase() }}<br>
            <nuxt-link to="/playground">Песочница</nuxt-link>
            <nuxt-link to="/tasks">Задания</nuxt-link>
        </h1>
        <h1 v-else>
            Пользователя не существует.
        </h1>
    </section>
</template>
<script>
export default {
    data : () => ({
        user : null
    }),
    methods : {
        getUser : function () {
            console.log(this.$route.params.id )
            this.$axios.$get("http://la4z.xyz:8080/api/user?id=" + this.$route.params.id)
                .then((response) => {
                    if (response.exists) {
                        this.user = response.user
                    }
                    console.log(response)
                })
                .catch((err) => {
                    console.log(err)
                })
        }
    },
    mounted : function() {
        this.$axios.setToken(localStorage.getItem("token"), 'Bearer')
        this.getUser()
    }
}
</script>