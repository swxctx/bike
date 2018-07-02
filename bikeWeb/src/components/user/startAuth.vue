<template>
    <div>
        <div class="login-wrap">
            <h3>实名认证</h3>
            <input type="text" placeholder="请输入真实姓名" v-model="card_name">
            <input type="text" placeholder="请输入身份证号" v-model="card_id">
            <button v-on:click="auth">开始认证</button>
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
            card_id: '',
            card_name: ''
        }
    },
    mounted(){},
    methods:{
        auth() {
            if(this.card_id == ""){
                alert("请输入真实姓名")
            }else{
                alert("确定认证吗?")
                let access_token = getCookie('access_token')
                if (access_token==""){
                    alert("登录已过期，请重新登录，即将跳转至登录页面")
                }
                let data = {'card_id':this.card_id,'card_name':this.card_name,'access_token':access_token}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8890/bike/user/v1/auth_card',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data.success ==true) {
                        setTimeout(function(){
                            this.$router.push('/auth')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
