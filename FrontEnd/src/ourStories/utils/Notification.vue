<template>
    <div>
        <argon-alert v-if="showNotification" @close="showNotification = false" :type="notificationType">
            <strong>{{ title }}</strong> {{ message }}
        </argon-alert>

    </div>
</template>

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
            notificationType: "", // You can define different types of notifications here (e.g., success, error, info, etc.)
            title: "",
        };
    },
    watch: {
        message(newMessage) {
            // Call a method to determine the notification type based on the message
            this.setNotificationType(newMessage);
            // Show the notification
            this.showNotification = true;
        },
    },
    methods: {
        setNotificationType(message) {
            // Here, you can implement logic to determine the notification type based on the message
            // For example, you can check the API response and set the notification type accordingly
            // For simplicity, let's assume we set the notification type based on a keyword in the message
            if (message.includes("register")) {
                this.notificationType = "success";
                // this.title = "Registration Successful!";
            } else {
                this.notificationType = "error";
                // this.title = "Error!";
            }
        },
    },
};
</script>