<template>
  <!-- 6.定义一个点击事件-->
  <button @click="handleClick('父组件')">点击</button>
  <!-- 3.在页面中使用dialog-component组件-->
  <!-- 4.设置v-if语句，通过动态改变Visiable值用来控制弹窗是否弹出-->
  <!-- 5.设置ref属性，相当于唯一标识dialog-component组件-->
  <dialog-component v-if="Visiable" ref="dialog"></dialog-component>
</template>

<script>
//	1.引入弹窗组件dialogComponent
import dialogComponent from "./dialogComponent.vue";
export default {
  //  2.在components中注册dialogComponent组件
  components:{
    dialogComponent
  },
  data(){
    return{
      Visiable:false
    }
  },
  methods:{
    // 7.实现点击事件，点击事件一定要包含以下内容
    handleClick(data){
      if (this.Visiable==false){
        this.Visiable=true;
        this.$nextTick(()=>{
          //这里的dialog与上面dialog-component组件里面的ref属性值是一致的
          //init调用的是dialog-component组件里面的init方法
          //data是传递给弹窗页面的值
          this.$refs.dialog.init(data);
        })
      }else {
        this.Visiable=false
      }
    },
  }
}
</script>