<script>
  const link = {
    redirect: "",
    url: "",
    random: false,
  };
</script>

<form action="">
  <legend>Create Link</legend>

  <fieldset class="grid gap-5">
    <input
      type="url"
      placeholder="Redirect to"
      name="redirect"
      bind:value={link.redirect}
      class="block"
    />
    <input
      type="text"
      placeholder="custom name"
      name="url"
      bind:value={link.url}
      class="block"
    />
    <label for="random" class="flex items-center gap-2">
      <input
        type="checkbox"
        name="random"
        id="random"
        bind:checked={link.random}
      />
      <span>Generate Random</span>
    </label>
    <p>{link.random}</p>

    <button
      type="submit"
      class="block text-sm text-white bg-teal-600 px-6 py-1 rounded-sm w-fit"
    >
      Create
    </button>
  </fieldset>
</form>
