<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>登录</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入用户名" v-model="username">
            <input type="password" placeholder="请输入密码" v-model="password">
            <button v-on:click="login">登录</button>
        </div>
    </div>
</template>

<style>
    .login-wrap{text-align:center;}
    input{display:block; width:250px; height:40px; line-height:40px; margin:0 auto; margin-bottom: 10px; outline:none; border:1px solid #888; padding:10px; box-sizing:border-box;}
    p{color:red;}
    button{display:block; width:250px; height:40px; line-height: 40px; margin:0 auto; border:none; background-color:#41b883; color:#fff; font-size:16px; margin-bottom:5px;cursor: pointer;}
    span{cursor:pointer;}
    span:hover{color:#41b883;}
</style>

<script>
import {setCookie,getCookie} from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            showLogin: true,
            showTishi: false,
            tishi: '',
            username: '',
            name: '',
            password: '',
            is_admin: ''
        }
    },
    mounted(){
    /*页面挂载获取cookie，如果存在username的cookie，则跳转到主页，不需登录*/
        if(getCookie('username')){
            this.$router.push('/home')
        }
    },
    methods:{
        ToLogin() {
            this.showRegister = false
            this.showLogin = true
        },
        login() {
            if(this.username == "" || this.password == ""){
                alert("请输入用户名或密码")
            }else{
                let data = {'username':this.username,'password':this.password}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8894/bike_mp/account/v1/login',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data.success ==true && res.data.result.username !="") {
                        this.tishi = "登录成功"
                        this.showTishi = true
                        setCookie('username',res.data.result.username,1000*60)
                        setCookie('name',res.data.result.name,1000*60)
                        setCookie('is_admin',res.data.result.is_admin,1000*60)
                        setTimeout(function(){
                            this.$router.push('/home')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
