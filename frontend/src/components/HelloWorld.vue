<script setup>
import {reactive, VueElement} from 'vue'
import {Open, Test} from '../../wailsjs/go/main/App'
function greet(name) {
  Open(name)
}
</script>
<template>
  <main>
    <p>
      <router-link to="/">退出</router-link>
    </p>
    <router-view></router-view>
    <button class="btn" v-for="item in buttonlist4" @click="greet(item),handleClick(item)">
      {{ item }}
    </button>
    <dialog-component v-if="Visiable" ref="dialog"></dialog-component>
  </main>
</template>

<script>
import {Search, Test2} from '../../wailsjs/go/main/App.js'
import dialogComponent from "./dialogComponent.vue";
import {ref} from "vue";
export default {
  name: "Mybutton",
  components:{
    dialogComponent
  },
  data() {
    return {
      Visiable:false,
      buttonlist4:["11","111"]
    }
  },
  created: async function(){
    this.buttonlist4=await Search()
  },
  methods: {
    // 7.实现点击事件，点击事件一定要包含以下内容
    handleClick(data) {
      if (this.Visiable == false) {
        this.Visiable = true;
        this.$nextTick(() => {
          //这里的dialog与上面dialog-component组件里面的ref属性值是一致的
          //init调用的是dialog-component组件里面的init方法
          //data是传递给弹窗页面的值
          this.$refs.dialog.init(data);
        })
      } else {
        this.Visiable = false
      }
    }
  }
}
</script>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.btn{
  width: auto;
}
</style>
