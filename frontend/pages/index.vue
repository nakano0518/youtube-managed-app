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
      <nav class="pagination">
        <a
          href.prevent="#"
          class="pagination-next"
          @click="loadMore"
        >
        More
        </a>
      </nav>
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
      },
      nextPageToken() {
        return this.$store.getters.getMeta.nextPageToken
      },
    },
    methods: {
      loadMore() {
        const payload = {
          uri: ROUTES.GET.POPULARS,
          params: {
            pageToken: this.nextPageToken
          }
        }
        this.$store.dispatch('fetchPopularVideos', payload)
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
  .pagination {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
  a.pagination-next {
    text-decoration: none;
    color: #333;
    font-weight: bold;
    border: 1px solid #333;
    border-radius: 0.5rem;
    padding: 0.3rem 0.5rem;
    background-color: #fff;

  }
</style>

