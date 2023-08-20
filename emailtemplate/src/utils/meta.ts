import { readFile } from "fs/promises"; 
import { join } from "path";
import { Meta, MetaFactors } from "../types/meta";

export async function getMetaFactors(): Promise<MetaFactors> {
  let commit_sha = "";
  try {
    const commit_sha_bytes = await readFile(join(__dirname, "..", "..", ".meta", "HASH"));
    commit_sha = commit_sha_bytes.toString("utf8");
  } catch (_e) {}

  let version = "";
  try {
    const version_bytes = await readFile(join(__dirname, "..", "..", ".meta", "VERSION"));
    version = version_bytes.toString("utf8");
  } catch (_e) {}

  commit_sha = commit_sha.trim() || "";
  version = version.trim() || "";
  
  return {
    startTime: new Date(),
    hash: commit_sha,
    version,
  } satisfies MetaFactors;
}

export function getMetaFromFactors(metaFactors: MetaFactors): Meta {
  return {
    uptime: Math.floor((new Date().getTime() - metaFactors.startTime.getTime()) / 1000),
    hash: metaFactors.hash,
    version: metaFactors.version,
  };
}
