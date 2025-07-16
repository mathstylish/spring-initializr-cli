#!/usr/bin/env node
import { spawn } from "child_process";
const yo = spawn("yo", ["spring-cli"], {
    stdio: "inherit",
    shell: true,
});
yo.on("exit", (code) => {
    process.exit(code ?? 0);
});
