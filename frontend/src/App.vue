<script setup>
import HelloWorld from './components/HelloWorld.vue'
import Father from "./components/Father.vue";
</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <Father/>
  <router-view></router-view>
</template>

<style>
#logo {
  display: block;
  width: 25%;
  height: 38%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
