import { Config } from "../types/config";

export function getConfig(): Config {
  const port = process.env.PORT;
  if(!port || isNaN(Number(port))) {
    throw new Error("PORT is not set or is not a number");
  }

  return {
    port: Number(port),
  };
}
