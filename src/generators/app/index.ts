import chalk from "chalk";
import fs from "fs/promises";
import os from "os";
import path from "path";
import Generator from "yeoman-generator";

import { downloadAndExtractProject } from "./download.js";
import { fetchMetadata } from "./fetchMetadata.js";

import {
  askAddMoreDeps,
  askBaseProjectInfo,
  askDependencies,
  askPackageName,
  askReuseLastDeps,
  askUseLastConfig,
} from "./prompts.js";

const CONFIG_PATH = path.join(os.homedir(), ".spring-cli-config.json");

async function loadConfig(): Promise<any> {
  try {
    const data = await fs.readFile(CONFIG_PATH, "utf-8");
    return JSON.parse(data);
  } catch {
    return {}; // arquivo não existe ou erro -> retorna vazio
  }
}

async function saveConfig(config: any): Promise<void> {
  await fs.writeFile(CONFIG_PATH, JSON.stringify(config, null, 2), "utf-8");
}

export default class extends Generator {
  private allowedProjectTypes = ["gradle-project", "gradle-project-kotlin", "maven-project"];

  async prompting(): Promise<void> {
    const initializr = await fetchMetadata();

    this.log(
      chalk.bold(`🌱 ${chalk.hex("#6DB33F")("Spring")} ${chalk.hex("#FFFFFF")("Initializr")} CLI\n`)
    );

    const savedConfig = await loadConfig();
    const lastAnswers = savedConfig.lastAnswers;

    if (lastAnswers) {
      const useLast = await askUseLastConfig();

      if (useLast) {
        this.log(chalk.greenBright("\n✅ Using last saved configuration:"));
        await downloadAndExtractProject(lastAnswers);
        return;
      }
    }

    const baseAnswers = await askBaseProjectInfo(initializr, this.allowedProjectTypes);
    const packageName = await askPackageName(`${baseAnswers.groupId}.${baseAnswers.artifactId}`);

    const selectedDependencies: { id: string; name: string }[] = [];

    const allDependenciesFlat = Object.values(initializr.dependencies.values).flatMap((category) =>
      category.values.map((dep) => ({
        name: dep.name,
        value: dep.id,
      }))
    );

    if (lastAnswers?.dependencies?.length) {
      const reuseLastDeps = await askReuseLastDeps();

      if (reuseLastDeps) {
        for (const id of lastAnswers.dependencies) {
          const found = allDependenciesFlat.find((d) => d.value === id);
          if (found) {
            selectedDependencies.push({ id: found.value, name: found.name });
          }
        }

        const addMore = await askAddMoreDeps();

        if (!addMore) {
          const answersToSave = {
            ...baseAnswers,
            packageName,
            dependencies: selectedDependencies.map((d) => d.id),
          };
          await saveConfig({ lastAnswers: answersToSave });
          await downloadAndExtractProject(answersToSave);
          return;
        }
      }
    }

    const finalSelectedDependencies = await askDependencies(allDependenciesFlat, selectedDependencies);

    const answersToSave = {
      ...baseAnswers,
      packageName,
      dependencies: finalSelectedDependencies.map((d) => d.id),
    };

    await saveConfig({ lastAnswers: answersToSave });

    await downloadAndExtractProject(answersToSave);
  }
}
