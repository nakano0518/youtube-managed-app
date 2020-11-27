<template>
  <section class="section">
      <div class="items">
          <div class="item" v-for="item in items" :key="item.id">
            <AppVideo
              :item="item"
              :video-id="item.id" 
            />
          </div>
      </div>
  </section>
</template>

<script>
  import ROUTES from '~/routes/api';
  import AppVideo from '../components/AppVideo';
  export default {
    components: {AppVideo},
    computed: {
      items() {
        //console.log(this.$store.getters.getPopularVideos)
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
  .item {
    margin-top: 1rem;
    background-color: #fff;
    padding: 1rem 1rem;
    max-width: 900px;
  }
</style>

