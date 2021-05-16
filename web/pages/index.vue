<template>
  <div style="margin:0;padding:0">
    <nav class="navbar navbar-light navbar-expand sticky-top bg-light navigation-clean">
		<div class="container">
			<div><a class="navbar-brand" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-gray-dark);">Пиши</a><a class="navbar-brand border rounded-0 border-light shadow-none" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-blue);">Кот</a><a class="navbar-brand" style="margin: 0;margin-right: 0px;width: 100px;height: 35px;padding: 0px;padding-top: 0px;margin-top: 0;margin-bottom: 0;color: var(--bs-gray-dark);">.</a></div>
			<button data-bs-toggle="collapse" class="navbar-toggler" data-bs-target="#navcol-1"></button>
			<div class="collapse navbar-collapse" id="navcol-1"></div>
			<button class="btn btn-primary" type="button" style="margin-right: 10px;">Начать</button>
		</div>
	</nav>
	<header class="text-center text-white masthead Head-img">
		<div class="overlay"></div>
		<div class="container">
			<div class="row">
				<div class="col-xl-9 mx-auto position-relative">
					<h1 class="text-white mb-5" style="color: var(--bs-white);text-shadow: 0px 5px 2px var(--bs-gray-dark);filter: brightness(100%) hue-rotate(0deg);border-color: var(--bs-light);">Создавать программы<br>&nbsp;не так уж и сложно!</h1></div>
			</div>
		</div>
	</header>
	<section class="showcase">
		<div class="container-fluid p-0">
			<div class="row g-0">
				<div class="col-lg-6 text-white order-lg-2 showcase-img SC-img1"><span></span></div>
				<div class="col-lg-6 my-auto order-lg-1 showcase-text">
					<h2>Fully Responsive Design</h2>
					<p class="lead mb-0">When you use a theme created with Bootstrap, you know that the theme will look great on any device, whether it's a phone, tablet, or desktop the page will behave responsively!</p>
				</div>
			</div>
			<div class="row g-0">
				<div class="col-lg-6 text-white showcase-img SC-img2"><span></span></div>
				<div class="col-lg-6 my-auto order-lg-1 showcase-text">
					<h2>Updated For Bootstrap 4</h2>
					<p class="lead mb-0">Newly improved, and full of great utility classes, Bootstrap 4 is leading the way in mobile responsive web development! All of the themes are now using Bootstrap 4!</p>
				</div>
			</div>
			<div class="row g-0">
				<div class="col-lg-6 text-white order-lg-2 showcase-img SC-img3"><span></span></div>
				<div class="col-lg-6 my-auto order-lg-1 showcase-text">
					<h2>Easy to Use &amp;&nbsp;Customize</h2>
					<p class="lead mb-0">Landing Page is just HTML and CSS with a splash of SCSS for users who demand some deeper customization options. Out of the box, just add your content and images, and your new landing page will be ready to go!</p>
				</div>
			</div>
		</div>
	</section>
	<footer class="bg-light footer">
		<div class="container">
			<div class="row">
				<div class="col-lg-6 text-center text-lg-start my-auto h-100">
					<ul class="list-inline mb-2">
						<li class="list-inline-item"><a href="#">About</a></li>
						<li class="list-inline-item"><span>⋅</span></li>
						<li class="list-inline-item"><a href="#">Contact</a></li>
						<li class="list-inline-item"><span>⋅</span></li>
						<li class="list-inline-item"><a href="#">Terms of &nbsp;Use</a></li>
						<li class="list-inline-item"><span>⋅</span></li>
						<li class="list-inline-item"><a href="#">Privacy Policy</a></li>
					</ul>
					<p class="text-muted small mb-4 mb-lg-0">© ПишиКот 2021. All Rights Reserved.</p>
				</div>
				<div class="col-lg-6 text-center text-lg-end my-auto h-100">
					<ul class="list-inline mb-0">
						<li class="list-inline-item"><a href="#"><i class="fa fa-facebook fa-2x fa-fw"></i></a></li>
						<li class="list-inline-item"><a href="#"><i class="fa fa-twitter fa-2x fa-fw"></i></a></li>
						<li class="list-inline-item"><a href="#"><i class="fa fa-instagram fa-2x fa-fw"></i></a></li>
					</ul>
				</div>
			</div>
		</div>
	</footer>
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
}

@media (min-width: 768px) {
  header.masthead {
    padding-top: 12rem;
    padding-bottom: 12rem;
  }
}
header.masthead {
  position: relative;
  background-color: #343a40;
  background-size: cover;
  padding-top: 8rem;
  padding-bottom: 8rem;
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

@media (min-width: 768px) {
  .showcase .showcase-text {
    padding: 7rem;
  }
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

.Head-img{
  background-image: url("~assets/img/Back.png");
  background-size: cover;
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
