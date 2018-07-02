<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>修改密码</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="password" placeholder="请输入密码" v-model="old_password">
            <input type="password" placeholder="请输入新密码" v-model="password">
            <input type="password" placeholder="请再次输入新密码" v-model="againpassword">
            <button v-on:click="addUser">确认修改</button>
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
import {setCookie,getCookie,delCookie} from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            showLogin: true,
            showTishi: false,
            tishi: '',
            old_password: '',
            password: '',
            againpassword: ''
        }
    },
    mounted(){
    },
    methods:{
        addUser() {
            if (this.old_password ==""){
                alert("请输入密码")
            }
            if(this.password==""||this.againpassword ==""){
                alert("请输入新密码及确认密码")
            }
            if (this.password != this.againpassword){
                alert("两次输入密码不一致")
            }
            else{
                let access_token = getCookie('access_token')
                if (access_token==""){
                    alert("登录已过期，请重新登录，即将跳转至登录页面")
                }
                let data = {'access_token':access_token,'password':this.password,"old_password": this.old_password}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8890/bike/user/v1/update_password',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data =="") {
                        this.tishi = "修改成功"
                        this.showTishi = true
                        /*删除cookie*/
                        delCookie('username')
                        delCookie('access_token')
                        setTimeout(function(){
                            this.$router.push('/login')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
