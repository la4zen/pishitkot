<template>
    <div class="container">
        <div class="loading" v-if="!loaded">
            Загрузочка...
        </div>
        <div v-else>
            <div class="editorMain">
                <div class="editorlines">
                    <div v-for="(_, index) in code.split('\n')">
                        {{ ++index }}
                    </div>
                </div>
                <div class="editor">
                    <textarea autocomplete="off" v-model="code" :disabled="editorLock"/>
                </div>
            </div>
            <button class="btn btn-primary" @click="checkTask()"> Проверить </button> 
            <div class="output">
                {{ output }}
            </div>
            <div class="error" v-if="error"> 
                В коде ошибка!
                <pre>{{ error }}</pre>
            </div>
            <div class="linter" v-if="stepScene < 17 && stepScene != 5 && stepScene != 9 && stepScene != 12 && stepScene != 14">
                <div v-if="stepScene == 0">
                    Привет, {{ user.first_name }}!
                    <button class="btn btn-primary" @click="stepScene++">
                        Привет
                    </button>
                </div>
                <div v-if="stepScene == 1">
                    Ты хочешь научиться программированию для создания разных интересных штук?
                    <button class="btn btn-primary" @click="stepScene++">
                        Конечно
                    </button>
                </div>
                <div v-if="stepScene == 2">
                    Тогда ты точно по адресу! Давай начнём.
                    <button class="btn btn-primary" @click="stepScene++">
                        Давай
                    </button>
                </div>
                <div v-if="stepScene == 3">
                    Перед тобой простая программа на языке программирования Python, которая выводит переменную x. Как видишь, мы назначили переменной x значение 3. Переменные используются постоянно, и могут содержать не только числа. 
                    <button class="btn btn-primary" @click="stepScene++">
                        Класс
                    </button>
                </div>
                <div v-if="stepScene == 4">
                    Давай попробуем запустить её и проверить как она работает. Попробуй нажать на кнопку "Проверить"
                    <button class="btn btn-primary" @click="stepScene++">
                        Ок
                    </button>
                </div>
                <div v-if="stepScene == 6">
                    Как видишь, программа успешно выполнилась и вывела результат. Его ты сможешь увидеть выше
                    <button class="btn btn-primary" @click="stepScene++">
                        Класс, а что дальше?
                    </button>
                </div>
                <div v-if="stepScene == 7">
                    Выводить переменные на экран это, конечно, круто. Python, как и любой уважающий себя язык программирования умеет делать арифметические операции, давай попробуем следующий пример
                    <button class="btn btn-primary" @click="stepScene++; code=`x = 3\ny = 4\nz = x + y\nprint(z)`">
                        А давай
                    </button>
                </div>
                <div v-if="stepScene == 8">
                    Как видишь, мы добавили ещё две переменных, первая равна 3-ём, вторая - 4-ём, а третьей присвоили сумму первой и второй. Попробуй проверить вывод этой программы. 
                    <button class="btn btn-primary" @click="stepScene++">
                        Окей
                    </button>
                </div>
                <div v-if="stepScene == 10">
                    Мы можем не только прибавлять, но и делать любые математические операции над переменными
                    <button class="btn btn-primary" @click="stepScene++; code=`x = 15\ny = 5\nz = x * \nprint(z)`">
                        А как?
                    </button>
                </div>
                <div v-if="stepScene == 11">
                    Ну вот тот же пример, только с умножением. Проверь, работает ли данный код
                    <button class="btn btn-primary" @click="stepScene++">
                        Сейчас проверю
                    </button>
                </div>
                <div v-if="stepScene == 13">
                    Оп! Кажется кто то забыл написать название переменной, отвечающую за второй множитель. Сможешь поправить?
                    <button class="btn btn-primary" @click="stepScene++; editorLock = false">
                        Конечно
                    </button>
                </div>
                <div v-if="stepScene == 15">
                    Спасибо что помог! Как видишь, наша программа функцианирует и запускается. Получается, "кодить" на Python легче, чем кажется, не правда ли?
                    <button class="btn btn-primary" @click="stepScene++">
                        Правда
                    </button>
                </div>
                <div v-if="stepScene == 16">
                    Сейчас может быть ты задашься вопросом "Что именно можно написать на этом языке"? Ответ - почти всё. Сайты, боты, нейросети, и многое другое - Python умеет. 
                    Всё что было перечисленно выше возможно тебя шокирует, но не бойся, со временем мы со всем этим разберемся на реальных примерах. 
                    А сейчас я предлагаю перейти в панель задач, где ты сможешь разобрать различные задания на реальной практике
                    <nuxt-link to="/tasks">Перейти</nuxt-link>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data : () => ({
        editorLock : true,
        user : null,
        loaded : false,
        code : "x = 5\nprint(x)",
        stepScene : 0,
        output : "",
        error : ""
    }),
    methods: {
        checkTask : function () {
            this.output = ""
            this.error = ""
            console.log(this.stepScene)
            if (this.stepScene == 5 || this.stepScene == 9 || this.stepScene == 12)  { this.stepScene++ } 
            console.log(this.stepScene)
            this.$axios.$post("http://la4z.xyz:8080/api/check_task", {
                code : this.code
            }).then(response => {
                console.log(response)
                if (!response.output) {
                    response.error.forEach((err, index) => {
                        if (err.message.includes("invalid syntax")) {
                            this.error += `Ошибка синтаксиса в ${err.line} линии.\nПовнимательнее!`
                        }
                        console.log(err)
                    })
                } else {
                    if (this.stepScene == 14) {this.stepScene++; this.editorLock = false}
                    this.output = "Результат выполнения : "+response.output
                }
            }).catch(err => {
                console.log(err)
            })
        }
    },
    mounted : function() {
        this.$axios.setToken(localStorage.getItem("token"), 'Bearer')
        this.$axios.$get("http://la4z.xyz:8080/api/user").then((response) => {
            this.user = response.user
            this.loaded = true
            console.log(this.user)
        }).catch((error) => {
            location.href = "/auth/login"
        })
    }
}
</script>

<style>
body{
      margin: 0;
  font-family: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, Noto Sans, sans-serif, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  font-family: var(--bs-font-sans-serif);
  font-size: 1rem;
  font-weight: 400;
  line-height: 1.5;
  color: #212529;
  background-color: #fff;
  -webkit-text-size-adjust: 100%;
  -webkit-tap-highlight-color: rgba(0,0,0,0);
  font-family: "Lato","Helvetica Neue",Helvetica,Arial,sans-serif;
  background-image: url("~assets/img/Back.jpg");
  background-repeat: no-repeat;
  background-size: 100%;
}
.editorlines {
    display: inline-block;
    justify-content: left;
    width: 25px;
    margin-top:2px;
}
.editor textarea{
    background-color: rgb(24, 33, 34);
    color: white;
    outline: none;
    width : 700px;
    height: 400px;
    display: inline-block;
    resize: none;
}
.editorMain {
    display: flex;
    flex-direction: row;
    width: 800px;
}
.linter {
    margin-top : 50px;
    border: mediumblue solid 3px;
    border-radius: 15px;
    padding: 5px;
    width: 700px;
}
</style>