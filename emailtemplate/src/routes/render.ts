import { Response } from 'express';
import { renderAsync } from '@react-email/render';
import { AppRequest } from '../types/appRequest';

type RenderRequestBody = {
  template_name: string,
  template_version: string,
  locale: string,
  params: Record<string, unknown>,
};

type RenderResponse = {
  status: "ok" | "error",
  error: string | null,
  data: {
    plain: string,
    html: string,
  } | null,
};

export async function renderTemplateRoute(req: AppRequest, res: Response) {
  const templateRepository = req.appContext.repositories.template;

  const body = req.body as unknown as RenderRequestBody;

  const templateName = body.template_name;
  const templateVersion = body.template_version;
  const params = body.params;
  const locale = body.locale;

  const template = templateRepository.get(templateName, templateVersion);

  if(!template) {
    return res.status(404).json({
      status: "error",
      error: "unknown template name or version",
      data: null,
    } satisfies RenderResponse);
  }

  const renderedPlainEmail = await renderAsync(template({
    locale,
    params,
  }), {
    pretty: false,
    plainText: true,
  });
  const renderedHtmlEmail = await renderAsync(template({
    locale,
    params,
  }), {
    pretty: false,
    plainText: false,
  });

  return res.status(200).json({
    status: "ok",
    error: null,
    data: {
      plain: renderedPlainEmail,
      html: renderedHtmlEmail,
    },
  } satisfies RenderResponse);
}
