import { EmailConfirmationTemplateV1 } from "./EmailConfirmationV1";

const versions = new Map<string, Function>();

versions.set("1", EmailConfirmationTemplateV1);

export default versions;
