<!--
翻页器组件
-->

<template>
        <div class="pager">

            <a :href="url1" @click="uppage">上一页</a>
            <a :href="url3"><input type="button" value="跳转" @click="tiao"></a><input type="text" style="width: 30px" :value="page" id="tiaozhuan">
            <a :href="url2" @click="dowpage">下一页</a>
        </div>

</template>

<script>
    export default {
        name: "pager",
        props: ['page'],  //当前页码从父组件传递过来
        data() {
            return {
                url1: "",  //上一页
                url2: "",  //下一页
                url3: "",  //跳转页
            }
        },
        methods: {
            uppage() {
                if (Number(this.page)<=1){
                    // 当前所在页对类别从仓库获取
                    this.url1 = this.$store.state.tag + 1
                }else {
                    this.url1 = this.$store.state.tag + (Number(this.page)-1)
                }
            },
            dowpage() {
                this.url2 = this.$store.state.tag + (Number(this.page)+1)
            },
            tiao(){
                this.url3 = this.$store.state.tag + document.getElementById("tiaozhuan").value;
                console.log(this.url3)
            }
        }
    }
</script>

<style scoped>
    .pager {
        position: fixed;
        left: 80%;
        top: 90%;
    }

</style>