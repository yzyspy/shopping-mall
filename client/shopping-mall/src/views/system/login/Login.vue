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
import service from "../../../api/request.ts";

export default {
  name: 'Person', // 组件名称
  setup() {
    // 数据
    let name = ref('张三') // 注意这样写name不是响应式数据
    let age = ref(12) // 注意这样写age不是响应式数据
    let tel = '138888888'

    // 方法
    function changeName() {
      queryUser().then(result => {
        const dataString = JSON.stringify(result);
        console.log('获取到的数据:', dataString);
        name.value = dataString;
      }).catch(error => {
            console.error('获取数据失败:', error);
          });



    }

    function changeAge() {
      age.value += 1
    }

    function showTel() {
      alert(tel)
    }

    async function queryUser() {
      try {
        const response = await service.post('/login/psw', {
          username: 'abc',
          password: '111'
        });
        console.log('创建成功:', response.data);
        return response.data;
      } catch (error) {
        console.error('创建失败:', error);
        throw error;
      }
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