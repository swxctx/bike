<template>
    <div>
        <span>我的余额： {{balance/100}}元</span>
        <h3>充值记录</h3>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>充值金额</td>
                <td>支付宝账号</td>
                <td>充值时间</td>
                <td>状态</td>
            </tr>
            <tr v-for="log in toplog">
                <td>{{ log.count  }}元</td>
                <td>{{ log.phone  }}</td>                
                <td>{{ log.top_ts  }}</td>
                <td>{{ log.status  }}</td>
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
            balance: '',
            toplog: [
                {
                    top_ts: '',
                    count: '',
                    status: '',
                    phone: ''
                }
            ]        
        }
    },
    mounted(){
        let access_token = getCookie('access_token')
        if (access_token==""){
            alert("登录已过期，请重新登录，即将跳转至登录页面")
        }
        let data = {'access_token':access_token}
        /*接口请求*/
        this.$http.post('http://127.0.0.1:8890/bike/bag/v1/detail',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.balance = res.data.result.balance
                this.toplog = res.data.result.list
            }
        })
    }
}
</script>
