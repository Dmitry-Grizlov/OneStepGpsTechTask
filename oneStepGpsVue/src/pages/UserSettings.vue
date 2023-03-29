<template>
    <div class="container">
        <div class="mt-4 p-4 row justify-content-center">
            <input type="hidden" v-model="model.Id">
            <div class="mx-4 alert alert-primary " role="alert" v-if="!hidden">
                {{ message }}
            </div>
            <form class="col-md-8 p-3 bg-white rounded">
                <div class="form-group mb-4">
                    <label for="userName" class="mb-2">Name</label>
                    <input type="text" id="userName" class="form-control" v-model="model.Name" />
                </div>

                <div class="form-group mb-4">
                    <label for="email" class="mb-2">Email</label>
                    <input type="email" id="email" class="form-control" v-model="model.Email" />
                </div>

                <div class="form-group mb-4">
                    <label for="apiKey" class="mb-2">Api key</label>
                    <input type="text" id="apiKey" class="form-control" v-model="model.ApiKey" />
                </div>
                <div class="btn btn-primary" @click="saveSettings">Save</div>
            </form>
        </div>
    </div>
</template>

<script>
import config from '../../config/dev.env'
import axios from 'axios'

export default {
    data() {
        return {
            model: '',
            hidden: true,
            message: '',
        }
    },

    async created() {
        await this.getSettings();
    },

    methods: {
        async getSettings() {
            let data = { id: localStorage.getItem('userId') };
            let resp = await axios.post(config.API_HOST + "user", data);
            if (resp.data == null) {
                this.showMessage("Could not get data");
                return;
            }
            this.model = resp.data;
        },

        async saveSettings() {
            let model = (this.model);
            let resp = await axios.post(config.API_HOST + "updateUser", model);
            console.log(resp);
            this.showMessage(resp.data.message);
        },

        showMessage(data) {
            this.message = data;
            this.hidden = false;
        }
    }
}
</script>