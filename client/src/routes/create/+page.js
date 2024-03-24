import { PUBLIC_API_URL } from "$env/static/public"

/** @type {import("./$types").PageLoad} */
export async function load({ fetch, url }) {
  const id = Number(url.searchParams.get("update"))

  if (!id) {
    return { link: {} }
  }

  try {
    const res = await fetch(`${PUBLIC_API_URL}/links/${id}`)
    const data = await res.json()
    return { link: data.link }
  } catch (_) {
    return { link: {} }
  }
}
