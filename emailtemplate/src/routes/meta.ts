import { Response } from "express";
import { AppRequest } from "../types/appRequest";
import { getMetaFromFactors } from "../utils/meta";
import { Meta } from "../types/meta";

type MetaResponse = {
  status: "ok" | "error",
  error: string | null,
  data: Meta,
};

export async function metaRoute(req: AppRequest, res: Response) {
  const meta = getMetaFromFactors(req.appContext.metaFactors);

  return res.status(200).json({
    status: "ok",
    error: null,
    data: meta,
  } satisfies MetaResponse);
}
