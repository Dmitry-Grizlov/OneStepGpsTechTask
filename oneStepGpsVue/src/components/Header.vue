<template>
    <nav class="navbar navbar-expand-lg bg-light">
        <div class="container-fluid">
            <router-link to="/" class="navbar-brand"><img src="../assets/logo.png" class="img" height="50px"
                    width="50px"></router-link>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown"
                aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNavDropdown">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <router-link :to="{name: 'index'}" class="nav-link">Map</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link to="/objects" class="nav-link">Objects</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link to="/userSettings" class="nav-link">Account settings</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link to="/login" class="nav-link" @click="logOut">Log out</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script>
export default {
    methods: {
        logOut() {
            localStorage.removeItem('key');
            localStorage.removeItem('userId');
        }
    }
}
</script>

<style>
.profile-image {
    border-radius: 50%;
}
</style>