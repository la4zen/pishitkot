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
            <button @click="checkTask()"> Проверить </button> 
            <div class="output">
                {{ output }}
            </div>
            <div class="error" v-if="error"> 
                В коде ошибка!
                <pre>{{ error }}</pre>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data : () => ({
        editorLock : false,
        user : null,
        loaded : false,
        code : "z = [1,2,3,4,5]\nx = 2\nfor i in z:\n   print(i*x)",
        output : "",
        error : ""
    }),
    methods: {
        checkTask : function () {
            this.editorLock = true
            this.output = ""
            this.error = ""
            this.$axios.$post("http://la4z.xyz:8080/api/check_task", {
                code : this.code
            }).then(response => {
                console.log(response)
                this.editorLock = false
                if (!response.output) {
                    response.error.forEach((err, index) => {
                        if (err.message.includes("invalid syntax")) {
                            this.error += `Ошибка синтаксиса в ${err.line} линии.\nПовнимательнее!`
                        }
                        console.log(err)
                    })
                } else {
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