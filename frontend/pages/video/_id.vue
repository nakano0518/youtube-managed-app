<template>
  <div class="columns">
      <div class="column is-6">
          <div class="block video-player">
              <youtube
                :video-id="this.$route.params.id"
                ref="youtube"
              ></youtube>
          </div>
          <div class="box">
              <p>
                  <span class="title is-4">{{item.snippet.title}}</span>
              </p>
              <div class="level">
                  <div class="level-left">
                    {{item.snippet.channelTitle}}
                    <br />
                    {{item.snippet.publishedAt}}
                  </div>
              </div>

              <hr>
              <p>{{item.snippet.description}}</p>
          </div>
      </div>
      <div class="column is-4">
          <div class="box">
              <p>
                  <span>関連動画</span>
              </p>
              <div v-for="relatedItem in relatedItems" :key="relatedItem.id">
                  <hr>
                  <nuxt-link
                    :to="`/${relatedItem.id.videoId}`"
                  >
                    <article>
                        <img :src="relatedItem.snippet.thumbnails.default.url" alt="thumbnail">
                        <p>
                            <strong>{{relatedItem.snippet.title}}</strong>
                            <small>{{relatedItem.snippet.channelTitle}}</small>
                        </p>
                    </article>
                  </nuxt-link>
              </div>
          </div>
      </div>
  </div>
</template>

<script>
import ROUTES from '~/routes/api'
export default {
    computed: {
        item() {
            console.log(this.$store.getters.getVideo)
            return this.$store.getters.getVideo
        },
        relatedItems() {
            return this.$store.getters.getRelatedVideos
        },
    },
    async fetch({store, route}) {
        await store.dispatch('findVideo', {
            uri: ROUTES.GET.VIDEO.replace(':id', route.params.id)
        })
        await store.dispatch('fetchRelatedVideos', {
            uri: ROUTES.GET.RELATED.replace(':id', route.params.id)
        })
    }
}
</script>

<style>
    
</style>