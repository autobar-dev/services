import templates from './templates';

export class TemplateRepository {
  templates: Map<string, Map<string, Function>>;

  constructor() {
      this.templates = templates;
  }

  get(name: string, version: string): Function | null {
    const templateVersions = this.templates.get(name);
    if (!templateVersions) {
      return null;
    }

    const template = templateVersions.get(version);
    if (!template) {
      return null;
    }

    return template;
  }
}
