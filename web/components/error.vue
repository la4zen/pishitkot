<template>
    <transition name="slide-fade">
        <p v-if="show">{{ textError}}</p>
    </transition>
</template>
<script>
export default {
    data : () => ({
        textError : null,
        show : false
    }),
    methods : {
        sendError : function (text) {
            console.log(`Sending error with text : ${text}`)
            this.textError = text
            this.show = true
            setTimeout(() => {
                this.show = false
            }, 15000)
        }
    }
}
</script>
<style>
.slide-fade-enter-active {
  transition: all .3s ease;
}
.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active до версии 2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
</style>