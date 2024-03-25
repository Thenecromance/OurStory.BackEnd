<template>

  <!-- <div class="row justify-content-center">
      <div class="col-lg-5 col-md-7 text-center">
          <h1 class="text-white">@ViewBag.StatusCode - @ViewBag.StatusCodeDescription</h1>
          <a class="text-light" href="#" title="Go Home">Go Home</a>
      </div>
  </div> -->
  <div class="container top-0 position-sticky z-index-sticky  ">
    <div class="row">
      <div class="col-12">
        <navbar isBtn="bg-gradient-light" isBlur="blur  border-radius-lg start-0 end-0 mx-4 shadow"/>
      </div>
    </div>
  </div>
  <!--  <navbar class="fixed-top" isBlur="blur  border-radius-lg my-3 py-2 start-0 end-0 mx-4 shadow"
          v-bind:darkMode="true" isBtn="bg-gradient-success" /> -->
  <main class="main-content ">
    <section class="error-section ">
      <div class="error-wrapper">
        <div class="title">404</div>
        <div class="subtitle">Oops! Page not found.</div>
        <div class="description">The page you are looking for might have been removed, had its name changed,
          or is temporarily unavailable.
        </div>
        <router-link to="/" class="btn btn-primary">Go back to Home</router-link>
      </div>
    </section>
  </main>

</template>

<script setup>
import {onBeforeUnmount, onBeforeMount} from "vue";
import {useStore} from "vuex";
import Navbar from "../PageLayout/navbars/Navbar.vue";

const body = document.getElementsByTagName("body")[0];
const store = useStore();
onBeforeMount(() => {
  store.state.hideConfigButton = true;
  store.state.showNavbar = false;
  store.state.showSidenav = false;
  store.state.showFooter = false;
  body.classList.remove("bg-gray-100");
});
onBeforeUnmount(() => {
  store.state.hideConfigButton = false;
  store.state.showNavbar = true;
  store.state.showSidenav = true;
  store.state.showFooter = true;
  body.classList.add("bg-gray-100");
});
</script>

<script>
export default {
  name: 'NotFound'
}
</script>

<style scoped>
.container {
  position: relative;
}

.fixed-top {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.main-content {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-image: url('https://th.bing.com/th?id=OHR.WaikatoWater_ZH-CN0417438809_1920x1080.webp&qlt=50');
  background-size: cover;
  background-position: top;
}

.error-section {
  text-align: center;
}

.error-wrapper {
  max-width: 400px;
}

.title {
  font-size: 100px;
  font-weight: bold;
  color: #5e72e4;
}

.subtitle {
  font-size: 24px;
  margin-top: 20px;
}

.description {
  font-size: 16px;
  margin-top: 20px;
}

.btn {
  display: inline-block;
  margin-top: 20px;
  padding: 10px 20px;
  border-radius: 4px;
  background-color: #5e72e4;
  color: #fff;
  text-decoration: none;
  transition: transform 0.3s ease;
}

.btn:hover {
  transform: translateY(-3px);
}
</style>