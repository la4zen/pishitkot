<template>
    <section class="container">
        Панель задач
        <div class="tasks"> 
            <li v-for="task in tasks">
                <nuxt-link :to="task[1]"> {{ task[0]}} </nuxt-link>
            </li>
        </div>
        <nuxt-link to="/playground"> Песочница </nuxt-link>
    </section>
</template>
<script>
export default {
    data : () => ({
        tasks : [
            ["Арифметические операции", "/tasks/1"],
            ["Условия", "/tasks/2"],
            ["Циклы", "/tasks/3"],
            ["Функции", "/tasks/4"],
            ["Фукции, возвращающие результат", "/tasks/5"]
        ]
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