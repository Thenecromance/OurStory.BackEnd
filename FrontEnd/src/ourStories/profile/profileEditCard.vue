<script setup>
import ArgonInput from "@/components/ArgonInput.vue";
import ArgonButton from "@/components/ArgonButton.vue";
</script>

<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header pb-0">
                        <div class="d-flex align-items-center">
                            <p class="mb-0">Edit Profile</p>
                            <argon-button color="success" size="sm" class="ms-auto">Settings</argon-button>
                        </div>
                    </div>

                    <div class="card-body">
                        <p class="text-uppercase text-sm">User Information</p>
                        <div class="row">
                            <div class="col-md-6">
                                <label for="example-text-input" class="form-control-label">Username</label>
                                <argon-input type="text" value="lucky.jesse" />
                            </div>
                            <div class="col-md-6">
                                <label for="example-text-input" class="form-control-label">Email address</label>
                                <argon-input type="email" value="jesse@example.com" />
                            </div>
                            <div class="col-md-6">
                                <label for="example-text-input" class="form-control-label">First name</label>
                                <input class="form-control" type="text" value="Jesse" />
                            </div>
                            <div class="col-md-6">
                                <label for="example-text-input" class="form-control-label">Last name</label>
                                <argon-input type="text" value="Lucky" />
                            </div>
                        </div>
                        <hr class="horizontal dark" />
                        <p class="text-uppercase text-sm">Contact Information</p>
                        <div class="row">
                            <div class="col-md-12">
                                <label for="example-text-input" class="form-control-label">Address</label>
                                <argon-input type="text" value="Bld Mihail Kogalniceanu, nr. 8 Bl 1, Sc 1, Ap 09" />
                            </div>
                            <div class="col-md-4">
                                <label for="example-text-input" class="form-control-label">City</label>
                                <argon-input type="text" value="New York" />
                            </div>
                            <div class="col-md-4">
                                <label for="example-text-input" class="form-control-label">Country</label>
                                <argon-input type="text" value="United States" />
                            </div>
                            <div class="col-md-4">
                                <label for="example-text-input" class="form-control-label">Postal code</label>
                                <argon-input type="text" value="437300" />
                            </div>
                        </div>
                        <hr class="horizontal dark" />
                        <p class="text-uppercase text-sm">About me</p>
                        <div class="row">
                            <div class="col-md-12">
                                <label for="example-text-input" class="form-control-label">About me</label>
                                <argon-input type="text"
                                    value="A beautiful Dashboard for Bootstrap 5. It is Free and Open Source." />
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>

</template>