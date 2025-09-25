<template>
  <!-- html -->
  <div class="person">
    <h2>姓名：{{ name }}</h2>
    <h2>年龄：{{ age }}</h2>
    <button @click="changeName">修改名字</button>
    <button @click="changeAge">修改年龄</button>
    <button @click="showTel">查看联系方式</button>
  </div>
</template>

<script lang="ts">
// JS 或 TS
import {ref} from "vue";

export default {
  name: 'Person', // 组件名称
  setup() {
    // 数据
    let name = '张三' // 注意这样写name不是响应式数据
    let age = ref(12) // 注意这样写age不是响应式数据
    let tel = '138888888'

    // 方法
    function changeName() {
      name = '李四'
    }

    function changeAge() {
      age.value += 1
    }

    function showTel() {
      alert(tel)
    }

    // 将数据，方法交出去，模版中才可以使用
    return {name, age, tel, changeName, changeAge, showTel}
  }
}
</script>

<style>
/* css样式 */
.person {
  background-color: skyblue;
  box-shadow: 0 0 10px;
  border-radius: 10px;
  padding: 20px;
}

button {
  margin: 0 5px;
}
</style>