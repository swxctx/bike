<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>添加单车</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <input type="text" placeholder="请输入单车ip" v-model="ip">
            <input type="text" placeholder="请输入单车端口" v-model="port">
            <button v-on:click="addBike">确认添加</button>
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
            ip: '',
            port: '',
        }
    },
    mounted(){
    },
    methods:{
        addBike() {
            if (this.ip ==""){
                alert("请输入单车ip")
            }
            if (this.port==""){
                alert("请输入单车端口")
            }
            else{
                let data = {'ip':this.ip,'port':this.port}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8894/bike_mp/bike/v1/add',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false) {
                        alert(res.data.error.content)
                    }
                    if (res.data =="") {
                        this.tishi = "添加成功"
                        this.showTishi = true
                        setTimeout(function(){
                            this.$router.push('/bike_list')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
