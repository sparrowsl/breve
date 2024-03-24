<script>
  import { goto } from "$app/navigation";
  import { PUBLIC_API_URL } from "$env/static/public";

  // const link = { redirect: "", url: "", random: false };

  /** @param {Event} e  */
  async function submitForm(e) {
    e.preventDefault();

    const res = await fetch(`${PUBLIC_API_URL}/links`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      // body: JSON.stringify(link),
    });
    // const data = await res.json();

    if (res.ok) {
      goto("/");
    }
  }

  /**
   * @param {Event} e
   * @param {Number} id
   */
  async function updateForm(e, id) {
    e.preventDefault();

    const res = await fetch(`${PUBLIC_API_URL}/links/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      // body: JSON.stringify(link),
    });
    // const data = await res.json();

    if (res.ok) {
      goto("/");
    }
  }

  const { data } = $props();
  const link = data?.link;
</script>

<form action="" onsubmit={submitForm} method="POST">
  <legend>Create Link</legend>

  <fieldset class="grid gap-5">
    <input
      type="url"
      placeholder="Redirect to"
      name="redirect"
      value={link?.redirect || ""}
      class="block rounded-md text-sm"
    />
    <input
      type="text"
      placeholder="custom name"
      name="url"
      value={link?.url || ""}
      class="block rounded-md text-sm"
    />
    <label for="random" class="flex items-center gap-2 block w-fit">
      <input
        type="checkbox"
        name="random"
        id="random"
        checked={link?.random || false}
        class="rounded"
      />
      <span>Generate Random</span>
    </label>

    {#if !link}
      <button
        type="submit"
        class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
      >
        Create
      </button>
    {:else}
      <button
        type="button"
        onclick={() => updateForm(e, link?.id)}
        class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
      >
        Update
      </button>
    {/if}
  </fieldset>
</form>
