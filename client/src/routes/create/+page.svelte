<script>
  import { goto } from "$app/navigation";
  import { PUBLIC_API_URL } from "$env/static/public";

  /** @param {Event} e  */
  async function submitForm(e) {
    e.preventDefault();

    const res = await fetch(`${PUBLIC_API_URL}/links`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        url: link?.url,
        redirect: link?.redirect,
        random: link?.random,
      }),
    });
    const data = await res.json();
    console.log(data);

    if (res.ok) {
      goto("/");
    }
  }

  /** @param {Event} e */
  async function updateForm(e) {
    e.preventDefault();

    const res = await fetch(`${PUBLIC_API_URL}/links/${link?.id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        url: link.url,
        redirect: link.redirect,
        random: link.random,
      }),
    });
    // const data = await res.json();

    if (res.ok) {
      goto("/");
    }
  }

  const { data } = $props();
  const link = $derived(data.link);
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
      disabled={link.random}
      bind:value={link.url}
      class="block rounded-md text-sm disabled:(cursor-not-allowed bg-gray-50)"
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

    {#if link?.blank}
      <button
        type="submit"
        class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
      >
        Create
      </button>
    {:else}
      <button
        type="button"
        onclick={updateForm}
        class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
      >
        Update
      </button>
    {/if}
  </fieldset>
</form>
