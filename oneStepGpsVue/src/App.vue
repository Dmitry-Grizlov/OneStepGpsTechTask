<template>
  <div id="app">
    <div v-if="showNavigation">
      <Header></Header>
    </div>
    <router-view />
  </div>
</template>

<script>
import Header from "./components/Header.vue";

export default {
  name: 'App',
  components: {
    Header
  },
  data() {
    return {
      key: this.getKey()
    }
  },

  methods: {
    showNavigation() {
      return this.$route.path.includes('/login');
    },

    getKey() {
      let value = localStorage.getItem("key");
      if (value == null) {
        this.$router.push(this.$route.query.redirect || '/login');
        return;
      }
      this.key = value;
      return this.key;
    }
  }
}
</script>

<style></style>
