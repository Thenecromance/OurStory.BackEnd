<script setup>


import TimelineList from "@/ourStories/Cards/TimelineList.vue";
import TimelineItem from "@/ourStories/Cards/TimelineItem.vue";

import axios from "axios";
import {onMounted} from "vue";

let travelList = null;

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
onMounted(
    async () => {
      const response = await axios.get("/api/travel?user=1");
      console.log(response.data);
      travelList = response.data.data;
      /* await axios.get("/api/travel?user=1").then(
           response=>{
             if (response.code != 0) {
               travelList = []
               return
             }
             travelList.value = response.data
           }
       )*/

    }
)

// eslint-disable-next-line no-unused-vars
function unixToTime(time_stamp) {
  const unixTimestamp = new Date(time_stamp * 1000);
  const commonTime = unixTimestamp.toLocaleString();
  return commonTime
}

</script>

<template>
  <timeline-list v-for="(item,i) in travelList" :key="i" title="Timeline with dotted line">

    <timeline-item
        :icon="{ component: 'ni ni-bell-55', color: 'success' }"
        title="{{item.location}}"
        date-time="{{unixToTime(item.stamp)}}"
        description="{{item.log}}"
    />
    <!--    <timeline-item :icon="{ component: 'ni ni-bell-55', color: 'success' }" title="$2400 Design changes"
                       date-time="22 DEC 7:20 PM"
                       description="People care about how you see the world, how you think, what motivates you, what you’re struggling with or afraid of."
                       :badges="['design']"/>
        <TimelineItem :icon="{ component: 'ni ni-html5', color: 'danger' }" title="New order #1832412"
                      date-time="21 DEC 11 PM"
                      description="People care about how you see the world, how you think, what motivates you, what you’re struggling with or afraid of."
                      :badges="['order', '#1832412']"/>
        <TimelineItem :icon="{ component: 'ni ni-cart', color: 'info' }" title="Server payments for April"
                      date-time="21 DEC 9:34 PM"
                      description="People care about how you see the world, how you think, what motivates you, what you’re struggling with or afraid of."
                      :badges="['server', 'payments']"/>-->
  </timeline-list>
</template>

<style scoped>

</style>