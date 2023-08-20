
import EmailConfirmationVersions from "./email-confirmation";

const templates = new Map<string, Map<string, Function>>();

templates.set("email-confirmation", EmailConfirmationVersions);

export default templates;
