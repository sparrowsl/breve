import { PUBLIC_API_URL } from "$env/static/public"

/** import("./$types").PageLoad */
export async function load({ fetch }) {
  const req = await fetch(`${PUBLIC_API_URL}/links`)
  const { links } = await req.json()

  return { links }
}
