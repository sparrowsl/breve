<script>
  import Card from "./Card.svelte";

  const { data } = $props();
</script>

<p>Total Links: {data.links.length}</p>
<section class="grid gap-5">
  {#each data.links as link (link.id)}
    <Card {link} />
  {/each}
</section>
