<!--
导航条组件
-->

<template>
    <div class="index">
        <div>
            <a
                    style="color: rgb(255, 255, 255); border-bothrefm: 2px solid rgb(255, 255, 255); font-weight: 600;"
                    href="/"> 首页
            </a>
            <a href=""></a>
            <a href="/"> 购买精选</a>
            <a href="/shoes"> 鞋类</a>
            <a href="/cloth"> 潮服</a>
            <a href="/package"> 箱包</a>
            <a href="/watch"> 手表</a>
            <a href="/lipstick"> 口红</a>
            <a href="/"> 关于我们</a>
        </div>
    </div>
</template>

<script>
    export default {
        name: "index"
    }
</script>

<style scoped>
    * {
        box-sizing: border-box;
    }

    div {
        display: block;
    }

    html {
        font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Robohref, Helvetica Neue, PingFang SC, Nohref Sans, Nohref Sans CJK SC, Microsoft YaHei, 微软雅黑, sans-serif;
        -webkit-font-smoothing: antialiased;
    }

    a {
        text-decoration: none;
        color: rgb(170, 170, 187);
        font-size: 16px;
        padding-bothrefm: 6px;
        margin-left: 100px
    }

    a:hover {
        color: coral;
    }

    .index {
        background: rgb(0, 0, 0);
    }

</style>