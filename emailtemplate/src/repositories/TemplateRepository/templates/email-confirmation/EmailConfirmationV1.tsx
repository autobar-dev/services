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
  Img,
} from "@react-email/components";

export type EmailConfirmationTemplateV1Params = {
  email_confirmation_url: string,
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
          <Img 
            alt="Autobar logo"
            src=""
            width="100"
            style={logoStyle}
          />
          <Text style={largeTextStyle}>Hello!</Text>
          <Text style={smallTextStyle}>Thank you for registering at Autobar. (not you? Let us know by replying to this email)</Text>
          <Text style={smallTextStyle}>Please confirm your email address by clicking the button below.</Text>
          <Link style={linkStyle} href={props.params.email_confirmation_url}>Confirm email address</Link>
          <Hr />
          <Text style={smallTextStyle}>Sincerely,</Text>
          <Text style={smallTextStyle}>Your folks at Autobar</Text>
        </Container>
      </Body>
    </Html>
  );
}

const bodyStyle = {
  fontFamily: "Roboto,sans-serif",
  backgroundColor: "#181818",
  padding: "0 20px",
} as React.CSSProperties;

const logoStyle = {
  margin: "0 auto",
} as React.CSSProperties;

const containerStyle = {
  margin: "0 auto",
  padding: "20px",
  backgroundColor: "#181818",
  border: "1px solid #303030",
  borderRadius: "10px",
} as React.CSSProperties;

const largeTextStyle = {
  fontSize: "24px",
  fontWeight: "bold",
  color: "#f8f8f8",
} as React.CSSProperties;

const smallTextStyle = {
  fontSize: "16px",
  color: "#f8f8f8",
} as React.CSSProperties;

const linkStyle = {
  fontSize: "16px",
  color: "#0000EE",
} as React.CSSProperties;
