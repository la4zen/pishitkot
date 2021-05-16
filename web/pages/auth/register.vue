<template>
    <div class="card card0">
        <div class="d-flex flex-lg-row flex-column-reverse">
            <div class="card card1">
                <div class="row justify-content-center my-auto">
                    <div class="col-md-8 col-10 my-5">
                        <h3 class="mb-5 text-center heading">Мы ПишиКот :d</h3>
                        <hr class="">
                        <h6 class="msg-info">Класс, вот ваша форма</h6>
                        <div class="form-group"> <label class="form-control-label text-muted">Логин</label> <input type="text" v-model="login" placeholder="Гениальный логин" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Пароль</label> <input type="password" v-model="password" placeholder="Ваш гиперсуперпупер сложный пароль" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Имя</label> <input type="text" v-model="first_name" placeholder="Красивое" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Фамилия</label> <input type="text" v-model="last_name" placeholder="Крутая" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Email</label> <input type="email" v-model="email" placeholder="Какие то символы" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Возраст</label> <input type="number" v-model.number="age" placeholder="Только честно" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Вы кто?</label> 
                            <select v-model="user_type" class="form-control"> 
                                <option selected value="Ученик">Ученик</option>
                                <option selected value="Учитель">Учитель</option>
                                <option selected value="Родитель">Родитель</option>
                            </select></div>
                        <div class="row justify-content-center my-3 px-3"> <button v-on:click="register()" class="btn-block btn-color">Войти</button> </div>
                    </div>
                </div>
                <div class="bottom text-center mb-5">
                    <p href="#" class="sm-text mx-auto mb-3">Есть аккаунт? <nuxt-link to="/auth/login" class="btn btn-white ml-2"> Тогда зачем вы сюда заходили?</nuxt-link></p>
                </div>
            </div>
            <div class="card card2">
                <div class="my-auto mx-md-5 px-md-5 right">
                    <h3 class="text-white">Развивайтесь с нами</h3> <small class="text-white">{{ funtext }}</small>
                </div>
            </div>
        </div>
        <error/>
    </div>
</template>
<script>
import error from "@/components/error"
export default {
    components : {
        error
    },
    data : () => ({
        funtext : null,
        login: null,
        email : null,
        password : null,
        first_name : null,
        last_name : null,
        age : null,
        user_type : null,
        phrases : ["Начинаем по новому"]
    }),
    methods : {
        register : function() {
            this.$axios.setToken(localStorage.getItem("token"), 'Bearer')
            this.$axios.$post("http://localhost:8080/api/auth/accessible").then((response) => {
                location.href = "/lk/users/"
                return
            })
            this.$axios.$post("http://localhost:8080/users/register", {
                login:this.login,
                age:this.age,
                first_name:this.first_name,
                last_name:this.last_name,
                email:this.email,
                password:this.password,
                user_type:this.user_type
            }).then((response) => {
                console.log(response)
                this.$axios.setToken(response.data.token, "Bearer")
                localStorage.setItem("token", response.data.token)
                localStorage.setItem("refresh_token", response.data.refresh_token)
                if (this.user_type === "Ученик") {
                    location.href = "/start"
                }
            }).catch((err) => {
                alert(err)
                error.methods.sendError("err")
            })
        },
    },
    mounted : function () {
        let i = 0
        this.phrases.forEach((phrase) => {
            setTimeout(() => {
                this.funtext = phrase
            }, i)
            i+=10000
        })
    }
}
</script>
<style>
.text2 {
    color: black;
    font: small-caps;
}

body {
    color: #619ADA;
    overflow-x: hidden;
    height: 100%;
    background-image: linear-gradient(to right, #619ADA, #619ADA);
    background-repeat: no-repeat
}

input,
textarea {
    background-color: #F3E5F5;
    border-radius: 50px !important;
    padding: 12px 15px 12px 15px !important;
    width: 100%;
    box-sizing: border-box;
    border: none !important;
    border: 1px solid #F3E5F5 !important;
    font-size: 16px !important;
    color: #000 !important;
    font-weight: 400
}

input:focus,
textarea:focus {
    -moz-box-shadow: none !important;
    -webkit-box-shadow: none !important;
    box-shadow: none !important;
    border: 1px solid #C25E55 !important;
    outline-width: 0;
    font-weight: 400
}

button:focus {
    -moz-box-shadow: none !important;
    -webkit-box-shadow: none !important;
    box-shadow: none !important;
    outline-width: 0
}

.card0 {
    border-radius: 1;
    border: none
}

.card1 {
    width: 50%;
    padding: 40px 30px 10px 30px
}

.card2 {
    width: 50%;
    margin: 4px;
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    background-image: linear-gradient(232deg, #dd4638, #619ADA)
}

.card2:hover{
    background-image: linear-gradient(297deg, #dd4638, #619ADA)
}

#logo {
    width: 70px;
    height: 60px
}

.heading {
    margin-bottom: 60px !important
}

::placeholder {
    color: #000 !important;
    opacity: 1
}

:-ms-input-placeholder {
    color: #000 !important
}

::-ms-input-placeholder {
    color: #000 !important
}

.form-control-label {
    font-size: 12px;
    margin-left: 15px
}

.msg-info {
    padding-left: 15px;
    margin-bottom: 30px
}

.btn-color {
    border-radius: 50px;
    color: #fff;
    padding: 15px;
    cursor: pointer;
    transition: 0.5s;
    background-image: linear-gradient(to right, #619ADA, #619ADA);
    border: none !important;
    margin-top: 40px
}

.btn-color:hover {
    color: #fff;
    background-image: linear-gradient(to right, #C25E55, #C25E55)
}

.sm-text{
    color: black;
}

.btn-white {
    border-radius: 50px;
    color: black;
    transition: 0.2s cubic-bezier(.17,.67,.91,.51);
    background-color: #fff;
    padding: 8px 40px;
    cursor: pointer;
    border: 2px solid #619ADA !important
}

.btn-white:hover {
    color: #fff;
    background-image: linear-gradient(to right, #C25E55, #C25E55)
}

a {
    color: #000
}

a:hover {
    color: #000
}

.bottom {
    width: 100%;
    margin-top: 50px !important
}

.sm-text {
    font-size: 15px
}

@media screen and (max-width: 992px) {
    .card1 {
        width: 100%;
        padding: 40px 30px 10px 30px
    }

    .card2 {
        width: 0%
    }

    .right {
        margin-top: 100px !important;
        margin-bottom: 100px !important
    }
}

@media screen and (max-width: 768px) {
    .container {
        padding: 10px !important
    }

    .card2 {
        padding: 50px
    }

    .right {
        margin-top: 50px !important;
        margin-bottom: 50px !important
    }
}
</style>