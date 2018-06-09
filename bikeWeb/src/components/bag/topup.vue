<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>充值</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入充值金额(元)" v-model="amount">
            <input type="text" placeholder="请输入手机号(支付宝绑定账号)" v-model="phone">
            <button v-on:click="topup">充值</button>
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
            amount: '',
            phone: ''
        }
    },
    mounted(){
        alert("注意!!! 手机号为支付宝绑定手机号，否则将充值失败")
    },
    methods:{
        ToUseBike() {
            this.showLogin = true
        },
        topup() {
            if(this.amount == ""){
                alert("输入金额不能为空")
            }if(this.phone==""){
                alert("请输入手机号")
            }else{
                alert("确定充值吗?")
                let access_token = getCookie('access_token')
                if (access_token==""){
                    alert("登录已过期，请重新登录，即将跳转至登录页面")
                }
                let data = {'access_token':access_token,"phone": this.phone,"amount":this.amount}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8890/bike/bag/v1/top_up',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data.success ==true) {
                        this.tishi = "充值成功"
                        this.showTishi = true
                        setTimeout(function(){
                            this.$router.push('/detail')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
