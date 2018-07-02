<template>
    <div>
        <div class="login-wrap" v-show="showLogin">
            <h3>骑行中...</h3>
            <p v-show="showTishi">{{tishi}}</p>
            <button v-on:click="finish">结束使用</button></br>
        </div>
        <div style="left:10%;margin-right:12%;height:500px;border:#ccc solid 1px;" id="allmap">
        </div>
    </div>
</template>

<style>
    .login-wrap{text-align:center;}
    input{display:block; width:250px; height:40px; line-height:40px; margin:0 auto; margin-bottom: 10px; outline:none; border:1px solid #888; padding:10px; box-sizing:border-box;}
    p{color:red;}
    button{display:block; width:250px; height:40px; line-height: 40px; margin:0 auto; border:none; background-color:#41b883; color:#fff; font-size:16px; margin-bottom:5px;cursor: pointer;}
    span{cursor:pointer;}
    span:hover{color:#41b883;}
</style>

<script>
import {setCookie,getCookie} from '../../assets/js/cookie.js'
export default{
    data() {
        return {
            showLogin: true,
            showTishi: false,
            tishi: '',
            bike_id: '',
            amount: ''
        }
    },
    mounted(){
        this.initMap()
    },
    methods:{
        //这几个地方加this
        initMap () {
            this.createMap() ; //创建地图 
        },
        createMap(){
        // 百度地图API功能
        var map = new BMap.Map("allmap");
        var point = new BMap.Point(120.328033,31.686802);
        map.centerAndZoom(point, 15);

        var geolocation = new BMap.Geolocation();
        geolocation.getCurrentPosition(function (r) {
                if (this.getStatus() == BMAP_STATUS_SUCCESS) {
                    var mk = new BMap.Marker(r.point);
                    map.addOverlay(mk);
                    map.panTo(r.point);
                }else {
                    alert('failed' + this.getStatus());
                }
            }, { enableHighAccuracy: true })
        },
        finish() {
            this.initMap()
            alert("确定结束吗?")
            let access_token = getCookie('access_token')
            if (access_token==""){
                alert("登录已过期，请重新登录，即将跳转至登录页面")
            }
            let point_lng = getCookie("point_lng")
            let point_lat = getCookie("point_lat")
            let bike_id = getCookie("use_bike_id")

            if (bike_id == ""){
                alert("结束失败，稍后再试")
            }
            let data = {'bike_id':bike_id,'access_token':access_token,"point_lng":point_lng,"point_lat": point_lat}
            /*接口请求*/
            this.$http.post('http://127.0.0.1:8890/bike/use/v1/finish_use',data,{"emulateJSON":true}).then((res)=>{
                console.log(res)
                if (res.data.success == false) {
                    alert(res.data.error.content)
                }
                if (res.data.success ==true) {
                    this.tishi = "您已完成本次骑行"
                    this.showTishi = true
                    this.amount = res.data.result.amount/100
                    let msg = "本次累计骑行"+(res.data.result.end_ts-res.data.result.start_ts)/60+"分钟，"+"总计消费"+this.amount+"元，已经从您的余额中自动扣除，Swxctx共享单车感谢您的使用，祝您生活愉快"
                    alert(msg)
                    setTimeout(function(){
                        this.$router.push('/map')
                    }.bind(this),1000)
                }
            })
        }
    }
}
</script>
