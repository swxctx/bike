<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>修改密码</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入用户名" v-model="username">
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
import {setCookie,getCookie} from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            showLogin: true,
            showTishi: false,
            tishi: '',
            username: '',
            old_password: '',
            password: '',
            againpassword: ''
        }
    },
    mounted(){
    },
    methods:{
        addUser() {
            if (this.username == ""){
                alert("请输入用户名")
            }
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
                let data = {'username':this.username,'password':this.password,"old_password": this.old_password}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8894/bike_mp/account/v1/update_password',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data =="") {
                        this.tishi = "修改成功"
                        this.showTishi = true
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
