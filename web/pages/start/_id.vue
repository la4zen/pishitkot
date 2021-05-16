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
                    <textarea autocomplete="off" v-model="code" disabled/>
                </div>
            </div>
            <div class="linter">
                <div v-if="stepScene == 0">
                    Привет, {{ user.first_name }}!
                    <button @click="stepScene++">
                        Привет
                    </button>
                </div>
                <div v-if="stepScene == 1">
                    
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data : () => ({
        user : null,
        loaded : false,
        code : "x = 5\ny = 3\nz = x + y\nprint(z)",
        stepScene : 0,
    }),
    methods: {
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

}
</style>