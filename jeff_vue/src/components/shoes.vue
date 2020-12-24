<template>
    <div class="goods-box">
        <a :href="data.Logo_url"><img :src="data.Logo_url" @click="boxClick"></a>
        <p><router-link :to="`/shop/info/${data.Spu_id}`" >{{data.Title}}</router-link></p>
        <p>价格：¥{{(data.Price)/100}}</p>
        <p>{{data.Sold_num}}人已付款</p>
    </div>
</template>

<script>
    export default {
        name: "GoodsBox",
        props: ["data"],
        methods: {
            boxClick() {
                this.$store.state.info = this.data
            }
        },
    }
</script>

<style scoped>
    .goods-box {
        width: 260px;
        height: 350px;
        border-radius: 5px;
        margin: 20px;
        float: left;
        /*background-color: black;*/
        text-align: center;
        vertical-align: bottom
    }

    img {
        width: 260px;
        text-decoration: none;
    }

    a {
        text-decoration: none;
    }

    a:hover {
        color: blueviolet;
    }

    p {
        font-size: 15px;
        font-weight: 100;
        line-height: 23px;
        text-transform: uppercase;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
    }

</style>