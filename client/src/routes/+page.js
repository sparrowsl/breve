/** import("./$types").PageLoad */
export async function load({ fetch }) {
  const req = await fetch("http://localhost:5000/links")
  const { links } = await req.json()

  return { links }
}
