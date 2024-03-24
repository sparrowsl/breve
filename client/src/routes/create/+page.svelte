<script>
  import { goto } from "$app/navigation";
  import { PUBLIC_API_URL } from "$env/static/public";

  const link = { redirect: "", url: "", random: false };

  /** @param {Event} e  */
  async function submitForm(e) {
    e.preventDefault();

    const res = await fetch(`${PUBLIC_API_URL}/links`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(link),
    });
    // const data = await res.json();

    if (res.ok) {
      goto("/");
    }
  }
</script>

<form action="" onsubmit={submitForm} method="POST">
  <legend>Create Link</legend>

  <fieldset class="grid gap-5">
    <input
      type="url"
      placeholder="Redirect to"
      name="redirect"
      bind:value={link.redirect}
      class="block rounded-md text-sm"
    />
    <input
      type="text"
      placeholder="custom name"
      name="url"
      bind:value={link.url}
      class="block rounded-md text-sm"
    />
    <label for="random" class="flex items-center gap-2 block w-fit">
      <input
        type="checkbox"
        name="random"
        id="random"
        bind:checked={link.random}
        class="rounded"
      />
      <span>Generate Random</span>
    </label>

    <button
      type="submit"
      class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
    >
      Create
    </button>
  </fieldset>
</form>
