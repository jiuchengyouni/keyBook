<template>
  <!--1.首先，弹窗页面中要有el-dialog组件即弹窗组件，我们把弹窗中的内容放在el-dialog组件中-->
  <!--2.设置:visible.sync属性，动态绑定一个布尔值，通过这个属性来控制弹窗是否弹出-->
  <el-dialog title="弹窗" :visible.sync="detailVisible" width="55%">
    <div>{{msg}}</div>
    <div>aaaaaaaa</div>
  </el-dialog>
</template>

<script>
import {computed} from 'vue'
export default {
  
  name: "dialogComponent",
  props:{
    msg:String
  },
  setup(props){
    return{

    }
  },
  methods:{
    //3.定义一个init函数，通过设置detailVisible值为true来让弹窗弹出，这个函数会在父组件的方法中被调用
    init(data){
      this.detailVisible=true;
      console.log(data);
    },
  },
  data(){
    let msg;
    return{
      msg,
      detailVisible:false
    }
  }
}
</script>