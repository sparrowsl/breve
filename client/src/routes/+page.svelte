<script>
  import Card from "./Card.svelte";

  let { data } = $props();
</script>

<h1 class="font-bold text-2xl py-5">Breve - URL Shortner in Go and Svelte</h1>
<section class="grid gap-5">
  {#each data.links as link (link.id)}
    <Card {link} />
  {/each}
</section>
