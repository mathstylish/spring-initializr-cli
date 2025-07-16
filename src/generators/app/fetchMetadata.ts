import axios from "axios";
import { SpringMetadata } from "./types.js";

export async function fetchMetadata(): Promise<SpringMetadata> {
  const res = await axios.get<SpringMetadata>("https://start.spring.io/metadata/client");
  return res.data;
}
