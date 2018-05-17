<template>
    <!--地图容器-->
      <div style="left:10%;margin-right:12%;height:450px;border:#ccc solid 1px;" id="allmap">
      </div>
</template>

<script>
  export default{
    data () {
      return {
      }
    },
    mounted () {
      this.initMap()
    },
    methods: {
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
                // alert('您的位置：' + r.point.lng + ',' + r.point.lat);
                alert('定位成功，可以开始使用了.')
            }
            else {
                alert('failed' + this.getStatus());
            }
        }, { enableHighAccuracy: true })
      }
    }
  }
</script>