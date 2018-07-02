<template>
    <div>
        <h3>我的足迹</h3>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>单车编号</td>
                <td>开始时间</td>
                <td>结束时间</td>
                <td>累计骑行</td>
                <td>消费金额</td>                
            </tr>
            <tr v-for="log in uselog">
                <td>{{ log.bike_id  }}</td>
                <td>{{ log.start_ts  }}</td>
                <td>{{ log.end_ts  }}</td>
                <td>{{ log.all_ts  }}</td>
                <td>{{ log.amount  }}</td>
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
            uselog: [
                {
                    bike_id: '',
                    start_ts: '',
                    end_ts: '',
                    all_ts: '',
                    amount: ''
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
        this.$http.post('http://127.0.0.1:8890/bike/use/v1/history_log',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            console.log(res.data.error)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.uselog = res.data.result.list
            }
        })
    }
}
</script>
