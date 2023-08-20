import * as React from 'react';
import {
  Html,
  Head,
  Body,
  Text,
  Hr,
  Link,
  Container,
  Preview,
} from "@react-email/components";

export type EmailConfirmationTemplateV1Params = {
  first_name: string,
  last_name: string,
  email_confirmation_code: string,
};

export function EmailConfirmationTemplateV1(props: {
  locale: string,
  params: EmailConfirmationTemplateV1Params,
}) {
  return (
    <Html lang={props.locale}>
      <Head />
      <Preview>Autobar - Email confirmation</Preview>
      <Body style={bodyStyle}>
        <Container style={containerStyle}>
          <Text>Hi {props.params.first_name} {props.params.last_name},</Text>
          <Text>Thank you for registering at Autobar. (not you? Let us know by replying to this email)</Text>
          <Text>Please confirm your email address by clicking on the link below.</Text>
          <Link href={`https://autobar.co/email-confirmation?code=${props.params.email_confirmation_code}`}>Confirm email address</Link>
          <Hr />
          <Text>Thanks,</Text>
          <Text>Autobar</Text>
        </Container>
      </Body>
    </Html>
  );
}

const bodyStyle = {
  fontFamily: "Roboto,sans-serif",
} as React.CSSProperties;

const containerStyle = {
  margin: "0 auto",
  padding: "20px 0 48px",
} as React.CSSProperties;
