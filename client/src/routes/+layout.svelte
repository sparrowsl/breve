<script>
  import "@unocss/reset/tailwind.css";
  import "uno.css";

  const { children } = $props();
</script>

<svelte:head>
  <title>Breve</title>
</svelte:head>

<main class="font-roboto container relative">
  <header class="my-3">
    <h1 class="font-bold text-2xl py-5"><a href="/">Breve - URL Shortner in Go and Svelte</a></h1>
    <a
      href="/create"
      class="block w-fit text-sm text-white bg-gray-800 px-4 py-1 font-bold rounded-sm"
    >
      Create Link
    </a>
  </header>
 
  {@render children?.()}
</main>
