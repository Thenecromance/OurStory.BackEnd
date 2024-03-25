<script setup>
import { onBeforeUnmount, onBeforeMount } from "vue";
import { useStore } from "vuex";

import Navbar from "../PageLayout/navbars/Navbar.vue";
import AppFooter from "@/examples/PageLayout/Footer.vue";
import ArgonInput from "@/components/ArgonInput.vue";
// import ArgonCheckbox from "@/components/ArgonCheckbox.vue";
import ArgonButton from "@/components/ArgonButton.vue";
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

import { ref } from "vue";
import axios from "axios";
const user = ref({
    username: '',
    email: '',
    password: '',
})


// Handle form submission
const submitForm = async () => {
    try {
        // console.log(user.value)
        const response = await axios.post('/api/user/register' , user.value)
        console.log(response.data)
    } catch (error) {
        console.error(error)
    }
}


</script>
<template>
    <div class="container top-0 position-sticky z-index-sticky">
        <div class="row">
            <div class="col-12">
                <navbar isBtn="bg-gradient-light" />
            </div>
        </div>
    </div>

    <main class="main-content mt-0">
        <div class="page-header align-items-start min-vh-50 pt-5 pb-11 m-3 border-radius-lg" style="
        background-image: url(&quot;https://raw.githubusercontent.com/creativetimofficial/public-assets/master/argon-dashboard-pro/assets/img/signup-cover.jpg&quot;);
        background-position: top;
      ">
            <span class="mask bg-gradient-dark opacity-6"></span>
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-lg-5 text-center mx-auto">
                        <h1 class="text-white mb-2 mt-5">Welcome!</h1>
                        <p class="text-lead text-white">
                            Use these awesome forms to login or create new account in your
                            project for free.
                        </p>
                    </div>
                </div>
            </div>
        </div>
        <div class="container">
            <div class="row mt-lg-n10 mt-md-n11 mt-n10 justify-content-center">
                <div class="col-xl-4 col-lg-5 col-md-7 mx-auto">
                    <div class="card z-index-0">
                        <div class="card-header text-center pt-4">
                            <h5>注册</h5>
                        </div>
                        <div class="card-body">

                            <form role="form" @submit.prevent="submitForm">
                                <argon-input v-model="user.username" id="username" type="text" placeholder="昵称"
                                    aria-label="Name" />
                                <argon-input v-model="user.email" id="email" type="email" placeholder="邮箱"
                                    aria-label="Email" />
                                <argon-input v-model="user.password" id="password" type="password" placeholder="密码"
                                    aria-label="Password" />
                                <!-- <argon-checkbox checked>
                                    <label class="form-check-label" for="flexCheckDefault">
                                        I agree the
                                        <a href="javascript:;" class="text-dark font-weight-bolder">Terms and Conditions</a></label>
                                </argon-checkbox> -->
                                <div class="text-center">
                                    <argon-button fullWidth color="dark" variant="gradient"
                                        class="my-4 mb-2">注册</argon-button>
                                </div>
                                <p class="text-sm mt-3 mb-0">
                                    已经有账号了？
                                    <a href="/login" class="text-dark font-weight-bolder"> 登陆</a>
                                </p>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
    <app-footer />
</template>
