<template>
    <div>
        <h3>单车告警信息</h3>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>单车编号</td>
                <td>上锁状态</td>
                <td>单车状态</td>
                <td>告警时间</td>
            </tr>
            <tr v-for="alarm in alarm_list">
                <td>{{ alarm.id  }}</td>
                <td>{{ alarm.by_id  }}</td>
                <td>{{ alarm.mh_count  }}</td>
                <td>{{ alarm.sw_count  }}</td>
                <td>{{ alarm.alarm_ts  }}</td>
            </tr>             
        </table>
    </div>
</template>

<script>
/*引入公共方法*/
export default{
    data(){
        return{
            alarm_list: [
                {
                    id: '',
                    by_id: '',
                    mh_count: '',
                    sw_count:'',
                    alarm_ts: ''
                }
            ]        
        }
    },
    mounted(){
        let data = {}
        /*接口请求*/
        this.$http.post('http://127.0.0.1:8894/bike_mp/bike/v1/alarm_list',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.alarm_list = res.data.result.list.list
            }
        })
    }
}
</script>
