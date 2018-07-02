<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>登录</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入用户名" v-model="username">
            <input type="password" placeholder="请输入密码" v-model="password">
            <button v-on:click="login">登录</button>
            <span v-on:click="ToRegister">没有账号？马上注册</span>
        </div>

        <div class="register-wrap" v-show="showRegister">
            <h3>注册</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入用户名" v-model="newUsername">
            <input type="text" placeholder="请输入手机号" v-model="phone">
            <input type="text" placeholder="请输入邮箱(选填)" v-model="email">
            <input type="text" placeholder="个性签名(选填)" v-model="bio">
            <span>选择性别</span>
            <select v-model="sex">  
                <option v-for="item in items" v-bind:value="item.value">{{item.text}}</option>  
            </select>  
            <input type="password" placeholder="请输入密码" v-model="newPassword">
            <input type="password" placeholder="请再次输入密码" v-model="againPassword">
            <button v-on:click="register">注册</button>
            <span v-on:click="ToLogin">已有账号？马上登录</span>
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
            items: [{text:'男',value:'1'},{text:'女',value:'2'}],  
            showLogin: true,
            showRegister: false,
            showTishi: false,
            tishi: '',
            username: '',
            newUsername: '',
            phone: '',
            password: '',
            newPassword: '',
            againPassword: '',
            email: '',
            sex: '',
            bio: ''
        }
    },
    mounted(){
    /*页面挂载获取cookie，如果存在username的cookie，则跳转到主页，不需登录*/
        if(getCookie('username') && getCookie('access_token')){
            this.$router.push('/home')
        }
    },
    methods:{
        ToRegister() {
        this.showRegister = true
        this.showLogin = false
        },
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
                this.$http.post('http://127.0.0.1:8890/bike/user/v1/login',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    // console.log(res.data.error)
                    if (res.data.success == false){
                        alert(res.data.error.content)
                    }
                    if(res.data.success == true && res.data.result.username !="" && res.data.result.access_token != ""){
                        console.log("success")
                        this.tishi = "登录成功"
                        this.showTishi = true
                        setCookie('username',res.data.result.username,1000*60)
                        setCookie('access_token',res.data.result.access_token,1000*60)
                        setTimeout(function(){
                            this.$router.push('/')
                        }.bind(this),1000)
                    }
                })
            }
        },
        register(){
            if(this.newUsername == "" || this.newPassword == ""){
                alert("请输入用户名或密码")
            }
            if(this.phone == ""){
                alert("电话号码为必填项，请输入")
            }
            else{
                let data = {'username':this.newUsername,'phone': this.phone,'password':this.newPassword,'email':this.email,"sex": this.sex,"bio": this.bio}
                this.$http.post('http://127.0.0.1:8890/bike/user/v1/register',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false){
                        alert(res.data.error.content)
                    }
                    if(res.data.success == true && res.data.result.username !="" && res.data.result.access_token != ""){
                        console.log("success")
                        this.tishi = "注册成功"
                        this.showTishi = true
                        setCookie('username',res.data.result.username,1000*60)
                        setCookie('access_token',res.data.result.access_token,1000*60)
                        setCookie('phone',res.data.result.phone,1000*60)
                        /*注册成功之后再跳回登录页*/
                        setTimeout(function(){
                            this.showRegister = false
                            this.showLogin = true
                            this.showTishi = false
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
