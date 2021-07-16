<template>
  <div class="fullscreen text-white text-center q-pa-md flex flex-center">
    <div>
      <div style="font-size: 30vh">
        404
      </div>

      <div class="text-h2 text-amber-1" style="opacity:.4">
        不好意思，发生了未知的错误！
      </div>

      <q-btn
        class="q-mt-xl"
        color="white"
        text-color="blue"
        unelevated
        to="/"
        label="返回主页"
        no-caps
      />
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Error404'
})
</script>

<style>
  body{
    background-image: url('https://w.wallhaven.cc/full/mp/wallhaven-mpe799.jpg')
  }
</style>
