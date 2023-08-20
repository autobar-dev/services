import { Config } from "./config";
import { MetaFactors } from "./meta";
import { TemplateRepository } from "../repositories/TemplateRepository";

export type AppContext = {
  metaFactors: MetaFactors,
  config: Config,
  repositories: {
    template: TemplateRepository,
  },
};
