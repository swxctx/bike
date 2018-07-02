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
            <tr v-for="user in userList">
                <td>{{ user.id  }}</td>
                <td>{{ user.username  }}</td>
                <td>{{ user.phone  }}</td>
                <td>{{ user.sex  }}</td>
                <td>{{ user.email  }}</td>
                <td>{{ user.bio  }}</td>
                <td>{{ user.card_name  }}</td>                
                <td>{{ user.card_id  }}</td>
                <td>{{ user.status  }}</td>
                <td>
                    <button v-on:click="changeStatus(user.id,1)" style="background-color:#41b883; color:#fff; font-size:14px; cursor: pointer;">封禁</button>
                    <button v-on:click="changeStatus(user.id,0)" style="background-color:#41b883; color:#fff; font-size:14px; cursor: pointer;">正常</button>
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
            userList: [
                {
                    id: '',
                    username: '',
                    phone: '',
                    email: '',
                    sex: '',
                    card_name: '',
                    card_id: '',
                    status: '',
                    bio: ''
                }
            ]        
        }
    },
    mounted(){
        let data = {}
        /*接口请求*/
        this.$http.post('http://127.0.0.1:8894/bike_mp/user/v1/get_list',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.userList = res.data.result.list.list
            }
        })
    },
    methods:{
        changeStatus(uid,status){
            // alert("确认操作?")
            let data = {'id':uid,'status':status}
            /*接口请求*/
            this.$http.post('http://127.0.0.1:8894/bike_mp/user/v1/forbid',data,{"emulateJSON":true}).then((res)=>{
                console.log(res)
                if (res.data.success == false) {
                    alert(res.data.error.content)
                }
                if (res.data =="") {
                    setTimeout(function(){
                        // alert("操作成功")
                        location.reload()
                    }.bind(this),1000)
                }
            })
        }
    }
}
</script>
