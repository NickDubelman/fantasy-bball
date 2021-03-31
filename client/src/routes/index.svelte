<svelte:head>
  <title>Fantasy Basketball</title>
</svelte:head>

<h1>Your leagues</h1>

<button type="button">Create a league</button>

<style lang="postcss">
  h1 {
    @apply text-4xl font-serif;
  }

  button {
    @apply bg-blue-500 hover:bg-blue-600 hover:shadow-lg focus:outline-none;
    @apply text-white text-sm mt-3 py-2.5 px-5 rounded-md;
  }
</style>
