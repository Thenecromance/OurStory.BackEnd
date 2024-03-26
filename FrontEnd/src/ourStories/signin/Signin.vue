<script setup>
import {onBeforeUnmount, onBeforeMount} from "vue";
import {useStore} from "vuex";
import Navbar from "../PageLayout/navbars/Navbar.vue";
import ArgonInput from "@/components/ArgonInput.vue";
// import ArgonSwitch from "@/components/ArgonSwitch.vue";
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


import {ref} from 'vue'
import axios from 'axios'
import Notification from "@/ourStories/components/Notification.vue";

// Create reactive references
const form = ref({
  username: '',
  password: ''
})

let notificationMessage = ref(null)
let color = ref(null)


// Handle form submission
const submitForm = async () => {
  try {
    await axios.post('/api/user/login', form.value).then(
        response => {
          console.log(response)
          let result
          result = response.data
          if (result.code === 0) {
            console.log("Login Complete!")
            //pop up normal alert
            notificationMessage.value = "登陆成功！" + result.result.username;
            color.value = 'success'


            //then redirect to home page

          } else if (result.code === 1) {
            notificationMessage.value = result.result;
            color.value = 'danger'
          } else {
            notificationMessage.value = "unknown response";
          }
        }
    )


  } catch (error) {
    console.error(error)
  }
}

</script>
<template>
  <div class="container top-0 position-sticky z-index-sticky">
    <div class="row">
      <div class="col-12">
        <navbar isBlur="blur border-radius-lg my-3 py-2 start-0 end-0 mx-4 shadow" v-bind:darkMode="true"
                isBtn="bg-gradient-success"/>
      </div>
    </div>
  </div>


  <main class="mt-0 main-content">

    <section>
      <notification :message="notificationMessage" :color="color"/>
      <div class="page-header min-vh-100">
        <div class="container">
          <div class="row">
            <div class="mx-auto col-xl-4 col-lg-5 col-md-7 d-flex flex-column mx-lg-0">
              <div class="card card-plain">
                <div class="pb-0 card-header text-start">
                  <h4 class="font-weight-bolder">登陆</h4>
                  <p class="mb-0">输入你的账号和密码来登陆</p>
                </div>
                <div class="card-body">
                  <form role="form" @submit.prevent="submitForm">
                    <div class="mb-3">
                      <!-- <argon-input id="email" type="email" placeholder="Email" name="email"
                          size="lg" v-model="form.username"/> -->
                      <argon-input id="username" type="username" placeholder="Username"
                                   name="username" size="lg" v-model="form.username"/>
                    </div>
                    <div class="mb-3">
                      <argon-input id="password" type="password" placeholder="Password"
                                   name="password" size="lg" v-model="form.password"/>
                    </div>
                    <!-- <argon-switch id="rememberMe" name="remember-me">记住我</argon-switch> -->

                    <div class="text-center">
                      <argon-button class="mt-4" variant="gradient" color="light" fullWidth
                                    size="lg">登陆
                      </argon-button>
                    </div>
                  </form>
                </div>
                <div class="px-1 pt-0 text-center card-footer px-lg-2">
                  <p class="mx-auto mb-4 text-sm">
                    没有账号？
                    <a href="/signup" class="text-success text-gradient font-weight-bold">注册啊</a>
                  </p>
                </div>
              </div>
            </div>
            <div
                class="top-0 my-auto text-center col-6 d-lg-flex d-none h-100 pe-0 position-absolute end-0 justify-content-center flex-column">
              <div
                  class="position-relative bg-gradient-primary h-100 m-3 px-7 border-radius-lg d-flex flex-column justify-content-center overflow-hidden"
                  style="
                  background-image: url(&quot;https://img2.imgtp.com/2024/03/21/xhKFnJV3.jpg&quot;);
                  background-size: cover;
                ">
                <span class="mask bg-gradient-success opacity-6"></span>
                <h4 class="mt-5 text-white font-weight-bolder position-relative">
                  "心灵毒鸡汤 标题"
                </h4>
                <p class="text-white position-relative">
                  心灵毒鸡汤内容
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>
