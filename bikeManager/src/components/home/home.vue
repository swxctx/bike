<template>
    <div>
        <img src="../../assets/welcome.jpg" width="50%">
        <h3>欢迎您 {{name}}</h3>
        <a href="#" @click="quit">注销登录</a>
    </div>
</template>

<script>
/*引入公共方法*/
import { setCookie,getCookie,delCookie } from '../../assets/js/cookie.js'
    export default{
        data(){
            return{
                name: ''
            }
        },
        mounted(){
            /*页面挂载获取保存的cookie值，渲染到页面上*/
            let uname = getCookie('username')
            this.name = uname
            /*如果cookie不存在，则跳转到登录页*/
            if(uname == ""){
                this.$router.push('/login')
            }else{
                
            }
        },
        methods:{
            quit(){
                this.tishi = "注销成功"
                this.showTishi = true
                /*删除cookie*/
                delCookie('username')
                setTimeout(function(){
                    this.$router.push('/login')
                }.bind(this),1000)
            }
        }
    }
</script>