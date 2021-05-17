<template>
  <div style="margin:0;padding:0">
      <nav class="navbar navbar-light navbar-expand sticky-top bg-light navigation-clean">
      <div class="container">
        <div><a class="navbar-brand" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-gray-dark);">Пиши</a><a class="navbar-brand border rounded-0 border-light shadow-none" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-blue);">Кот</a><a class="navbar-brand" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-gray-dark);">.</a></div>
        <button data-bs-toggle="collapse" class="navbar-toggler" data-bs-target="#navcol-1"></button>
        <div class="collapse navbar-collapse" id="navcol-1"></div>
        <nuxt-link class="btn btn-primary" type="button" style="margin-right: 10px;" to='/auth/login'>Начать</nuxt-link>
      </div>
    </nav>
    <header class="text-center text-white masthead">
      <div class="overlay"></div>
      <div class="container">
        <div class="row">
          <div class="col-xl-9 mx-auto position-relative">
            <h1 class="text-white mb-5" style="color: var(--bs-white);text-shadow: 0px 5px 2px var(--bs-gray-dark);filter: brightness(100%) hue-rotate(0deg);border-color: var(--bs-light);">Создавать программы<br>&nbsp;не так уж и сложно!</h1></div>
        </div>
      </div>
    </header>
  </div>
</template>

<script>
import Default from '../layouts/default.vue'
export default {
  components: {
    Default
    
  }
}
</script>

<style>
body {
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

@media (min-width: 768px) {
  header.masthead {
    padding-top: 12rem;
  }
}

@media (min-width: 992px) {
  .col-lg-6 {
    flex: 0 0 auto;
    width: 50%;
  }
}

@media (min-width: 992px) {
  .order-lg-1 {
    order: 1 !important;
  }
}

.row {
  --bs-gutter-x: 1.5rem;
  --bs-gutter-y: 0;
  display: flex;
  flex-wrap: wrap;
  margin-top: calc(var(--bs-gutter-y) * -1);
  margin-right: calc(var(--bs-gutter-x) / -2);
  margin-left: calc(var(--bs-gutter-x) / -2);
}


h1, .h1, h2, .h2, h3, .h3, h4, .h4, h5, .h5, h6, .h6 {
  font-family: "Lato","Helvetica Neue",Helvetica,Arial,sans-serif;
  font-weight: 700;
}

.showcase .showcase-img {
  min-height: 30rem;
  background-size: cover;
}

.showcase .showcase-text {
  padding: 3rem;
}


@media (min-width: 992px) {
  .col-lg-4 {
    flex: 0 0 auto;
    width: 33.3333333333%;
  }
}

.features-icons .features-icons-item .features-icons-icon i {
  font-size: 4.5rem;
}

.SC-img1{
  background-image:url("~/assets/img/bg-showcase-1.jpg");
}
.SC-img2{
  background-image:url("~/assets/img/bg-showcase-2.jpg");
}
.SC-img3{
  background-image:url("~/assets/img/bg-showcase-3.jpg");
}

footer.footer {
  padding-top: 4rem;
  padding-bottom: 4rem;
}

.bg-light {
  background-color: #f8f9fa !important;
}

</style>
