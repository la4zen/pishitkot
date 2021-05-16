<template>
    <div class="container px-4 py-5 mx-auto">
    <div class="card card">
        <div class="d-flex flex-lg-row flex-column-reverse">
            <div class="card card1">
                <div class="row justify-content-center my-auto">
                    <div class="col-md-8 col-10 my-5">
                        <h3 class="mb-5 text-center heading">Мы ПишиКот</h3>
                        <hr class="">
                        <h6 class="msg-info">Зайдите в свой аккаунт</h6>
                        <div class="form-group">
                            <label class="form-control-label text-muted">Логин</label>
                            <input type="text" v-model="login" placeholder="Тут точно логин" class="form-control"> 
                        </div>
                        <div class="form-group">
                            <label class="form-control-label text-muted">Пароль</label>
                        <input type="password" v-model="password" placeholder="Наверное пароль" class="form-control"> </div>
                        <div class="row justify-content-center my-3 px-3"> 
                            <button v-on:click="auth()" class="btn btn-white ml-2 mt-3">Войти</button> 
                        </div>
                        <div class="row justify-content-center my-2">
                            <a href="#"><small class="text-muted">Забыли пароль? А у нас эндпоинтов на это нет...</small></a>
                        </div>
                    </div>
                </div>
                <div class="bottom text-center mb-5">
                    <p href="#" class="sm-text mx-auto mb-3">Нет аккаунта? </p>
                    <nuxt-link to="/auth/register" class="btn btn-white ml-2"> Скорее регистрируйтесь!</nuxt-link>
                </div>
            </div>
            <div class="card card2">
                <div class="my-auto mx-md-5 px-md-5 right">
                    <h3 class="text-white">Развивайтесь с нами</h3> 
                    <small class="text-white">
                        {{ funtext }}
                    </small>
                </div>
            </div>
        </div>
    </div>
</div>
</template>
<script>
export default {
    data : () => ({
        funtext : null,
        login: null,
        password : null,
        phrases : ["Объясняем", "Пробуем новое", "Много практикуемся", "Просто попробуйте!"]
    }),
    methods : {
        auth : function() {
            this.$axios.setToken(localStorage.getItem("token"), 'Bearer')
            this.$axios.$post("http://localhost:8080/api/auth/accessible").then((response) => {
                location.href = "/lk/users/"
                return
            })
            this.$axios.$post("http://la4z.xyz:8080/users/login", {
                login:this.login,
                password:this.password
            }).then((response) => {
                console.log(response)
                this.$axios.setToken(response.token, "Bearer")
                localStorage.setItem("token", response.token)
                localStorage.setItem("refresh_token", response.refresh_token)
                location.href = "/lk/users"
            }).catch((err) => {
                console.log(err)
                alert(err)
            })
        }
    },
    mounted : function () {
        let i = 0
        this.phrases.forEach((phrase) => {
            setTimeout(() => {
                this.funtext = phrase
            }, i)
            i+=5000
        })
    }
}
</script>