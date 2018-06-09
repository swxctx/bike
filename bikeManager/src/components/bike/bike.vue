<template>
    <div>
        <router-link tag='a' :to="'/add_bike'" ><span>新增单车</span></router-link></br></br>
        <table border="1px" width="1000px" align="center" cellpadding="0" cellspacing="0" style="border-collapse:collapse">           
            <tr>
                <td>编号</td>
                <td>单车编号</td>
                <td>IP</td>                
                <td>经度</td>
                <td>纬度</td>
                <td>状态</td>
                <td>操作</td>
            </tr>
            <tr v-for="bike in bikes">
                <td>{{ bike.id  }}</td>
                <td>{{ bike.by_id  }}</td>
                <td>{{ bike.ip  }}</td>
                <td>{{ bike.longi_tude  }}</td>
                <td>{{ bike.lati_tude  }}</td>
                <td>{{ bike.status  }}</td>
                <td>
                    <button v-on:click="changeStatus(bike.by_id,1)" style="background-color:#41b883; color:#fff; font-size:14px; cursor: pointer;">封禁</button>
                    <button v-on:click="changeStatus(bike.by_id,0)" style="background-color:#41b883; color:#fff; font-size:14px; cursor: pointer;">正常</button>
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
            bikes: [
                {
                    id: '',
                    by_id: '',
                    ip: '',
                    longi_tude: '',
                    lati_tude: '',
                    status: ''
                }
            ]        
        }
    },
    mounted(){
        let data = {}
        /*接口请求*/
        this.$http.post('http://127.0.0.1:8894/bike_mp/bike/v1/get_list',data,{"emulateJSON":true}).then((res)=>{
            console.log(res)
            if (res.data.success == false){
                alert(res.data.error.content)
            }
            if(res.data.success == true){
                this.bikes = res.data.result.list
            }
        })
    },
    methods:{
        changeStatus(bike_id,status){
            // alert("确认操作?")
            let data = {'bike_id':bike_id,'status':status}
            /*接口请求*/
            this.$http.post('http://127.0.0.1:8894/bike_mp/bike/v1/change',data,{"emulateJSON":true}).then((res)=>{
                console.log(res)
                if (res.data.success == false) {
                    alert(res.data.error.content)
                }
                if (res.data =="") {
                    setTimeout(function(){
                        // this.$router.push('/bike_list')
                        alert("操作成功")
                    }.bind(this),1000)
                }
            })
        }
    }
}
</script>
