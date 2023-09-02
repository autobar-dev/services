import express from 'express';
import dotenv from 'dotenv';

import { metaRoute } from "./routes/meta";
import { renderTemplateRoute } from "./routes/render";

import { getConfig } from './utils/config';
import { AppContext } from './types/appContext';
import { getMetaFactors } from './utils/meta';
import { TemplateRepository } from './repositories/TemplateRepository';
import { AppRequest } from './types/appRequest';

new Promise(async (_res, _rej) => {
  dotenv.config();
  const config = getConfig();

  const metaFactors = await getMetaFactors();

  const appContext: AppContext = {
    metaFactors,
    config,
    repositories: {
      template: new TemplateRepository(),
    },
  };

  const app = express();

  app.use(express.json());
  app.use((req: any, _res, next) => {
    req.appContext = appContext;
    next();
  });

  app.get("/meta", (req, res) => metaRoute(req as AppRequest, res));
  app.post("/render", (req, res) => renderTemplateRoute(req as AppRequest, res));

  app.listen(config.port, () => {
    console.log(`Listening on http://localhost:${config.port}`);
  });
});
