<template>
        <div class="login-wrap">
            <!-- <button v-on:click="login">登录</button> -->
            <textarea cols ="50" rows = "5" v-model="feedtext"></textarea>
            <button v-on:click="submit">提交</button>
        </div>
</template>

<style>
    .login-wrap{text-align:center;}
    input{display:block; width:250px; height:40px; line-height:40px; margin:0 auto; margin-bottom: 10px; outline:none; border:1px solid #888; padding:10px; box-sizing:border-box;}
    p{color:red;}
    button{display:block; width:50px; height:40px; line-height: 40px; margin:0 auto; border:none; background-color:#41b883; color:#fff; font-size:16px; margin-bottom:5px;cursor: pointer;}
    span{cursor:pointer;}
    span:hover{color:#41b883;}
</style>

<script>
import { setCookie,getCookie,delCookie } from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            feedtext: '',
            access_token: ''
        }
    },
    mounted(){
        let accessToken = getCookie('access_token')
        this.access_token = accessToken
        console.log(accessToken)
        if(accessToken == ""){
            alert("您的登录已过期，请重新登录")
            this.$router.push('/login')
        }
    },
    methods:{
        submit() {
            console.log(this.feedtext)
        }
    }
}
</script>
