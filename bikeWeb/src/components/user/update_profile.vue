<template>
    <div>
        <div class="login-wrap">
            <h3>资料修改</h3>
            <input type="text" placeholder="" v-model="username">
            <input type="text" placeholder="" v-model="phone">
            <span>选择性别</span>
            <select v-model="sex">  
                <option v-for="item in items" v-bind:value="item.value">{{item.text}}</option>  
            </select>  
            <input type="text" placeholder="" v-model="email">
            <input type="text" placeholder="" v-model="bio">
            <button v-on:click="update_profile">保存</button>
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
            username: '',
            phone: '',
            email: '',
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
                        if (res.data.result.sex =="0"){
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
                    }
                }
            })
        }
    },
    methods:{
        update_profile() {
            alert("确定修改吗?")
            let access_token = getCookie('access_token')
            if (access_token==""){
                alert("登录已过期，请重新登录，即将跳转至登录页面")
            }
            let data = {'access_token':access_token,"username":this.username,"phone": this.phone,"email": this.email,"bio": this.bio, "sex": this.sex}
            /*接口请求*/
            this.$http.post('http://127.0.0.1:8890/bike/user/v1/update_profile',data,{"emulateJSON":true}).then((res)=>{
                console.log(res)
                if (res.data.success == false) {
                    alert(res.data.error.content)
                }
                if (res.data.success ==true) {
                    setTimeout(function(){
                        this.$router.push('/profile')
                    }.bind(this),1000)
                }
            })
        }
    }
}
</script>
