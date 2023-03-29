<template>
    <section class="text-center">
        <div class="p-5 bg-image bg-img"></div>
        <div class="card mx-4 mx-md-5 shadow-5-strong login-card">
            <div class="card-body py-5 px-md-5">
                <div class="row d-flex justify-content-center">
                    <div class="col-lg-8">
                        <h2 class="fw-bold mb-5">Log in to OneStepGPS app</h2>
                        <form>
                            <div class="form-outline mb-4">
                                <input type="text" class="form-control" id="form3Example3" v-model="userInput">
                                <label class="form-label" for="form3Example3">API key</label>
                            </div>
                            <div class="btn btn-primary btn-block mb-4" @click="login">
                                Log in
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import config from "../../config/dev.env"
import axios from "axios"

export default {
    data() {
        return {
            error: "",
            userInput: ""
        };
    },

    mounted() {
        localStorage.removeItem('key');
        localStorage.removeItem('id');
    },

    methods: {
        async login() {
            let key = this.userInput;
            let data = { "apiKey": key };
            let response = await axios.post(config.API_HOST + "login", data);
            console.log(response);
            if (response.status == 200 && response.data.id > 0) {
                localStorage.setItem("key", response.data.apiKey);
                localStorage.setItem("userId", response.data.id);
                this.$router.push('/')
            }
        }
    }
}
</script>

<style>
.bg-img {
    background-image: url('../assets/loginBg.jpg');
    height: 300px;
}

.login-card {
    margin-top: -100px;
    background: hsla(0, 0%, 100%, 0.8);
    backdrop-filter: blur(30px);
}
</style>