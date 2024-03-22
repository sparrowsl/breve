<script>
  const { link } = $props();
</script>

<div class="rounded-md border-2 p-4 max-w-md">
  <p>Breve: http://localhost:5000/r/{link?.url}</p>
  <p>Redirect: {link?.redirect}</p>
  <p>Clicked: {link?.clicked}</p>
</div>
