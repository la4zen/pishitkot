<template>
    <div class="container">
        <div class="loading" v-if="!loaded">
            
        </div>
        <div v-else class="editorMain">
            <div class="editorlines">
                <div v-for="(_, index) in code.split('\n')">
                    {{ ++index }}
                </div>
            </div>
            <div class="editor">
                <textarea autocomplete="off" v-model="code"/>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data : () => ({
        user : null,
        loaded : true,
        code : "x = 5\ny = 3\nprint(x+y)",
    }),
    methods: {

    },
    mounted : () => ({

    })
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
</style>