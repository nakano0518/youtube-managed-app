<template>
    <nuxt-link
        :to="`/video/${videoId}`"
    >
        <article>
            <img :src="item.snippet.thumbnails.medium.url" alt="Image">
            <p>
                <strong>{{item.snippet.title}}</strong>
                <br>
                <small>{{item.snippet.publishedAt}}</small>
                <small>{{item.snippet.channelTitle}}</small>
                <br>
                {{ item.snippet.description | omit }}
            </p>
        </article>
    </nuxt-link>
</template>

<script>
export default {
    props: {
        item: {
            type: Object
        },
        videoId: {
            type: String
        }
    },
    filters: {
        omit: (value) => { //動画の説明120文字超えたら...で省略して表示
            if (!value) {
                return '';
            }
            if (value.length > 120) {
                return value.substr(0, 120) + '...';
            }
            return value;
        }
    }
};
</script>

<style scoped>
    a {
        text-decoration: none;
        color: #333;
    }
    a:hover {
        color: #aaa;
    }
    article {
        display: flex;
    }
    img {
        width: 220px;
        height: 120px;
        border: 1px solid #333;
    }
    p {
        padding-left: 1.5rem;
    }
</style>