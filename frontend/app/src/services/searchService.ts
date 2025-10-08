import { searchFormData } from "@/types/search"

export async function fetchMapInfo(data: searchFormData) {
  const res = await fetch("/api/search", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    throw new Error("Search API failed");
  }

  return await res.json();
}
