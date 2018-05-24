<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>添加用户</h3>
            <span>选择用户类型</span>
            <select v-model="selected">  
                <option v-for="item in items" v-bind:value="item.value">{{item.text}}</option>  
            </select>  
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入用户名" v-model="username">
            <input type="text" placeholder="请输入真实姓名" v-model="name">
            <input type="password" placeholder="请输入密码" v-model="password">
            <input type="password" placeholder="请输入密码" v-model="againPassword">
            <button v-on:click="addUser">确认添加</button>
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
            is_admin: '',
            general_user: '',
            items: [{text:'管理员',value:'1'},{text:'普通用户',value:'0'}],  
            selected: ''
        }
    },
    mounted(){
    },
    methods:{
        addUser() {
            if (this.selected == ""){
                alert("请选择用户类型")
            }
            if(this.username == "" || this.password == ""){
                alert("请输入用户名或密码")
            }
            if (this.name ==""){
                alert("请输入用户姓名")
            }
            else{
                let data = {'username':this.username,'password':this.password,'name': this.name,'is_admin': selected}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8894/bike_mp/user/v1/login',data,{"emulateJSON":true}).then((res)=>{
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
