<template>
    <div class="container px-4 py-5 mx-auto">
    <div class="card card0">
        <div class="d-flex flex-lg-row flex-column-reverse">
            <div class="card card1">
                <div class="row justify-content-center my-auto">
                    <div class="col-md-8 col-10 my-5">
                        <h3 class="mb-5 text-center heading">Мы ПишиКот :D</h3>
                        <h6 class="msg-info">Зайдите в свой аккаунт</h6>
                        <div class="form-group"> <label class="form-control-label text-muted">Логин</label> <input type="text" v-model="login" placeholder="Тут точно логин" class="form-control"> </div>
                        <div class="form-group"> <label class="form-control-label text-muted">Пароль</label> <input type="password" v-model="password" placeholder="Наверное пароль" class="form-control"> </div>
                        
                        <div class="row justify-content-center my-3 px-3"> <button v-on:click="auth()" class="btn-block btn-color">Войти</button> </div>
                        <div class="row justify-content-center my-2"> <a href="#"><small class="text-muted">Забыли пароль? А у нас эндпоинтов на это нет...</small></a> </div>
                    </div>
                </div>
                <div class="bottom text-center mb-5">
                    <p href="#" class="sm-text mx-auto mb-3">Нет аккаунта? <nuxt-link to="/auth/register" class="btn btn-white ml-2"> Скорее регистрируйтесь!</nuxt-link></p>
                </div>
            </div>
            <div class="card card2">
                <div class="my-auto mx-md-5 px-md-5 right">
                    <h3 class="text-white">Развивайтесь с нами</h3> <small class="text-white">{{ funtext }}</small>
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
        phrases : ["Тут точно должен быть какой то текст. Пока не знаем какой :D", "Стараемся работать продуктивнее", "Писал это с vue routing'ом в 4 утра :Р",
            "Хочу спать....", "Но мне ещё вебсокеты делать", "пойду кофе сделаю", "регистрацию закончил, щас буду связывать бекенд с фронтом", "axios не умеет слать json?",
            "оказывается умеет, лол", "а кто такие промисы?", "шутка.", "с авторизацией закончили, приступаем к веб сокетам...", "Добро пожаловать!"]
    }),
    methods : {
        auth : function() {
            this.$axios.$post("http://la4z.xyz:8080/users/login", {
                login:this.login,
                password:this.password
            }).then((response) => {
                console.log(response)
                this.$axios.setToken(response.token, "Bearer")
                localStorage.setItem("token", response.token)
                localStorage.setItem("refresh_token", response.refresh_token)
            }).catch((err) => {
                console.log(err)
                alert("Что то пошло не так.")
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
<style>
body {
    color: #000;
    overflow-x: hidden;
    height: 100%;
    background-image: linear-gradient(to right, #D500F9, #FFD54F);
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
    border: 1px solid #D500F9 !important;
    outline-width: 0;
    font-weight: 400
}

button:focus {
    -moz-box-shadow: none !important;
    -webkit-box-shadow: none !important;
    box-shadow: none !important;
    outline-width: 0
}

.card {
    border-radius: 0;
    border: none
}

.card1 {
    width: 50%;
    padding: 40px 30px 10px 30px
}

.card2 {
    width: 50%;
    background-image: linear-gradient(to right, blue, green)
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
    background-image: linear-gradient(to right, #FFD54F, #D500F9);
    padding: 15px;
    cursor: pointer;
    border: none !important;
    margin-top: 40px
}

.btn-color:hover {
    color: #fff;
    background-image: linear-gradient(to right, #D500F9, #FFD54F)
}

.btn-white {
    border-radius: 50px;
    color: #D500F9;
    background-color: #fff;
    padding: 8px 40px;
    cursor: pointer;
    border: 2px solid #D500F9 !important
}

.btn-white:hover {
    color: #fff;
    background-image: linear-gradient(to right, #FFD54F, #D500F9)
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
        width: 100%
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