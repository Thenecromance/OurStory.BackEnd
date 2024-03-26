<script setup>


import TimelineList from "@/ourStories/Cards/TimelineList.vue";
import TimelineItem from "@/ourStories/Cards/TimelineItem.vue";

import axios from "axios";

import {onMounted, ref} from "vue";

let travelList = ref(null);

onMounted(
    async () => {
      const response = await axios.get("/api/travel?user=1");
      console.log(response.data);
      travelList.value = response.data.data;
    }
)

// eslint-disable-next-line no-unused-vars
function unixToTime(time_stamp) {
  const date = new Date(time_stamp);
  let Y = date.getFullYear() + '-';
  let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
  let D = date.getDate();
  return Y + M + D;
}

</script>

<template>
  <timeline-list title="旅行日记">

    <timeline-item v-for="(item,i) in travelList" :key="i"
                   :icon="{ component: 'ni ni-bell-55', color: 'success' }"
                   color="primary"
                   :title="item.location"
                   :date-time="unixToTime(item.stamp)"
                   :description="item.logs"
                   :badges="['准备中']"
    />

  </timeline-list>
</template>

<style scoped>

</style>