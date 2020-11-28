<template>
    <nav class="navbar">
        <div class="navbar-brand">
            <nuxt-link
                class="navbar-item subtitle is-4"
                to="/"
            >
            <strong> Youtube Player</strong>
            </nuxt-link>
        </div>

        <div class="navbar-menu">
            <div class="navbar-start">
                <div class="navbar-item">
                <div class="field has-addons">
                    <form action="/search" method="get">
                    <div class="field has-addons">
                        <div class="control">
                        <input name="q" class="input" type="text" placeholder="検索">
                        </div>
                        <div class="control">
                        <button class="button">
                            Search
                        </button>
                        </div>
                    </div>
                    </form>
                </div>
                </div>
            </div>
        </div>
    </nav>
</template>  

<style scoped>
  .navbar {
    background: #fff
  }
  strong {
      color: #555;
      font-size: 1.5rem;
  }
</style>