<template>
    <div>
      <!-- Register button -->
      <button @click="registerUser">Register</button>
      <!-- Notification component -->
      <notification :message="notificationMessage" />
    </div>
  </template>
  
  <script>
  import Notification from "./Notification.vue"; // Import the notification component
  
  export default {
    components: {
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