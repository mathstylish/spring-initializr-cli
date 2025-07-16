import AdmZip from "adm-zip";
import axios from "axios";
import chalk from "chalk";
import fs from "fs";
import path from "path";
export async function downloadAndExtractProject(options) {
    if (!options) {
        console.error(chalk.redBright("❌ Invalid options object passed to downloadAndExtractProject"));
        return;
    }
    try {
        console.log(chalk.blueBright(`\n⏳ Downloading project zip for artifactId: ${options.artifactId} ...`));
        const params = new URLSearchParams();
        params.append("type", options.type);
        params.append("language", options.language);
        params.append("bootVersion", options.bootVersion);
        params.append("baseDir", options.artifactId);
        params.append("groupId", options.groupId);
        params.append("artifactId", options.artifactId);
        params.append("name", options.name);
        params.append("description", options.description);
        params.append("packageName", options.packageName);
        params.append("packaging", options.packagingType);
        params.append("javaVersion", options.javaVersion);
        if (options.dependencies.length > 0) {
            params.append("dependencies", options.dependencies.join(","));
        }
        const url = `https://start.spring.io/starter.zip?${params.toString()}`;
        const response = await axios.get(url, { responseType: "arraybuffer" });
        const zipFilePath = path.resolve(process.cwd(), `${options.artifactId}.zip`);
        fs.writeFileSync(zipFilePath, Buffer.from(response.data));
        console.log(chalk.greenBright(`✅ Download complete. Extracting to ./${options.artifactId}/`));
        const zip = new AdmZip(zipFilePath);
        zip.extractAllTo(process.cwd(), true);
        fs.unlinkSync(zipFilePath);
        console.log(chalk.greenBright(`✅ Project extracted to ./${options.artifactId}`));
    }
    catch (error) {
        const message = error instanceof Error ? error.message : String(error);
        console.error(chalk.redBright(`❌ Error when trying to download or extract the project: ${message}`));
    }
}
