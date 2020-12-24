<template>
    <div>
        <shoes v-for="data in img_list" :data="data"/>
        <pager :page="page" />  <!-- 父组件把当前页传递给子组件-->
    </div>
</template>

<script>
    import shoes from "../components/shoes"
    import pager from "../components/pager"
    export default {
        components: {
            shoes,
            pager,
        },
        data() {
            return {
                img_list: "",
                page:'',
                tag:''
            }
        },
        created() {
            this.page = Number(this.$route.params.page)
            this.tag = this.$route.params.tag+'/'
            this.$store.state.tag = "/"+this.tag  //当前页的类别记录仓库
            // ajax请求后台
            this.$axios({
                url: "http://127.0.0.1:8080/"+this.tag+ this.page,  //动态url匹配
                method: "GET",
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            }).then(response => {
                    this.img_list = response.data['msg']
                }
            ).catch(error => {
                    console.log(error)
                }
            );
        },
    }


</script>

<style scoped>
    body{
        background-color: aliceblue;
    }

</style>