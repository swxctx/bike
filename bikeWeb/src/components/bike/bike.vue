<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>单车开锁</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入单车编号" v-model="bike_id">
            <button v-on:click="login">开始开锁</button>
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
            bike_id: '',
        }
    },
    mounted(){},
    methods:{
        ToUseBike() {
            this.showLogin = true
        },
        login() {
            if(this.bike_id == ""){
                alert("请输入单车编号")
            }else{
                alert("确定开锁吗?")
                let access_token = getCookie('access_token')
                if (access_token==""){
                    alert("登录已过期，请重新登录，即将跳转至登录页面")
                }
                let data = {'bike_id':this.bike_id,'access_token':access_token}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8890/bike/use/v1/start_use',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data.success ==true) {
                        this.tishi = "开锁成功,可以开始使用啦"
                        this.showTishi = true
                        setTimeout(function(){
                            this.$router.push('/map')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
