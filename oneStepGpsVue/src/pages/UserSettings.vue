<template>
    <div>
        <Header></Header>
        <div class="container">
            <div class="mt-4 p-4 row justify-content-center">
                <input type="hidden" v-model="model.Id">
                <div class="mx-4 alert alert-primary" role="alert" v-if="!hidden">
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
                        <input type="text" id="apiKey" class="form-control" v-model="model.ApiKey" required />
                    </div>
                    <div class="btn btn-primary" @click="saveSettings">Save</div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import Header from '@/components/Header';
import config from '../../config/dev.env'
import axios from 'axios'
import common from '@/common'

export default {
    components: {
        Header
    },

    data() {
        return {
            model: '',
            hidden: true,
            message: '',
        }
    },

    async created() {
        if (!common.isAuthorised()) {
            this.$router.push('/login');
            return;
        }

        await this.getSettings();
    },

    methods: {
        async getSettings() {
            let data = { id: localStorage.getItem('userId') };
            await axios.post(config.API_HOST + "user", data)
                .then((response) => {
                    this.model = response.data.Data;
                })
                .catch((error) => {
                    this.showMessage(resp.data.Message);
                })
        },

        async saveSettings() {
            await axios.post(config.API_HOST + "updateUser", this.model)
                .then((response) => {
                    this.showMessage("Information successfully updated");
                })
                .catch((error) => {
                    this.showMessage(error.response.data.Message)
                });
        },

        showMessage(data) {
            this.message = data;
            this.hidden = false;
        }
    }
}
</script>