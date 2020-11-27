<!-- layout --> 

<template>
  <div id="app">
    <TheHeader />
    <div class="row">
      <div class="col-2">
        <TheSidebar />
      </div>
      <div class="col-10">
        <nuxt />
      </div>
    </div>
  </div>
</template>

<script>
import TheHeader from '~/components/TheHeader';
import TheSidebar from '~/components/TheSidebar';
export default {
  components: {
    TheHeader,
    TheSidebar,
  }
}
</script>

<style scoped>
  #app {
    min-height: 100vh;
    width: 100vw;
    background: #ccc
  }
</style>


