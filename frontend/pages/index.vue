<template>
  <section class="section">
    <div class="container">
      <div class="block">
          <div class="block video-block" v-for="item in items" :key="item.id">
            <AppVideo
              :item="item"
              :video-id="item.id" 
            />
          </div>
      </div>
    </div>
  </section>
</template>

<script>
  import ROUTES from '~/routes/api';
  import AppVideo from '~/components/AppVideo';
  export default {
    components: {AppVideo},
    computed: {
      items() {
        return this.$store.getters.getPopularVideos
      }
    },
    async fetch({store}) {
      const payload = {
        uri: ROUTES.GET.POPULARS
      }
      if (store.getters.getPopularVideos && store.getters.getPopularVideos.length > 0) {
        return
      }
      await store.dispatch('fetchPopularVideos', payload)
    }
  }
</script>

<style scoped>
  .video-block {
    max-width: 900px;
  }
</style>

