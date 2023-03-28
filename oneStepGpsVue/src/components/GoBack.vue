<template>
    <span class="go-back">
        <button @click="goBack" class="btn btn-primary">Back</button>
    </span>
</template>

<script>
export default {
    methods: {
        goBack() {
            return this.$router.go(-1);
        }
    }
};
</script>