import { API_BASE } from "@/constants/api";
import { searchFormData } from "@/types/search"

export async function fetchMapInfo(data: searchFormData) {
  console.log(`${JSON.stringify(data)}`);
  if (data.address === "")
      return null;
  const res = await fetch(`${API_BASE}/search`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  });
  if (!res.ok) throw new Error("APIコール失敗");

  return res.json();
}
