<template>
    <div>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>单车编号</td>
                <td>经度</td>
                <td>维度</td>
                <td>状态</td>
                <td>操作</td>
            </tr>
            <tr v-for="person in people">
                <td>{{ person.id  }}</td>
                <td>{{ person.by_id  }}</td>
                <td>{{ person.longi_tude  }}</td>
                <td>{{ person.lati_tude  }}</td>
                <td>{{ person.status  }}</td>
                <td>
                    <button v-on:click="login" style="background-color:#41b883; color:#fff; font-size:14px;">封禁</button>
                    <button v-on:click="login" style="background-color:#41b883; color:#fff; font-size:14px;">正常</button>
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
            people: [
                {
                    name: 'xiechuang',
                    sex: '男'
                }
            ]        
        }
    }
}
</script>
