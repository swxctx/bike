<template>
    <div>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>昵称</td>
                <td>手机号</td>
                <td>性别</td>
                <td>邮箱</td>
                <td>签名</td>
                <td>真实姓名</td>
                <td>身份证号</td>
                <td>状态</td>
                <td>操作</td>
            </tr>
            <tr v-for="person in people">
                <td>{{ person.id  }}</td>
                <td>{{ person.name  }}</td>
                <td>{{ person.phone  }}</td>
                <td>{{ person.sex  }}</td>
                <td>{{ person.emial  }}</td>
                <td>{{ person.bio  }}</td>
                <td>{{ person.card_name  }}</td>                
                <td>{{ person.card_id  }}</td>
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
                    id: '1',
                    name: 'xiechuang',
                    phone: '15776869842',
                    emial: 'swxctx@yeah.net',
                    bio: 'swxctx',
                    card_name: '解闯',
                    card_id: '530381...',
                    status: '正常',
                    sex: '男'
                }
            ]        
        }
    }
}
</script>
