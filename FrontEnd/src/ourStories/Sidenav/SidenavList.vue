<script setup>
import {useRoute} from "vue-router";
import SidenavItem from "./SidenavItem.vue";
// import SidenavCard from "./SidenavCard.vue";


const getRoute = () => {
  const route = useRoute();
  const routeArr = route.path.split("/");
  return routeArr[1];
};


import {ref, onMounted} from "vue";

let navList = ref(null);

/* import axios from "axios";
onMounted(
  async () => {
    try {

      const response = await axios.get('/agronDash/sidenav');
      navList.value = response.data;
    } catch (error) {
      // console.error(error);
      navList.value = []; // 或者你的默认值
    }
  }
); */


function NavSpliter(msg) {
  return {
    to: "",
    navText: "",
    icon: {
      component: "",
    },
    Msg: msg,
    Split: true,
  }

}

function NavItem(to, NavText, icon) {
  return {
    to: to,
    navText: NavText,
    icon: {
      component: icon,
    },
    Msg: "",
    Split: false,
  }

}

onMounted(
    () => {
      navList.value = [
        NavSpliter("Our stories Start"),
        NavItem("/dashboard-default", "dashboard-default", "ni ni-tv-2 text-primary text-sm opacity-10"),
        NavItem("/", "myDash", "ni ni-tv-2 text-primary text-sm opacity-10"),
        NavItem("/login", "login", "ni ni-tv-2 text-primary text-sm opacity-10"),
        NavItem("/register", "register", "ni ni-tv-2 text-primary text-sm opacity-10"),
        NavItem("/404", "404", ""),
        NavItem("/test", "test", ""),
        NavSpliter("---------------------"),
      ]
    }
)


</script>
<template>
  <div class="collapse navbar-collapse w-auto h-auto h-100" id="sidenav-collapse-main">

    <!--Load Navigation Items from remote Server dynamically start-->
    <ul class="navbar-nav" v-for="(navList, index) in  navList" :key="index">
      <li v-if="navList.Split === false" class="nav-item">
        <sidenav-item :to="navList.to" :class="getRoute() === navList.to ? 'active' : ''" :navText="navList.navText">
          <template v-slot:icon>
            <i :class="navList.icon.component"></i>
          </template>
        </sidenav-item>
      </li>
      <li v-else class="mt-3 nav-item">
        <h6 class="text-xs ps-4 text-uppercase font-weight-bolder opacity-6" :class="'ms-2'">
          {{ navList.Msg }}
        </h6>
      </li>
    </ul>
    <!--Load Navigation Items from remote Server dynamically finished-->

    <ul class="navbar-nav">

      <li class="nav-item">
        <sidenav-item to="/tables" :class="getRoute() === 'tables' ? 'active' : ''" :navText="'Tables'">
          <template v-slot:icon>
            <i class="ni ni-calendar-grid-58 text-warning text-sm opacity-10"></i>
          </template>
        </sidenav-item>
      </li>

      <li class="nav-item">
        <sidenav-item to="/billing" :class="getRoute() === 'billing' ? 'active' : ''" :navText="'Billing'">
          <template v-slot:icon>
            <i class="ni ni-credit-card text-success text-sm opacity-10"></i>
          </template>
        </sidenav-item>
      </li>

      <!--  <li class="nav-item">
         <sidenav-item to="/virtual-reality" :class="getRoute() === 'virtual-reality' ? 'active' : ''"
           :navText="'Virtual Reality'">
           <template v-slot:icon>
             <i class="ni ni-app text-info text-sm opacity-10"></i>
           </template>
         </sidenav-item>
       </li> -->

      <!--  <li class="nav-item">
         <sidenav-item to="/rtl-page" :class="getRoute() === 'rtl-page' ? 'active' : ''" navText="RTL">
           <template v-slot:icon>
             <i class="ni ni-world-2 text-danger text-sm opacity-10"></i>
           </template>
         </sidenav-item>
       </li> -->

      <li class="mt-3 nav-item">
        <h6 class="text-xs ps-4 text-uppercase font-weight-bolder opacity-6" :class="'ms-2'">
          ACCOUNT PAGES
        </h6>
      </li>

      <li class="nav-item">
        <sidenav-item to="/profile" :class="getRoute() === 'profile' ? 'active' : ''" :navText="'Profile'">
          <template v-slot:icon>
            <i class="ni ni-single-02 text-dark text-sm opacity-10"></i>
          </template>
        </sidenav-item>
      </li>

      <!-- <li class="nav-item">
        <sidenav-item to="/signin" :class="getRoute() === 'signin' ? 'active' : ''" :navText="'Sign In'">
          <template v-slot:icon>
            <i class="ni ni-single-copy-04 text-danger text-sm opacity-10"></i>
          </template>
        </sidenav-item>
      </li>

      <li class="nav-item">
        <sidenav-item to="/signup" :class="getRoute() === 'signup' ? 'active' : ''" :navText="'Sign Up'">
          <template v-slot:icon>
            <i class="ni ni-collection text-info text-sm opacity-10"></i>
          </template>
        </sidenav-item>
      </li> -->
    </ul>
  </div>

  <!--  <div class="pt-3 mx-3 mt-3 sidenav-footer">-->
  <!--    <sidenav-card :card="{-->
  <!--      title: 'Need Help?',-->
  <!--      description: 'Please check our docs',-->
  <!--      links: [-->
  <!--        {-->
  <!--          label: 'Documentation',-->
  <!--          route:-->
  <!--            'https://www.creative-tim.com/learning-lab/vue/overview/argon-dashboard/',-->
  <!--          color: 'dark',-->
  <!--        },-->
  <!--        {-->
  <!--          label: 'Buy now',-->
  <!--          route:-->
  <!--            'https://www.creative-tim.com/product/vue-argon-dashboard-pro?ref=vadp',-->
  <!--          color: 'success',-->
  <!--        },-->
  <!--      ],-->
  <!--    }" />-->
  <!--  </div>-->
</template>
