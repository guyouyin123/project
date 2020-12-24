<template>
    <div class="info">
        <p>1231312312</p>
        <p><img :src="data.Logo_url" alt=""></p>
        <p><a :href="data.Logo_url">{{data.Title}}</a></p>
        <p>标题:{{data.Sub_title}}</p>
        <p>价格：¥{{(data.Price)/100}}</p>
        <p>{{data.Sold_num}}人已付款</p>
    </div>

</template>

<script>
    export default {
        name: "shopinfo",
        data() {
            return {
                data: "",
                brand: "",
            }
        },
        created() {
            this.data = this.$store.state.info
            console.log(this.data)
        },
    }
</script>

<style scoped>
    body{
        background-color: aliceblue;
    }
</style>