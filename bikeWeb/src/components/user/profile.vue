<template>
    <div class="login-wrap">
        <router-link tag='a' :to="'/use'" ><img src="../../assets/mkicon.png"></router-link></br></br>
        <span>用户名： {{username}}</span></br></br>
        <span>手机号： {{phone}}</span></br></br>
        <span>邮箱： {{email}}</span></br></br>
        <span>性别： {{sex}}</span></br></br>
        <span>签名： {{bio}}</span></br></br>
    <router-link tag='a' :to="'/update_profile'" >资料修改</router-link>
    </div>
</template>

<style>
    .login-wrap{text-align:center;}
    span{cursor:pointer;}
    span:hover{color:#41b883;}
</style>


<script>
/*引入公共方法*/
import { setCookie,getCookie,delCookie } from '../../assets/js/cookie.js'
    export default{
        data(){
            return{
                username: '',
                phone: '',
                emial: '',
                sex: '',
                bio: '',
            }       
        },
        mounted(){
            let uname = getCookie('username')
            this.name = uname
            /*如果cookie不存在，则跳转到登录页*/
            if(uname == ""){
                this.$router.push('/')
            }else{
                let access_token = getCookie('access_token')
                if (access_token==""){
                    alert("登录已过期，请重新登录，即将跳转至登录页面")
                }
                let data = {'access_token':access_token}
                /*接口请求*/
                this.$http.post('http://127.0.0.1:8890/bike/user/v1/get_profile',data,{"emulateJSON":true}).then((res)=>{
                    console.log(res)
                    console.log(res.data.error)
                    if (res.data.success == false){
                        alert(res.data.error.content)
                    }
                    if(res.data.success == true){
                        if (res.data.result.username != "" && res.data.result.phone !=""){
                            this.username = res.data.result.username
                            this.phone = res.data.result.phone
                            this.email = res.data.result.email
                            this.bio = res.data.result.bio
                            if (res.data.result.sex =="1"){
                                this.sex = "男"
                            }else{
                                this.sex = "女"
                            }
                            if (this.email == ""){
                                this.email = "暂未设置邮箱"
                            }
                            if (this.bio == ""){
                                this.bio = "暂未设置签名"
                            }
                        }else{
                            setTimeout(function(){
                                this.$router.push('/')
                            }.bind(this),1000)
                        }
                    }
                })
            }
        },
        methods:{}
    }
</script>