<template>
    <div>
        <h3>单车告警信息</h3>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>单车编号</td>
                <td>红外数据</td>
                <td>震动数据</td>
                <td>告警时间</td>
            </tr>
            <tr v-for="person in people">
                <td>{{ person.id  }}</td>
                <td>{{ person.by_id  }}</td>
                <td>{{ person.mh_count  }}</td>
                <td>{{ person.sw_count  }}</td>
                <td>{{ person.alarm_ts  }}</td>
            </tr>             
        </table>
    </div>
</template>

<script>
/*引入公共方法*/
export default{
    data(){
        return{
            people: [
                {
                    id: '1',
                    by_id: '100001',
                    mh_count: '异常',
                    sw_count: '异常',
                    alarm_ts: '2018-05-30 17:00:00'
                },
                {
                    id: '2',
                    by_id: '100002',
                    mh_count: '异常',
                    sw_count: '正常',
                    alarm_ts: '2018-05-30 17:01:00'
                }
            ]        
        }
    }
}
</script>
