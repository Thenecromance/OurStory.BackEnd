<script setup>
import MiniStatisticsCard from "@/ourStories/Cards/MiniStatisticsCard.vue";

import { ref, onBeforeMount } from 'vue';

let cardData = ref(null);


/*
{
  "title": "Card Title",
  "value": "Card Value",
  "description": "Card Description",
  "icon": {
    "component": "ni ni-money-coins",
    "background": "bg-gradient-primary",
    "shape": "rounded-circle"
  }
}
*/
import axios from 'axios';

onBeforeMount(
    async () => {
        try {
            const response = await axios.get('/agronDash/topCard');
            cardData.value = response.data;
           /*  console.info(response.data) */
        } catch (error) {
            // console.error(error);
            cardData.value = []; // default just need to be empty 
        }
    }
);


// onBeforeMount(() => {
//     // simulate data
//     cardData.value = [
//         {
//             title: "Card 1",
//             value: "Value 1",
//             description: "Description 1",
//             icon: {
//                 component: "ni ni-money-coins",
//                 background: "bg-gradient-primary",
//                 shape: "rounded-circle"
//             }
//         },
//         {
//             title: "Card 2",
//             value: "Value 2",
//             description: "Description 2",
//             icon: {
//                 component: "ni ni-money-coins",
//                 background: "bg-gradient-danger",
//                 shape: "rounded-circle"
//             }
//         },
//     ];
// });
</script>

<template>

    <div class="py-4 container-fluid">

        <div class="row">
            <div class="col-lg-3 col-md-6 col-12" v-for="(card, index) in cardData" :key="index">
                <MiniStatisticsCard :title="card.title" :value="card.value" :description="card.description"
                    :icon="card.icon" />
            </div>
        </div>
    </div>
</template>
