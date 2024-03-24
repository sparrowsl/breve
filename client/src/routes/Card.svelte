<script>
  import { PUBLIC_API_URL } from "$env/static/public";

  const { link } = $props();
</script>

<section class="rounded-md border-2 p-4 max-w-md">
  <div>
    <p>Breve: {PUBLIC_API_URL}/links/r/{link?.url}</p>
    <p>Redirect: {link?.redirect}</p>
    <p>Clicked: {link?.clicked}</p>
  </div>
  <div
    class="flex items-center justify-between mt-5 text-sm text-white font-bold"
  >
    <a href="{PUBLIC_API_URL}/links/r/{link?.url}" class="block">
      <i class="gg-link text-black"></i>
    </a>
    <button class="block bg-teal-600 px-4 py-1 rounded-md">Update</button>
    <button class="block bg-red-500 px-4 py-1 rounded-md">Delete</button>
  </div>
</section>
