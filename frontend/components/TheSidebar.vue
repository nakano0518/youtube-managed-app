<template>
    <div class="section">
        <aside class="menu">
            <ul class="menu-list">
                <li>
                    <nuxt-link
                    to="/"
                    >
                    ホーム
                    </nuxt-link> 
                </li>
            </ul>
        </aside>
    </div>
</template>   

<style scoped>
    .section {
        margin-top: 1rem;
    }
    a {
        color: #333;
        text-decoration: none;
    }
    a:hover {
        color: #aaa;
    }
    li {
        padding: 0;
        list-style-type: none;
    }
</style>