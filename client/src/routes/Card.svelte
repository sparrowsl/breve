<script>
  import { invalidateAll } from "$app/navigation";
  import { PUBLIC_API_URL } from "$env/static/public";

  /** @param {Number} id  */
  async function deleteLink(id) {
    const deleted = confirm("Sure you want to delete??");
    if (!deleted) {
      return;
    }

    const res = await fetch(`${PUBLIC_API_URL}/links/${id}`, {
      method: "DELETE",
    });
    if (!res.ok) {
      return;
    }

    invalidateAll();
  }

  const { link } = $props();
</script>

<section class="rounded-md border-2 p-4 max-w-md">
  <div>
    <p>Breve: {PUBLIC_API_URL}/links/r/{link?.url}</p>
    <p>Redirect: {link?.redirect}</p>
    <p>Clicked: {link?.clicked}</p>
  </div>
  <div class="flex items-center mt-5 text-sm text-white font-bold gap-6">
    <a
      href="{PUBLIC_API_URL}/links/r/{link?.url}"
      target="_blank"
      class="block mr-auto"
    >
      <i class="gg-link text-black"></i>
    </a>
    <button class="block bg-teal-600 px-4 py-1 rounded-md">Update</button>
    <button
      class="block bg-red-500 px-4 py-1 rounded-md"
      onclick={() => deleteLink(link?.id)}>Delete</button
    >
  </div>
</section>
