import axios from "axios";
export async function fetchMetadata() {
    const res = await axios.get("https://start.spring.io/metadata/client");
    return res.data;
}
