<script>

// import notification_card from "@/ourStories/utils/Test.vue";
import Notification from "@/ourStories/utils/Notification.vue";
import TravelLog from "@/ourStories/Travel/TraveTimeLine.vue";

export default {
  components: {
    TravelLog,
    Notification,
  },
  data() {
    return {
      notificationMessage: "",
    };
  },
  methods: {
    registerUser() {
      // Simulate API response
      const apiResponse = {
        result: {
          code: "ok",
          method: "register",
        },
        data: {
          username: "test",
          email: "123@qq.com",
        },
      };

      // Set notification message based on API response
      if (apiResponse.result.code === "ok" && apiResponse.result.method === "register") {
        this.notificationMessage = `Register Success!`;
      } else {
        this.notificationMessage = "Registration failed!";
      }
    },
  },
};
</script>

<template>
  <div>
    <!-- Register button -->
    <button @click="registerUser">Register</button>
    <!-- Notification component -->
    <notification :message="notificationMessage"/>

    <TravelLog/>
  </div>
</template>