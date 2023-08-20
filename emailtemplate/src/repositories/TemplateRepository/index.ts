import templates from './templates';

export class TemplateRepository {
  templates: Map<string, Map<string, Function>>;

  constructor() {
      this.templates = templates;
  }

  get(name: string, version: string | null): Function | null {
    const templateVersions = this.templates.get(name);
    if (!templateVersions) {
      return null;
    }

    if(!version) {
      const latest_template = Array.from(templateVersions.values()).pop();
      if(!latest_template) {
        return null;
      }

      return latest_template;
    }

    const template = templateVersions.get(version);
    if (!template) {
      return null;
    }

    return template;
  }
}
