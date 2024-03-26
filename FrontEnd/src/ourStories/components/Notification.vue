<script>

import ArgonAlert from "@/components/ArgonAlert.vue";

export default {
  components: {
    ArgonAlert,
  },
  props: {
    message: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      showNotification: false,
      title: "",
      color: ""
    };
  },
  watch: {
  /*  message(newMessage) {
      // Call a method to determine the notification type based on the message
      // this.setNotificationType(newMessage);
      // Show the notification
      this.title = newMessage; // temp look like this
      this.showNotification = true;
    },*/

    title(newTitle) {
      this.title = newTitle;
      this.showNotification = true;
    },
    color(newColor) {
      this.color = newColor;
      this.showNotification =true;

    }
  },
  methods: {
    /*setNotificationType(message) {
      // Here, you can implement logic to determine the notification type based on the message
      // For example, you can check the API response and set the notification type accordingly
      // For simplicity, let's assume we set the notification type based on a keyword in the message
      /!* if (message.includes("register")) {
         this.notificationType = "success";
         // this.title = "Registration Successful!";
       } else {
         this.notificationType = "error";
         // this.title = "Error!";
       }*!/
    },*/
  },
};
</script>

<template>
  <argon-alert v-if="showNotification" @close="showNotification = false" :color="color">
    <strong>{{ title }}</strong> {{ message }}
  </argon-alert>
</template>

<style scoped>

</style>