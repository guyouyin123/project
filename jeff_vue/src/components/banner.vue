<!--
轮播图组件
-->
<template>
    <div class="block">
        <span class="demonstration">精品推荐</span>
        <el-carousel style="height: 250px">
            <el-carousel-item v-for="img in img_list" :key="img">
                <img :src="img" alt="">
            </el-carousel-item>
        </el-carousel>
    </div>
</template>

<script>
    export default {
        name: "pager",
        data() {
            return {
                img_list:[
                    'https://cdn.poizon.com/source-img/origin-img/20201205/22184f4520af401e991e68de68f3129d.jpg',
                    'https://cdn.poizon.com/source-img/origin-img/20201206/674ab33271e642079cd46cdfc3d9a9cd.jpg',
                    'https://cdn.poizon.com/source-img/origin-img/20201205/bdddaa2d8bb34340897b1b21b503775d.jpg',
                    'https://cdn.poizon.com/source-img/origin-img/20201205/210b39261172473883a43128c1e45627.jpg'
                ]
            }
        },
    }
</script>

<style>
    .el-carousel__item h3 {
        color: #475669;
        font-size: 14px;
        opacity: 0.75;
        line-height: 150px;
        margin: 0;
    }

    .el-carousel__container {
        position: relative;
        height: 600px;
    }
    img{
        height: 250px;
    }
    .block{
        position: relative;
        top: 100px;
        background-color: rgba(0,0,0,0.5);
    }
    .demonstration{
        height: 60px;
        font-size: 30px;
        color: aliceblue;
        font-weight: 100;
        /*line-height: 23px;*/
        text-transform: uppercase;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
    }
</style>