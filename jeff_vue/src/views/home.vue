<template>
    <div class="home">
        <banner/>
    </div>
</template>
<script>
    import banner from "../components/banner";

    export default {
        components: {
            banner,
        },
    }
</script>

<style scoped>
    .home {
        text-align: center;
        background: url("../assets/img/home.jpg");
        height: 800px;
    }
    img{
        height: 800px;
    }
</style>