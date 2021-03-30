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
  import Nav from '../components/Nav.svelte'

  const { session } = stores()

  export let currPage: string
  let currPageEncoded = ''

  if (typeof window !== 'undefined') {
    currPageEncoded = btoa(currPage)
  }
</script>

<Nav />

<main>
  {#if $session.user}
    <slot />
  {:else}
    <a href="/auth/login?next={currPageEncoded}">Login</a>
  {/if}
</main>

<style lang="postcss">
  main {
    position: relative;
    max-width: 56em;
    background-color: white;
    padding: 2em;
    margin: 0 auto;
    box-sizing: border-box;
  }

  a {
    @apply text-white text-xl py-3.5 px-6 rounded-md bg-purple-600;
  }
</style>
