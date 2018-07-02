<template>
    <div>
        <h3>待处理充值</h3>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>用户ID</td>
                <td>手机号</td>
                <td>支付宝账号</td>
                <td>充值金额</td>
                <td>状态</td>    
                <td>操作</td>    
            </tr>
            <tr v-for="top in topList">
                <td>{{ top.id  }}</td>
                <td>{{ top.uid  }}</td>
                <td>{{ top.phone  }}</td>
                <td>{{ top.count  }}</td>
                <td>{{ top.created_at  }}</td>
                <td>未处理</td>
                <td>
                    <button v-on:click="changeStatus(top.id)" style="background-color:#41b883; color:#fff; font-size:14px; cursor: pointer;">确认</button>
                </td>
            </tr>             
        </table>
    </div>
</template>

<script>
/*引入公共方法*/
export default{
    data(){
        return{
            topList: [
                {
                    id: '',
                    uid: '',
                    phone: '',
                    count: '',
                    created_at: '',
                    status: ''
                }
            ]        
        }
    },
    mounted(){
        let data = {}
        /*接口请求*/
        this.$http.post('http://127.0.0.1:8894/bike_mp/order/v1/get_list',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.topList = res.data.result.list
            }
        })
    },
    methods:{
        changeStatus(tid){
            alert("确认充值吗?")
            let data = {'id':tid}
            /*接口请求*/
            this.$http.post('http://127.0.0.1:8894/bike_mp/order/v1/change',data,{"emulateJSON":true}).then((res)=>{
                console.log(res)
                if (res.data.success == false) {
                    alert(res.data.error.content)
                }
                if (res.data =="") {
                    setTimeout(function(){
                        alert("操作成功")
                    }.bind(this),1000)
                }
            })
        }
    }
}
</script>
