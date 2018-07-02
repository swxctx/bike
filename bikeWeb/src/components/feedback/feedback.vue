<template>
        <div class="login-wrap">
            <!-- <button v-on:click="login">登录</button> -->
            <textarea cols ="50" rows = "5" v-model="feedtext"></textarea>
            <button v-on:click="dofeed">提交</button>
        </div>
</template>

<style>
    .login-wrap{text-align:center;}
    input{display:block; width:250px; height:40px; line-height:40px; margin:0 auto; margin-bottom: 10px; outline:none; border:1px solid #888; padding:10px; box-sizing:border-box;}
    p{color:red;}
    button{display:block; width:70px; height:40px; line-height: 40px; margin:0 auto; border:none; background-color:#41b883; color:#fff; font-size:16px; margin-bottom:5px;cursor: pointer;}
    span{cursor:pointer;}
    span:hover{color:#41b883;}
</style>

<script>
import {getCookie} from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            feedtext: ''
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
        dofeed() {
            if (this.feedtext=="") {
                alert("输入内容不能为空哦")
            }
            let access_token = getCookie("access_token")
            if (access_token=="") {
                alert("登录已过期，请重新登录后再提交")
                this.$router.push('/login')
            }else{
                let data = {'access_token': access_token,'content': this.feedtext}
                this.$http.post('http://127.0.0.1:8890/bike/feedback/v1/report',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    if (res.data.success == false){
                        alert(res.data.error.content)
                    }
                    if(res.data.success == true){
                        /*提交成功回到主页*/
                        setTimeout(function(){
                            this.$router.push('/')
                        }.bind(this),1000)
                    }
                })
            }
        }
    }
}
</script>
