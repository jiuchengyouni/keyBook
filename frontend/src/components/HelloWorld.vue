<script setup>
import {reactive, VueElement} from 'vue'
import {Open, Test} from '../../wailsjs/go/main/App'
function greet(name) {
  Open(name)
}
</script>

<template>
  <main>
    <button class="btn" @click="greet">Greet</button>
    <button class="btn" v-for="item in buttonlist4" @click="greet(item)">
      {{ item }}
    </button>
  </main>
</template>

<script>
import {Search, Test2} from '../../wailsjs/go/main/App.js'
export default {
  name: "Mybutton",
  data() {
    return {
      buttonlist4:["11","111"],
      buttonlist3:[{
        name:"1",
        targetPath:"aaa"
      },{
        name:"2",
        targetPath:"aaa"
      }]
    }
  },
  created: async function(){
    //this.buttonlist3[0].name="111111"
    this.buttonlist4=await Search()
    //this.buttonlist3=await Search()
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
.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
