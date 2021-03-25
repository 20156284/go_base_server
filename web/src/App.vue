<template>
  <div id="app">
    <router-view/>
  </div>

</template>

<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
  name: 'App',
})
</script>

<style lang="less">
#app{
  text-align: center;
}
</style>
