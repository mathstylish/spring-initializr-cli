import chalk from "chalk";
import inquirer from "inquirer";
import autocomplete from "inquirer-autocomplete-prompt";
inquirer.registerPrompt("autocomplete", autocomplete);
export async function askUseLastConfig() {
    const { useLast } = await inquirer.prompt([
        {
            type: "list",
            name: "useLast",
            message: "Choose a configuration mode:",
            choices: [
                { name: "🕘 Use latest config created", value: true },
                { name: "📦 New config", value: false },
            ],
        },
    ]);
    return useLast;
}
export async function askBaseProjectInfo(initializr, allowedProjectTypes) {
    const projectTypes = {
        type: "list",
        name: "type",
        message: "Project",
        choices: initializr.type.values
            .filter((value) => allowedProjectTypes.includes(value.id))
            .map((value) => ({
            name: value.name,
            value: value.id,
        })),
        default: "maven-project",
    };
    const language = {
        type: "list",
        name: "language",
        message: "Language",
        choices: initializr.language.values.map((value) => ({
            name: value.name,
            value: value.id,
        })),
        default: initializr.language.default,
    };
    const bootVersion = {
        type: "list",
        name: "bootVersion",
        message: "Spring Boot",
        choices: initializr.bootVersion.values.map((value) => ({
            name: value.name,
            value: value.id,
        })),
        default: initializr.bootVersion.default,
    };
    const groupId = {
        type: "input",
        name: "groupId",
        message: "Group:",
        default: initializr.groupId.default,
    };
    const artifactId = {
        type: "input",
        name: "artifactId",
        message: "Artifact:",
        default: initializr.artifactId.default,
    };
    const name = {
        type: "input",
        name: "name",
        message: "Name:",
        default: initializr.name.default,
    };
    const description = {
        type: "input",
        name: "description",
        message: "Description:",
        default: initializr.description.default,
    };
    const packagingType = {
        type: "list",
        name: "packagingType",
        message: "Packaging",
        choices: initializr.packaging.values.map((value) => ({
            name: value.name,
            value: value.id,
        })),
        default: initializr.packaging.default,
    };
    const javaVersion = {
        type: "list",
        name: "javaVersion",
        message: "Java",
        choices: initializr.javaVersion.values.map((value) => ({
            name: value.name,
            value: value.id,
        })),
        default: initializr.javaVersion.default,
    };
    const baseAnswers = await inquirer.prompt([
        projectTypes,
        language,
        bootVersion,
        groupId,
        artifactId,
        name,
        description,
        packagingType,
        javaVersion,
    ]);
    return baseAnswers;
}
export async function askPackageName(defaultPackageName) {
    const { packageName } = await inquirer.prompt([
        {
            type: "input",
            name: "packageName",
            message: "Package:",
            default: defaultPackageName,
        },
    ]);
    return packageName;
}
export async function askReuseLastDeps() {
    const { reuseLastDeps } = await inquirer.prompt([
        {
            type: "confirm",
            name: "reuseLastDeps",
            message: "Use latest dependencies?",
            default: false,
        },
    ]);
    return reuseLastDeps;
}
export async function askAddMoreDeps() {
    const { addMore } = await inquirer.prompt([
        {
            type: "confirm",
            name: "addMore",
            message: "Add more dependencies?",
            default: false,
        },
    ]);
    return addMore;
}
/**
 * Função que retorna a lista filtrada para autocomplete no prompt de dependências
 */
export function createSearchDependencies(allDependenciesFlat, selectedDependencies) {
    return async (_, input = "") => {
        const lowerInput = input.toLowerCase();
        return allDependenciesFlat
            .filter((dep) => !selectedDependencies.find((d) => d.id === dep.value) &&
            (dep.name.toLowerCase().includes(lowerInput) ||
                dep.value.toLowerCase().includes(lowerInput)))
            .map((dep) => ({
            name: dep.name,
            value: dep.value,
        }));
    };
}
export async function askDependencies(allDependenciesFlat, selectedDependencies) {
    const searchDependencies = createSearchDependencies(allDependenciesFlat, selectedDependencies);
    const showSelectedDependencies = () => selectedDependencies.map((d) => `- ${chalk.cyan(d.name)}`).join("\n  ");
    let done = false;
    while (!done) {
        console.clear();
        console.log(chalk.green.bold("📦 Selected dependencies:"));
        console.log("  " + showSelectedDependencies());
        const { dependency } = await inquirer.prompt([
            {
                type: "autocomplete",
                name: "dependency",
                message: `➕ Add another dependency (or press Enter to finish)`,
                source: searchDependencies,
                pageSize: 6,
                suggestOnly: true,
            },
        ]);
        const trimmed = dependency?.trim();
        if (!trimmed) {
            done = true;
            break;
        }
        const depMatch = allDependenciesFlat.find((dep) => dep.value.toLowerCase() === trimmed.toLowerCase() ||
            dep.name.toLowerCase() === trimmed.toLowerCase() ||
            dep.name.toLowerCase().includes(trimmed.toLowerCase()) ||
            dep.value.toLowerCase().includes(trimmed.toLowerCase()));
        if (depMatch && !selectedDependencies.find((d) => d.id === depMatch.value)) {
            selectedDependencies.push({ id: depMatch.value, name: depMatch.name });
        }
    }
    return selectedDependencies;
}
