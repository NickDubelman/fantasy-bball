<script context="module">
  export const preload = async function (page, session) {
    let currPage = page.path

    if (Object.keys(page.query).length > 0) {
      const search = new URLSearchParams(page.query)
      currPage += `?${search.toString()}`
    }

    return { currPage }
  }
</script>

<script lang="ts">
  import { stores } from '@sapper/app'

  const { session } = stores()

  export let currPage: string
  let currPageEncoded = ''

  if (typeof window !== 'undefined') {
    currPageEncoded = btoa(currPage)
  }
</script>

<main>
  {#if $session.user}
    <slot />
  {:else}
    <a href="/auth/login?next={currPageEncoded}">Login</a>
  {/if}
</main>

<style>
  main {
    position: relative;
    max-width: 56em;
    background-color: white;
    padding: 2em;
    margin: 0 auto;
    box-sizing: border-box;
  }

  a {
    background-color: #5f14a5;
    color: white;
    padding: 1em 1.4em;
    text-decoration: none;
    text-transform: uppercase;
    border-radius: 4px;
  }
</style>
