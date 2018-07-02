<template>
    <div>
        <h3>欢迎 {{name}}</h3>
        <!-- <span>真实姓名: {{name}}</span></br>
        <span>身份证号: {{name}}</span> -->
         <table border="1px" width="600px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>真实姓名</td>
                <td>身份证号</td>
            </tr>
                <td>{{ card_name  }}</td>
                <td>{{ card_id  }}</td>
            </tr>             
        </table>
    </div>
</template>

<script>
/*引入公共方法*/
import { setCookie,getCookie,delCookie } from '../../assets/js/cookie.js'
    export default{
        data(){
            return{
                name: '',
                card_name: '',
                card_id: ''
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
                        if (res.data.result.card_name != "" && res.data.result.card_id !=""){
                            this.card_name = res.data.result.card_name
                            this.card_id = res.data.result.card_id
                        }else{
                            setTimeout(function(){
                                this.$router.push('/start_auth')
                            }.bind(this),1000)
                        }
                    }
                })
            }
        },
        methods:{}
    }
</script>